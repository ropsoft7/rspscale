// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package wgengine provides the Rspscale WireGuard engine interface.
package wgengine

import (
	"errors"
	"net/netip"
	"time"

	"scale.ropsoft.cloud/ipn/ipnstate"
	"scale.ropsoft.cloud/net/dns"
	"scale.ropsoft.cloud/tailcfg"
	"scale.ropsoft.cloud/types/key"
	"scale.ropsoft.cloud/types/netmap"
	"scale.ropsoft.cloud/wgengine/capture"
	"scale.ropsoft.cloud/wgengine/filter"
	"scale.ropsoft.cloud/wgengine/router"
	"scale.ropsoft.cloud/wgengine/wgcfg"
	"scale.ropsoft.cloud/wgengine/wgint"
)

// Status is the Engine status.
//
// TODO(bradfitz): remove this, subset of ipnstate? Need to migrate users.
type Status struct {
	AsOf       time.Time // the time at which the status was calculated
	Peers      []ipnstate.PeerStatusLite
	LocalAddrs []tailcfg.Endpoint // the set of possible endpoints for the magic conn
	DERPs      int                // number of active DERP connections
}

// StatusCallback is the type of status callbacks used by
// Engine.SetStatusCallback.
//
// Exactly one of Status or error is non-nil.
type StatusCallback func(*Status, error)

// NetworkMapCallback is the type used by callbacks that hook
// into network map updates.
type NetworkMapCallback func(*netmap.NetworkMap)

// ErrNoChanges is returned by Engine.Reconfig if no changes were made.
var ErrNoChanges = errors.New("no changes made to Engine config")

// PeerForIP is the type returned by Engine.PeerForIP.
type PeerForIP struct {
	// Node is the matched node. It's always a valid value when
	// Engine.PeerForIP returns ok==true.
	Node tailcfg.NodeView

	// IsSelf is whether the Node is the local process.
	IsSelf bool

	// Route is the route that matched the IP provided
	// to Engine.PeerForIP.
	Route netip.Prefix
}

// Engine is the Rspscale WireGuard engine interface.
type Engine interface {
	// Reconfig reconfigures WireGuard and makes sure it's running.
	// This also handles setting up any kernel routes.
	//
	// This is called whenever tailcontrol (the control plane)
	// sends an updated network map.
	//
	// The returned error is ErrNoChanges if no changes were made.
	Reconfig(*wgcfg.Config, *router.Config, *dns.Config) error

	// PeerForIP returns the node to which the provided IP routes,
	// if any. If none is found, (nil, false) is returned.
	PeerForIP(netip.Addr) (_ PeerForIP, ok bool)

	// GetFilter returns the current packet filter, if any.
	GetFilter() *filter.Filter

	// SetFilter updates the packet filter.
	SetFilter(*filter.Filter)

	// GetJailedFilter returns the current packet filter for jailed nodes,
	// if any.
	GetJailedFilter() *filter.Filter

	// SetJailedFilter updates the packet filter for jailed nodes.
	SetJailedFilter(*filter.Filter)

	// SetStatusCallback sets the function to call when the
	// WireGuard status changes.
	SetStatusCallback(StatusCallback)

	// RequestStatus requests a WireGuard status update right
	// away, sent to the callback registered via SetStatusCallback.
	RequestStatus()

	// PeerByKey returns the WireGuard status of the provided peer.
	// If the peer is not found, ok is false.
	PeerByKey(key.NodePublic) (_ wgint.Peer, ok bool)

	// Close shuts down this wireguard instance, remove any routes
	// it added, etc. To bring it up again later, you'll need a
	// new Engine.
	Close()

	// Done returns a channel that is closed when the Engine's
	// Close method is called, the engine aborts with an error,
	// or it shuts down due to the closure of the underlying device.
	// You don't have to call this.
	Done() <-chan struct{}

	// SetNetworkMap informs the engine of the latest network map
	// from the server. The network map's DERPMap field should be
	// ignored as as it might be disabled; get it from SetDERPMap
	// instead.
	// The network map should only be read from.
	SetNetworkMap(*netmap.NetworkMap)

	// UpdateStatus populates the network state using the provided
	// status builder.
	UpdateStatus(*ipnstate.StatusBuilder)

	// Ping is a request to start a ping of the given message size to the peer
	// handling the given IP, then call cb with its ping latency & method.
	//
	// If size is zero too small, it is ignored. See rspscale.PingOpts for details.
	Ping(ip netip.Addr, pingType tailcfg.PingType, size int, cb func(*ipnstate.PingResult))

	// InstallCaptureHook registers a function to be called to capture
	// packets traversing the data path. The hook can be uninstalled by
	// calling this function with a nil value.
	InstallCaptureHook(capture.Callback)
}
