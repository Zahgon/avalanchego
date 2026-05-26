// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ips

import (
	"errors"
	"net/netip"
)

var errNoIPsFound = errors.New("no IPs found")

// Lookup attempts to resolve a hostname to a single IP. If multiple IPs are
// found, then lookup will attempt to return an IPv4 address, otherwise it will
// pick any of the IPs.
//
// Note: IPv4 is preferred because `net.Listen` prefers IPv4.
func Lookup(hostname string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}
