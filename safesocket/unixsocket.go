// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows && !js && !plan9

package safesocket

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func connect(ctx context.Context, path string) (net.Conn, error) {
	var std net.Dialer
	return std.DialContext(ctx, "unix", path)
}

func listen(path string) (net.Listener, error) {
	// Unix sockets hang around in the filesystem even after nobody
	// is listening on them. (Which is really unfortunate but long-
	// entrenched semantics.) Try connecting first; if it works, then
	// the socket is still live, so let's not replace it. If it doesn't
	// work, then replace it.
	//
	// Note that there's a race condition between these two steps. A
	// "proper" daemon usually uses a dance involving pidfiles to first
	// ensure that no other instances of itself are running, but that's
	// beyond the scope of our simple socket library.
	c, err := net.Dial("unix", path)
	if err == nil {
		c.Close()
		if rspscaledRunningUnderLaunchd() {
			return nil, fmt.Errorf("%v: address already in use; rspscaled already running under launchd (to stop, run: $ sudo launchctl stop com.rspscale.rspscaled)", path)
		}
		return nil, fmt.Errorf("%v: address already in use", path)
	}
	_ = os.Remove(path)

	perm := socketPermissionsForOS()

	sockDir := filepath.Dir(path)
	if _, err := os.Stat(sockDir); os.IsNotExist(err) {
		os.MkdirAll(sockDir, 0755) // best effort

		// If we're on a platform where we want the socket
		// world-readable, open up the permissions on the
		// just-created directory too, in case a umask ate
		// it. This primarily affects running rspscaled by
		// hand as root in a shell, as there is no umask when
		// running under systemd.
		if perm == 0666 {
			if fi, err := os.Stat(sockDir); err == nil && fi.Mode()&0077 == 0 {
				if err := os.Chmod(sockDir, 0755); err != nil {
					log.Print(err)
				}
			}
		}
	}
	pipe, err := net.Listen("unix", path)
	if err != nil {
		return nil, err
	}
	os.Chmod(path, perm)
	return pipe, err
}

func rspscaledRunningUnderLaunchd() bool {
	if runtime.GOOS != "darwin" {
		return false
	}
	plist, err := exec.Command("launchctl", "list", "com.rspscale.rspscaled").Output()
	_ = plist // parse it? https://github.com/DHowett/go-plist if we need something.
	running := err == nil
	return running
}

// socketPermissionsForOS returns the permissions to use for the
// rspscaled.sock.
func socketPermissionsForOS() os.FileMode {
	if PlatformUsesPeerCreds() {
		return 0666
	}
	// Otherwise, root only.
	return 0600
}
