#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   11/13/19   Check SELinux mode

# nix_selinux_mode_chk.sh

# XCCDF_VALUE_REGEX="Enforcing" #<- Example XCCDF_VALUE_REGEX
# XCCDF_VALUE_REGEX="Permissive|"Enforcing"

a_output=() a_output2=()
l_getenforce="$(readlink -f /sbin/getenforce || readlink -f /usr/sbin/getenforce)"
l_status="$("$l_getenforce" 2>&1)"

if grep -Psiq -- "$XCCDF_VALUE_REGEX" <<< "$l_status"; then
   a_output+=("  - SELinux mode is: \"$l_status\"")
else
   a_output2+=("  - SELinux mode is: \"$l_status\"" \
   "    and should be: \"${XCCDF_VALUE_REGEX//|/ OR }\"")
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