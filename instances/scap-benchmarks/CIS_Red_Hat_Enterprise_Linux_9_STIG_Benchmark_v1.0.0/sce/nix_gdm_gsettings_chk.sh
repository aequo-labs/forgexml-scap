#!/usr/bin/env bash

# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/08/25   Check GDM gsettings setting
# E. Pinnell   05/14/25   Modified to use "readlink -e" instead of "readlink -f"
#
# XCCDF_VALUE_REGEX="org.gnome.login-screen banner-message-enable:true:get" #<- Example XCCDF_VALUE_REGEX

a_output=() a_output2=() l_out=""
l_gsettings="$(readlink -e /bin/gsettings || readlink -e /usr/bin/gsettings)"
if command -v "$l_gsettings" &>/dev/null; then
   # if grep -Psiq '(user-db|system-db)' /etc/dconf/profile/*/*; then
      while IFS=: read -r l_setting l_value l_query; do
         # Get the current values of the GSettings keys
         l_out="$("$l_gsettings" "$l_query" $l_setting 2>&1)"
         if grep -Psiq -- ''"$l_value"'\b' <<< "$l_out"; then
            a_output+=(" - \"$l_setting\" is correctly set to: \"$l_out\"")
         else
            a_output2+=(" - \"$l_setting\" is incorrectly set to: \"$l_out\"" \
            "    Should be set to \"$l_value\"")
         fi
      done <<< "$XCCDF_VALUE_REGEX"
   # else
   #    a_output2+=(" - User profile does not exist")
   # fi
else
   a_output+=(" - The gsettings command not found on the system")
fi

# Send test results and output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi