// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin

package hostinfo

import (
	"os"
	"path/filepath"
)

func init() {
	packageType = packageTypeDarwin
}

func packageTypeDarwin() string {
	// Using rspscaled or IPNExtension?
	exe, _ := os.Executable()
	return filepath.Base(exe)
}
