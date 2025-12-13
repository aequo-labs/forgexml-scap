#!/usr/bin/env sh

#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# E. Pinnell   12/26/19   Check if filesystem is installed with modpobe --showconfig

modprobe --showconfig | grep -Eq "^\s*install\s+$REGEX\s+\/bin\/(true|false)\s*$" && passing=true


# If the regex matched, output would be generated.  If so, we pass
if [ "$passing" = true ] ; then
	echo "filesystem \"$REGEX\" is not loadable"
	exit "${XCCDF_RESULT_PASS:-101}"
else
    # print the reason why we are failing
	echo "filesystem \"$REGEX\" is loadable"
	exit "${XCCDF_RESULT_FAIL:-102}"
fi
