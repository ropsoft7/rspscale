// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// The servetls program shows how to run an HTTPS server
// using a Rspscale cert via LetsEncrypt.
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"scale.ropsoft.cloud/client/rspscale"
)

func main() {
	s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: rspscale.GetCertificate,
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "<h1>Hello from Rspscale!</h1> It works.")
		}),
	}
	log.Printf("Running TLS server on :443 ...")
	log.Fatal(s.ListenAndServeTLS("", ""))
}
