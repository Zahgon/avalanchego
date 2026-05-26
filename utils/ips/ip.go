// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ips

import "net/netip"

// IsPublic returns true if the provided address is considered to be a public
// IP.
func IsPublic(addr netip.Addr) bool { _ = "STUB: not implemented"; return false }

// ParseAddr returns the IP address from the provided string. If the string
// represents an IPv4 address in an IPv6 address, the IPv4 address is returned.
func ParseAddr(s string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// ParseAddrPort returns the IP:port address from the provided string. If the
// string represents an IPv4 address in an IPv6 address, the IPv4 address is
// returned.
func ParseAddrPort(s string) (netip.AddrPort, error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil
}

// AddrFromSlice returns the IP address from the provided byte slice. If the
// byte slice represents an IPv4 address in an IPv6 address, the IPv4 address is
// returned.
func AddrFromSlice(b []byte) (netip.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), false
}
