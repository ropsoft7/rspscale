// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"flag"
	"runtime"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/version/distro"
)

var configureCmd = &ffcli.Command{
	Name:       "configure",
	ShortUsage: "rspscale configure <subcommand>",
	ShortHelp:  "[ALPHA] Configure the host to enable more Rspscale features",
	LongHelp: strings.TrimSpace(`
The 'configure' set of commands are intended to provide a way to enable different
services on the host to use Rspscale in more ways.
`),
	FlagSet: (func() *flag.FlagSet {
		fs := newFlagSet("configure")
		return fs
	})(),
	Subcommands: configureSubcommands(),
}

func configureSubcommands() (out []*ffcli.Command) {
	if runtime.GOOS == "linux" && distro.Get() == distro.Synology {
		out = append(out, synologyConfigureCmd)
		out = append(out, synologyConfigureCertCmd)
	}
	return out
}
