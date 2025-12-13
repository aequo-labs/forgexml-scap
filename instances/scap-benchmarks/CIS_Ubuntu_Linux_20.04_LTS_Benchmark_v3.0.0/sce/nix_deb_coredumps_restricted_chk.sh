#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R.Bejar      01/29/25   Ensure core dumps are restricted

{
   a_output=() a_output2=()
   l_config_file="/etc/systemd/coredump.conf"
   l_config_params=("Storage=none" "ProcessSizeMax=0")
   # check if package is installed
   if ! dpkg -l | grep -q "systemd-coredump"; then
      a_output+=("systemd-coredump not installed")
      exit 0
   fi

   # check config file for settings   
   if [[ -f "$l_config_file" ]]; then
      for param in "${l_config_params[@]}"; do
         if grep -q "^$param" "$l_config_file"; then
            a_output+=("$param is set in $l_config_file")         
         else
            a_output2+=("$param not found in $l_config_file")
         fi
      done
   else
      a_output2+=("$l_config_file not found")
   fi               
    
   # Audit Results
   if [ "${#a_output2[@]}" -le 0 ]; then
      printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
	  exit "${XCCDF_RESULT_PASS:-101}"
   else
      printf '%s\n' "" "- Audit Result:" "  ** FAIL **" "  * Reasons for audit failure *" "${a_output2[@]}" ""
      [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}" ""
	  exit "${XCCDF_RESULT_FAIL:-102}"
   fi 
}