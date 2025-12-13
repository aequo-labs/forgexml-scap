#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# rbejar       11/20/2025  pam-config tool checks for PAM Module pwhistory

# XCCDF_VALUE_REGEX="remember\s*=\s*(2[4-9]|[3-9][0-9])" # <- Example XCCDF_VALUE_REGEX variable

# Define variables
l_output="" l_output2=""

# Query the pam-config module
module_query=$(pam-config --query --pwhistory 2>/dev/null)
# Check if pwhistory module is enabled and setting
if echo -e "$module_query" | grep -q "password"; then
   if echo -e "$module_query" | grep -Pq "$XCCDF_VALUE_REGEX"; then
      l_output="pwhistory setting meets the required conditions"
   else
      l_output2="pwhistory setting does not meet the required conditions"
    fi
else
   l_output="pwhistory module is not enabled in pam-config"
fi         
     
# CIS-CAT Assessment Evidence output
# if "#l_output2 is empty, we pass"
if [ -z "$l_output2" ]; then
   echo -e "\n- Audit Result:\n  *** PASS ***\n- * Correctly configured * :\n$l_output\n"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   echo -e "\n- Audit Result:\n  ** FAIL **\n - * Reasons for audit failure * :\n$l_output2\n"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi