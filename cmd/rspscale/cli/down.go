// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/ipn"
)

var downCmd = &ffcli.Command{
	Name:       "down",
	ShortUsage: "rspscale down",
	ShortHelp:  "Disconnect from Rspscale",

	Exec:    runDown,
	FlagSet: newDownFlagSet(),
}

var downArgs struct {
	acceptedRisks string
}

func newDownFlagSet() *flag.FlagSet {
	downf := newFlagSet("down")
	registerAcceptRiskFlag(downf, &downArgs.acceptedRisks)
	return downf
}

func runDown(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many non-flag arguments: %q", args)
	}

	if isSSHOverRspscale() {
		if err := presentRiskToUser(riskLoseSSH, `You are connected over Rspscale; this action will disable Rspscale and result in your session disconnecting.`, downArgs.acceptedRisks); err != nil {
			return err
		}
	}

	st, err := localClient.Status(ctx)
	if err != nil {
		return fmt.Errorf("error fetching current status: %w", err)
	}
	if st.BackendState == "Stopped" {
		fmt.Fprintf(Stderr, "Rspscale was already stopped.\n")
		return nil
	}
	_, err = localClient.EditPrefs(ctx, &ipn.MaskedPrefs{
		Prefs: ipn.Prefs{
			WantRunning: false,
		},
		WantRunningSet: true,
	})
	return err
}
