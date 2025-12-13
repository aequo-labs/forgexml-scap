#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/10/25   Check local interactive user home directory exist check
# E. Pinnell   04/21/25   Modified to correct bug found in testing

# File: nix_user_home_directory_exist_chk.sh

a_output=() a_output2=()

while IFS=: read -r l_user l_directory; do
   [[ -n "$l_user" && -n "$l_directory" ]] && a_output2+=("  - $l_user: $l_directory")
done < <(pwck -r | grep -v pwck:)

# Provide output from checks
if [ "${#a_output2[@]}" -le 0 ]; then
   a_output=(" - All local interactive user's home directories exist")
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi