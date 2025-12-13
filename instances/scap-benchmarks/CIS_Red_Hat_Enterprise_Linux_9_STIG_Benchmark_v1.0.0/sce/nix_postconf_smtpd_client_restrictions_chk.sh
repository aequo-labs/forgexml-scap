#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
# 
# Name           Date       Description
# ------------------------------------------------------------------------
# E. Pinnell     04/07/25   Check  postconf smtpd_client_restrictions

a_output=() a_output2=() l_out=""
if ! rpm -q postfix &>/dev/null; then
   a_output+=(" - The postfix package is not installed. This check in not applicable")
else
   l_out="$(postconf -n smtpd_client_restrictions 2>&1)"
   if grep -Psiq '^\h*smtpd_client_restrictions\h*=\h*permit_mynetworks,reject\h*$' <<< "$l_out"; then
      a_output+=(" - smtpd_client_restrictions is correctly set to:" \
      "    \"$(cut -d= -f2 <<< "$l_out" | xargs)\"")
   else
      a_output2+=(" - smtpd_client_restrictions is incorrectly set to:" \
      "    \"$(cut -d= -f2 <<< "$l_out" | xargs)\"")
   fi
fi

# Provide output to CIS-CAT
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}" ""
   exit "${XCCDF_RESULT_FAIL:-102}"
fi