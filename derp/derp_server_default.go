// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux

package derp

import "context"

func (c *sclient) startStatsLoop(ctx context.Context) {
	// Nothing to do
	return
}
