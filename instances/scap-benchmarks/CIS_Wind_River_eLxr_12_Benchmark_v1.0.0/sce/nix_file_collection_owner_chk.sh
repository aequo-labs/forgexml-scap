#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   08/18/23   check file collection owner check (Three part colon separated variable. {DIRECTORY}:{OWNER}:{EXTENSION}) example: /usr/sbin:root:*.conf
# E. Pinnell   04/07/25   Modified to improve efficiency and assessment evidence output

# XCCDF_VALUE_REGEX="/etc/systemd:(alma|adm):*.conf" #<- example XCCDF_VALUE_REGEX variable
# XCCDF_VALUE_REGEX="/etc/:root:cron*"

a_output=() a_output2=() # l_count="0"

while IFS=: read -r l_directory l_owner l_ext; do

   while IFS= read -r -d $'\0' l_fname; do
      l_file_owner="$(stat -Lc '%U' "$l_fname")"
      if grep -Psiq "\b$l_owner\b" <<< "$l_file_owner"; then
         a_output+=(" - file: \"$l_fname\" is correctly owned by: \"$l_file_owner\"" \
         "    (should be owned by: \"${l_owner/|/ or }\")")
      else
         a_output2+=(" - file: \"$l_fname\" is owned by: \"$l_file_owner\"" \
         "    (should be owned by: \"${l_owner/|/ or }\")")
      fi
   done < <(find "$l_directory" -type f -name "$l_ext" -print0)

   # Account for assessment evidence output limitations
   if [ "${#a_output[@]}" -gt 25 ]; then
      if [ "${#a_output2[@]}" -le 0 ]; then
         a_output=(" - All files matching \"$l_ext\" in: \"$l_directory\"" \
         "    are owned by: \"${l_owner/|/ or }\"")
      else
         a_output=(" - All files not listed in audit failure list; matching \"$l_ext\" in: \"$l_directory\"" \
         "    are owned by: \"${l_owner/|/ or }\"")
      fi
   fi

   # Account for no matching files exist
   [[ "${#a_output[@]}" -le 0 && "${#a_output2[@]}" -le 0 ]] && \
   a_output=("  - No files matching \"$l_ext\" exist in: \"$l_directory\"")

done <<< "$XCCDF_VALUE_REGEX"

# Provide output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi