#!/usr/bin/env bash
#
#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# rbejar       11/20/2025   SUSE Ensure a single firewall configuration utility is in use

{
   active_firewall=() firewalls=("firewalld" "susefirewall2")
   # Determine which firewall is in use
   for firewall in "${firewalls[@]}"; do
      case $firewall in
         firewalld|susefirewall2)
            cmd=$firewall
      esac          
      if command -v $cmd &> /dev/null && systemctl is-enabled --quiet $firewall && systemctl is-active --quiet $firewall; then
         active_firewall+=("$firewall")
      fi
   done
   # Display audit results
   if [ ${#active_firewall[@]} -eq 1 ]; then
      printf '%s\n' "" "Audit Results:" " ** PASS **" " - A single firewall is in use follow the recommendation in ${active_firewall[0]} subsection ONLY"
      exit "${XCCDF_RESULT_PASS:-101}"
   elif [ ${#active_firewall[@]} -eq 0 ]; then
      printf '%s\n' "" " Audit Results:" " ** FAIL **" "- No firewall in use or unable to determine firewall status"
      exit "${XCCDF_RESULT_FAIL:-102}"  
   else
      printf '%s\n' "" " Audit Results:" " ** FAIL **" " - Multiple firewalls are in use: ${active_firewall[*]}"
      exit "${XCCDF_RESULT_FAIL:-102}"
   fi  
}


