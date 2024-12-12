// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"testing"

	"scale.ropsoft.cloud/tstest/deptest"
)

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"testing":                            "do not use testing package in production code",
			"gvisor.dev/gvisor/pkg/buffer":       "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/cpuid":        "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip":        "https://github.com/ropsoft7/rspscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip/header": "https://github.com/ropsoft7/rspscale/issues/9756",
			"scale.ropsoft.cloud/wgengine/filter":      "brings in bart, etc",
			"github.com/bits-and-blooms/bitset":  "unneeded in CLI",
			"github.com/gaissmai/bart":           "unneeded in CLI",
			"scale.ropsoft.cloud/net/ipset":            "unneeded in CLI",
		},
	}.Check(t)
}
