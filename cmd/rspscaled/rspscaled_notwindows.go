// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows && go1.19

package main // import "scale.ropsoft.cloud/cmd/rspscaled"

import "scale.ropsoft.cloud/logpolicy"

func isWindowsService() bool { return false }

func runWindowsService(pol *logpolicy.Policy) error { panic("unreachable") }

func beWindowsSubprocess() bool { return false }
