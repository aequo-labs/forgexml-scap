#!/usr/bin/env bash
#
# CIS-CAT Script Check Engine
#
# Name         Date       Description
# -------------------------------------------------------------------
# R. Bejar     04/01/25   STIG REC:UBTU-22-271030 - disable the x86 Ctrl-Alt-Delete key sequence if a graphical user interface is installed.

{
    # Get the current value of the GSettings key
    l_daemon_plugins_media_keys=$(gsettings get org.gnome.settings-daemon.plugins.media-keys logout)

    # Check if the logout setting is []
    if [[ "$l_daemon_plugins_media_keys" == "@"*'[]' ]]; then
        echo -e "Audit Result:\n ** PASS **\n org.gnome.settings-daemon.plugins.media-keys is set to logout."
		exit "${XCCDF_RESULT_PASS:-101}"

    else
        echo -e "Audit Result:\n ** FAIL **\n org.gnome.settings-daemon.plugins.media-keys is not set to logout."
		exit "${XCCDF_RESULT_FAIL:-102}"

    fi
}