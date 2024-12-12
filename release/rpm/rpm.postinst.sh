# $1 == 1 for initial installation.
# $1 == 2 for upgrades.

if [ $1 -eq 1 ] ; then
    # Normally, the rspscale-relay package would request shutdown of
    # its service before uninstallation. Unfortunately, the
    # rspscale-relay package we distributed doesn't have those
    # scriptlets. We definitely want relaynode to be stopped when
    # installing rspscaled though, so we blindly try to turn off
    # relaynode here.
    #
    # However, we also want this package installation to look like an
    # upgrade from relaynode! Therefore, if relaynode is currently
    # enabled, we want to also enable rspscaled. If relaynode is
    # currently running, we also want to start rspscaled.
    #
    # If there doesn't seem to be an active or enabled relaynode on
    # the system, we follow the RPM convention for package installs,
    # which is to not enable or start the service.
    relaynode_enabled=0
    relaynode_running=0
    if systemctl is-enabled rspscale-relay.service >/dev/null 2>&1; then
        relaynode_enabled=1
    fi
    if systemctl is-active rspscale-relay.service >/dev/null 2>&1; then
        relaynode_running=1
    fi

    systemctl --no-reload disable rspscale-relay.service >/dev/null 2>&1 || :
    systemctl stop rspscale-relay.service >/dev/null 2>&1 || :

    if [ $relaynode_enabled -eq 1 ]; then
        systemctl enable rspscaled.service >/dev/null 2>&1 || :
    else
        systemctl preset rspscaled.service >/dev/null 2>&1 || : 
    fi

    if [ $relaynode_running -eq 1 ]; then
        systemctl start rspscaled.service >/dev/null 2>&1 || :
    fi
fi 
