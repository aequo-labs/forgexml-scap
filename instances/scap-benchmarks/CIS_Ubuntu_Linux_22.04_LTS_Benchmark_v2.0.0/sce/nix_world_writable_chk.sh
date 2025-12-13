#!/usr/bin/env bash

#
# CIS-CAT Script Check Engine
#
# Name       Date       Description
# -------------------------------------------------------------------
# E. Pinnell 03/27/23   Check world writable files and directories
# E. Pinnell 08/25/23   Modified to add "/sys/fs/selinux/*" to the excluded path list
# E. Pinnell 12/08/23   Modified to add "/sys/*" to the excluded path list (Supersedes the need for /sys/{ANY_DIRECTORY})
# E. Pinnell 02/09/24   Modified to add vfat to excluded fstype amd cleanup a_path array list
# E. Pinnell 03/08/24   Modified to improve efficiency 
# E. Pinnell 05/05/25   Modified to exclude "*/kubelet/*", all of "/run/*", and use arrays for output

a_output=() a_output2=() a_file=() a_dir=()
l_smask='01000' l_limit="50" # Set report output limit

# Initialize array
a_path=(! -path "/run/*" -a ! -path "/proc/*" -a ! -path "*/containerd/*" -a ! -path "*/kubelet/*" -a ! -path "/sys/*" -a ! -path "/snap/*")

# Collect relevant files and directories 
while IFS= read -r l_mount; do
   while IFS= read -r -d $'\0' l_file; do
      if [ -e "$l_file" ]; then
         [ -f "$l_file" ] && a_file+=("$l_file") # Add WR files
         if [ -d "$l_file" ]; then # Add directories w/o sticky bit
            l_mode="$(stat -Lc '%#a' "$l_file")"
            [ ! $(( $l_mode & $l_smask )) -gt 0 ] && a_dir+=("$l_file")
         fi
      fi
   done < <(find "$l_mount" -xdev \( "${a_path[@]}" \) \( -type f -o -type d \) -perm -0002 -print0 2> /dev/null)
done < <(findmnt -Dkerno fstype,target | awk '($1 !~ /^\s*(nfs|proc|cifs|smb|vfat|iso9660|efivarfs|selinuxfs)/ && $2 !~ /^(\/run\/user\/|\/tmp|\/var\/tmp)/){print $2}')

# Create output arrays
if [ "${#a_file[@]}" -gt "$l_limit" ]; then
   a_output2+=("    ** NOTE: **" "    More than \"$l_limit\" world writable files exist" \
   "    only the first \"$l_limit\" will be listed")
fi

if [ "${#a_dir[@]}" -gt "$l_limit" ]; then
   a_output2+=("    ** NOTE: **" "    More than \"$l_limit\" World writable directories without the sticky bit exist" \
   "    only the first \"$l_limit\" will be listed")
fi

if [ "${#a_file[@]}" -le 0 ]; then
   a_output+=("  - No world writable files exist on the local filesystem.")
else
   a_output2+=("" " - There are \"${#a_file[@]}\" World writable files on the system." \
   "   - The following is a list of World writable files:" \
   "${a_file[@]:0:$l_limit}" "   - end of list")
fi

if [ "${#a_dir[@]}" -le 0 ]; then
   a_output+=("  - Sticky bit is set on world writable directories on the local filesystem.")
else
   a_output2+=("" " - There are \"${#a_dir[@]}\" World writable directories without the sticky bit on the system." \
   "   - The following is a list of World writable directories without the sticky bit:" \
   "${a_dir[@]:0:$l_limit}" "   - end of list")
fi 

# Remove arrays
unset a_path; unset a_file; unset a_dir

# Provide output from checks
if [ "${#a_output2[@]}" -le 0 ]; then
   printf '%s\n' "" "- Audit Result:" "  ** PASS **" "${a_output[@]}"
   exit "${XCCDF_RESULT_PASS:-101}"
else
   printf '%s\n' "" "- Audit Result:" "  ** FAIL **" "  * Reasons for audit failure *" "${a_output2[@]}" ""
   [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "- Correctly set:" "${a_output[@]}"
   exit "${XCCDF_RESULT_FAIL:-102}"
fi