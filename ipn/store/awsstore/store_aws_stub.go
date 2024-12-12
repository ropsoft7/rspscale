// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux || ts_omit_aws

package awsstore

import (
	"fmt"
	"runtime"

	"scale.ropsoft.cloud/ipn"
	"scale.ropsoft.cloud/types/logger"
)

func New(logger.Logf, string) (ipn.StateStore, error) {
	return nil, fmt.Errorf("AWS store is not supported on %v", runtime.GOOS)
}
