#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/11/25   Check faillock log security context

# nix_faillock_log_context_chk.sh

a_output=() a_output2=()

if [ ! -e "/etc/security/faillock.conf" ]; then
   a_output2+=("  - Faillock config file not found")
else
   l_logdir="$(awk -F= '$1 ~ /^[ :space:]*dir[ :space: ]*$/ {print $2}' /etc/security/faillock.conf | xargs)"
   if [ ! -d "$l_logdir" ]; then
      a_output2+=("  - Faillock config file does not include the dir argument")
   else
      l_context="$(ls -Zd "$l_logdir")"
      if grep -Psiq ':faillog_t:' <<< "$l_context"; then
         a_output+=("  - faillock log file directory: \"$l_logdir\"" \
         "    security context is set correctly to:"
         "    \"$l_context\"")
      else
         a_output2+=("  - faillock log file directory: \"$l_logdir\"" \
         "    security context is incorrectly set correctly to:"
         "    \"$l_context\"")
      fi
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