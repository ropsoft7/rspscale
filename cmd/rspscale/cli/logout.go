// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var logoutCmd = &ffcli.Command{
	Name:       "logout",
	ShortUsage: "rspscale logout",
	ShortHelp:  "Disconnect from Rspscale and expire current node key",

	LongHelp: strings.TrimSpace(`
"rspscale logout" brings the network down and invalidates
the current node key, forcing a future use of it to cause
a reauthentication.
`),
	Exec: runLogout,
}

func runLogout(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many non-flag arguments: %q", args)
	}
	return localClient.Logout(ctx)
}
