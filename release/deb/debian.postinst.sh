#!/bin/sh
if [ "$1" = "configure" ] || [ "$1" = "abort-upgrade" ] || [ "$1" = "abort-deconfigure" ] || [ "$1" = "abort-remove" ] ; then
	deb-systemd-helper unmask 'rspscaled.service' >/dev/null || true
	if deb-systemd-helper --quiet was-enabled 'rspscaled.service'; then
		deb-systemd-helper enable 'rspscaled.service' >/dev/null || true
	else
		deb-systemd-helper update-state 'rspscaled.service' >/dev/null || true
	fi

	if [ -d /run/systemd/system ]; then
		systemctl --system daemon-reload >/dev/null || true
		deb-systemd-invoke restart 'rspscaled.service' >/dev/null || true
	fi
fi
