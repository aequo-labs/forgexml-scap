#!/usr/bin/env sh

#
# CIS-CAT Script Check Engine
#
# Name       Date       Description
# -------------------------------------------------------------------
# E. Pinnell 03/30/21   Check symbolic link permissions
#

passing="" output="" output2="" output3=""

if [ -e "$REGEX" ]; then
   if stat -Lc "%a" "$REGEX" | grep -Eq '[640][40][40]'; then
      passing=true
   else
      passing=false && output="File: \"$(readlink -f "$REGEX")\" has permissions: \"$(stat -Lc "%a" "$REGEX")\""
   fi

   if [ "$(stat -Lc "%u" "$REGEX")" = 0 ]; then
      [ "$passing" != false ] && passing=true
   else
      passing=false && output2="File: \"$(readlink -f "$REGEX")\" is owned by \"$(stat -Lc "%U" "$REGEX")\""
   fi

   if [ "$(stat -Lc "%g" "$REGEX")" = 0 ]; then
      [ "$passing" != false ] && passing=true
   else
      passing=false && output3="File: \"$(readlink -f "$REGEX")\" belongs to group  \"$(stat -Lc "%G" "$REGEX")\""
   fi
fi


# If passing is true, we pass
if [ "$passing" = true ] ; then
   echo "PASSED!"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   # print the reason why we are failing
   echo "FAILED!"
   [ -n "$output" ] && echo "$output"
   [ -n "$output2" ] && echo "$output2"
   [ -n "$output3" ] && echo "$output3"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi