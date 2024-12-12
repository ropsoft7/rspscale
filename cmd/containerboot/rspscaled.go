// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux

package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"scale.ropsoft.cloud/client/rspscale"
)

func startRspscaled(ctx context.Context, cfg *settings) (*rspscale.LocalClient, *os.Process, error) {
	args := rspscaledArgs(cfg)
	// rspscaled runs without context, since it needs to persist
	// beyond the startup timeout in ctx.
	cmd := exec.Command("rspscaled", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	log.Printf("Starting rspscaled")
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("starting rspscaled failed: %v", err)
	}

	// Wait for the socket file to appear, otherwise API ops will racily fail.
	log.Printf("Waiting for rspscaled socket")
	for {
		if ctx.Err() != nil {
			log.Fatalf("Timed out waiting for rspscaled socket")
		}
		_, err := os.Stat(cfg.Socket)
		if errors.Is(err, fs.ErrNotExist) {
			time.Sleep(100 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalf("Waiting for rspscaled socket: %v", err)
		}
		break
	}

	tsClient := &rspscale.LocalClient{
		Socket:        cfg.Socket,
		UseSocketOnly: true,
	}

	return tsClient, cmd.Process, nil
}

// rspscaledArgs uses cfg to construct the argv for rspscaled.
func rspscaledArgs(cfg *settings) []string {
	args := []string{"--socket=" + cfg.Socket}
	switch {
	case cfg.InKubernetes && cfg.KubeSecret != "":
		args = append(args, "--state=kube:"+cfg.KubeSecret)
		if cfg.StateDir == "" {
			cfg.StateDir = "/tmp"
		}
		fallthrough
	case cfg.StateDir != "":
		args = append(args, "--statedir="+cfg.StateDir)
	default:
		args = append(args, "--state=mem:", "--statedir=/tmp")
	}

	if cfg.UserspaceMode {
		args = append(args, "--tun=userspace-networking")
	} else if err := ensureTunFile(cfg.Root); err != nil {
		log.Fatalf("ensuring that /dev/net/tun exists: %v", err)
	}

	if cfg.SOCKSProxyAddr != "" {
		args = append(args, "--socks5-server="+cfg.SOCKSProxyAddr)
	}
	if cfg.HTTPProxyAddr != "" {
		args = append(args, "--outbound-http-proxy-listen="+cfg.HTTPProxyAddr)
	}
	if cfg.RspscaledConfigFilePath != "" {
		args = append(args, "--config="+cfg.RspscaledConfigFilePath)
	}
	// Once enough proxy versions have been released for all the supported
	// versions to understand this cfg setting, the operator can stop
	// setting TS_RSPSCALED_EXTRA_ARGS for the debug flag.
	if cfg.DebugAddrPort != "" && !strings.Contains(cfg.DaemonExtraArgs, cfg.DebugAddrPort) {
		args = append(args, "--debug="+cfg.DebugAddrPort)
	}
	if cfg.DaemonExtraArgs != "" {
		args = append(args, strings.Fields(cfg.DaemonExtraArgs)...)
	}
	return args
}

// rspscaleUp uses cfg to run 'rspscale up' everytime containerboot starts, or
// if TS_AUTH_ONCE is set, only the first time containerboot starts.
func rspscaleUp(ctx context.Context, cfg *settings) error {
	args := []string{"--socket=" + cfg.Socket, "up"}
	if cfg.AcceptDNS != nil && *cfg.AcceptDNS {
		args = append(args, "--accept-dns=true")
	} else {
		args = append(args, "--accept-dns=false")
	}
	if cfg.AuthKey != "" {
		args = append(args, "--authkey="+cfg.AuthKey)
	}
	// --advertise-routes can be passed an empty string to configure a
	// device (that might have previously advertised subnet routes) to not
	// advertise any routes. Respect an empty string passed by a user and
	// use it to explicitly unset the routes.
	if cfg.Routes != nil {
		args = append(args, "--advertise-routes="+*cfg.Routes)
	}
	if cfg.Hostname != "" {
		args = append(args, "--hostname="+cfg.Hostname)
	}
	if cfg.ExtraArgs != "" {
		args = append(args, strings.Fields(cfg.ExtraArgs)...)
	}
	log.Printf("Running 'rspscale up'")
	cmd := exec.CommandContext(ctx, "rspscale", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rspscale up failed: %v", err)
	}
	return nil
}

// rspscaleSet uses cfg to run 'rspscale set' to set any known configuration
// options that are passed in via environment variables. This is run after the
// node is in Running state and only if TS_AUTH_ONCE is set.
func rspscaleSet(ctx context.Context, cfg *settings) error {
	args := []string{"--socket=" + cfg.Socket, "set"}
	if cfg.AcceptDNS != nil && *cfg.AcceptDNS {
		args = append(args, "--accept-dns=true")
	} else {
		args = append(args, "--accept-dns=false")
	}
	// --advertise-routes can be passed an empty string to configure a
	// device (that might have previously advertised subnet routes) to not
	// advertise any routes. Respect an empty string passed by a user and
	// use it to explicitly unset the routes.
	if cfg.Routes != nil {
		args = append(args, "--advertise-routes="+*cfg.Routes)
	}
	if cfg.Hostname != "" {
		args = append(args, "--hostname="+cfg.Hostname)
	}
	log.Printf("Running 'rspscale set'")
	cmd := exec.CommandContext(ctx, "rspscale", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rspscale set failed: %v", err)
	}
	return nil
}
