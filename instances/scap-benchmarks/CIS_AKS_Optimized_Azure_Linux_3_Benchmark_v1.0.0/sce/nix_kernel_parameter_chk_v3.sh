#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   07/01/24   Check system kernel parameters
# E. Pinnell   07/22/24   Modified to use more efficient check
# E. Pinnell   11/21/24   Modified to work on systems without the unified filesystem
# E. Pinnell   05/08/25   Modified to change "readlink -f" to "readlink -e" which resolves a potential error
# E. Pinnell   06/13/25   Modified to account for potential false positive

# Note:
# - Allows simple regex to allow for more than one acceptable value
# - Supersedes the deprecated nix_kernel_parameter_chk_v2.sh

# XCCDF_VALUE_REGEX="kernel.yama.ptrace_scope=(1|2|3)" #<- Example XCCDF_VALUE_REGEX variable

#Initialize arrays and variables
a_output=() a_output2=() a_files=() l_ipv6_disabled=""

# Set variables
l_systemdsysctl="$(readlink -e /lib/systemd/systemd-sysctl || readlink -e /usr/lib/systemd/systemd-sysctl)"
l_ufw_file="$([ -f /etc/default/ufw ] && awk -F= '/^\s*IPT_SYSCTL=/ {print $2}' /etc/default/ufw)"

# build array of current files containing kernel parameter parameters in order of presence
f_build_files_array()
{
   a_files=()
   [ -f "$(readlink -e "$l_ufw_file")" ] && a_files+=("$l_ufw_file")
   a_files+=("/etc/sysctl.conf")
   while IFS= read -r l_fname; do
      l_file="$(readlink -e "${l_fname//# /}")"
      [ -n "$l_file" ] && ! grep -Psiq -- '(^|\h+)'"$l_file"'\b' <<< "${a_files[*]}" && a_files+=("$l_file")
   done < <("$l_systemdsysctl" --cat-config | tac | grep -Pio '^\h*#\h*\/[^#\n\r\h]+\.conf\b')
}

# Function to check if IPv6 is enabled
f_ipv6_chk()
{
   l_ipv6_disabled="no"
   ! grep -Pqs -- '^\h*0\b' /sys/module/ipv6/parameters/disable && l_ipv6_disabled="yes"
   if sysctl net.ipv6.conf.all.disable_ipv6 | grep -Pqs -- "^\h*net\.ipv6\.conf\.all\.disable_ipv6\h*=\h*1\b" && \
      sysctl net.ipv6.conf.default.disable_ipv6 | grep -Pqs -- "^\h*net\.ipv6\.conf\.default\.disable_ipv6\h*=\h*1\b"; then
      l_ipv6_disabled="yes"
   fi
}

# Function to check kernel parameter values 
f_kernel_parameter_chk()
{
   # Check kernel parameter in the running configuration
   l_running_parameter_value="$(sysctl "$l_parameter_name" | awk -F= '{print $2}' | xargs)"
   if [ "$l_running_parameter_value" = "$l_parameter_value" ] || \
   grep -Pq -- '\b'"$l_parameter_value"'\b' <<< "$l_running_parameter_value"; then
      a_output+=("  - Parameter: \"$l_parameter_name\"" \
      "    correctly set to \"$l_running_parameter_value\" in the running configuration")
   else
      a_output2+=("  - Parameter: \"$l_parameter_name\"" \
      "    is incorrectly set to \"$l_running_parameter_value\" in the running configuration" \
      "    Should be set to: \"$l_value_out\"")
   fi

   # Check kernel parameter value loaded from the configuration files
   [ "${#a_files[@]}" -le "0" ] && f_build_files_array
   l_grep="${l_parameter_name//./(\\.|\\/)}"
   a_out=() a_out2=()
   for l_file in "${a_files[@]}"; do
      l_opt="$(grep -Poi '^\h*'"$l_grep"'\h*=\h*\H+\b' "$l_file" | tail -n 1)"
      l_option_value="$(cut -d= -f2 <<< "$l_opt" | xargs)"
      if [ -n "$l_option_value" ]; then
         if [ "$l_option_value" = "$l_parameter_value" ] || \
         grep -Pq -- '\b'"$l_parameter_value"'\b' <<< "$l_option_value"; then
            a_out+=(" - \"$l_parameter_name\" is correctly set to: \"$l_option_value\"" \
            "    in: \"$l_file\"")
         else
            a_out2+=(" - \"$l_parameter_name\" is incorrectly set to: \"$l_option_value\"" \
            "    in: \"$l_file\"" \
            "    and should be set to: \"$l_value_out\"")
         fi
         break
      fi
   done
   if [[ "${#a_out[@]}" -eq 0 && "${#a_out2[@]}" -eq 0 ]]; then
      a_output2+=("  - Parameter: \"$l_parameter_name\" is not set in an included file" \
      "  *** Note: \"$l_parameter_name\" May be set in a file that's ignored by load procedure ***")
   else
      [ "${#a_out[@]}" -gt 0 ] && a_output+=("${a_out[@]}")
      [ "${#a_out2[@]}" -gt 0 ] && a_output2+=("${a_out2[@]}")
   fi
}

# Main loop for XCCDF_VALUE_REGEX value
while IFS="=" read -r l_parameter_name l_parameter_value; do # Check parameters
   l_parameter_name="${l_parameter_name// /}"; l_parameter_value="${l_parameter_value// /}"
   l_value_out="${l_parameter_value//-/ through }"; l_value_out="${l_value_out//|/ or }"
   l_value_out="$(tr -d '(){}' <<< "$l_value_out")"
   if grep -q '^net.ipv6.' <<< "$l_parameter_name"; then
      [ -z "$l_ipv6_disabled" ] && f_ipv6_chk
      if [ "$l_ipv6_disabled" = "yes" ]; then
         a_output+=(" - IPv6 is disabled on the system, \"$l_parameter_name\" is not applicable")
      else
         f_kernel_parameter_chk
      fi
   else
      f_kernel_parameter_chk
   fi
done <<< "$XCCDF_VALUE_REGEX"

# Send test results and output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi