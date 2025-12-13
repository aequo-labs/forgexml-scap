#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   02/13/25   Auditd privileged commands check
# Note:
# Replaces deprecated:
# - auditd_privileged_commands_rules_file.sh
# - auditd_privileged_commands.sh
# - nix_auditd_privileged_commands_rule_chk.sh

# Set variables
a_output=() a_output2=() a_out2=() l_running_config=""
a_option_check=("-a\h+(always,exit|exit,always)" "-F\h+perm=x" "-F\h+auid>=$(awk '/^\s*UID_MIN/{print $2}' /etc/login.defs)" "-F\h+auid!=(unset|-1|4294967295)")

# auditd rules check function
f_auditd_rules_chk()
{
   if [ -n "$l_check_line" ]; then
      for l_option_check in "${a_option_check[@]}"; do
         if ! grep -Pq -- "$l_option_check" <<< "$l_check_line"; then
            a_out2+=("  - option: \"$l_option_check\" is missing")
         fi
      done
   else
      a_out2+=("  - auditd rule doesn't exist in $l_check_type")
   fi
}

# Check for auditctl status
f_auditctl_chk()
{
   l_running_config=""
   if command -v auditctl &>/dev/null; then
      # Build array with auditd rules in the running config (run once instead of multiple times)
      a_auditd_rule=()
      while IFS= read -r l_auditd_rule; do
         [ -n "$l_auditd_rule" ] && a_auditd_rule+=("$l_auditd_rule")
      done < <(auditctl -l)
      if (( "${#a_auditd_rule[@]}" != 0 )); then
         l_running_config="Y"
      else
         l_running_config="  - No rules are loaded in the auditd running configuration"
      fi
   else
      l_running_config="  - auditctl command not found on the system"
   fi
}

# Main check script
#Build privileged files path exclusion array
a_path=(! -path \"/run/user/*\")
while IFS= read -r l_exclude_path; do
   a_path+=( -a ! -path \""$l_exclude_path"/*\")
#done < <(findmnt -krn | awk '/(noexec|nodev)/{print $1}')
done < <(findmnt -Dkerno fstype,target,options | awk '($1~/^(tmpfs|vfat|fuse)/ || $3~/noexec/){print $2}')
#printf '%s ' "${a_path[@]}"
# Check if rule exists for file
while IFS= read -r -d $'\0' l_file; do
   a_out2=()
   l_check_type="an auditd rules file"
   l_check_line="$(awk '/path='"${l_file//\//\\/}"'/{print}' /etc/audit/rules.d/*.rules 2>/dev/null)"
   f_auditd_rules_chk
   [ -z "$l_running_config" ] && f_auditctl_chk
   if [ "$l_running_config" = "Y" ]; then
      l_check_type="the auditd running configuration"
      l_check_line="$(awk '/path='"${l_file//\//\\/}"'/{print}' <<< "${a_auditd_rule[@]}" 2>/dev/null)"
      f_auditd_rules_chk
   else
      a_out2+=("$l_running_config")
   fi
   # Create output for file
   if [ "${#a_out2[@]}" -gt 0 ]; then
      a_output2+=(" - Privileged file: \"$l_file\":${a_out2[@]}" "")
   else
      a_output+=(" - Privileged file: \"$l_file\" auditd rule:" \
      "  - exists in a rules file" "  - exists in the running configuration" "")
   fi
done < <(find / \( "${a_path[@]}" \) -xdev -perm /6000 -type f -print0 2> /dev/null)
#done < <(findmnt -Dkerno fstype,target,options | awk '($1!~/^(tmpfs|vfat|fuse)/ && $3!~/noexec/){print $2}')

# Send check results and output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi