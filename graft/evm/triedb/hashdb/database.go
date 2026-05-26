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
// Copyright 2018 The go-ethereum Authors
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

package hashdb

import (
	"reflect"
	"sync"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/stateconf"
	"github.com/ava-labs/libevm/metrics"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/trie/triestate"
	"github.com/ava-labs/libevm/triedb"
	"github.com/ava-labs/libevm/triedb/database"

	// Force libevm metrics of the same name to be registered first.
	_ "github.com/ava-labs/libevm/triedb/hashdb"
)

const (
	cacheStatsUpdateFrequency = 1000 // update trie cache stats once per 1000 ops
)

// ====== If resolving merge conflicts ======
//
// All calls to metrics.NewRegistered*() for metrics also defined in libevm/triedb/hashdb
// have been replaced with metrics.GetOrRegister*() to get metrics already registered in
// libevm/triedb/hashdb or register them here otherwise. These replacements ensure the same
// metrics are shared between the two packages.
var (
	memcacheCleanHitMeter   = metrics.GetOrRegisterMeter("hashdb/memcache/clean/hit", nil)
	memcacheCleanMissMeter  = metrics.GetOrRegisterMeter("hashdb/memcache/clean/miss", nil)
	memcacheCleanReadMeter  = metrics.GetOrRegisterMeter("hashdb/memcache/clean/read", nil)
	memcacheCleanWriteMeter = metrics.GetOrRegisterMeter("hashdb/memcache/clean/write", nil)

	memcacheDirtyHitMeter   = metrics.GetOrRegisterMeter("hashdb/memcache/dirty/hit", nil)
	memcacheDirtyMissMeter  = metrics.GetOrRegisterMeter("hashdb/memcache/dirty/miss", nil)
	memcacheDirtyReadMeter  = metrics.GetOrRegisterMeter("hashdb/memcache/dirty/read", nil)
	memcacheDirtyWriteMeter = metrics.GetOrRegisterMeter("hashdb/memcache/dirty/write", nil)

	memcacheDirtySizeGauge      = metrics.GetOrRegisterGaugeFloat64("hashdb/memcache/dirty/size", nil)
	memcacheDirtyChildSizeGauge = metrics.GetOrRegisterGaugeFloat64("hashdb/memcache/dirty/childsize", nil)
	memcacheDirtyNodesGauge     = metrics.GetOrRegisterGauge("hashdb/memcache/dirty/nodes", nil)

	memcacheFlushMeter         = metrics.GetOrRegisterMeter("hashdb/memcache/flush/count", nil)
	memcacheFlushTimeTimer     = metrics.GetOrRegisterResettingTimer("hashdb/memcache/flush/time", nil)
	memcacheFlushLockTimeTimer = metrics.GetOrRegisterResettingTimer("hashdb/memcache/flush/locktime", nil)
	memcacheFlushNodesMeter    = metrics.GetOrRegisterMeter("hashdb/memcache/flush/nodes", nil)
	memcacheFlushBytesMeter    = metrics.GetOrRegisterMeter("hashdb/memcache/flush/bytes", nil)

	memcacheGCTimeTimer  = metrics.GetOrRegisterResettingTimer("hashdb/memcache/gc/time", nil)
	memcacheGCNodesMeter = metrics.GetOrRegisterMeter("hashdb/memcache/gc/nodes", nil)
	memcacheGCBytesMeter = metrics.GetOrRegisterMeter("hashdb/memcache/gc/bytes", nil)

	memcacheCommitMeter         = metrics.GetOrRegisterMeter("hashdb/memcache/commit/count", nil)
	memcacheCommitTimeTimer     = metrics.GetOrRegisterResettingTimer("hashdb/memcache/commit/time", nil)
	memcacheCommitLockTimeTimer = metrics.GetOrRegisterResettingTimer("hashdb/memcache/commit/locktime", nil)
	memcacheCommitNodesMeter    = metrics.GetOrRegisterMeter("hashdb/memcache/commit/nodes", nil)
	memcacheCommitBytesMeter    = metrics.GetOrRegisterMeter("hashdb/memcache/commit/bytes", nil)
)

// ChildResolver defines the required method to decode the provided
// trie node and iterate the children on top.
type ChildResolver interface {
	ForEach(node []byte, onChild func(common.Hash))
}

type cache interface {
	HasGet([]byte, []byte) ([]byte, bool)
	Del([]byte)
	Set([]byte, []byte)
	Reset()
	SaveToFileConcurrent(string, int) error
}

// Config contains the settings for database.
type Config struct {
	CleanCacheSize                  int    // Maximum memory allowance (in bytes) for caching clean nodes
	StatsPrefix                     string // Prefix for cache stats (disabled if empty)
	ReferenceRootAtomicallyOnUpdate bool   // Whether to reference the root node on update
}

func (c Config) BackendConstructor(diskdb ethdb.Database) triedb.DBOverride {
	_ = "STUB: not implemented"
	return *new(triedb.DBOverride)
}

// Defaults is the default setting for database if it's not specified.
// Notably, clean cache is disabled explicitly,
var Defaults = &Config{
	// Explicitly set clean cache size to 0 to avoid creating fastcache,
	// otherwise database must be closed when it's no longer needed to
	// prevent memory leak.
	CleanCacheSize: 0,
}

// Database is an intermediate write layer between the trie data structures and
// the disk database. The aim is to accumulate trie writes in-memory and only
// periodically flush a couple tries to disk, garbage collecting the remainder.
//
// The trie Database is thread-safe in its mutations and is thread-safe in providing individual,
// independent node access.
type Database struct {
	diskdb   ethdb.Database // Persistent storage for matured trie nodes
	resolver ChildResolver  // The handler to resolve children of nodes

	cleans  cache                       // GC friendly memory cache of clean node RLPs
	dirties map[common.Hash]*cachedNode // Data and references relationships of dirty trie nodes
	oldest  common.Hash                 // Oldest tracked node, flush-list head
	newest  common.Hash                 // Newest tracked node, flush-list tail

	gctime  time.Duration      // Time spent on garbage collection since last commit
	gcnodes uint64             // Nodes garbage collected since last commit
	gcsize  common.StorageSize // Data storage garbage collected since last commit

	flushtime  time.Duration      // Time spent on data flushing since last commit
	flushnodes uint64             // Nodes flushed since last commit
	flushsize  common.StorageSize // Data storage flushed since last commit

	dirtiesSize  common.StorageSize // Storage size of the dirty node cache (exc. metadata)
	childrenSize common.StorageSize // Storage size of the external children tracking

	lock sync.RWMutex

	referenceRoot bool
}

// cachedNode is all the information we know about a single cached trie node
// in the memory database write layer.
type cachedNode struct {
	node      []byte                   // Encoded node blob, immutable
	parents   uint32                   // Number of live nodes referencing this one
	external  map[common.Hash]struct{} // The set of external children
	flushPrev common.Hash              // Previous node in the flush-list
	flushNext common.Hash              // Next node in the flush-list
}

// cachedNodeSize is the raw size of a cachedNode data structure without any
// node data included. It's an approximate size, but should be a lot better
// than not counting them.
var cachedNodeSize = int(reflect.TypeOf(cachedNode{}).Size())

// forChildren invokes the callback for all the tracked children of this node,
// both the implicit ones from inside the node as well as the explicit ones
// from outside the node.
func (n *cachedNode) forChildren(resolver ChildResolver, onChild func(hash common.Hash)) {
	_ = "STUB: not implemented"
	return
}

// New initializes the hash-based node database.
func New(diskdb ethdb.Database, config *Config, resolver ChildResolver) *Database {
	_ = "STUB: not implemented"
	return nil
}

// insert inserts a trie node into the memory database. All nodes inserted by
// this function will be reference tracked. This function assumes the lock is
// already held.
func (db *Database) insert(hash common.Hash, node []byte) {
	_ = "STUB: not implemented"
	// If the node's already cached, skip
	return
}

// Create the cached entry for this node

// Update the flush-list endpoints

// node retrieves an encoded cached trie node from memory. If it cannot be found
// cached, the method queries the persistent database for the content.
func (db *Database) node(hash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	// It doesn't make sense to retrieve the metaroot
	return nil, nil
}

// Retrieve the node from the clean cache if available

// Delete anything from cache that may have been added incorrectly
//
// This will prevent a panic as callers of this function assume the raw
// or cached node is populated.

// Retrieve the node from the dirty cache if available.

// Return the cached node if it's found in the dirty set.
// The dirty.node field is immutable and safe to read it
// even without lock guard.

// Content unavailable in memory, attempt to retrieve from disk

// Reference adds a new reference from a parent node to a child node.
// This function is used to add reference between internal trie node
// and external node(e.g. storage trie root), all internal trie nodes
// are referenced together by database itself.
func (db *Database) Reference(child common.Hash, parent common.Hash) {
	_ = "STUB: not implemented"
	return
}

// reference is the private locked version of Reference.
func (db *Database) reference(child common.Hash, parent common.Hash) {
	_ = "STUB: not implemented"
	// If the node does not exist, it's a node pulled from disk, skip
	return
}

// The reference is for state root, increase the reference counter.

// The reference is for external storage trie, don't duplicate if
// the reference is already existent.

// Dereference removes an existing reference from a root node.
func (db *Database) Dereference(root common.Hash) {
	_ = "STUB: not implemented"
	// Sanity check to ensure that the meta-root is not removed
	return
}

// dereference is the private locked version of Dereference.
func (db *Database) dereference(hash common.Hash) {
	_ = "STUB: not implemented"
	// If the node does not exist, it's a previously committed node.
	return
}

// If there are no more references to the node, delete it and cascade

// This is a special cornercase where a node loaded from disk (i.e. not in the
// memcache any more) gets reinjected as a new node (short node split into full,
// then reverted into short), causing a cached node to have no parents. That is
// no problem in itself, but don't make maxint parents out of it.

// Remove the node from the flush-list

// Dereference all children and delete the node

// flushItem is used to track all [cachedNode]s that must be written to disk
type flushItem struct {
	hash common.Hash
	node *cachedNode
	rlp  []byte
}

// writeFlushItems writes all items in [toFlush] to disk in batches of
// [ethdb.IdealBatchSize]. This function does not access any variables inside
// of [Database] and does not need to be synchronized.
func (db *Database) writeFlushItems(toFlush []*flushItem) error {
	_ = "STUB: not implemented"
	return nil
}

// If we exceeded the ideal batch size, commit and reset

// Flush out any remainder data from the last batch

// Cap iteratively flushes old but still referenced trie nodes until the total
// memory usage goes below the given threshold.
func (db *Database) Cap(limit common.StorageSize) error {
	_ = "STUB: not implemented"

	// It is important that outside code doesn't see an inconsistent state
	// (referenced data removed from memory cache during commit but not yet
	// in persistent storage). This is ensured by only uncaching existing
	// data when the database write finalizes.
	return nil
}

// db.dirtiesSize only contains the useful data in the cache, but when reporting
// the total memory consumption, the maintenance metadata is also needed to be
// counted.

// Keep removing nodes from the flush-list until we're below allowance

// Fetch the oldest referenced node and push into the batch

// Iterate to the next flush item, or abort if the size cap was achieved. Size
// is the total size, including the useful cached data (hash -> blob), the
// cache item metadata, as well as external children mappings.

// Write nodes to disk

// Flush all written items from dirites
//
// NOTE: The order of the flushlist may have changed while the lock was not
// held, so we cannot just iterate to [oldest].

// [item.rlp] is populated in [writeFlushItems]

// Commit iterates over all the children of a particular node, writes them out
// to disk, forcefully tearing down all references in both directions. As a side
// effect, all pre-images accumulated up to this point are also written.
func (db *Database) Commit(node common.Hash, report bool) error {
	_ = "STUB: not implemented"
	return nil

	// It is important that outside code doesn't see an inconsistent state (referenced
	// data removed from memory cache during commit but not yet in persistent storage).
	// This is ensured by only uncaching existing data when the database write finalizes.
}

// Write nodes to disk

// Flush all written items from dirites

// [item.rlp] is populated in [writeFlushItems]

// Reset the garbage collection statistics

// commit is the private locked version of Commit. This function does not
// mutate any data, rather it collects all data that should be committed.
func (db *Database) commit(hash common.Hash, toFlush []*flushItem) ([]*flushItem, error) {
	_ = "STUB: not implemented"
	// If the node does not exist, it's a previously committed node
	return nil, nil
}

// By processing the children of each node before the node itself, we ensure
// that children are committed before their parents (an invariant of this
// package).

// removeFromDirties is invoked after database writes and implements dirty data uncaching.
//
// This is the post-processing step of a commit operation where the already persisted trie is
// removed from the dirty cache and moved into the clean cache. The reason behind
// the two-phase commit is to ensure data availability while moving from memory
// to disk.
//
// It is assumed the caller holds the [dirtiesLock] when this function is
// called.
func (db *Database) removeFromDirties(hash common.Hash, rlp []byte) {
	_ = "STUB: not implemented"
	// If the node does not exist, we're done on this path. This could happen if
	// nodes are capped to disk while another thread is committing those same
	// nodes.
	return
}

// Node still exists, remove it from the flush-list

// Remove the node from the dirty cache

// Move the flushed node into the clean cache to prevent insta-reloads

// Initialized returns an indicator if state data is already initialized
// in hash-based scheme by checking the presence of genesis state.
func (db *Database) Initialized(genesisRoot common.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

// Update inserts the dirty nodes in provided nodeset into database and link the
// account trie with multiple storage tries if necessary.
// If ReferenceRootAtomicallyOnUpdate was enabled in the config, it will also add a reference from
// the root to the metaroot while holding the db's lock.
func (db *Database) Update(root common.Hash, parent common.Hash, block uint64, nodes *trienode.MergedNodeSet, states *triestate.Set, _ ...stateconf.TrieDBUpdateOption) error {
	_ = "STUB: not implemented"
	// Ensure the parent state is present and signal a warning if not.
	return nil
}

func (db *Database) update(root common.Hash, parent common.Hash, nodes *trienode.MergedNodeSet) error {
	_ = "STUB: not implemented"
	// Insert dirty nodes into the database. In the same tree, it must be
	// ensured that children are inserted first, then parent so that children
	// can be linked with their parent correctly.
	//
	// Note, the storage tries must be flushed before the account trie to
	// retain the invariant that children go into the dirty cache first.
	return nil
}

// ignore deletion

// Link up the account trie and storage trie if the node points
// to an account trie leaf.

// Size returns the current storage size of the memory cache in front of the
// persistent database layer.
//
// The first return will always be 0, representing the memory stored in unbounded
// diff layers above the dirty cache. This is only available in pathdb.
func (db *Database) Size() (common.StorageSize, common.StorageSize) {
	_ = "STUB: not implemented"
	return *new(common.StorageSize), *new(common.StorageSize)
}

// db.dirtiesSize only contains the useful data in the cache, but when reporting
// the total memory consumption, the maintenance metadata is also needed to be
// counted.

// Close closes the trie database and releases all held resources.
func (db *Database) Close() error { _ = "STUB: not implemented"; return nil }

// Scheme returns the node scheme used in the database.
func (db *Database) Scheme() string { _ = "STUB: not implemented"; return "" }

// Reader retrieves a node reader belonging to the given state root.
// An error will be returned if the requested state is not available.
func (db *Database) Reader(root common.Hash) (database.Reader, error) {
	_ = "STUB: not implemented"
	return *new(database.Reader), nil
}

// reader is a state reader of Database which implements the Reader interface.
type reader struct {
	db *Database
}

// Node retrieves the trie node with the given node hash. No error will be
// returned if the node is not found.
func (reader *reader) Node(owner common.Hash, path []byte, hash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
