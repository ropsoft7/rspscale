// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// The scale.ropsoft.cloudmand is the Rspscale command-line client. It interacts
// with the rspscaled node agent.
package main // import "scale.ropsoft.cloud/cmd/rspscale"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"scale.ropsoft.cloud/cmd/rspscale/cli"
)

func main() {
	args := os.Args[1:]
	if name, _ := os.Executable(); strings.HasSuffix(filepath.Base(name), ".cgi") {
		args = []string{"web", "-cgi"}
	}
	if err := cli.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
