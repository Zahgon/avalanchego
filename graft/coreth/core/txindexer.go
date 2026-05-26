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
// Copyright 2024 The go-ethereum Authors
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
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>

package core

import (
	"github.com/ava-labs/libevm/ethdb"
)

// txIndexer is the module responsible for maintaining transaction indexes
// according to the configured indexing range by users.
type txIndexer struct {
	// limit is the maximum number of blocks from head whose tx indexes
	// are reserved:
	//  * 0: means the entire chain should be indexed
	//  * N: means the latest N blocks [HEAD-N+1, HEAD] should be indexed
	//       and all others shouldn't.
	limit  uint64
	db     ethdb.Database
	term   chan chan struct{}
	closed chan struct{}

	chain *BlockChain
}

// newTxIndexer initializes the transaction indexer.
func newTxIndexer(limit uint64, chain *BlockChain) *txIndexer {
	_ = "STUB: not implemented"
	return nil
}

// run executes the scheduled indexing/unindexing task in a separate thread.
// If the stop channel is closed, the task should be terminated as soon as
// possible, the done channel will be closed once the task is finished.
func (indexer *txIndexer) run(tail *uint64, head uint64, stop chan struct{}, done chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// Short circuit if chain is empty and nothing to index.

// Defensively ensure tail is not nil.

// use intermediate variable to avoid modifying the pointer

// Unindex a part of stale indices and forward index tail to HEAD-limit

// loop is the scheduler of the indexer, assigning indexing/unindexing tasks depending
// on the received chain event.
func (indexer *txIndexer) loop(chain *BlockChain) { _ = "STUB: not implemented"; return }

// Listening to chain events and manipulate the transaction indexes.

// Non-nil if background routine is active.
// Non-nil if background routine is active.
// The latest announced chain head (whose tx indexes are assumed created)
// The head number being processed in background.

// startRun launches the background unindexing task.

// Launch the initial processing if chain is not empty (head != genesis).
// This step is useful in these scenarios that chain has no progress.

// If no background task is running, start a new one.
// We cannot block on the subscription channel because it can
// cause a fatal error.

// If there is a new head arrived during the last run, start a new one.

// close shutdown the indexer. Safe to be called for multiple times.
func (indexer *txIndexer) close() { _ = "STUB: not implemented"; return }

// lockedRun runs the indexing/unindexing task in a locked manner. It reads
// the current tail index from the database.
func (indexer *txIndexer) lockedRun(head uint64, stop chan struct{}, done chan struct{}) {
	_ = "STUB: not implemented"
	return
}
