// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build ios || android || js

package magicsock

import (
	"context"
	"net/netip"

	"scale.ropsoft.cloud/types/logger"
)

type cloudInfo struct{}

func newCloudInfo(_ logger.Logf) *cloudInfo {
	return &cloudInfo{}
}

func (ci *cloudInfo) GetPublicIPs(_ context.Context) ([]netip.Addr, error) {
	return nil, nil
}
