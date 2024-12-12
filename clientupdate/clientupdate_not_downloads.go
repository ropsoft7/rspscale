// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !((linux && !android) || windows)

package clientupdate

func (up *Updater) downloadURLToFile(pathSrc, fileDst string) (ret error) {
	panic("unreachable")
}
