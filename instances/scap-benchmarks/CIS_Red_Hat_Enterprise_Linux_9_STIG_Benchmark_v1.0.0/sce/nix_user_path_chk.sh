#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/09/25   Check local interactive user initialization file executable search path
# File: nix_user_path_chk.sh

{
   a_output=() a_output2=()
   while IFS= read -r -d $'\0' l_file; do
      if [ ! "$l_file" = ".bash_history" ]; then
         l_path="$(grep -Psio 'path=\H+' "$l_file" | cut -d= -f2)"
         l_path2="${l_path//\"/}"; l_path2="${l_path2//:/\\n}"; a_out2=()
         while IFS= read -r l_part; do
            if [ -n "$l_part" ]; then
               if ! grep -Psiq '^\h*\$(HOME|PATH)\b' <<< "$l_part"; then
                  a_out2+=(" - File \"$l_file\" includes adding \"$l_part\" to the user's path")
               fi
            fi
         done < <(echo -e "$l_path2")
      fi
      if [ "${#a_out2[@]}" -gt 0 ]; then
         a_output2+=("${a_out2[@]}")
      else
         [ -n "$l_path" ] && a_output+=(" - PATH is correctly set to: $l_path" \
         "    in the file: \"$l_file\"")
      fi
   done < <(find /home -maxdepth 2 -type f -name ".[^.]*" -print0)
   [[ "${#a_output[@]}" -le 0 && "${#a_output2[@]}" -le 0 ]] && \
   a_output=(" - PATH is not set in any dot files in the /home directory")

   # Provide output from checks
   if [ "${#a_output2[@]}" -le 0 ]; then
      printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
      exit "${XCCDF_RESULT_PASS:-101}"
   else
      printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
      [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
      exit "${XCCDF_RESULT_FAIL:-102}"
   fi
}