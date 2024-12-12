// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux && !windows && !darwin

package cli

import "fmt"

// The github.com/mitchellh/go-ps package doesn't work on all platforms,
// so just don't diagnose connect failures.

func fixRspscaledConnectError(origErr error) error {
	return fmt.Errorf("failed to connect to local rspscaled process (is it running?); got: %w", origErr)
}
