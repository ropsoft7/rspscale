if [ "$1" = "configure" ] || [ "$1" = "abort-upgrade" ] || [ "$1" = "abort-deconfigure" ] || [ "$1" = "abort-remove" ] ; then
	  deb-systemd-helper unmask 'rspscale.nginx-auth.socket' >/dev/null || true
	  if deb-systemd-helper --quiet was-enabled 'rspscale.nginx-auth.socket'; then
		    deb-systemd-helper enable 'rspscale.nginx-auth.socket' >/dev/null || true
	  else
		    deb-systemd-helper update-state 'rspscale.nginx-auth.socket' >/dev/null || true
	  fi

    if systemctl is-active rspscale.nginx-auth.socket >/dev/null; then
        systemctl --system daemon-reload >/dev/null || true
        deb-systemd-invoke stop 'rspscale.nginx-auth.service' >/dev/null || true
        deb-systemd-invoke restart 'rspscale.nginx-auth.socket' >/dev/null || true
    fi
fi
