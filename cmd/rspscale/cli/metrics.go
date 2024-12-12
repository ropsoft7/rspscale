// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/atomicfile"
)

var metricsCmd = &ffcli.Command{
	Name:      "metrics",
	ShortHelp: "Show Rspscale metrics",
	LongHelp: strings.TrimSpace(`

The 'rspscale metrics' command shows Rspscale user-facing metrics (as opposed
to internal metrics printed by 'rspscale debug metrics').

For more information about Rspscale metrics, refer to
https://scale.ropsoft.cloud/s/client-metrics

`),
	ShortUsage: "rspscale metrics <subcommand> [flags]",
	UsageFunc:  usageFuncNoDefaultValues,
	Exec:       runMetricsNoSubcommand,
	Subcommands: []*ffcli.Command{
		{
			Name:       "print",
			ShortUsage: "rspscale metrics print",
			Exec:       runMetricsPrint,
			ShortHelp:  "Prints current metric values in the Prometheus text exposition format",
		},
		{
			Name:       "write",
			ShortUsage: "rspscale metrics write <path>",
			Exec:       runMetricsWrite,
			ShortHelp:  "Writes metric values to a file",
			LongHelp: strings.TrimSpace(`

The 'rspscale metrics write' command writes metric values to a text file provided as its
only argument. It's meant to be used alongside Prometheus node exporter, allowing Rspscale
metrics to be consumed and exported by the textfile collector.

As an example, to export Rspscale metrics on an Ubuntu system running node exporter, you
can regularly run 'rspscale metrics write /var/lib/prometheus/node-exporter/rspscaled.prom'
using cron or a systemd timer.

	`),
		},
	},
}

// runMetricsNoSubcommand prints metric values if no subcommand is specified.
func runMetricsNoSubcommand(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("rspscale metrics: unknown subcommand: %s", args[0])
	}

	return runMetricsPrint(ctx, args)
}

// runMetricsPrint prints metric values to stdout.
func runMetricsPrint(ctx context.Context, args []string) error {
	out, err := localClient.UserMetrics(ctx)
	if err != nil {
		return err
	}
	Stdout.Write(out)
	return nil
}

// runMetricsWrite writes metric values to a file.
func runMetricsWrite(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: rspscale metrics write <path>")
	}
	path := args[0]
	out, err := localClient.UserMetrics(ctx)
	if err != nil {
		return err
	}
	return atomicfile.WriteFile(path, out, 0644)
}
