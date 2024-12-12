// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build plan9 || aix

package tstun

import (
	"github.com/tailscale/wireguard-go/tun"
	"scale.ropsoft.cloud/types/logger"
)

func New(logf logger.Logf, tunName string) (tun.Device, string, error) {
	panic("not implemented")
}

func Diagnose(logf logger.Logf, tunName string, err error) {
	panic("not implemented")
}
