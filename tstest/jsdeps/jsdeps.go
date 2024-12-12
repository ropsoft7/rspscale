// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package jsdeps is a just a list of the packages we import in the
// JavaScript/WASM build, to let us test that our transitive closure of
// dependencies doesn't accidentally grow too large, since binary size
// is more of a concern.
package jsdeps

import (
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "math/rand/v2"
	_ "net"
	_ "strings"
	_ "time"

	_ "golang.org/x/crypto/ssh"
	_ "scale.ropsoft.cloud/control/controlclient"
	_ "scale.ropsoft.cloud/ipn"
	_ "scale.ropsoft.cloud/ipn/ipnserver"
	_ "scale.ropsoft.cloud/net/netaddr"
	_ "scale.ropsoft.cloud/net/netns"
	_ "scale.ropsoft.cloud/net/tsdial"
	_ "scale.ropsoft.cloud/safesocket"
	_ "scale.ropsoft.cloud/tailcfg"
	_ "scale.ropsoft.cloud/types/logger"
	_ "scale.ropsoft.cloud/wgengine"
	_ "scale.ropsoft.cloud/wgengine/netstack"
	_ "scale.ropsoft.cloud/words"
)
