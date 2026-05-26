// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package snapshot implements a journalled, dynamic state dump.
package snapshot

import (
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/libevm/common"
	ethsnapshot "github.com/ava-labs/libevm/core/state/snapshot"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/stateconf"
	"github.com/ava-labs/libevm/metrics"
	"github.com/ava-labs/libevm/triedb"
)

const (
	// skipGenThreshold is the minimum time that must have elapsed since the
	// creation of the previous disk layer to start snapshot generation on a new
	// disk layer.
	//
	// If disk layers are being discarded at a frequency greater than this threshold,
	// starting snapshot generation is not worth it (will be aborted before meaningful
	// work can be done).
	skipGenThreshold = 500 * time.Millisecond
)

// ====== If resolving merge conflicts ======
//
// All calls to metrics.NewRegistered*() for metrics also defined in libevm/core/state/snapshot
// have been replaced with metrics.GetOrRegister*() to get metrics already registered in
// libevm/core/state/snapshot or register them here otherwise. These replacements ensure the
// same metrics are shared between the two packages.
var (
	snapshotCleanAccountHitMeter   = metrics.GetOrRegisterMeter("state/snapshot/clean/account/hit", nil)
	snapshotCleanAccountMissMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/account/miss", nil)
	snapshotCleanAccountInexMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/account/inex", nil)
	snapshotCleanAccountReadMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/account/read", nil)
	snapshotCleanAccountWriteMeter = metrics.GetOrRegisterMeter("state/snapshot/clean/account/write", nil)

	snapshotCleanStorageHitMeter   = metrics.GetOrRegisterMeter("state/snapshot/clean/storage/hit", nil)
	snapshotCleanStorageMissMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/storage/miss", nil)
	snapshotCleanStorageInexMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/storage/inex", nil)
	snapshotCleanStorageReadMeter  = metrics.GetOrRegisterMeter("state/snapshot/clean/storage/read", nil)
	snapshotCleanStorageWriteMeter = metrics.GetOrRegisterMeter("state/snapshot/clean/storage/write", nil)

	snapshotDirtyAccountHitMeter   = metrics.GetOrRegisterMeter("state/snapshot/dirty/account/hit", nil)
	snapshotDirtyAccountMissMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/account/miss", nil)
	snapshotDirtyAccountInexMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/account/inex", nil)
	snapshotDirtyAccountReadMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/account/read", nil)
	snapshotDirtyAccountWriteMeter = metrics.GetOrRegisterMeter("state/snapshot/dirty/account/write", nil)

	snapshotDirtyStorageHitMeter   = metrics.GetOrRegisterMeter("state/snapshot/dirty/storage/hit", nil)
	snapshotDirtyStorageMissMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/storage/miss", nil)
	snapshotDirtyStorageInexMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/storage/inex", nil)
	snapshotDirtyStorageReadMeter  = metrics.GetOrRegisterMeter("state/snapshot/dirty/storage/read", nil)
	snapshotDirtyStorageWriteMeter = metrics.GetOrRegisterMeter("state/snapshot/dirty/storage/write", nil)

	snapshotDirtyAccountHitDepthHist = metrics.GetOrRegisterHistogram("state/snapshot/dirty/account/hit/depth", nil, metrics.NewExpDecaySample(1028, 0.015))
	snapshotDirtyStorageHitDepthHist = metrics.GetOrRegisterHistogram("state/snapshot/dirty/storage/hit/depth", nil, metrics.NewExpDecaySample(1028, 0.015))

	snapshotFlushAccountItemMeter = metrics.GetOrRegisterMeter("state/snapshot/flush/account/item", nil)
	snapshotFlushAccountSizeMeter = metrics.GetOrRegisterMeter("state/snapshot/flush/account/size", nil)
	snapshotFlushStorageItemMeter = metrics.GetOrRegisterMeter("state/snapshot/flush/storage/item", nil)
	snapshotFlushStorageSizeMeter = metrics.GetOrRegisterMeter("state/snapshot/flush/storage/size", nil)

	snapshotBloomIndexTimer = metrics.GetOrRegisterResettingTimer("state/snapshot/bloom/index", nil)
	snapshotBloomErrorGauge = metrics.GetOrRegisterGaugeFloat64("state/snapshot/bloom/error", nil)

	snapshotBloomAccountTrueHitMeter  = metrics.GetOrRegisterMeter("state/snapshot/bloom/account/truehit", nil)
	snapshotBloomAccountFalseHitMeter = metrics.GetOrRegisterMeter("state/snapshot/bloom/account/falsehit", nil)
	snapshotBloomAccountMissMeter     = metrics.GetOrRegisterMeter("state/snapshot/bloom/account/miss", nil)

	snapshotBloomStorageTrueHitMeter  = metrics.GetOrRegisterMeter("state/snapshot/bloom/storage/truehit", nil)
	snapshotBloomStorageFalseHitMeter = metrics.GetOrRegisterMeter("state/snapshot/bloom/storage/falsehit", nil)
	snapshotBloomStorageMissMeter     = metrics.GetOrRegisterMeter("state/snapshot/bloom/storage/miss", nil)

	// ErrSnapshotStale is returned from data accessors if the underlying snapshot
	// layer had been invalidated due to the chain progressing forward far enough
	// to not maintain the layer's original state.
	ErrSnapshotStale = errors.New("snapshot stale")

	// ErrStaleParentLayer is returned when Flatten attempts to flatten a diff layer into
	// a stale parent.
	ErrStaleParentLayer = errors.New("parent disk layer is stale")

	// ErrNotCoveredYet is returned from data accessors if the underlying snapshot
	// is being generated currently and the requested data item is not yet in the
	// range of accounts covered.
	ErrNotCoveredYet = errors.New("not covered yet")

	// ErrNotConstructed is returned if the callers want to iterate the snapshot
	// while the generation is not finished yet.
	ErrNotConstructed = errors.New("snapshot is not constructed")
)

// Snapshot represents the functionality supported by a snapshot storage layer.
type Snapshot = ethsnapshot.Snapshot

// snapshot is the internal version of the snapshot data layer that supports some
// additional methods compared to the public API.
type snapshot interface {
	Snapshot

	BlockHash() common.Hash

	// Parent returns the subsequent layer of a snapshot, or nil if the base was
	// reached.
	//
	// Note, the method is an internal helper to avoid type switching between the
	// disk and diff layers. There is no locking involved.
	Parent() snapshot

	// Update creates a new layer on top of the existing snapshot diff tree with
	// the specified data items.
	//
	// Note, the maps are retained by the method to avoid copying everything.
	Update(blockHash, blockRoot common.Hash, destructs map[common.Hash]struct{}, accounts map[common.Hash][]byte, storage map[common.Hash]map[common.Hash][]byte) *diffLayer

	// Stale return whether this layer has become stale (was flattened across) or
	// if it's still live.
	Stale() bool

	// AccountIterator creates an account iterator over an arbitrary layer.
	AccountIterator(seek common.Hash) AccountIterator

	// StorageIterator creates a storage iterator over an arbitrary layer.
	StorageIterator(account common.Hash, seek common.Hash) (StorageIterator, bool)
}

// Config includes the configurations for snapshots.
type Config struct {
	CacheSize  int  // Megabytes permitted to use for read caches
	NoBuild    bool // Indicator that the snapshots generation is disallowed
	AsyncBuild bool // The snapshot generation is allowed to be constructed asynchronously
	SkipVerify bool // Indicator that all verification should be bypassed
}

// Tree is an Ethereum state snapshot tree. It consists of one persistent base
// layer backed by a key-value store, on top of which arbitrarily many in-memory
// diff layers are topped. The memory diffs can form a tree with branching, but
// the disk layer is singleton and common to all. If a reorg goes deeper than the
// disk layer, everything needs to be deleted.
//
// The goal of a state snapshot is twofold: to allow direct access to account and
// storage data to avoid expensive multi-level trie lookups; and to allow sorted,
// cheap iteration of the account/storage tries for sync aid.
type Tree struct {
	config Config              // Snapshots configurations
	diskdb ethdb.KeyValueStore // Persistent database to store the snapshot
	triedb *triedb.Database    // In-memory cache to access the trie through
	// Collection of all known layers
	// blockHash -> snapshot
	blockLayers map[common.Hash]snapshot
	// stateRoot -> blockHash -> snapshot
	// Update creates a new block layer with a parent taken from the blockHash -> snapshot map
	// we can support grabbing a read only Snapshot by getting any one from the state root based map
	stateLayers map[common.Hash]map[common.Hash]snapshot
	verified    bool // Indicates if snapshot integrity has been verified
	lock        sync.RWMutex

	// Test hooks
	onFlatten func() // Hook invoked when the bottom most diff layers are flattened
}

// New attempts to load an already existing snapshot from a persistent key-value
// store (with a number of memory layers from a journal), ensuring that the head
// of the snapshot matches the expected one.
//
// If the snapshot is missing or the disk layer is broken, the snapshot will be
// reconstructed using both the existing data and the state trie.
// The repair happens on a background thread.
func New(config Config, diskdb ethdb.KeyValueStore, triedb *triedb.Database, blockHash, root common.Hash) (*Tree, error) {
	_ = "STUB: not implemented"
	// Create a new, empty snapshot tree
	return nil, nil
}

// if SkipVerify is true, all verification will be bypassed

// Attempt to load a previously persisted snapshot and rebuild one if failed

// Bail out the error, don't rebuild automatically.

// Existing snapshot loaded, seed all the layers
// It is unnecessary to grab the lock here, since it was created within this function
// call, but we grab it nevertheless to follow the spec for insertSnap.

// Verify any synchronously generated or loaded snapshot from disk

// insertSnap inserts [snap] into the tree.
// Assumes the lock is held.
func (t *Tree) insertSnap(snap snapshot) { _ = "STUB: not implemented"; return }

// Snapshot retrieves a snapshot belonging to the given state root, or nil if no
// snapshot is maintained for that state root.
func (t *Tree) Snapshot(stateRoot common.Hash) Snapshot {
	_ = "STUB: not implemented"
	return *new(Snapshot)
}

// getSnapshot retrieves a Snapshot by its state root. If the caller already holds the
// snapTree lock when callthing this function, [holdsTreeLock] should be set to true.
func (t *Tree) getSnapshot(stateRoot common.Hash, holdsTreeLock bool) snapshot {
	_ = "STUB: not implemented"
	return *new(snapshot)
}

// Snapshots returns all visited layers from the topmost layer with specific
// root and traverses downward. The layer amount is limited by the given number.
// If nodisk is set, then disk layer is excluded.
func (t *Tree) Snapshots(blockHash common.Hash, limits int, nodisk bool) []Snapshot {
	_ = "STUB: not implemented"
	return nil
}

type blockHashes struct {
	blockHash       common.Hash
	parentBlockHash common.Hash
}

func WithBlockHashes(blockHash, parentBlockHash common.Hash) stateconf.SnapshotUpdateOption {
	_ = "STUB: not implemented"
	return *new(stateconf.SnapshotUpdateOption)
}

// Update adds a new snapshot into the tree, if that can be linked to an existing
// old parent. It is disallowed to insert a disk layer (the origin of all).
func (t *Tree) Update(
	blockRoot common.Hash,
	parentRoot common.Hash,
	destructs map[common.Hash]struct{},
	accounts map[common.Hash][]byte,
	storage map[common.Hash]map[common.Hash][]byte,
	opts ...stateconf.SnapshotUpdateOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) UpdateWithBlockHashes(
	blockHash, blockRoot, parentBlockHash common.Hash,
	destructs map[common.Hash]struct{},
	accounts map[common.Hash][]byte,
	storage map[common.Hash]map[common.Hash][]byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Grab the parent snapshot based on the parent block hash, not the parent state root

// verifyIntegrity performs an integrity check on the current snapshot using
// verify. Most importantly, verifyIntegrity ensures verify is called at
// most once during the entire lifetime of [Tree], returning immediately if
// already invoked. If [waitBuild] is true, verifyIntegrity will wait for
// generation of the snapshot to finish before verifying.
//
// It is assumed that the caller holds the [snapTree] lock
// when calling this function.
func (t *Tree) verifyIntegrity(base *diskLayer, waitBuild bool) error {
	_ = "STUB: not implemented"
	// Find the rebuild termination channel and wait until
	// the snapshot is generated
	return nil
}

func (t *Tree) Cap(root common.Hash, layers int) error {
	_ = "STUB: not implemented"
	// No-op as this code uses Flatten on block accept instead
	return nil
}

// Flatten flattens the snapshot for [blockHash] into its parent. if its
// parent is not a disk layer, Flatten will return an error.
// Note: a blockHash is used instead of a state root so that the exact state
// transition between the two states is well defined. This is intended to
// prevent the following edge case
//
//	  A
//	 /  \
//	B    C
//	     |
//	     D
//
// In this scenario, it's possible For (A, B) and (A, C, D) to be two
// different paths to the resulting state. We use block hashes and parent
// block hashes to ensure that the exact path through which we flatten
// diffLayers is well defined.
func (t *Tree) Flatten(blockHash common.Hash) error { _ = "STUB: not implemented"; return nil }

// Invoke the hook if it's registered. Ugly hack.

// Remove parent layer

// We created a new diskLayer [base] to replace [diff], so we need to replace
// it in both maps and replace all pointers to it.

// stateSnaps must already be initialized here, since we are replacing
// an existing snapshot instead of adding a new one.

// Replace the parent pointers for any snapshot that referenced
// the replaced diffLayer.

// TODO add tracking of children to the snapshots to reduce overhead here.

// If the disk layer was modified, regenerate all the cumulative blooms

// Length returns the number of snapshot layers that is currently being maintained.
func (t *Tree) NumStateLayers() int { _ = "STUB: not implemented"; return 0 }

func (t *Tree) NumBlockLayers() int { _ = "STUB: not implemented"; return 0 }

// Discard removes layers that we no longer need
func (t *Tree) Discard(blockHash common.Hash) error { _ = "STUB: not implemented"; return nil }

// discard removes the snapshot associated with [blockHash] from the
// snapshot tree.
// If [force] is true, discard may delete the disk layer. This should
// only be called within Flatten, when a new disk layer is being created.
// Assumes the lock is held.
func (t *Tree) discard(blockHash common.Hash, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Never discard the disk layer

// Discard the block from the map. If there are no more blocks
// mapping to the same state remove it from [stateLayers] as well.

// AbortGeneration aborts an ongoing snapshot generation process (if it hasn't
// stopped already).
//
// It is not required to manually abort snapshot generation. If generation has not
// been manually aborted prior to invoking [diffToDisk], it will be aborted anyways.
//
// It is safe to call this method multiple times and when there is no snapshot
// generation currently underway.
func (t *Tree) AbortGeneration() { _ = "STUB: not implemented"; return }

// diffToDisk merges a bottom-most diff into the persistent disk layer underneath
// it. The method will panic if called onto a non-bottom-most diff layer.
//
// The disk layer persistence should be operated in an atomic way. All updates should
// be discarded if the whole transition if not finished.
func diffToDisk(bottom *diffLayer) (*diskLayer, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Attempt to abort generation (if not already aborted)

// Put the deletion in the batch writer, flush all updates in the final step.

// Mark the original base as stale as we're going to create a new wrapper

// we've committed into the same base from two children, boo

// Destroy all the destructed accounts from the database

// Skip any account not covered yet by the snapshot

// Remove all storage slots

// Ensure we don't delete too much data blindly (contract can be
// huge). It's ok to flush, the root will go missing in case of a
// crash and we'll detect and regenerate the snapshot.

// Push all updated accounts into the database

// Skip any account not covered yet by the snapshot

// Push the account to disk

// Ensure we don't write too much data blindly. It's ok to flush, the
// root will go missing in case of a crash and we'll detect and regen
// the snapshot.

// Push all the storage slots into the database

// Skip any account not covered yet by the snapshot

// Generation might be mid-account, track that case too

// Skip any slot not covered yet by the snapshot

// Update the snapshot block marker and write any remainder data

// Write out the generator progress marker and report

// Flush all the updates in the single db operation. Ensure the
// disk layer transition is atomic.

// If snapshot generation hasn't finished yet, port over all the starts and
// continue where the previous round left off.
//
// Note, the `base.genPending` comparison is not used normally, it's checked
// to allow the tests to play with the marker without triggering this path.

// If the diskLayer we are about to discard is not very old, we skip
// generation on the next layer (assuming generation will just get canceled
// before doing meaningful work anyways).

// Release releases resources
func (t *Tree) Release() { _ = "STUB: not implemented"; return }

// Rebuild wipes all available snapshot data from the persistent database and
// discard all caches and diff layers. Afterwards, it starts a new snapshot
// generator with the given root hash.
func (t *Tree) Rebuild(blockHash, root common.Hash) { _ = "STUB: not implemented"; return }

// Track whether there's a wipe currently running and keep it alive if so

// Iterate over and mark all layers stale

// If the base layer is generating, abort it and save

// Layer should be inactive now, mark it as stale

// If the layer is a simple diff, simply mark as stale

// Start generating a new snapshot from scratch on a background thread. The
// generator will run a wiper first if there's not one running right now.

// AccountIterator creates a new account iterator for the specified root hash and
// seeks to a starting account hash. When [force] is true, a new account
// iterator is created without acquiring the [snapTree] lock and without
// confirming that the snapshot on the disk layer is fully generated.
func (t *Tree) AccountIterator(root common.Hash, seek common.Hash, force bool) (AccountIterator, error) {
	_ = "STUB: not implemented"
	return *new(AccountIterator), nil
}

// StorageIterator creates a new storage iterator for the specified root hash and
// account. The iterator will be move to the specific start position. When [force]
// is true, a new account iterator is created without acquiring the [snapTree]
// lock and without confirming that the snapshot on the disk layer is fully generated.
func (t *Tree) StorageIterator(root common.Hash, account common.Hash, seek common.Hash) (StorageIterator, error) {
	_ = "STUB: not implemented"
	return *new(StorageIterator), nil
}

func (t *Tree) StorageIteratorWithForce(root common.Hash, account common.Hash, seek common.Hash, force bool) (StorageIterator, error) {
	_ = "STUB: not implemented"
	return *new(StorageIterator), nil
}

// Verify iterates the whole state(all the accounts as well as the corresponding storages)
// with the specific root and compares the re-computed hash with the original one.
func (t *Tree) Verify(root common.Hash) error { _ = "STUB: not implemented"; return nil }

// verify iterates the whole state(all the accounts as well as the corresponding storages)
// with the specific root and compares the re-computed hash with the original one.
// When [force] is true, it is assumed that the caller has confirmed that the
// snapshot is generated and that they hold the snapTree lock.
func (t *Tree) verify(root common.Hash, force bool) error { _ = "STUB: not implemented"; return nil }

// disklayer is an internal helper function to return the disk layer.
// The lock of snapTree is assumed to be held already.
func (t *Tree) disklayer() *diskLayer { _ = "STUB: not implemented"; return nil }

// diskRoot is a internal helper function to return the disk layer root.
// The lock of snapTree is assumed to be held already.
func (t *Tree) diskRoot() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

// generating is an internal helper function which reports whether the snapshot
// is still under the construction.
func (t *Tree) generating() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// DiskRoot is an external helper function to return the disk layer root.
func (t *Tree) DiskRoot() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

// Size returns the memory usage of the diff layers above the disk layer and the
// dirty nodes buffered in the disk layer. Currently, the implementation uses a
// special diff layer (the first) as an aggregator simulating a dirty buffer, so
// the second return will always be 0. However, this will be made consistent with
// the pathdb, which will require a second return.
func (t *Tree) Size() (diffs common.StorageSize, buf common.StorageSize) {
	_ = "STUB: not implemented"
	return *new(common.StorageSize), *new(common.StorageSize)
}
