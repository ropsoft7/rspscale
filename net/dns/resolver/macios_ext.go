// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build ts_macext && (darwin || ios)

package resolver

import (
	"errors"
	"net"

	"scale.ropsoft.cloud/net/netmon"
	"scale.ropsoft.cloud/net/netns"
)

func init() {
	initListenConfig = initListenConfigNetworkExtension
}

func initListenConfigNetworkExtension(nc *net.ListenConfig, netMon *netmon.Monitor, tunName string) error {
	nif, ok := netMon.InterfaceState().Interface[tunName]
	if !ok {
		return errors.New("utun not found")
	}
	return netns.SetListenConfigInterfaceIndex(nc, nif.Interface.Index)
}
