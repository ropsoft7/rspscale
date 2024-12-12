// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package permissions

import "testing"

func TestPermissionsImpl(t *testing.T) {
	if err := permissionsImpl(t.Logf); err != nil {
		t.Error(err)
	}
}
