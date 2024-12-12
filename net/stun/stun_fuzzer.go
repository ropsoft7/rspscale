// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause
//go:build gofuzz

package stun

func FuzzStunParser(data []byte) int {
	_, _, _ = ParseResponse(data)

	_, _ = ParseBindingRequest(data)
	return 1
}
