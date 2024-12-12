#!/bin/sh
set -e
if [ -d /run/systemd/system ] ; then
	  systemctl --system daemon-reload >/dev/null || true
fi

if [ -x "/usr/bin/deb-systemd-helper" ]; then
    if [ "$1" = "remove" ]; then
		    deb-systemd-helper mask 'rspscale.nginx-auth.socket' >/dev/null || true
		    deb-systemd-helper mask 'rspscale.nginx-auth.service' >/dev/null || true
	  fi

    if [ "$1" = "purge" ]; then
		    deb-systemd-helper purge 'rspscale.nginx-auth.socket' >/dev/null || true
		    deb-systemd-helper unmask 'rspscale.nginx-auth.socket' >/dev/null || true
		    deb-systemd-helper purge 'rspscale.nginx-auth.service' >/dev/null || true
		    deb-systemd-helper unmask 'rspscale.nginx-auth.service' >/dev/null || true
	  fi
fi
