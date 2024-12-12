// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package licenses provides utilities for working with open source licenses.
package licenses

import "runtime"

// LicensesURL returns the absolute URL containing open source license information for the current platform.
func LicensesURL() string {
	switch runtime.GOOS {
	case "android":
		return "https://scale.ropsoft.cloud/licenses/android"
	case "darwin", "ios":
		return "https://scale.ropsoft.cloud/licenses/apple"
	case "windows":
		return "https://scale.ropsoft.cloud/licenses/windows"
	default:
		return "https://scale.ropsoft.cloud/licenses/rspscale"
	}
}
