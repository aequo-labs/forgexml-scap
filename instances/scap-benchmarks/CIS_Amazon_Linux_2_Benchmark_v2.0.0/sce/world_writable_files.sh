#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
# 
# Name       Date       Description
# -------------------------------------------------------------------
# B. Munyan  7/13/16    Ensure no world-writable files exist
# B. Munyan  2/07/17    Eliminate /sys from this check as well
# B. Munyan  02/04/19   Unix line endings
# E. Pinnell 01/30/23   changed environment to bash, updated audit for better efficiency, improved output
# E. Pinnell 03/27/23   Changed approach, avoided limitation with xargs. Caped output size to help avoid potential heap memory issue
# E. Pinnell 05/06/25   Modified to follow approach of new nix_world_writable_chk.sh

# Note: this SEC script is deprecated

a_output=() a_output2=() a_file=()
l_limit="50" # Set report output limit

# Initialize array
a_path=(! -path "/run/*" -a ! -path "/proc/*" -a ! -path "*/containerd/*" -a ! -path "*/kubelet/*" -a ! -path "/sys/*" -a ! -path "/snap/*")

# Collect relevant files and directories 
while IFS= read -r l_mount; do
   while IFS= read -r -d $'\0' l_file; do
      [ -e "$l_file" ] && a_file+=("$(readlink -f "$l_file")") # Add WR files
   done < <(find "$l_mount" -xdev \( "${a_path[@]}" \) -type f  -perm -0002 -print0 2> /dev/null)
done < <(findmnt -Dkerno fstype,target | awk '($1 !~ /^\s*(nfs|proc|cifs|smb|vfat|iso9660|efivarfs|selinuxfs)/ && $2 !~ /^(\/run\/user\/|\/tmp|\/var\/tmp)/){print $2}')

# Create output arrays
if [ "${#a_file[@]}" -gt "$l_limit" ]; then
   a_output2+=("    ** NOTE: **" "    More than \"$l_limit\" world writable files exist" \
   "    only the first \"$l_limit\" will be listed")
fi

if [ "${#a_file[@]}" -le 0 ]; then
   a_output+=("  - No world writable files exist on the local filesystem.")
else
   a_output2+=("" " - There are \"${#a_file[@]}\" World writable files on the system." \
   "   - The following is a list of World writable files:" \
   "${a_file[@]:0:$l_limit}" "   - end of list")
fi

# Remove arrays
unset a_path; unset a_file

# Provide output from checks
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" "  * Reasons for audit failure *" "${a_output2[@]}" ""
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi