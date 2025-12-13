#!/usr/bin/env bash
#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R. Bejar     02/12/25   SCE Ensure GDM login banner is configured
#
# supersedes nix_gdm_login_banner_configured_chk.sh and nix_gdm_login_banner_configured_chk_v2.sh

{
    if grep -Psiq "user-db|system-db" /etc/dconf/profile/*/*; then	
		# Get the current values of the GSettings keys
        l_banner_enabled=$(gsettings get org.gnome.login-screen banner-message-enable 2>/dev/null)
        l_banner_text_message=$(gsettings get org.gnome.login-screen banner-message-text 2>/dev/null)

        # Check if the banner-message-enable is true
        if [ "$l_banner_enabled" == "true" ]; then
            # Check if the banner text is not empty
            if [ -n "$l_banner_text_message" ] && [ "$l_banner_text_message" != "''" ]; then
                echo -e "Audit Result:\n ** PASS **\n org.gnome.login-screen banner settings are correctly set"
                exit "${XCCDF_RESULT_PASS:-101}"
            else
                echo -e "Audit Result:\n  ** FAIL **\n Banner message text is not set or empty."
                exit "${XCCDF_RESULT_FAIL:-102}"
            fi
        else
            echo -e "Audit Result:\n  ** FAIL **\n Banner message enable is not set to true."   
            exit "${XCCDF_RESULT_FAIL:-102}"
        fi
	else
		echo -e "Audit Result:\n ** FAIL **\n User profile does not exist."
        exit "${XCCDF_RESULT_FAIL:-102}"
	fi 	
}