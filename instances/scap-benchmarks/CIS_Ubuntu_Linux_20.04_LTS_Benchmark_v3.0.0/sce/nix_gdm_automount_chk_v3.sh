#!/usr/bin/env bash
#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R. Bejar     02/12/25   SCE Ensure GDM automatic mounting of removable media is disabled
#
# supersedes nix_gdm_automount_chk.sh and nix_gdm_automount_chk_v2.sh

{
   if grep -Psi "user-db|system-db" /etc/dconf/profile/*/*; then
      # Get the current values of the GSettings keys
      l_automount=$(gsettings get org.gnome.desktop.media-handling automount)
      l_automount_open=$(gsettings get org.gnome.desktop.media-handling automount-open)

      # Check if the automount is false
      if [ "$l_automount" == "false" ]; then
         # Check if the automount-open is false
         if [ "$l_automount_open" = "false" ]; then
            echo -e "Audit Result:\n ** PASS **\n automount & automount-open is set correctly to false"
            exit "${XCCDF_RESULT_PASS:-101}"
         else
            echo -e "Audit Result:\n  ** FAIL **\n automount-open is set to true."
            exit "${XCCDF_RESULT_FAIL:-102}"
         fi
      else
         echo -e "Audit Result:\n  ** FAIL **\n automount is set to true."  
         exit "${XCCDF_RESULT_FAIL:-102}" 
      fi
   else
		echo -e "Audit Result:\n ** FAIL **\n User profile does not exist."
        exit "${XCCDF_RESULT_FAIL:-102}"
	fi 	     
}