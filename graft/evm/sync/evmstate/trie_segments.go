// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"context"
	"fmt"
	"sync"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/trie"

	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/leaf"
)

var (
	_ leaf.SyncTask = (*trieSegment)(nil)
	_ fmt.Stringer  = (*trieSegment)(nil)
)

// trieToSync keeps the state of a single trie syncing
// this can be a storage or the main trie.
type trieToSync struct {
	root    common.Hash
	account common.Hash

	// The trie consists of a slice of segments. each
	// segment has a start and end range of keys, and
	// contains a pointer back to this struct.
	segments []*trieSegment

	// These fields are used to hash the segments in
	// order, even though they may finish syncing out
	// of order or concurrently.
	lock              sync.Mutex
	segmentsDone      map[int]struct{}
	segmentToHashNext int

	// We use a stack trie to hash the leafs and have
	// a batch used for writing it to disk.
	batch     ethdb.Batch
	stackTrie *trie.StackTrie

	// We keep a pointer to the overall sync operation,
	// used to add segments to the work queue and to
	// update the eta.
	sync *stateSync

	// task implements the syncTask interface with methods
	// containing logic specific to the main trie or storage
	// tries.
	task       syncTask
	isMainTrie bool
}

// NewTrieToSync initializes a trieToSync and restores any previously started segments.
func NewTrieToSync(sync *stateSync, root common.Hash, account common.Hash, syncTask syncTask) (*trieToSync, error) {
	_ = "STUB: not implemented"
	return nil,
		// TODO: migrate state sync to use database schemes.
		nil
}

// loadSegments reads persistent storage and initializes trieSegments that
// had been previously started and need to be resumed.
func (t *trieToSync) loadSegments() error {
	_ = "STUB: not implemented"
	// Get an iterator for segments for t.root and see if we find anything.
	// This lets us check if this trie was previously segmented, in which
	// case we need to restore the same segments on resume.
	return nil
}

// Track the previously added segment as we loop over persisted values.

// If we find any persisted segments with the specified
// prefix, we add a new segment to the trie here.
// The segment we add represents a segment ending at the
// key immediately prior to the segment we found on disk.
// This is because we do not persist the beginning of
// the first segment.

// keep tracking the previous segment

// this creates the last segment if any were found in the loop
// and also handles the case where there were no segments persisted to disk.

// for each segment we need to find the last key already persisted
// so syncing can begin at the subsequent key

// don't go past the end of the segment

// syncing will start from this key

// startSyncing adds the trieToSync's segments to the work queue.
func (t *trieToSync) startSyncing(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// addSegment appends a newly created segment specified by [start] and
// [end] to [t.segments] and returns it.
// note: addSegment does not take a lock and therefore is called only
// before multiple segments are syncing concurrently.
func (t *trieToSync) addSegment(start, end []byte) *trieSegment {
	_ = "STUB: not implemented"
	return nil
}

// segmentFinished is called when one the trie segment with index [idx] finishes syncing.
// creates intermediary hash nodes for the trie up to the last contiguous segment received from start.
func (t *trieToSync) segmentFinished(ctx context.Context, idx int) error {
	_ = "STUB: not implemented"
	return nil
}

// if not the next contiguous segment from the beginning of the trie
// don't do anything.

// persist any items in the batch as they will be iterated below.

// reset the batch to free memory (even though it is no longer used)

// iterate all the items from the start of the segment (end is checked in the loop)

// don't go past the end of the segment. (data belongs to the next segment)

// update the stack trie and cap the batch it writes to.

// trie not complete

// when the trie is finished, this hashes any remaining nodes in the stack
// trie and creates the root

// the batch containing the main trie's root will be committed on
// sync completion.

// remove all segments for this root from persistent storage

// createSegmentsIfNeeded is called from the leaf handler. In case the trie syncing only has
// one segment but a large number of leafs ([t.estimateSize() > segmentThreshold], it will
// create [numSegments-1] additional segments to sync the trie.
func (t *trieToSync) createSegmentsIfNeeded(ctx context.Context, numSegments int) error {
	_ = "STUB: not implemented"
	return nil
}

// shouldSegment returns true if a trie should be separated into segments.
func (t *trieToSync) shouldSegment() bool { _ = "STUB: not implemented"; return false }

// Return false if the trie has already been segmented.

// Return true iff the estimated size of the trie exceeds [segmentThreshold].
// Note: at this point there is only a single segment (loadSegments guarantees there
// is at least one segment).

// divide the key space into [numSegments] consecutive segments.
// we use 2 bytes to build the ranges and fill the rest with
// ones or zeroes accordingly.
// this represents the step between the first 2 bytes of the start
// key of consecutive segments.
// createSegments should only be called once when there is only one
// thread accessing this trie, such that there is no need to hold a lock.
func (t *trieToSync) createSegments(ctx context.Context, numSegments int) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip any portion of the trie that has already been synced.

// since the first segment is already syncing,
// it does not need to be added to the task queue.
// instead, we update its end and move on to creating
// the next segment

// create the segments

// add the newly created segments to the task queue
// after creating them. We skip the first one, as it
// is already syncing.
// this avoids concurrent access to [t.segments].

// trieSegment keeps the state of syncing one segment of a [trieToSync]
// struct and keeps a pointer to the [trieToSync] it is syncing.
// each trieSegment is accessed by its own goroutine, so locks are not
// needed to access its fields
type trieSegment struct {
	start []byte
	pos   []byte
	end   []byte

	trie  *trieToSync // points back to the trie the segment belongs to
	idx   int         // index of this segment in the trie's segment slice
	batch ethdb.Batch // batch for writing leafs to
	leafs uint64      // number of leafs added to the segment
}

func (t *trieSegment) String() string { _ = "STUB: not implemented"; return "" }

// these functions implement the LeafSyncTask interface.
func (t *trieSegment) Root() common.Hash    { _ = "STUB: not implemented"; return *new(common.Hash) }
func (t *trieSegment) Account() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }
func (t *trieSegment) End() []byte          { _ = "STUB: not implemented"; return nil }
func (*trieSegment) NodeType() message.NodeType {
	_ = "STUB: not implemented"
	return *new(message.NodeType)
}
func (t *trieSegment) OnStart() (bool, error)             { _ = "STUB: not implemented"; return false, nil }
func (t *trieSegment) OnFinish(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *trieSegment) Start() []byte { _ = "STUB: not implemented"; return nil }

func (t *trieSegment) OnLeafs(ctx context.Context, keys, vals [][]byte) error {
	_ = "STUB: not implemented"
	// invoke the onLeafs callback
	return nil
}

// cap the segment's batch

// remember the position, used in estimating trie size

// update eta

// estimateSize calculates an estimate of the number of leafs and returns it,
// this assumes the trie has uniform key density.
// Note: returns 0 if there has been no progress in syncing the trie.
func (t *trieSegment) estimateSize() uint64 { _ = "STUB: not implemented"; return 0 }

// this should not occur since estimateSize is called after processing
// a batch of leafs, which sets [pos].
// avoid division by 0 out of caution.

// addPadding returns a []byte of length [common.Hash], starting with the BigEndian
// representation of [pos], and the rest filled with [padding].
func addPadding(pos uint16, padding byte) []byte { _ = "STUB: not implemented"; return nil }
