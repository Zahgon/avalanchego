// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"github.com/ava-labs/avalanchego/ids"
)

// NoPrecondition can be supplied to [Set.Sample] to indicate that all peers
// are eligible to be returned in the sample.
func NoPrecondition(*Peer) bool {
	_ = "STUB: not implemented"

	// Set contains a group of peers.
	return false
}

type Set struct {
	peersMap   map[ids.NodeID]int // nodeID -> peer's index in peersSlice
	peersSlice []*Peer            // invariant: len(peersSlice) == len(peersMap)
}

// NewSet returns a set that does not internally manage synchronization.
//
// Only [Set.Add] and [Set.Remove] require exclusion on the data structure. The
// remaining methods are safe for concurrent use.
func NewSet() *Set { _ = "STUB: not implemented"; return nil }

// Add this peer to the set.
//
// If a peer with the same [Peer.ID] is already in the set, then the new
// peer instance will replace the old peer instance.
//
// Add does not change the [Peer.ID] returned from calls to [Set.GetByIndex].
func (s *Set) Add(peer *Peer) { _ = "STUB: not implemented"; return }

// GetByID attempts to fetch a [Peer] whose [Peer.ID] is equal to nodeID.
// If no such peer exists in the set, then [false] will be returned.
func (s *Set) GetByID(nodeID ids.NodeID) (*Peer, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetByIndex attempts to fetch a [Peer] who has been allocated index.
// If index < 0 or index >= [Set.Len], then false will be returned.
func (s *Set) GetByIndex(index int) (*Peer, bool) { _ = "STUB: not implemented"; return nil, false }

// Remove any [Peer] whose [Peer.ID] is equal to nodeID from the set.
func (s *Set) Remove(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// Len returns the number of peers currently in this set.
func (s *Set) Len() int { _ = "STUB: not implemented"; return 0 }

// Sample attempts to return a random slice of peers with length n. The
// slice will not include any duplicates. Only peers that cause the
// precondition to return true will be returned in the slice.
func (s *Set) Sample(n int, precondition func(*Peer) bool) []*Peer {
	_ = "STUB: not implemented"
	return nil
}

// We have run out of peers to attempt to sample.

// AllInfo returns information about all the peers.
func (s *Set) AllInfo() []Info { _ = "STUB: not implemented"; return nil }

// Info returns information about the requested peers if they are in the set.
func (s *Set) Info(nodeIDs []ids.NodeID) []Info { _ = "STUB: not implemented"; return nil }
