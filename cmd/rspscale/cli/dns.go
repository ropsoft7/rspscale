// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var dnsCmd = &ffcli.Command{
	Name:       "dns",
	ShortHelp:  "Diagnose the internal DNS forwarder",
	LongHelp:   dnsCmdLongHelp(),
	ShortUsage: "rspscale dns <subcommand> [flags]",
	UsageFunc:  usageFuncNoDefaultValues,
	Subcommands: []*ffcli.Command{
		{
			Name:       "status",
			ShortUsage: "rspscale dns status [--all]",
			Exec:       runDNSStatus,
			ShortHelp:  "Prints the current DNS status and configuration",
			LongHelp:   dnsStatusLongHelp(),
			FlagSet: (func() *flag.FlagSet {
				fs := newFlagSet("status")
				fs.BoolVar(&dnsStatusArgs.all, "all", false, "outputs advanced debugging information (fallback resolvers, nameservers, cert domains, extra records, and exit node filtered set)")
				return fs
			})(),
		},
		{
			Name:       "query",
			ShortUsage: "rspscale dns query <name> [a|aaaa|cname|mx|ns|opt|ptr|srv|txt]",
			Exec:       runDNSQuery,
			ShortHelp:  "Perform a DNS query",
			LongHelp:   "The 'rspscale dns query' subcommand performs a DNS query for the specified name using the internal DNS forwarder (100.100.100.100).\n\nIt also provides information about the resolver(s) used to resolve the query.",
		},

		// TODO: implement `rspscale log` here

		// The above work is tracked in https://github.com/ropsoft7/rspscale/issues/13326
	},
}

func dnsCmdLongHelp() string {
	return `The 'rspscale dns' subcommand provides tools for diagnosing the internal DNS forwarder (100.100.100.100).
	
For more information about the DNS functionality built into Rspscale, refer to https://scale.ropsoft.cloud/kb/1054/dns.`
}
