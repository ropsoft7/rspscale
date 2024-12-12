// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package cli

import (
	"cmp"
	"context"
	"crypto/tls"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cgi"
	"net/netip"
	"os"
	"os/signal"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"scale.ropsoft.cloud/client/web"
	"scale.ropsoft.cloud/ipn"
)

var webCmd = &ffcli.Command{
	Name:       "web",
	ShortUsage: "rspscale web [flags]",
	ShortHelp:  "Run a web server for controlling Rspscale",

	LongHelp: strings.TrimSpace(`
"rspscale web" runs a webserver for controlling the Rspscale daemon.

It's primarily intended for use on Synology, QNAP, and other
NAS devices where a web interface is the natural place to control
Rspscale, as opposed to a CLI or a native app.
`),

	FlagSet: (func() *flag.FlagSet {
		webf := newFlagSet("web")
		webf.StringVar(&webArgs.listen, "listen", "localhost:8088", "listen address; use port 0 for automatic")
		webf.BoolVar(&webArgs.cgi, "cgi", false, "run as CGI script")
		webf.StringVar(&webArgs.prefix, "prefix", "", "URL prefix added to requests (for cgi or reverse proxies)")
		webf.BoolVar(&webArgs.readonly, "readonly", false, "run web UI in read-only mode")
		return webf
	})(),
	Exec: runWeb,
}

var webArgs struct {
	listen   string
	cgi      bool
	prefix   string
	readonly bool
}

func tlsConfigFromEnvironment() *tls.Config {
	crt := os.Getenv("TLS_CRT_PEM")
	key := os.Getenv("TLS_KEY_PEM")
	if crt == "" || key == "" {
		return nil
	}

	// We support passing in the complete certificate and key from environment
	// variables because pfSense stores its cert+key in the PHP config. We populate
	// TLS_CRT_PEM and TLS_KEY_PEM from PHP code before starting rspscale web.
	// These are the PEM-encoded Certificate and Private Key.

	cert, err := tls.X509KeyPair([]byte(crt), []byte(key))
	if err != nil {
		log.Printf("tlsConfigFromEnvironment: %v", err)

		// Fallback to unencrypted HTTP.
		return nil
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

func runWeb(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if len(args) > 0 {
		return fmt.Errorf("too many non-flag arguments: %q", args)
	}

	var selfIP netip.Addr
	st, err := localClient.StatusWithoutPeers(ctx)
	if err == nil && st.Self != nil && len(st.Self.RspscaleIPs) > 0 {
		selfIP = st.Self.RspscaleIPs[0]
	}

	var existingWebClient bool
	if prefs, err := localClient.GetPrefs(ctx); err == nil {
		existingWebClient = prefs.RunWebClient
	}
	var startedManagementClient bool // we started the management client
	if !existingWebClient && !webArgs.readonly {
		// Also start full client in rspscaled.
		log.Printf("starting rspscaled web client at http://%s\n", netip.AddrPortFrom(selfIP, web.ListenPort))
		if err := setRunWebClient(ctx, true); err != nil {
			return fmt.Errorf("starting web client in rspscaled: %w", err)
		}
		startedManagementClient = true
	}

	opts := web.ServerOpts{
		Mode:        web.LoginServerMode,
		CGIMode:     webArgs.cgi,
		PathPrefix:  webArgs.prefix,
		LocalClient: &localClient,
	}
	if webArgs.readonly {
		opts.Mode = web.ReadOnlyServerMode
	}
	webServer, err := web.NewServer(opts)
	if err != nil {
		log.Printf("rspscale.web: %v", err)
		return err
	}
	go func() {
		select {
		case <-ctx.Done():
			// Shutdown the server.
			webServer.Shutdown()
			if !webArgs.cgi && startedManagementClient {
				log.Println("stopping rspscaled web client")
				// When not in cgi mode, shut down the rspscaled
				// web client on cli termination if we started it.
				if err := setRunWebClient(context.Background(), false); err != nil {
					log.Printf("stopping rspscaled web client: %v", err)
				}
			}
		}
		os.Exit(0)
	}()

	if webArgs.cgi {
		if err := cgi.Serve(webServer); err != nil {
			log.Printf("rspscale.cgi: %v", err)
		}
		return nil
	} else if tlsConfig := tlsConfigFromEnvironment(); tlsConfig != nil {
		server := &http.Server{
			Addr:      webArgs.listen,
			TLSConfig: tlsConfig,
			Handler:   webServer,
		}
		defer server.Shutdown(ctx)
		log.Printf("web server running on: https://%s", server.Addr)
		return server.ListenAndServeTLS("", "")
	} else {
		log.Printf("web server running on: %s", urlOfListenAddr(webArgs.listen))
		return http.ListenAndServe(webArgs.listen, webServer)
	}
}

func setRunWebClient(ctx context.Context, val bool) error {
	_, err := localClient.EditPrefs(ctx, &ipn.MaskedPrefs{
		Prefs:           ipn.Prefs{RunWebClient: val},
		RunWebClientSet: true,
	})
	return err
}

// urlOfListenAddr parses a given listen address into a formatted URL
func urlOfListenAddr(addr string) string {
	host, port, _ := net.SplitHostPort(addr)
	return fmt.Sprintf("http://%s", net.JoinHostPort(cmp.Or(host, "127.0.0.1"), port))
}
