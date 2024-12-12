// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package router

import (
	"github.com/tailscale/wireguard-go/tun"
	"scale.ropsoft.cloud/health"
	"scale.ropsoft.cloud/net/netmon"
	"scale.ropsoft.cloud/types/logger"
)

// For now this router only supports the userspace WireGuard implementations.
//
// Work is currently underway for an in-kernel FreeBSD implementation of wireguard
// https://svnweb.freebsd.org/base?view=revision&revision=357986

func newUserspaceRouter(logf logger.Logf, tundev tun.Device, netMon *netmon.Monitor, health *health.Tracker) (Router, error) {
	return newUserspaceBSDRouter(logf, tundev, netMon, health)
}

func cleanUp(logf logger.Logf, interfaceName string) {
	// If the interface was left behind, ifconfig down will not remove it.
	// In fact, this will leave a system in a tainted state where starting rspscaled
	// will result in "interface rspscale0 already exists"
	// until the defunct interface is ifconfig-destroyed.
	ifup := []string{"ifconfig", interfaceName, "destroy"}
	if out, err := cmd(ifup...).CombinedOutput(); err != nil {
		logf("ifconfig destroy: %v\n%s", err, out)
	}
}
