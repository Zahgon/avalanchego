// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"errors"

	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/utils"
)

const (
	NodeIDPrefix = "NodeID-"
	NodeIDLen    = ShortIDLen
)

var (
	EmptyNodeID = NodeID{}

	errShortNodeID = errors.New("insufficient NodeID length")

	_ utils.Sortable[NodeID] = NodeID{}
)

type NodeID ShortID

func (id NodeID) String() string { _ = "STUB: not implemented"; return "" }

func (id NodeID) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (id NodeID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id NodeID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *NodeID) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If "null", do nothing

func (id *NodeID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (id NodeID) Compare(other NodeID) int { _ = "STUB: not implemented"; return 0 }

// ToNodeID attempt to convert a byte slice into a node id
func ToNodeID(bytes []byte) (NodeID, error) { _ = "STUB: not implemented"; return *new(NodeID), nil }

func NodeIDFromCert(cert *staking.Certificate) NodeID {
	_ = "STUB: not implemented"
	return *new(NodeID)
}

// NodeIDFromString is the inverse of NodeID.String()
func NodeIDFromString(nodeIDStr string) (NodeID, error) {
	_ = "STUB: not implemented"
	return *new(NodeID), nil
}
