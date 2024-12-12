// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package paths returns platform and user-specific default paths to
// Rspscale files and directories.
package paths

import (
	"os"
	"path/filepath"
	"runtime"

	"scale.ropsoft.cloud/syncs"
	"scale.ropsoft.cloud/version/distro"
)

// AppSharedDir is a string set by the iOS or Android app on start
// containing a directory we can read/write in.
var AppSharedDir syncs.AtomicValue[string]

// DefaultRspscaledSocket returns the path to the rspscaled Unix socket
// or the empty string if there's no reasonable default.
func DefaultRspscaledSocket() string {
	if runtime.GOOS == "windows" {
		return `\\.\pipe\ProtectedPrefix\Administrators\Rspscale\rspscaled`
	}
	if runtime.GOOS == "darwin" {
		return "/var/run/rspscaled.socket"
	}
	if runtime.GOOS == "plan9" {
		return "/srv/rspscaled.sock"
	}
	switch distro.Get() {
	case distro.Synology:
		if distro.DSMVersion() == 6 {
			return "/var/packages/Rspscale/etc/rspscaled.sock"
		}
		// DSM 7 (and higher? or failure to detect.)
		return "/var/packages/Rspscale/var/rspscaled.sock"
	case distro.Gokrazy:
		return "/perm/rspscaled/rspscaled.sock"
	case distro.QNAP:
		return "/tmp/ropsoft7/rspscaled.sock"
	}
	if fi, err := os.Stat("/var/run"); err == nil && fi.IsDir() {
		return "/var/run/ropsoft7/rspscaled.sock"
	}
	return "rspscaled.sock"
}

// Overridden in init by OS-specific files.
var (
	stateFileFunc func() string

	// ensureStateDirPerms applies a restrictive ACL/chmod
	// to the provided directory.
	ensureStateDirPerms = func(string) error { return nil }
)

// DefaultRspscaledStateFile returns the default path to the
// rspscaled state file, or the empty string if there's no reasonable
// default value.
func DefaultRspscaledStateFile() string {
	if f := stateFileFunc; f != nil {
		return f()
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("ProgramData"), "Rspscale", "server-state.conf")
	}
	return ""
}

// MkStateDir ensures that dirPath, the daemon's configuration directory
// containing machine keys etc, both exists and has the correct permissions.
// We want it to only be accessible to the user the daemon is running under.
func MkStateDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, 0700); err != nil {
		return err
	}
	return ensureStateDirPerms(dirPath)
}

// LegacyStateFilePath returns the legacy path to the state file when
// it was stored under the current user's %LocalAppData%.
//
// It is only called on Windows.
func LegacyStateFilePath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("LocalAppData"), "Rspscale", "server-state.conf")
	}
	return ""
}
