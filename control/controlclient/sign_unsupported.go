// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package controlclient

import (
	"scale.ropsoft.cloud/tailcfg"
	"scale.ropsoft.cloud/types/key"
)

// signRegisterRequest on non-supported platforms always returns errNoCertStore.
func signRegisterRequest(req *tailcfg.RegisterRequest, serverURL string, serverPubKey, machinePubKey key.MachinePublic) error {
	return errNoCertStore
}
