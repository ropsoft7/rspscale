// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package posture

import (
	"fmt"

	"scale.ropsoft.cloud/types/logger"
	"scale.ropsoft.cloud/util/syspolicy"
)

// GetSerialNumbers returns the serial number of the iOS/tvOS device as reported by an
// MDM solution. It requires configuration via the DeviceSerialNumber system policy.
// This is the only way to gather serial numbers on iOS and tvOS.
func GetSerialNumbers(_ logger.Logf) ([]string, error) {
	s, err := syspolicy.GetString(syspolicy.DeviceSerialNumber, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get serial number from MDM: %v", err)
	}
	if s != "" {
		return []string{s}, nil
	}
	return nil, nil
}
