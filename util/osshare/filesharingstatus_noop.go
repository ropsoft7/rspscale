// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package osshare

import (
	"scale.ropsoft.cloud/types/logger"
)

func SetFileSharingEnabled(enabled bool, logf logger.Logf) {}
