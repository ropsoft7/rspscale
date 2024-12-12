// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// get-authkey allocates an authkey using an OAuth API client
// https://scale.ropsoft.cloud/s/oauth-clients and prints it
// to stdout for scripts to capture and use.
package main

import (
	"cmp"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
	"scale.ropsoft.cloud/client/rspscale"
)

func main() {
	// Required to use our client API. We're fine with the instability since the
	// client lives in the same repo as this code.
	rspscale.I_Acknowledge_This_API_Is_Unstable = true

	reusable := flag.Bool("reusable", false, "allocate a reusable authkey")
	ephemeral := flag.Bool("ephemeral", false, "allocate an ephemeral authkey")
	preauth := flag.Bool("preauth", true, "set the authkey as pre-authorized")
	tags := flag.String("tags", "", "comma-separated list of tags to apply to the authkey")
	flag.Parse()

	clientID := os.Getenv("TS_API_CLIENT_ID")
	clientSecret := os.Getenv("TS_API_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("TS_API_CLIENT_ID and TS_API_CLIENT_SECRET must be set")
	}

	if *tags == "" {
		log.Fatal("at least one tag must be specified")
	}

	baseURL := cmp.Or(os.Getenv("TS_BASE_URL"), "https://api.scale.ropsoft.cloud")

	credentials := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     baseURL + "/api/v2/oauth/token",
	}

	ctx := context.Background()
	tsClient := rspscale.NewClient("-", nil)
	tsClient.UserAgent = "rspscale-get-authkey"
	tsClient.HTTPClient = credentials.Client(ctx)
	tsClient.BaseURL = baseURL

	caps := rspscale.KeyCapabilities{
		Devices: rspscale.KeyDeviceCapabilities{
			Create: rspscale.KeyDeviceCreateCapabilities{
				Reusable:      *reusable,
				Ephemeral:     *ephemeral,
				Preauthorized: *preauth,
				Tags:          strings.Split(*tags, ","),
			},
		},
	}

	authkey, _, err := tsClient.CreateKey(ctx, caps)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(authkey)
}
