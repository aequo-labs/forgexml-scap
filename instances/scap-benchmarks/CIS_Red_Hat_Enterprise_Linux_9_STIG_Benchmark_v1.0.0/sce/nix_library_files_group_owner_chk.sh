#!/usr/bin/env bash

# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/08/25   Check software libraries group owner
#
# nix_library_files_group_owner_chk.sh

a_output=() a_output2=() a_output3=()
a_directories=("/bin /sbin" "/usr/bin" "/usr/sbin" "/usr/libexec" "/usr/local/bin" "/usr/local/sbin")
l_interactive_shells="^($( awk -F\/ '$NF != "nologin" {print}' /etc/shells | sed -rn '/^\//{s,/,\\\/,g;p}' | paste -s -d '|' - ))$"

while IFS= read -r -d $'\0' l_file_name; do
   while IFS=: read -r l_group l_gid; do
      l_user="$(awk -v pat="$l_gid" -F: '$4 == pat {print $1}' /etc/passwd)"
      if [ -n "$l_user" ]; then
         if grep -Psiq -- "$l_interactive_shells" <<< "$(awk -v pat="$l_user" -F: '$1 == pat {print $(NF)}' /etc/passwd)"; then
            a_output2+=("  - File: \"$l_file_name\" belongs to group: \"$l_group\"" \
            "    which is the primary group of interactive user: \"$l_user\"")
         else
            a_output3+=("  - File: \"$l_file_name\" belongs to group: \"$l_group\"" \
            "    which is the primary group of system account: \"$l_user\"")
         fi
      else
         a_output3+=("  - File: \"$l_file_name\" belongs to group: \"$l_group\"" \
         "    and is not the primary group of any users")
      fi
   done < <(stat -Lc '%G:%g' "$l_file_name")
done < <(find "${a_directories[@]}" ! -group root -print0)

# Provide output to CIS-CAT
[ "${#a_output3[@]}" -gt 0 ] && printf '%s\n' "" "  ** NOTE **" "${a_output3[@]:0:50}" 
if [ "${#a_output2[@]}" -le 0 ]; then
   a_output+=("  - all files in \"${a_directories[*]}\"" \
   "    are correctly group owned by root, a service account's primary group, or is not any user's primary group")
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]:0:50}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi