// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"net/netip"

	"github.com/ava-labs/avalanchego/ids"
)

var _ Beacon = (*beacon)(nil)

type Beacon interface {
	ID() ids.NodeID
	IP() netip.AddrPort
}

type beacon struct {
	id ids.NodeID
	ip netip.AddrPort
}

func New(id ids.NodeID, ip netip.AddrPort) Beacon { _ = "STUB: not implemented"; return *new(Beacon) }

func (b *beacon) ID() ids.NodeID { _ = "STUB: not implemented"; return *new(ids.NodeID) }

func (b *beacon) IP() netip.AddrPort { _ = "STUB: not implemented"; return *new(netip.AddrPort) }
