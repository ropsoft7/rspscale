// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package magicsock

import (
	"scale.ropsoft.cloud/types/logger"
	"scale.ropsoft.cloud/types/nettype"
)

func trySetUDPSocketOptions(pconn nettype.PacketConn, logf logger.Logf) {}
