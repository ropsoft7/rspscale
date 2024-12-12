// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build rspscale_go && android

package version

import "fmt"

func init() {
	// For official Android builds using the rspscale_go toolchain,
	// panic if the builder is screwed up and we fail to stamp a valid
	// version string.
	if !isValidLongWithTwoRepos(Long()) {
		panic(fmt.Sprintf("malformed version.Long value %q", Long()))
	}
}
