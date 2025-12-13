#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   12/11/19   Check ufw open ports
# E. Pinnell   03/12/21   Modified to allow for entries without protocol
# E. Pinnell   06/21/22   Modified to allow for IPv6 and simplification 

output=""

ufw_out="$(ufw status verbose)"
for lpn in $(ss -tuln | awk '($5!~/%lo:/ && $5!~/127.0.0.1:/ && $5!~/::1/) {split($5, a, ":"); print a[2]}' | sort | uniq); do
	! grep -Pq "^\h*$lpn\b" <<< "$ufw_out" && output="$output\n- Port: \"$lpn\" is missing a firewall rule"
done

# for i in $( ss -4tuln | grep LISTEN | grep -Ev "(127\.0\.0\.1|\:\:1)" | sed -E "s/^(\s*)(tcp|udp)(\s+\S+\s+\S+\s+\S+\s+\S+:)(\S+)(\s+\S+\s*$)/\4/") ; do
#	 ufw status | grep -Eq -- "$i(\/(tcp|udp))?\s+.*(ALLOW|DENY)" || passing=""
# done

# If the regex matched, output would be generated.  If so, we pass
if [ -z "$output" ] ; then
	echo -e "PASS:\nAll listening ports have a firewall rule"
	exit "${XCCDF_RESULT_PASS:-101}"
else
    # print the reason why we are failing
	echo -e "FAIL:\n$output"
	exit "${XCCDF_RESULT_FAIL:-102}"
fi