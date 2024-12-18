// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package testenv

import (
	"testing"

	"scale.ropsoft.cloud/tstest/deptest"
)

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"testing": "see pkg docs",
		},
	}.Check(t)
}
