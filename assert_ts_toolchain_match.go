// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build rspscale_go

package rspscaleroot

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	tsRev, ok := rspscaleToolchainRev()
	if !ok {
		panic("binary built with rspscale_go build tag but failed to read build info or find rspscale.toolchain.rev in build info")
	}
	want := strings.TrimSpace(GoToolchainRev)
	if tsRev != want {
		if os.Getenv("TS_PERMIT_TOOLCHAIN_MISMATCH") == "1" {
			fmt.Fprintf(os.Stderr, "rspscale.toolchain.rev = %q, want %q; but ignoring due to TS_PERMIT_TOOLCHAIN_MISMATCH=1\n", tsRev, want)
			return
		}
		panic(fmt.Sprintf("binary built with rspscale_go build tag but Go toolchain %q doesn't match github.com/ropsoft7/rspscale expected value %q; override this failure with TS_PERMIT_TOOLCHAIN_MISMATCH=1", tsRev, want))
	}
}
