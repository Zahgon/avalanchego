// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dynamicip

import (
	"context"
	"net/netip"
)

var _ Resolver = (*ifConfigResolver)(nil)

// ifConfigResolver resolves our public IP using ifconfig's format.
type ifConfigResolver struct {
	url string
}

func (r *ifConfigResolver) Resolve(ctx context.Context) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

//nolint:bodyclose // body is closed via rpc.CleanlyCloseBody

// Drop any error to report the original error
