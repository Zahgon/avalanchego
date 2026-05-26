// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/state"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

// Extender is the sync extender for the atomic VM.
type Extender struct {
	backend     *state.AtomicBackend
	trie        *state.AtomicTrie
	requestSize uint16 // maximum number of leaves to sync in a single request
}

// Initialize initializes the sync extender with the backend and trie and request size.
func (e *Extender) Initialize(backend *state.AtomicBackend, trie *state.AtomicTrie, requestSize uint16) {
	_ = "STUB: not implemented"
	return
}

// CreateSyncer creates the atomic syncer with the given client and verDB.
func (e *Extender) CreateSyncer(client types.LeafClient, verDB *versiondb.Database, summary message.Syncable) (types.Syncer, error) {
	_ = "STUB: not implemented"
	return *new(types.Syncer), nil
}

// OnFinishBeforeCommit implements the sync.Extender interface by marking the previously last accepted block for the shared memory cursor.
func (e *Extender) OnFinishBeforeCommit(lastAcceptedHeight uint64, summary message.Syncable) error {
	_ = "STUB: not implemented"
	// Mark the previously last accepted block for the shared memory cursor, so that we will execute shared
	// memory operations from the previously last accepted block when ApplyToSharedMemory
	// is called.
	return nil
}

// OnFinishAfterCommit implements the sync.Extender interface by applying the atomic trie to the shared memory.
func (e *Extender) OnFinishAfterCommit(summaryHeight uint64) error {
	_ = "STUB: not implemented"
	// the chain state is already restored, and, from this point on,
	// the block synced to is the accepted block. The last operation
	// is updating shared memory with the atomic trie.
	// ApplyToSharedMemory does this, and, even if the VM is stopped
	// (gracefully or ungracefully), since MarkApplyToSharedMemoryCursor
	// is called, VM will resume ApplyToSharedMemory on Initialize.
	return nil
}
