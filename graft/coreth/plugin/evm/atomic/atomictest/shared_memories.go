// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"testing"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
)

var removeValue = []byte{0x1}

type SharedMemories struct {
	ThisChain   atomic.SharedMemory
	PeerChain   atomic.SharedMemory
	thisChainID ids.ID
	peerChainID ids.ID
}

func (s *SharedMemories) AddItemsToBeRemovedToPeerChain(ops map[ids.ID]*atomic.Requests) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedMemories) AssertOpsApplied(t *testing.T, ops map[ids.ID]*atomic.Requests) {
	_ = "STUB: not implemented"
	return
}

// should be able to get put requests

// should not be able to get remove requests

func (s *SharedMemories) AssertOpsNotApplied(t *testing.T, ops map[ids.ID]*atomic.Requests) {
	_ = "STUB: not implemented"
	return
}

// should not be able to get put requests

// should be able to get remove requests (these were previously added as puts on peerChain)

func NewSharedMemories(atomicMemory *atomic.Memory, thisChainID, peerChainID ids.ID) *SharedMemories {
	_ = "STUB: not implemented"
	return nil
}

func TestSharedMemory() atomic.SharedMemory {
	_ = "STUB: not implemented"
	return *new(atomic.SharedMemory)
}
