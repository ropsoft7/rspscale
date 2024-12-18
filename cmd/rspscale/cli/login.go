// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var loginArgs upArgsT

var loginCmd = &ffcli.Command{
	Name:       "login",
	ShortUsage: "rspscale login [flags]",
	ShortHelp:  "Log in to a Rspscale account",
	LongHelp: `"rspscale login" logs this machine in to your Rspscale network.
This command is currently in alpha and may change in the future.`,
	FlagSet: func() *flag.FlagSet {
		return newUpFlagSet(effectiveGOOS(), &loginArgs, "login")
	}(),
	Exec: func(ctx context.Context, args []string) error {
		if err := localClient.SwitchToEmptyProfile(ctx); err != nil {
			return err
		}
		return runUp(ctx, "login", args, loginArgs)
	},
}
