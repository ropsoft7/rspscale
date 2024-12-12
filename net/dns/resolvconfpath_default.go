// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !gokrazy

package dns

const (
	resolvConf = "/etc/resolv.conf"
	backupConf = "/etc/resolv.pre-rspscale-backup.conf"
)
