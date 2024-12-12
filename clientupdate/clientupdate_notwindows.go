// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package clientupdate

func (up *Updater) updateWindows() error {
	panic("unreachable")
}
