#!/usr/bin/env bash

# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R. Bejar     02/12/25   Ensure GDM screen locks cannot be overridden - updated to include lock-enabled setting
#
# supersedes nix_gdm_screen_lock_override_chk.sh, nix_gdm_screen_lock_override_chk_v2.sh, and nix_gdm_screen_lock_override_chk_v3.sh

{
   a_output=() a_output2=()

   # Function to check and report if a specific setting is locked
   f_check_setting() 
   {
      grep -Psrilq "^\h*$1\h*=\h*uint32\h+\d+\b|true" /etc/dconf/db/*/ && grep -Prilq "$2" /etc/dconf/db/*/locks && echo "- \"$3\" is locked" || echo "- \"$3\" is not locked or not set"
   }
   # Array of settings to check
   declare -A settings=(
      ["idle-delay"]="/org/gnome/desktop/session/idle-delay"
      ["lock-delay"]="/org/gnome/desktop/screensaver/lock-delay"
      ["lock-enabled"]="/org/gnome/desktop/screensaver/lock-enabled"
   )

   # Check GNOME Desktop Manager configurations
   for setting in "${!settings[@]}"; do
      result=$(f_check_setting "$setting" "${settings[$setting]}" "$setting")
      if [[ $result == *"is not locked"* || $result == *"not set to false"* ]]; then
         a_output2+=("$result")
      else
         a_output+=("$result")
      fi
   done

   # Report results
   printf '%s\n' "" "- Audit Result:"
   if [ "${#a_output2[@]}" -gt 0 ]; then
      printf '%s\n' "  ** FAIL **" " - Reason(s) for audit failure:" "${a_output2[@]}"
      [ "${#a_output[@]}" -gt 0 ] && printf '%s\n' "" "- Correctly set:" "${a_output[@]}"
      exit "${XCCDF_RESULT_FAIL:-102}"
   else
      printf '%s\n' "  ** PASS **" "${a_output[@]}"
      exit "${XCCDF_RESULT_PASS:-101}"
   fi
}
