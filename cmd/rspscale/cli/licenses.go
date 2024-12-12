// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/licenses"
)

var licensesCmd = &ffcli.Command{
	Name:       "licenses",
	ShortUsage: "rspscale licenses",
	ShortHelp:  "Get open source license information",
	LongHelp:   "Get open source license information",
	Exec:       runLicenses,
}

func runLicenses(ctx context.Context, args []string) error {
	url := licenses.LicensesURL()
	outln(`
Rspscale wouldn't be possible without the contributions of thousands of open
source developers. To see the open source packages included in Rspscale and
their respective license information, visit:

    ` + url)
	return nil
}
