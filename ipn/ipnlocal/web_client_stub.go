// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build ios || android

package ipnlocal

import (
	"errors"
	"net"

	"scale.ropsoft.cloud/client/rspscale"
)

const webClientPort = 5252

type webClient struct{}

func (b *LocalBackend) ConfigureWebClient(lc *rspscale.LocalClient) {}

func (b *LocalBackend) webClientGetOrInit() error {
	return errors.New("not implemented")
}

func (b *LocalBackend) webClientShutdown() {}

func (b *LocalBackend) handleWebClientConn(c net.Conn) error {
	return errors.New("not implemented")
}
func (b *LocalBackend) updateWebClientListenersLocked() {}
