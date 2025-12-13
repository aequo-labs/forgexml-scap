#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   12/18/24   Check if apt package is latest version available

# Script name: nix_deb_latest_package_version_chk.sh
# XCCDF_VALUE_REGEX="aide" #<- Example XCCDF_VALUE_REGEX

a_output=() a_output2=()
if ! dpkg-query -s "$XCCDF_VALUE_REGEX" &>/dev/null; then
   a_output2+=(" - Package: \"$XCCDF_VALUE_REGEX\" does not exist on the system")
else
   while IFS=: read -r l_installed l_available; do
      a_output2+=("  - Package \"$XCCDF_VALUE_REGEX\" is version: \"${l_installed//]/}\"" \
      "    and has a newer version: \"${l_available//]/} available")
   done < <(apt list --upgradable 2>&1 | awk '$1~/^'"$XCCDF_VALUE_REGEX"'/{print $6":"$2}')
fi

if [ "${#a_output2[@]}" -le 0 ]; then
   a_output+=("  - Package \"$XCCDF_VALUE_REGEX\" is installed and the latest version")
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi