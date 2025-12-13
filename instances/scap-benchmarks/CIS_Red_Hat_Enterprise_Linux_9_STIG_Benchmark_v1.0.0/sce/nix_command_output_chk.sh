#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   04/18/25   Check command output

# nix_command_output_chk.sh

# XCCDF_VALUE_REGEX="{COMMAND}:{COMMAND_OPTION(S)}:{REGEX_TO_MATCH_EXPECTED_OUTPUT}"
# XCCDF_VALUE_REGEX="systemctl:get-default:^\h*multi-user\.target\b" #<- Example variable

a_output=() a_output2=()

while IFS=: read -r l_command l_option l_regex_match; do
   if ! command -v "$l_command" &>/dev/null; then
      a_output2+=(" Command: \"$l_command\" not found")
   else
      l_out="$($l_command "$l_option")"
      if grep -Psiq -- "$l_regex_match" <<< "$l_out"; then
         a_output+=("  - Command: \"$l_command $l_option\" correctly returned: \"$l_out\"")
      else
         a_output2+=("  - Command: \"$l_command $l_option\" incorrectly returned: \"$l_out\"" \
         "    Return should have matched regex \"$l_regex_match\"")
      fi
   fi
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