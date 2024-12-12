// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build (linux && !android) || windows

package clientupdate

import (
	"context"

	"scale.ropsoft.cloud/clientupdate/distsign"
)

func (up *Updater) downloadURLToFile(pathSrc, fileDst string) (ret error) {
	c, err := distsign.NewClient(up.Logf, up.PkgsAddr)
	if err != nil {
		return err
	}
	return c.Download(context.Background(), pathSrc, fileDst)
}