// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !(linux || windows)

package ipnlocal

import (
	"scale.ropsoft.cloud/ipn"
)

func (b *LocalBackend) stopOfflineAutoUpdate() {
	// Not supported on this platform.
}

func (b *LocalBackend) maybeStartOfflineAutoUpdate(prefs ipn.PrefsView) {
	// Not supported on this platform.
}
