// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by rspscale/cmd/viewer; DO NOT EDIT.

package ipn

import (
	"encoding/json"
	"errors"
	"net/netip"

	"scale.ropsoft.cloud/drive"
	"scale.ropsoft.cloud/tailcfg"
	"scale.ropsoft.cloud/types/opt"
	"scale.ropsoft.cloud/types/persist"
	"scale.ropsoft.cloud/types/preftype"
	"scale.ropsoft.cloud/types/views"
)

//go:generate go run scale.ropsoft.cloud/cmd/cloner  -clonefunc=false -type=Prefs,ServeConfig,ServiceConfig,TCPPortHandler,HTTPHandler,WebServerConfig

// View returns a readonly view of Prefs.
func (p *Prefs) View() PrefsView {
	return PrefsView{ж: p}
}

// PrefsView provides a read-only view over Prefs.
//
// Its methods should only be called if `Valid()` returns true.
type PrefsView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *Prefs
}

// Valid reports whether underlying value is non-nil.
func (v PrefsView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v PrefsView) AsStruct() *Prefs {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v PrefsView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *PrefsView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x Prefs
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v PrefsView) ControlURL() string                          { return v.ж.ControlURL }
func (v PrefsView) RouteAll() bool                              { return v.ж.RouteAll }
func (v PrefsView) ExitNodeID() tailcfg.StableNodeID            { return v.ж.ExitNodeID }
func (v PrefsView) ExitNodeIP() netip.Addr                      { return v.ж.ExitNodeIP }
func (v PrefsView) InternalExitNodePrior() tailcfg.StableNodeID { return v.ж.InternalExitNodePrior }
func (v PrefsView) ExitNodeAllowLANAccess() bool                { return v.ж.ExitNodeAllowLANAccess }
func (v PrefsView) CorpDNS() bool                               { return v.ж.CorpDNS }
func (v PrefsView) RunSSH() bool                                { return v.ж.RunSSH }
func (v PrefsView) RunWebClient() bool                          { return v.ж.RunWebClient }
func (v PrefsView) WantRunning() bool                           { return v.ж.WantRunning }
func (v PrefsView) LoggedOut() bool                             { return v.ж.LoggedOut }
func (v PrefsView) ShieldsUp() bool                             { return v.ж.ShieldsUp }
func (v PrefsView) AdvertiseTags() views.Slice[string]          { return views.SliceOf(v.ж.AdvertiseTags) }
func (v PrefsView) Hostname() string                            { return v.ж.Hostname }
func (v PrefsView) NotepadURLs() bool                           { return v.ж.NotepadURLs }
func (v PrefsView) ForceDaemon() bool                           { return v.ж.ForceDaemon }
func (v PrefsView) Egg() bool                                   { return v.ж.Egg }
func (v PrefsView) AdvertiseRoutes() views.Slice[netip.Prefix] {
	return views.SliceOf(v.ж.AdvertiseRoutes)
}
func (v PrefsView) AdvertiseServices() views.Slice[string] {
	return views.SliceOf(v.ж.AdvertiseServices)
}
func (v PrefsView) NoSNAT() bool                          { return v.ж.NoSNAT }
func (v PrefsView) NoStatefulFiltering() opt.Bool         { return v.ж.NoStatefulFiltering }
func (v PrefsView) NetfilterMode() preftype.NetfilterMode { return v.ж.NetfilterMode }
func (v PrefsView) OperatorUser() string                  { return v.ж.OperatorUser }
func (v PrefsView) ProfileName() string                   { return v.ж.ProfileName }
func (v PrefsView) AutoUpdate() AutoUpdatePrefs           { return v.ж.AutoUpdate }
func (v PrefsView) AppConnector() AppConnectorPrefs       { return v.ж.AppConnector }
func (v PrefsView) PostureChecking() bool                 { return v.ж.PostureChecking }
func (v PrefsView) NetfilterKind() string                 { return v.ж.NetfilterKind }
func (v PrefsView) DriveShares() views.SliceView[*drive.Share, drive.ShareView] {
	return views.SliceOfViews[*drive.Share, drive.ShareView](v.ж.DriveShares)
}
func (v PrefsView) AllowSingleHosts() marshalAsTrueInJSON { return v.ж.AllowSingleHosts }
func (v PrefsView) Persist() persist.PersistView          { return v.ж.Persist.View() }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _PrefsViewNeedsRegeneration = Prefs(struct {
	ControlURL             string
	RouteAll               bool
	ExitNodeID             tailcfg.StableNodeID
	ExitNodeIP             netip.Addr
	InternalExitNodePrior  tailcfg.StableNodeID
	ExitNodeAllowLANAccess bool
	CorpDNS                bool
	RunSSH                 bool
	RunWebClient           bool
	WantRunning            bool
	LoggedOut              bool
	ShieldsUp              bool
	AdvertiseTags          []string
	Hostname               string
	NotepadURLs            bool
	ForceDaemon            bool
	Egg                    bool
	AdvertiseRoutes        []netip.Prefix
	AdvertiseServices      []string
	NoSNAT                 bool
	NoStatefulFiltering    opt.Bool
	NetfilterMode          preftype.NetfilterMode
	OperatorUser           string
	ProfileName            string
	AutoUpdate             AutoUpdatePrefs
	AppConnector           AppConnectorPrefs
	PostureChecking        bool
	NetfilterKind          string
	DriveShares            []*drive.Share
	AllowSingleHosts       marshalAsTrueInJSON
	Persist                *persist.Persist
}{})

// View returns a readonly view of ServeConfig.
func (p *ServeConfig) View() ServeConfigView {
	return ServeConfigView{ж: p}
}

// ServeConfigView provides a read-only view over ServeConfig.
//
// Its methods should only be called if `Valid()` returns true.
type ServeConfigView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *ServeConfig
}

// Valid reports whether underlying value is non-nil.
func (v ServeConfigView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v ServeConfigView) AsStruct() *ServeConfig {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v ServeConfigView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *ServeConfigView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x ServeConfig
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v ServeConfigView) TCP() views.MapFn[uint16, *TCPPortHandler, TCPPortHandlerView] {
	return views.MapFnOf(v.ж.TCP, func(t *TCPPortHandler) TCPPortHandlerView {
		return t.View()
	})
}

func (v ServeConfigView) Web() views.MapFn[HostPort, *WebServerConfig, WebServerConfigView] {
	return views.MapFnOf(v.ж.Web, func(t *WebServerConfig) WebServerConfigView {
		return t.View()
	})
}

func (v ServeConfigView) Services() views.MapFn[string, *ServiceConfig, ServiceConfigView] {
	return views.MapFnOf(v.ж.Services, func(t *ServiceConfig) ServiceConfigView {
		return t.View()
	})
}

func (v ServeConfigView) AllowFunnel() views.Map[HostPort, bool] {
	return views.MapOf(v.ж.AllowFunnel)
}

func (v ServeConfigView) Foreground() views.MapFn[string, *ServeConfig, ServeConfigView] {
	return views.MapFnOf(v.ж.Foreground, func(t *ServeConfig) ServeConfigView {
		return t.View()
	})
}
func (v ServeConfigView) ETag() string { return v.ж.ETag }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _ServeConfigViewNeedsRegeneration = ServeConfig(struct {
	TCP         map[uint16]*TCPPortHandler
	Web         map[HostPort]*WebServerConfig
	Services    map[string]*ServiceConfig
	AllowFunnel map[HostPort]bool
	Foreground  map[string]*ServeConfig
	ETag        string
}{})

// View returns a readonly view of ServiceConfig.
func (p *ServiceConfig) View() ServiceConfigView {
	return ServiceConfigView{ж: p}
}

// ServiceConfigView provides a read-only view over ServiceConfig.
//
// Its methods should only be called if `Valid()` returns true.
type ServiceConfigView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *ServiceConfig
}

// Valid reports whether underlying value is non-nil.
func (v ServiceConfigView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v ServiceConfigView) AsStruct() *ServiceConfig {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v ServiceConfigView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *ServiceConfigView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x ServiceConfig
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v ServiceConfigView) TCP() views.MapFn[uint16, *TCPPortHandler, TCPPortHandlerView] {
	return views.MapFnOf(v.ж.TCP, func(t *TCPPortHandler) TCPPortHandlerView {
		return t.View()
	})
}

func (v ServiceConfigView) Web() views.MapFn[HostPort, *WebServerConfig, WebServerConfigView] {
	return views.MapFnOf(v.ж.Web, func(t *WebServerConfig) WebServerConfigView {
		return t.View()
	})
}
func (v ServiceConfigView) Tun() bool { return v.ж.Tun }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _ServiceConfigViewNeedsRegeneration = ServiceConfig(struct {
	TCP map[uint16]*TCPPortHandler
	Web map[HostPort]*WebServerConfig
	Tun bool
}{})

// View returns a readonly view of TCPPortHandler.
func (p *TCPPortHandler) View() TCPPortHandlerView {
	return TCPPortHandlerView{ж: p}
}

// TCPPortHandlerView provides a read-only view over TCPPortHandler.
//
// Its methods should only be called if `Valid()` returns true.
type TCPPortHandlerView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TCPPortHandler
}

// Valid reports whether underlying value is non-nil.
func (v TCPPortHandlerView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TCPPortHandlerView) AsStruct() *TCPPortHandler {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TCPPortHandlerView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *TCPPortHandlerView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TCPPortHandler
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TCPPortHandlerView) HTTPS() bool          { return v.ж.HTTPS }
func (v TCPPortHandlerView) HTTP() bool           { return v.ж.HTTP }
func (v TCPPortHandlerView) TCPForward() string   { return v.ж.TCPForward }
func (v TCPPortHandlerView) TerminateTLS() string { return v.ж.TerminateTLS }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TCPPortHandlerViewNeedsRegeneration = TCPPortHandler(struct {
	HTTPS        bool
	HTTP         bool
	TCPForward   string
	TerminateTLS string
}{})

// View returns a readonly view of HTTPHandler.
func (p *HTTPHandler) View() HTTPHandlerView {
	return HTTPHandlerView{ж: p}
}

// HTTPHandlerView provides a read-only view over HTTPHandler.
//
// Its methods should only be called if `Valid()` returns true.
type HTTPHandlerView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *HTTPHandler
}

// Valid reports whether underlying value is non-nil.
func (v HTTPHandlerView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v HTTPHandlerView) AsStruct() *HTTPHandler {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v HTTPHandlerView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *HTTPHandlerView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x HTTPHandler
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v HTTPHandlerView) Path() string  { return v.ж.Path }
func (v HTTPHandlerView) Proxy() string { return v.ж.Proxy }
func (v HTTPHandlerView) Text() string  { return v.ж.Text }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _HTTPHandlerViewNeedsRegeneration = HTTPHandler(struct {
	Path  string
	Proxy string
	Text  string
}{})

// View returns a readonly view of WebServerConfig.
func (p *WebServerConfig) View() WebServerConfigView {
	return WebServerConfigView{ж: p}
}

// WebServerConfigView provides a read-only view over WebServerConfig.
//
// Its methods should only be called if `Valid()` returns true.
type WebServerConfigView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *WebServerConfig
}

// Valid reports whether underlying value is non-nil.
func (v WebServerConfigView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v WebServerConfigView) AsStruct() *WebServerConfig {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v WebServerConfigView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v *WebServerConfigView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x WebServerConfig
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v WebServerConfigView) Handlers() views.MapFn[string, *HTTPHandler, HTTPHandlerView] {
	return views.MapFnOf(v.ж.Handlers, func(t *HTTPHandler) HTTPHandlerView {
		return t.View()
	})
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _WebServerConfigViewNeedsRegeneration = WebServerConfig(struct {
	Handlers map[string]*HTTPHandler
}{})