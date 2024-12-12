// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// The web-client command demonstrates serving the Rspscale web client over tsnet.
package main

import (
	"flag"
	"log"
	"net/http"

	"scale.ropsoft.cloud/client/web"
	"scale.ropsoft.cloud/tsnet"
)

var (
	addr = flag.String("addr", "localhost:8060", "address of Rspscale web client")
)

func main() {
	flag.Parse()

	s := &tsnet.Server{RunWebClient: true}
	defer s.Close()

	lc, err := s.LocalClient()
	if err != nil {
		log.Fatal(err)
	}

	// Serve the Rspscale web client.
	ws, err := web.NewServer(web.ServerOpts{
		Mode:        web.LoginServerMode,
		LocalClient: lc,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Shutdown()
	log.Printf("Serving Rspscale web client on http://%s", *addr)
	if err := http.ListenAndServe(*addr, ws); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
