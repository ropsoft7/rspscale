// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux

package magicsock

import (
	"errors"
	"fmt"
	"io"

	"scale.ropsoft.cloud/types/logger"
	"scale.ropsoft.cloud/types/nettype"
)

func (c *Conn) listenRawDisco(family string) (io.Closer, error) {
	return nil, fmt.Errorf("raw disco listening not supported on this OS: %w", errors.ErrUnsupported)
}

func trySetSocketBuffer(pconn nettype.PacketConn, logf logger.Logf) {
	portableTrySetSocketBuffer(pconn, logf)
}

const (
	controlMessageSize = 0
)
