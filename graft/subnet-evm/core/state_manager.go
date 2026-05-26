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
// Copyright 2014 The go-ethereum Authors
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

package core

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// flushWindow is the distance to the [commitInterval] when we start
// optimistically flushing trie nodes to disk (only applicable in [pruning]
// mode).
//
// We perform this optimistic flushing to reduce synchronized database IO at the
// [commitInterval].
const flushWindow = 768

type TrieWriter interface {
	InsertTrie(block *types.Block) error // Handle inserted trie reference of [root]
	AcceptTrie(block *types.Block) error // Mark [root] as part of an accepted block
	RejectTrie(block *types.Block) error // Notify TrieWriter that the block containing [root] has been rejected
	Shutdown() error
}

type TrieDB interface {
	Dereference(root common.Hash) error
	Commit(root common.Hash, report bool) error
	Size() (common.StorageSize, common.StorageSize, common.StorageSize)
	Cap(limit common.StorageSize) error
}

func NewTrieWriter(db TrieDB, config *CacheConfig) TrieWriter {
	_ = "STUB: not implemented"
	return *new(TrieWriter)
}

// PathDB tracks recent all states implicitly, and should not ever be committed explicitly.
// Journaling happens at blockchain shutdown.

// Firewood tracks recent states via config, so tip buffer is not needed in pruning mode.

type noPruningTrieWriter struct {
	TrieDB
}

func (np *noPruningTrieWriter) InsertTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	// We don't attempt to [Cap] here because we should never have
	// a significant amount of [TrieDB.Dirties] (we commit each block).
	return nil
}

func (np *noPruningTrieWriter) AcceptTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	// We don't need to call [Dereference] on the block root at the end of this
	// function because it is removed from the [TrieDB.Dirties] map in [Commit].
	return nil
}

func (np *noPruningTrieWriter) RejectTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (np *noPruningTrieWriter) Shutdown() error { _ = "STUB: not implemented"; return nil }

type cappedMemoryTrieWriter struct {
	TrieDB
	memoryCap        common.StorageSize
	targetCommitSize common.StorageSize
	flushStepSize    common.StorageSize
	imageCap         common.StorageSize
	commitInterval   uint64

	tipBuffer *BoundedBuffer[common.Hash]
}

func (cm *cappedMemoryTrieWriter) InsertTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	// The use of [Cap] in [InsertTrie] prevents exceeding the configured memory
	// limit (and OOM) in case there is a large backlog of processing (unaccepted) blocks.
	return nil
}

// all memory is contained within the nodes return for hashdb

func (cm *cappedMemoryTrieWriter) AcceptTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil

	// Attempt to dereference roots at least [tipBufferSize] old (so queries at tip
	// can still be completed).
	//
	// Note: It is safe to dereference roots that have been committed to disk
	// (they are no-ops).
}

// Commit this root if we have reached the [commitInterval].

// Write at least [flushStepSize] of the oldest nodes in the trie database
// dirty cache to disk as we approach the [commitInterval] to reduce the number of trie nodes
// that will need to be written at once on [Commit] (to roughly [targetCommitSize]).
//
// To reduce the number of useless trie nodes that are committed during this
// capping, we only optimistically flush within the [flushWindow]. During
// this period, the [targetMemory] decreases stepwise by [flushStepSize]
// as we get closer to the commit boundary.
//
// Most trie nodes are 300B, so we will write at least ~1000 trie nodes in
// a single optimistic flush (with the default [flushStepSize]=312KB).
// this cannot be 0

func (cm *cappedMemoryTrieWriter) RejectTrie(block *types.Block) error {
	_ = "STUB: not implemented"
	return nil
}

func (cm *cappedMemoryTrieWriter) Shutdown() error {
	_ = "STUB: not implemented"
	// If [tipBuffer] entry is empty, no need to do any cleanup on
	// shutdown.
	return nil
}

// Attempt to commit last item added to [dereferenceQueue] on shutdown to avoid
// re-processing the state on the next startup.

type noopWriter struct{}

func (noopWriter) InsertTrie(*types.Block) error { _ = "STUB: not implemented"; return nil }

func (noopWriter) AcceptTrie(*types.Block) error { _ = "STUB: not implemented"; return nil }

func (noopWriter) RejectTrie(*types.Block) error { _ = "STUB: not implemented"; return nil }

func (noopWriter) Shutdown() error { _ = "STUB: not implemented"; return nil }
