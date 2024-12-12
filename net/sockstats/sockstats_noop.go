// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !rspscale_go || !(darwin || ios || android || ts_enable_sockstats)

package sockstats

import (
	"context"

	"scale.ropsoft.cloud/net/netmon"
	"scale.ropsoft.cloud/types/logger"
)

const IsAvailable = false

func withSockStats(ctx context.Context, label Label, logf logger.Logf) context.Context {
	return ctx
}

func get() *SockStats {
	return nil
}

func getInterfaces() *InterfaceSockStats {
	return nil
}

func getValidation() *ValidationSockStats {
	return nil
}

func setNetMon(netMon *netmon.Monitor) {
}

func debugInfo() string {
	return ""
}
