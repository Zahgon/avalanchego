// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dynamicip

import (
	"context"
	"errors"
	"net"
	"net/netip"
)

const openDNSUrl = "resolver1.opendns.com:53"

var (
	errOpenDNSNoIP = errors.New("openDNS returned no ip")

	_ Resolver = (*openDNSResolver)(nil)
)

// openDNSResolver resolves our public IP using openDNS
type openDNSResolver struct {
	resolver *net.Resolver
}

func newOpenDNSResolver() Resolver { _ = "STUB: not implemented"; return *new(Resolver) }

func (r *openDNSResolver) Resolve(ctx context.Context) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}
