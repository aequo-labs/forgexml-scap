#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
# 
# Name           Date       Description
# ------------------------------------------------------------------------
# E. Pinnell     03/05/25   Check systemd parameter version 4
# E. Pinnell     05/08/25   Modified to change "readlink -f" to "readlink -e" which resolves a potential error
#
# Note: 
# - (XCCDF_VALUE_REGEX colon separated parameter then absolute path to file)
#  - Field 1 - main systemd configuration file
#  - Field 2 - Systemd configuration file block name (DO NOT INCLUDE SQUARE BRACKETS)
#  - Filed 3 - Option name
#  - Field 4 - Option expression
#    - Less than {lt}
#    - Less than or equal to {le}
#    - Greater than {gt}
#    - Greater than or equal to {ge}
#    - Equal to {eq} (This can be an integer or a string)
# - Limitation, regex can not include a colon
# - Supersedes deprecated nix_systemd_parameter_chk_v2.sh and nix_systemd_parameter_chk_v3.sh

# XCCDF_VALUE_REGEX="systemd/logind.conf:Login:StopIdleSessionSec:le:900" # <- Example
# XCCDF_VALUE_REGEX="systemd/journald.conf:Journal:Compress:eq:yes"
# XCCDF_VALUE_REGEX="systemd/logind.conf:Login:KillUserProcesses:eq:yes"
# XCCDF_VALUE_REGEX="systemd/coredump.conf:Coredump:ProcessSizeMax:eq:0"

a_output=() a_output2=() a_output3=() l_out="" l_out2="" l_opt=""
l_analyze_cmd="$(readlink -e /bin/systemd-analyze || readlink -e /usr/bin/systemd-analyze)"

f_pass_output()
{
   if [ "${#a_output[@]}" -gt 0 ] || [ "${#a_output2[@]}" -gt 0 ]; then
      a_output3+=("  - $l_option is correctly set to $l_option_value" "    in $l_file" \
      "    but this setting will be ignored do to load preference")
   else
      if [ -n "$l_out2" ]; then
         a_output+=("$l_out2" "  - Default for $l_option is correctly set to $l_option_value" "$l_out")
      else
         a_output+=("  - $l_option is correctly set to $l_option_value" "    in $l_file" "$l_out")
      fi
   fi
}

f_fail_output()
{
   if [ "${#a_output[@]}" -gt 0 ] || [ "${#a_output2[@]}" -gt 0 ]; then
      a_output3+=("  - $l_option is incorrectly set to $l_option_value" "    in $l_file" \
      "    but this setting will be ignored do to load preference")
   else
      if [ -n "$l_out2" ]; then
         a_output2+=("$l_out2" "  - Default for $l_option is incorrectly set to $l_option_value" "$l_out")
      else
         a_output2+=("  - $l_option is incorrectly set to $l_option_value" "    in $l_file" "$l_out")
      fi
   fi
}

f_option_chk()
{
   case "${l_exp,,}" in
      lt )
         l_out="    and should be less than $l_value"
         if [ "$l_option_value" -lt "$l_value" ]; then
            f_pass_output
         else
            f_fail_output
         fi ;;
      le )
         l_out="    and should be less than or equal to $l_value"
         if [ "$l_option_value" -le "$l_value" ]; then
            f_pass_output
         else
            f_fail_output
         fi ;;
      gt )
         l_out="    and should be greater than $l_value"
         if [ "$l_option_value" -gt "$l_value" ]; then
            f_pass_output
         else
            f_fail_output
         fi ;;
      ge )
         l_out="    and should be greater than or equal to $l_value"
         if [ "$l_option_value" -ge "$l_value" ]; then
            f_pass_output
         else
            f_fail_output
         fi ;;
      * )
         l_out="    and should be equal to $l_value"
         if [ "$l_option_value" = "$l_value" ]; then
            f_pass_output
         else
            f_fail_output
         fi ;;
   esac
}

while IFS=: read -r l_conf_file l_block l_option l_exp l_value; do
   while IFS= read -r l_file; do
      l_file="${l_file//# /}"
      l_opt="$(awk '/\['"$l_block"'\]/{a=1;next}/\[/{a=0}a' "$l_file" 2>/dev/null | grep -Poi '^\h*'"$l_option"'\h*=\h*\H+\b' | tail -n 1)"
      if [ -n "$l_opt" ]; then
         l_option_value="$(cut -d= -f2 <<< "$l_opt" | xargs)"
         f_option_chk
      fi
   done < <("$l_analyze_cmd" cat-config "$l_conf_file" | tac | grep -Pio '^\h*#\h*\/[^#\n\r\h]+\.conf\b')
   # If nothing is explicitly set, check default
   if [ "${#a_output[@]}" -le 0 ] && [ "${#a_output2[@]}" -le 0 ]; then
      l_file="/etc/$l_conf_file"
      l_opt="$(awk '/\['"$l_block"'\]/{a=1;next}/\[/{a=0}a' "$l_file" 2>/dev/null | grep -Poim 1 '^(\h*#)?\h*'"$l_option"'\h*=\h*\H+\b')"
      if [ -n "$l_opt" ]; then
         l_option_value="$(cut -d= -f2 <<< "${l_opt//# /}" | xargs)"
         l_out2=" - Note: default value \"${l_opt//#/}\" is being used in the configuration"
         f_option_chk
      fi
   fi
done <<< "$XCCDF_VALUE_REGEX"

# Send check results and output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   [ "${#a_output3[@]}" -gt 0 ] && printf '%s\n' "  ** Note: **" "${a_output3[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   [ "${#a_output3[@]}" -gt 0 ] && printf '%s\n' "  ** Note: **" "${a_output3[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi