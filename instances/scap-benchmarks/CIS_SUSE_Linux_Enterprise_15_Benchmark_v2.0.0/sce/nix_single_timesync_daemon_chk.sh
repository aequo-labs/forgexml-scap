#!/usr/bin/env bash
#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R.Bejar      12/10/24   Ensure a single time synchronization daemon is in use
#

{
   active_enabled_service=() services=("systemd-timesyncd.service" "chrony.service" "chronyd.service")
   # Determine which time synchronization daemon is in use
   for service in "${services[@]}"; do         
      if systemctl is-enabled --quiet $service && systemctl is-active --quiet $service; then
         active_enabled_service+=("$service")
      fi
   done
   # Display audit results
   if [ ${#active_enabled_service[@]} -eq 1 ]; then
      printf '%s\n' "" "Audit Results:" " ** PASS **" " - A single time synchronization daemon is in use follow the recommendation in ${active_enabled_service[0]} subsection ONLY"
      exit "${XCCDF_RESULT_PASS:-101}"  
   elif [ ${#active_enabled_service[@]} -eq 0 ]; then
      printf '%s\n' "" " Audit Results:" " ** FAIL **" "- No time synchronization daemon in use or unable to determine time synchronization daemon status"
      exit "${XCCDF_RESULT_FAIL:-102}"
   else
      printf '%s\n' "" " Audit Results:" " ** FAIL **" " - Multiple services are in use: ${active_enabled_service[*]}"
      exit "${XCCDF_RESULT_FAIL:-102}"
   fi  
}   

