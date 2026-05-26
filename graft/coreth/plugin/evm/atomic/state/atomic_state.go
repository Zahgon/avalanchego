// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

// atomicState implements the AtomicState interface using
// a pointer to the atomicBackend.
type atomicState struct {
	backend     *AtomicBackend
	blockHash   common.Hash
	blockHeight uint64
	txs         []*atomic.Tx
	atomicOps   map[ids.ID]*avalancheatomic.Requests
	atomicRoot  common.Hash
}

func (a *atomicState) Root() common.Hash {
	_ = "STUB: not implemented"
	return *

	// Accept writes the atomic operations to the database and
	// updates the last accepted block in the atomic backend.
	// It also commits the `commitBatch` to the shared memory.
	new(common.Hash)
}

func (a *atomicState) Accept(commitBatch database.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the atomic tx repository. Note it is necessary to invoke
// the correct method taking bonus blocks into consideration.

// Accept the root of this atomic trie (will be persisted if at a commit interval)

// Update the last accepted block to this block and remove it from
// the map tracking undecided blocks.

// get changes from the atomic trie and repository in a batch
// to be committed atomically with [commitBatch] and shared memory.

// If this is a bonus block, write [commitBatch] without applying atomic ops
// to shared memory.

// Otherwise, atomically commit pending changes in the version db with
// atomic ops to shared memory.

// Reject frees memory associated with the state change.
func (a *atomicState) Reject() error {
	_ = "STUB: not implemented"
	// Remove the block from the map of undecided blocks.
	return nil
}

// Unpin the rejected atomic trie root from memory.
