// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package router

import (
	"github.com/tailscale/wireguard-go/tun"
	"scale.ropsoft.cloud/health"
	"scale.ropsoft.cloud/net/netmon"
	"scale.ropsoft.cloud/types/logger"
)

func newUserspaceRouter(logf logger.Logf, tundev tun.Device, netMon *netmon.Monitor, health *health.Tracker) (Router, error) {
	return newUserspaceBSDRouter(logf, tundev, netMon, health)
}

func cleanUp(logger.Logf, string) {
	// Nothing to do.
}
