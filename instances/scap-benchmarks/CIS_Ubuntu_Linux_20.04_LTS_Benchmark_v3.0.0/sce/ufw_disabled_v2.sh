#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R.Bejar      02/18/25    Ensure ufw is uninstalled or disabled with nftables & iptables
#

# supersedes ufw_disabled.sh

{
	a_failures=()
	# check if UFW is installed and then check its status
	if command -v ufw &>/dev/null; then
	   # check if UFW is active
	   ufw status | grep -q "Status: active" && a_failures+=("ufw is active")
	   # check if ufw service is masked
	   systemctl is-enabled ufw.service 2>/dev/null | grep -vq "masked" && a_failures+=("ufw service is not masked")
	   if [ ${#a_failures[@]} -gt 0 ]; then
		  for failure in "${a_failures[@]}"; do
			 echo -e "\n- Audit Result:\n ** FAIL **\n $failure"
             exit "${XCCDF_RESULT_FAIL:-102}"
		  done
		else
		   echo -e "\n- Audit Result:\ ** PASS **\n ufw is installed, inactive, and masked"
           exit "${XCCDF_RESULT_PASS:-101}"
		fi
	else
	   echo -e "\n- Audit Result:\ ** PASS **\ UFW is not installed"
       exit "${XCCDF_RESULT_PASS:-101}"
	fi
}