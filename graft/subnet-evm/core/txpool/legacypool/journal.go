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
// Copyright 2017 The go-ethereum Authors
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

package legacypool

import (
	"errors"
	"io"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// errNoActiveJournal is returned if a transaction is attempted to be inserted
// into the journal, but no such file is currently open.
var errNoActiveJournal = errors.New("no active journal")

// devNull is a WriteCloser that just discards anything written into it. Its
// goal is to allow the transaction journal to write into a fake journal when
// loading transactions on startup without printing warnings due to no file
// being read for write.
type devNull struct{}

func (*devNull) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
func (*devNull) Close() error                      { _ = "STUB: not implemented"; return nil }

// journal is a rotating log of transactions with the aim of storing locally
// created transactions to allow non-executed ones to survive node restarts.
type journal struct {
	path   string         // Filesystem path to store the transactions at
	writer io.WriteCloser // Output stream to write new transactions into
}

// newTxJournal creates a new transaction journal to
func newTxJournal(path string) *journal { _ = "STUB: not implemented"; return nil }

// load parses a transaction journal dump from disk, loading its contents into
// the specified pool.
func (journal *journal) load(add func([]*types.Transaction) []error) error {
	_ = "STUB: not implemented"
	// Open the journal for loading any past transactions
	return nil
}

// Skip the parsing if the journal file doesn't exist at all

// Temporarily discard any journal additions (don't double add on load)

// Inject all transactions from the journal into the pool

// Create a method to load a limited batch of transactions and bump the
// appropriate progress counters. Then use this method to load all the
// journaled transactions in small-ish batches.

// Parse the next transaction and terminate on error

// New transaction parsed, queue up for later, import if threshold is reached

// insert adds the specified transaction to the local disk journal.
func (journal *journal) insert(tx *types.Transaction) error { _ = "STUB: not implemented"; return nil }

// rotate regenerates the transaction journal based on the current contents of
// the transaction pool.
func (journal *journal) rotate(all map[common.Address]types.Transactions) error {
	_ = "STUB: not implemented"
	// Close the current journal (if any is open)
	return nil
}

// Generate a new journal with the contents of the current pool

// Replace the live journal with the newly generated one

// close flushes the transaction journal contents to disk and closes the file.
func (journal *journal) close() error { _ = "STUB: not implemented"; return nil }
