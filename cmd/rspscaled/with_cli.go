// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build ts_include_cli

package main

import (
	"fmt"
	"os"

	"scale.ropsoft.cloud/cmd/rspscale/cli"
)

func init() {
	beCLI = func() {
		args := os.Args[1:]
		if err := cli.Run(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
