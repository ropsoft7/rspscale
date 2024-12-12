// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"testing"

	"scale.ropsoft.cloud/tstest/deptest"
)

func TestOmitSSH(t *testing.T) {
	const msg = "unexpected with ts_omit_ssh"
	deptest.DepChecker{
		GOOS:   "linux",
		GOARCH: "amd64",
		Tags:   "ts_omit_ssh",
		BadDeps: map[string]string{
			"scale.ropsoft.cloud/ssh/tailssh":            msg,
			"golang.org/x/crypto/ssh":              msg,
			"scale.ropsoft.cloud/sessionrecording":       msg,
			"github.com/anmitsu/go-shlex":          msg,
			"github.com/creack/pty":                msg,
			"github.com/kr/fs":                     msg,
			"github.com/pkg/sftp":                  msg,
			"github.com/u-root/u-root/pkg/termios": msg,
			"tempfork/gliderlabs/ssh":              msg,
		},
	}.Check(t)
}