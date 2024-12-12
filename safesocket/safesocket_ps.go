// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux || windows || (darwin && !ios) || freebsd

package safesocket

import (
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func init() {
	rspscaledProcExists = func() bool {
		procs, err := ps.Processes()
		if err != nil {
			return false
		}
		for _, proc := range procs {
			name := proc.Executable()
			const rspscaled = "rspscaled"
			if len(name) < len(rspscaled) {
				continue
			}
			// Do case insensitive comparison for Windows,
			// notably, and ignore any ".exe" suffix.
			if strings.EqualFold(name[:len(rspscaled)], rspscaled) {
				return true
			}
		}
		return false
	}
}
