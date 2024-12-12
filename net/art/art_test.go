// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package art

import (
	"os"
	"testing"

	"scale.ropsoft.cloud/util/cibuild"
)

func TestMain(m *testing.M) {
	if cibuild.On() {
		// Skip CI on GitHub for now
		// TODO: https://github.com/ropsoft7/rspscale/issues/7866
		os.Exit(0)
	}
	os.Exit(m.Run())
}
