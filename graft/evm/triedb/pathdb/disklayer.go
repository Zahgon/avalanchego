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
// Copyright 2022 The go-ethereum Authors
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

package pathdb

import (
	"sync"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/trie/triestate"
	"golang.org/x/crypto/sha3"
)

// diskLayer is a low level persistent layer built on top of a key-value store.
type diskLayer struct {
	root   common.Hash      // Immutable, root hash to which this layer was made for
	id     uint64           // Immutable, corresponding state id
	db     *Database        // Path-based trie database
	cleans *fastcache.Cache // GC friendly memory cache of clean node RLPs
	buffer *nodebuffer      // Node buffer to aggregate writes
	stale  bool             // Signals that the layer became stale (state progressed)
	lock   sync.RWMutex     // Lock used to protect stale flag
}

// newDiskLayer creates a new disk layer based on the passing arguments.
func newDiskLayer(root common.Hash, id uint64, db *Database, cleans *fastcache.Cache, buffer *nodebuffer) *diskLayer {
	_ = "STUB: not implemented"
	// Initialize a clean cache if the memory allowance is not zero
	// or reuse the provided cache if it is not nil (inherited from
	// the original disk layer).
	return nil
}

// root implements the layer interface, returning root hash of corresponding state.
func (dl *diskLayer) rootHash() common.Hash {
	_ = "STUB: not implemented"

	// stateID implements the layer interface, returning the state id of disk layer.
	return *new(common.Hash)
}

func (dl *diskLayer) stateID() uint64 {
	_ = "STUB: not implemented"

	// parent implements the layer interface, returning nil as there's no layer
	// below the disk.
	return 0
}

func (dl *diskLayer) parentLayer() layer {
	_ = "STUB: not implemented"

	// isStale return whether this layer has become stale (was flattened across) or if
	// it's still live.
	return *new(layer)
}

func (dl *diskLayer) isStale() bool { _ = "STUB: not implemented"; return false }

// markStale sets the stale flag as true.
func (dl *diskLayer) markStale() { _ = "STUB: not implemented"; return }

// we've committed into the same base from two children, boom

// Node implements the layer interface, retrieving the trie node with the
// provided node info. No error will be returned if the node is not found.
func (dl *diskLayer) Node(owner common.Hash, path []byte, hash common.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to retrieve the trie node from the not-yet-written
// node buffer first. Note the buffer is lock free since
// it's impossible to mutate the buffer before tagging the
// layer as stale.

// Try to retrieve the trie node from the clean memory cache

// Try to retrieve the trie node from the disk.

// update implements the layer interface, returning a new diff layer on top
// with the given state set.
func (dl *diskLayer) update(root common.Hash, id uint64, block uint64, nodes map[common.Hash]map[string]*trienode.Node, states *triestate.Set) *diffLayer {
	_ = "STUB: not implemented"
	return nil
}

// commit merges the given bottom-most diff layer into the node buffer
// and returns a newly constructed disk layer. Note the current disk
// layer must be tagged as stale first to prevent re-access.
func (dl *diskLayer) commit(bottom *diffLayer, force bool) (*diskLayer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Construct and store the state history first. If crash happens after storing
// the state history but without flushing the corresponding states(journal),
// the stored state history will be truncated from head in the next restart.

// NOTE(freezer): This is disabled since we do not have a freezer.
// if dl.db.freezer != nil {
// 	err := writeHistory(dl.db.freezer, bottom)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Determine if the persisted history object has exceeded the configured
// 	// limitation, set the overflow as true if so.
// 	tail, err := dl.db.freezer.Tail()
// 	if err != nil {
// 		return nil, err
// 	}
// 	limit := dl.db.config.StateHistory
// 	if limit != 0 && bottom.stateID()-tail > limit {
// 		overflow = true
// 		oldest = bottom.stateID() - limit + 1 // track the id of history **after truncation**
// 	}
// }
// Mark the diskLayer as stale before applying any mutations on top.

// Store the root->id lookup afterwards. All stored lookups are identified
// by the **unique** state root. It's impossible that in the same chain
// blocks are not adjacent but have the same root.

// Construct a new disk layer by merging the nodes from the provided diff
// layer, and flush the content in disk layer if there are too many nodes
// cached. The clean cache is inherited from the original disk layer.

// In a unique scenario where the ID of the oldest history object (after tail
// truncation) surpasses the persisted state ID, we take the necessary action
// of forcibly committing the cached dirty nodes to ensure that the persisted
// state ID remains higher.

// To remove outdated history objects from the end, we set the 'tail' parameter
// to 'oldest-1' due to the offset between the freezer index and the history ID.

// NOTE(freezer): This is disabled since we do not have a freezer.
// 	pruned, err := truncateFromTail(ndl.db.diskdb, ndl.db.freezer, oldest-1)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Debug("Pruned state history", "items", pruned, "tailid", oldest)

// nolint: unused
// revert applies the given state history and return a reverted disk layer.
func (dl *diskLayer) revert(h *history, loader triestate.TrieLoader) (*diskLayer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject if the provided state history is incomplete. It's due to
// a large construct SELF-DESTRUCT which can't be handled because
// of memory limitation.

// Apply the reverse state changes upon the current state. This must
// be done before holding the lock in order to access state in "this"
// layer.

// Mark the diskLayer as stale before applying any mutations on top.

// State change may be applied to node buffer, or the persistent
// state, depends on if node buffer is empty or not. If the node
// buffer is not empty, it means that the state transition that
// needs to be reverted is not yet flushed and cached in node
// buffer, otherwise, manipulate persistent state directly.

// setBufferSize sets the node buffer size to the provided value.
func (dl *diskLayer) setBufferSize(size int) error { _ = "STUB: not implemented"; return nil }

// size returns the approximate size of cached nodes in the disk layer.
func (dl *diskLayer) size() common.StorageSize {
	_ = "STUB: not implemented"
	return *new(common.StorageSize)
}

// resetCache releases the memory held by clean cache to prevent memory leak.
func (dl *diskLayer) resetCache() { _ = "STUB: not implemented"; return }

// Stale disk layer loses the ownership of clean cache.

// hasher is used to compute the sha256 hash of the provided data.
type hasher struct{ sha crypto.KeccakState }

var hasherPool = sync.Pool{
	New: func() interface{} { return &hasher{sha: sha3.NewLegacyKeccak256().(crypto.KeccakState)} },
}

func newHasher() *hasher { _ = "STUB: not implemented"; return nil }

func (h *hasher) hash(data []byte) common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

func (h *hasher) release() { _ = "STUB: not implemented"; return }
