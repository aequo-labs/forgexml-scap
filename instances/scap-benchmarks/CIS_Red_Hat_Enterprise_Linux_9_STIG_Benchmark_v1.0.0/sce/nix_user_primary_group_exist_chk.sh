#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/14/25   interactive users primary group exists check

# nix_user_primary_group_exist_chk.sh

a_output=() a_output2=()
l_valid_shells="^($( awk -F\/ '$NF != "nologin" {print}' /etc/shells | sed -rn '/^\//{s,/,\\\\/,g;p}' | paste -s -d '|' - ))$"
l_users="$(awk -v pat="$l_valid_shells" -F: '$(NF) ~ pat { print $1 " " $(NF-1) }' /etc/passwd | wc -l)"

[ "$l_users" -gt 10000 ] && printf '%s\n' "" "  ** INFO **" \
"  $l_users Local interactive users found on the system" "  This may be a long running check" "  **********"

while IFS=" " read -r l_user l_gid; do
   if [ -z "$l_gid" ]; then
      a_output2+=("  - User: \"$l_user\" primary group is not assigned")
   else
      if ! grep -Psq -- '^[^:\n\r]+:[^:\n\r]*:'"$l_gid"':' /etc/group; then
         a_output2+=("  - User: \"$l_user\" primary group gid \"$l_gid\" does not exist")
      fi
   fi
done <<< "$(awk -v pat="$l_valid_shells" -F: '$(NF) ~ pat { print $1 " " $(NF-3) }' /etc/passwd)"

if [ "${#a_output2[@]}" -le 0 ]; then
   a_output=("  - All local interactive users primary group exists")
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi