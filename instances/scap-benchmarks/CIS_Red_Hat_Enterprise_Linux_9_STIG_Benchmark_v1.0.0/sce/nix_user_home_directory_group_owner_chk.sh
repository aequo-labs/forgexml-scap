#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/10/25   Check local interactive user home directory group owner

# nix_user_home_directory_group_owner_chk.sh

a_output=() a_output2=() a_output3=()
l_valid_shells="^($( awk -F\/ '$NF != "nologin" {print}' /etc/shells | sed -rn '/^\//{s,/,\\\\/,g;p}' | paste -s -d '|' - ))$"
l_users="$(awk -v pat="$l_valid_shells" -F: '$(NF) ~ pat { print $1 " " $(NF-1) }' /etc/passwd | wc -l)"

[ "$l_users" -gt 10000 ] && printf '%s\n' "" "  ** INFO **" \
"  $l_users Local interactive users found on the system" "  This may be a long running check" "  **********"

while IFS=" " read -r l_user l_user_gid l_home; do
   if [ -z "$l_home" ]; then
      a_output3+=("  - User: \"$l_user\" Home Directory is not assigned")
   elif [ ! -d "$l_home" ]; then
      a_output3+=("  - User: \"$l_user\" Home Directory does not exist")
   else
      l_gid="$(stat -Lc '%g' "$l_home")"
      [ ! "$l_gid" = "$l_user_gid" ] && a_output2+=("  - User: \"$l_user\" home directory: \"$l_home\"" \
      "    is group owned by: \"$(awk -F: '$3=='"$l_gid"' {print $1}' /etc/group)\"" \
      "    and should be group owned by: \"$(awk -F: '$3=='"$l_user_gid"' {print $1}' /etc/group)\"\"")
   fi
done <<< "$(awk -v pat="$l_valid_shells" -F: '$(NF) ~ pat { print $1 " " $4 " " $(NF-1) }' /etc/passwd)"

[ "${#a_output3[@]}" -gt 0 ] && printf '%s\n' "" "  ** WARNING **" "${a_output3[@]}"
if [ "${#a_output2[@]}" -le 0 ]; then
   a_output=("  - All local interactive users home directories sre group owned by their primary group")
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi