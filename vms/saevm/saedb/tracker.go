// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saedb

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/state/snapshot"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
)

// Config allows parameterization of the TrieDB and when state is committed.
type Config struct {
	// TODO(alarso16): move minimal elements to config and construct in method.
	TrieDBConfig       *triedb.Config
	Archival           bool // if true, will store every state on disk
	TrieCommitInterval uint64
}

// defaultCommitInterval is the default number of blocks between commits of the
// state trie to disk.
const defaultCommitInterval = 4096

// CommitInterval returns the trie commit interval.
func (c Config) CommitInterval() uint64 { _ = "STUB: not implemented"; return 0 }

// SnapshotCacheSizeMB is the snapshot cache size used by a [Tracker].
// TODO(alarso16): move to config
const SnapshotCacheSizeMB = 128

var _ StateDBOpener = (*Tracker)(nil)

// Tracker provides an abstraction to state-related operations, managing all
// database operations not exposed by the [state.StateDB] itself.
//
// All methods are safe to be called even after [Tracker.Close], but state
// will be unavailable.
type Tracker struct {
	snaps    *snapshot.Tree
	cache    state.Database
	isHashDB bool
	config   Config
	log      logging.Logger
}

// NewTracker provides a new [Tracker] on the underlying database.
func NewTracker(db ethdb.Database, c Config, lastExecuted common.Hash, log logging.Logger) (*Tracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Track tracks the root and may commit the trie associated with the root
// to the database if [Config.ShouldCommitTrieDB] returns true, or the [Config]
// specifies that the node is archival.
//
// This state will be available in memory until [Tracker.Untrack] has been
// called for the root as many times as [Tracker.Track] has been called.
func (t *Tracker) Track(root common.Hash) { _ = "STUB: not implemented"; return }

// Never returns an error because of the above check.

// MaybeCommit potentially calls [triedb.Database.Commit], based on the
// following priorities:
//
// 1. If [Config.Archival] is true, then `executionRoot` will be committed.
// 2. If [ShouldCommitTrieDB] based on `height`, `settledRoot` is committed.
// 3. Otherwise, nothing is committed.
//
// This does NOT change in-memory tracking.
func (t *Tracker) MaybeCommit(settledRoot, executionRoot common.Hash, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

/* log */

// LastHeightWithExecutionRootCommitted returns the greatest block height for
// which [Tracker.MaybeCommit] called [triedb.Database.Commit] with the
// post-execution state root of the block.
func LastHeightWithExecutionRootCommitted(db ethdb.Database, c Config, hooks hook.Points, lastSynchronous uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// Untrack informs the [Tracker] that the state corresponding
// with `root` can have its reference count reduced. If the reference
// count is 0, the state will be removed from memory.
//
// This should be called on each block after its state is no longer
// needed. If the state is already on disk, no operation is performed.
func (t *Tracker) Untrack(root common.Hash) { _ = "STUB: not implemented"; return }

// Never returns an error because of the above check.

// StateDB provides a [state.StateDB] at the given root.
//
// Each [state.StateDB] can be constructed and used concurrently.
// However, the right to call [state.StateDB.Commit] is reserved
// for canonical blocks, as any other use could result in a memory
// leak or state corruption.
func (t *Tracker) StateDB(root common.Hash) (*state.StateDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close releases all resources associated with the `[triedb.Database]`
// and persists `lastRoot` to the snapshot layer. `lastRoot` should be a
// recent state root.
func (t *Tracker) Close(lastRoot common.Hash) (errs error) { _ = "STUB: not implemented"; return nil }

// We don't use [snapshot.Tree.Journal] because re-orgs are impossible under
// SAE so we don't mind flattening all snapshot layers to disk. Note that
// calling `Cap([disk root], 0)` returns an error when it's actually a
// no-op, so we ensure there are changes.
