// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nat

import (
	"errors"
	"net/netip"
	"time"
)

var (
	_ Router = (*noRouter)(nil)

	errNoRouterCantMapPorts = errors.New("can't map ports without a known router")
	errFetchingIP           = errors.New("getting outbound IP failed")
)

const googleDNSServer = "8.8.8.8:80"

type noRouter struct {
	ip    netip.Addr
	ipErr error
}

func (noRouter) SupportsNAT() bool { _ = "STUB: not implemented"; return false }

func (noRouter) MapPort(uint16, uint16, string, time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (noRouter) UnmapPort(uint16, uint16) error { _ = "STUB: not implemented"; return nil }

func (r noRouter) ExternalIP() (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func getOutboundIP() (netip.Addr, error) { _ = "STUB: not implemented"; return *new(netip.Addr), nil }

// NewNoRouter returns a router that assumes the network is public
func NewNoRouter() Router { _ = "STUB: not implemented"; return *new(Router) }
