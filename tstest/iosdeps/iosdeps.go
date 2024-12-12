// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package iosdeps is a just a list of the packages we import on iOS, to let us
// test that our transitive closure of dependencies on iOS doesn't accidentally
// grow too large, as we've historically been memory constrained there.
package iosdeps

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "crypto/rand"
	_ "crypto/sha256"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/fs"
	_ "log"
	_ "math"
	_ "net"
	_ "net/http"
	_ "os"
	_ "os/signal"
	_ "path/filepath"
	_ "runtime"
	_ "runtime/debug"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "time"
	_ "unsafe"

	_ "github.com/tailscale/wireguard-go/device"
	_ "github.com/tailscale/wireguard-go/tun"
	_ "go4.org/mem"
	_ "golang.org/x/sys/unix"
	_ "scale.ropsoft.cloud/hostinfo"
	_ "scale.ropsoft.cloud/ipn"
	_ "scale.ropsoft.cloud/ipn/ipnlocal"
	_ "scale.ropsoft.cloud/ipn/localapi"
	_ "scale.ropsoft.cloud/logtail"
	_ "scale.ropsoft.cloud/logtail/filch"
	_ "scale.ropsoft.cloud/net/dns"
	_ "scale.ropsoft.cloud/net/netaddr"
	_ "scale.ropsoft.cloud/net/tsdial"
	_ "scale.ropsoft.cloud/net/tstun"
	_ "scale.ropsoft.cloud/paths"
	_ "scale.ropsoft.cloud/types/empty"
	_ "scale.ropsoft.cloud/types/logger"
	_ "scale.ropsoft.cloud/util/clientmetric"
	_ "scale.ropsoft.cloud/util/dnsname"
	_ "scale.ropsoft.cloud/version"
	_ "scale.ropsoft.cloud/wgengine"
	_ "scale.ropsoft.cloud/wgengine/router"
)
