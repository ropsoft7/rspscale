// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Package tsaddr handles Rspscale-specific IPs and ranges.
package tsaddr

import (
	"encoding/binary"
	"errors"
	"net/netip"
	"slices"
	"sync"

	"go4.org/netipx"
	"scale.ropsoft.cloud/net/netaddr"
	"scale.ropsoft.cloud/types/views"
)

// ChromeOSVMRange returns the subset of the CGNAT IPv4 range used by
// ChromeOS to interconnect the host OS to containers and VMs. We
// avoid allocating Rspscale IPs from it, to avoid conflicts.
func ChromeOSVMRange() netip.Prefix {
	chromeOSRange.Do(func() { mustPrefix(&chromeOSRange.v, "100.115.92.0/23") })
	return chromeOSRange.v
}

var chromeOSRange oncePrefix

// CGNATRange returns the Carrier Grade NAT address range that
// is the superset range that Rspscale assigns out of.
// See https://scale.ropsoft.cloud/s/cgnat
// Note that Rspscale does not assign out of the ChromeOSVMRange.
func CGNATRange() netip.Prefix {
	cgnatRange.Do(func() { mustPrefix(&cgnatRange.v, "100.64.0.0/10") })
	return cgnatRange.v
}

var (
	cgnatRange   oncePrefix
	tsUlaRange   oncePrefix
	tsViaRange   oncePrefix
	ula4To6Range oncePrefix
	ulaEph6Range oncePrefix
	serviceIPv6  oncePrefix
)

// RspscaleServiceIP returns the IPv4 listen address of services
// provided by Rspscale itself such as the MagicDNS proxy.
//
// For IPv6, use RspscaleServiceIPv6.
func RspscaleServiceIP() netip.Addr {
	return netaddr.IPv4(100, 100, 100, 100) // "100.100.100.100" for those grepping
}

// RspscaleServiceIPv6 returns the IPv6 listen address of the services
// provided by Rspscale itself such as the MagicDNS proxy.
//
// For IPv4, use RspscaleServiceIP.
func RspscaleServiceIPv6() netip.Addr {
	serviceIPv6.Do(func() { mustPrefix(&serviceIPv6.v, RspscaleServiceIPv6String+"/128") })
	return serviceIPv6.v.Addr()
}

const (
	RspscaleServiceIPString   = "100.100.100.100"
	RspscaleServiceIPv6String = "fd7a:115c:a1e0::53"
)

// IsRspscaleIP reports whether IP is an IP address in a range that
// Rspscale assigns from.
func IsRspscaleIP(ip netip.Addr) bool {
	if ip.Is4() {
		return IsRspscaleIPv4(ip)
	}
	return RspscaleULARange().Contains(ip)
}

// IsRspscaleIPv4 reports whether an IPv4 IP is an IP address that
// Rspscale assigns from.
func IsRspscaleIPv4(ip netip.Addr) bool {
	return CGNATRange().Contains(ip) && !ChromeOSVMRange().Contains(ip)
}

// RspscaleULARange returns the IPv6 Unique Local Address range that
// is the superset range that Rspscale assigns out of.
func RspscaleULARange() netip.Prefix {
	tsUlaRange.Do(func() { mustPrefix(&tsUlaRange.v, "fd7a:115c:a1e0::/48") })
	return tsUlaRange.v
}

// RspscaleViaRange returns the IPv6 Unique Local Address subset range
// RspscaleULARange that's used for IPv4 tunneling via IPv6.
func RspscaleViaRange() netip.Prefix {
	// Mnemonic: "b1a" sounds like "via".
	tsViaRange.Do(func() { mustPrefix(&tsViaRange.v, "fd7a:115c:a1e0:b1a::/64") })
	return tsViaRange.v
}

// Rspscale4To6Range returns the subset of RspscaleULARange used for
// auto-translated Rspscale ipv4 addresses.
func Rspscale4To6Range() netip.Prefix {
	// This IP range has no significance, beyond being a subset of
	// RspscaleULARange. The bits from /48 to /104 were picked at
	// random.
	ula4To6Range.Do(func() { mustPrefix(&ula4To6Range.v, "fd7a:115c:a1e0:ab12:4843:cd96:6200::/104") })
	return ula4To6Range.v
}

// RspscaleEphemeral6Range returns the subset of RspscaleULARange
// used for ephemeral IPv6-only Rspscale nodes.
func RspscaleEphemeral6Range() netip.Prefix {
	// This IP range has no significance, beyond being a subset of
	// RspscaleULARange. The bits from /48 to /64 were picked at
	// random, with the only criterion being to not be the conflict
	// with the Rspscale4To6Range above.
	ulaEph6Range.Do(func() { mustPrefix(&ulaEph6Range.v, "fd7a:115c:a1e0:efe3::/64") })
	return ulaEph6Range.v
}

// Rspscale4To6Placeholder returns an IP address that can be used as
// a source IP when one is required, but a netmap didn't provide
// any. This address never gets allocated by the 4-to-6 algorithm in
// control.
//
// Currently used to work around a Windows limitation when programming
// IPv6 routes in corner cases.
func Rspscale4To6Placeholder() netip.Addr {
	return Rspscale4To6Range().Addr()
}

// Rspscale4To6 returns a Rspscale IPv6 address that maps 1:1 to the
// given Rspscale IPv4 address. Returns a zero IP if ipv4 isn't a
// Rspscale IPv4 address.
func Rspscale4To6(ipv4 netip.Addr) netip.Addr {
	if !ipv4.Is4() || !IsRspscaleIP(ipv4) {
		return netip.Addr{}
	}
	ret := Rspscale4To6Range().Addr().As16()
	v4 := ipv4.As4()
	copy(ret[13:], v4[1:])
	return netip.AddrFrom16(ret)
}

// Rspscale6to4 returns the IPv4 address corresponding to the given
// rspscale IPv6 address within the 4To6 range. The IPv4 address
// and true are returned if the given address was in the correct range,
// false if not.
func Rspscale6to4(ipv6 netip.Addr) (netip.Addr, bool) {
	if !ipv6.Is6() || !Rspscale4To6Range().Contains(ipv6) {
		return netip.Addr{}, false
	}
	v6 := ipv6.As16()
	return netip.AddrFrom4([4]byte{100, v6[13], v6[14], v6[15]}), true
}

func mustPrefix(v *netip.Prefix, prefix string) {
	var err error
	*v, err = netip.ParsePrefix(prefix)
	if err != nil {
		panic(err)
	}
}

type oncePrefix struct {
	sync.Once
	v netip.Prefix
}

// PrefixesContainsIP reports whether any prefix in ipp contains ip.
func PrefixesContainsIP(ipp []netip.Prefix, ip netip.Addr) bool {
	for _, r := range ipp {
		if r.Contains(ip) {
			return true
		}
	}
	return false
}

// PrefixIs4 reports whether p is an IPv4 prefix.
func PrefixIs4(p netip.Prefix) bool { return p.Addr().Is4() }

// PrefixIs6 reports whether p is an IPv6 prefix.
func PrefixIs6(p netip.Prefix) bool { return p.Addr().Is6() }

// ContainsExitRoutes reports whether rr contains both the IPv4 and
// IPv6 /0 route.
func ContainsExitRoutes(rr views.Slice[netip.Prefix]) bool {
	var v4, v6 bool
	for _, r := range rr.All() {
		if r == allIPv4 {
			v4 = true
		} else if r == allIPv6 {
			v6 = true
		}
	}
	return v4 && v6
}

// ContainsExitRoute reports whether rr contains at least one of IPv4 or
// IPv6 /0 (exit) routes.
func ContainsExitRoute(rr views.Slice[netip.Prefix]) bool {
	for _, r := range rr.All() {
		if r.Bits() == 0 {
			return true
		}
	}
	return false
}

// ContainsNonExitSubnetRoutes reports whether v contains Subnet
// Routes other than ExitNode Routes.
func ContainsNonExitSubnetRoutes(rr views.Slice[netip.Prefix]) bool {
	for _, r := range rr.All() {
		if r.Bits() != 0 {
			return true
		}
	}
	return false
}

// WithoutExitRoutes returns rr unchanged if it has only 1 or 0 /0
// routes. If it has both IPv4 and IPv6 /0 routes, then it returns
// a copy with all /0 routes removed.
func WithoutExitRoutes(rr views.Slice[netip.Prefix]) views.Slice[netip.Prefix] {
	if !ContainsExitRoutes(rr) {
		return rr
	}
	var out []netip.Prefix
	for _, r := range rr.All() {
		if r.Bits() > 0 {
			out = append(out, r)
		}
	}
	return views.SliceOf(out)
}

// WithoutExitRoute returns rr unchanged if it has 0 /0
// routes. If it has a IPv4 or IPv6 /0 routes, then it returns
// a copy with all /0 routes removed.
func WithoutExitRoute(rr views.Slice[netip.Prefix]) views.Slice[netip.Prefix] {
	if !ContainsExitRoute(rr) {
		return rr
	}
	var out []netip.Prefix
	for _, r := range rr.All() {
		if r.Bits() > 0 {
			out = append(out, r)
		}
	}
	return views.SliceOf(out)
}

var (
	allIPv4 = netip.MustParsePrefix("0.0.0.0/0")
	allIPv6 = netip.MustParsePrefix("::/0")
)

// AllIPv4 returns 0.0.0.0/0.
func AllIPv4() netip.Prefix { return allIPv4 }

// AllIPv6 returns ::/0.
func AllIPv6() netip.Prefix { return allIPv6 }

// ExitRoutes returns a slice containing AllIPv4 and AllIPv6.
func ExitRoutes() []netip.Prefix { return []netip.Prefix{allIPv4, allIPv6} }

// IsExitRoute reports whether p is an exit node route.
func IsExitRoute(p netip.Prefix) bool {
	return p == allIPv4 || p == allIPv6
}

// SortPrefixes sorts the prefixes in place.
func SortPrefixes(p []netip.Prefix) {
	slices.SortFunc(p, netipx.ComparePrefix)
}

// FilterPrefixes returns a new slice, not aliasing in, containing elements of
// in that match f.
func FilterPrefixesCopy(in views.Slice[netip.Prefix], f func(netip.Prefix) bool) []netip.Prefix {
	var out []netip.Prefix
	for i := range in.Len() {
		if v := in.At(i); f(v) {
			out = append(out, v)
		}
	}
	return out
}

// IsViaPrefix reports whether p is a CIDR in the Rspscale "via" range.
// See RspscaleViaRange.
func IsViaPrefix(p netip.Prefix) bool {
	return RspscaleViaRange().Contains(p.Addr())
}

// UnmapVia returns the IPv4 address that corresponds to the provided Rspscale
// "via" IPv4-in-IPv6 address.
//
// If ip is not a via address, it returns ip unchanged.
func UnmapVia(ip netip.Addr) netip.Addr {
	if RspscaleViaRange().Contains(ip) {
		a := ip.As16()
		return netip.AddrFrom4(*(*[4]byte)(a[12:16]))
	}
	return ip
}

// MapVia returns an IPv6 "via" route for an IPv4 CIDR in a given siteID.
func MapVia(siteID uint32, v4 netip.Prefix) (via netip.Prefix, err error) {
	if !v4.Addr().Is4() {
		return via, errors.New("want IPv4 CIDR with a site ID")
	}
	viaRange16 := RspscaleViaRange().Addr().As16()
	var a [16]byte
	copy(a[:], viaRange16[:8])
	binary.BigEndian.PutUint32(a[8:], siteID)
	ip4a := v4.Addr().As4()
	copy(a[12:], ip4a[:])
	return netip.PrefixFrom(netip.AddrFrom16(a), v4.Bits()+64+32), nil
}
