// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"errors"
	"net/netip"

	"github.com/ava-labs/avalanchego/ids"
)

var (
	_ Set = (*set)(nil)

	errDuplicateID = errors.New("duplicated ID")
	errDuplicateIP = errors.New("duplicated IP")

	errUnknownID = errors.New("unknown ID")
	errUnknownIP = errors.New("unknown IP")
)

type Set interface {
	Add(Beacon) error

	RemoveByID(ids.NodeID) error
	RemoveByIP(netip.AddrPort) error

	Len() int

	IDsArg() string
	IPsArg() string
}

type set struct {
	ids     map[ids.NodeID]int
	ips     map[netip.AddrPort]int
	beacons []Beacon
}

func NewSet() Set { _ = "STUB: not implemented"; return *new(Set) }

func (s *set) Add(b Beacon) error { _ = "STUB: not implemented"; return nil }

func (s *set) RemoveByID(idToRemove ids.NodeID) error { _ = "STUB: not implemented"; return nil }

func (s *set) RemoveByIP(ip netip.AddrPort) error { _ = "STUB: not implemented"; return nil }

func (s *set) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *set) IDsArg() string { _ = "STUB: not implemented"; return "" }

func (s *set) IPsArg() string { _ = "STUB: not implemented"; return "" }
