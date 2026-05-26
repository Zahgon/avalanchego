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

package snapshot

import (
	"github.com/ava-labs/libevm/common"
)

// binaryIterator is a simplistic iterator to step over the accounts or storage
// in a snapshot, which may or may not be composed of multiple layers. Performance
// wise this iterator is slow, it's meant for cross validating the fast one,
type binaryIterator struct {
	a               Iterator
	b               Iterator
	aDone           bool
	bDone           bool
	accountIterator bool
	k               common.Hash
	account         common.Hash
	fail            error
}

// initBinaryAccountIterator creates a simplistic iterator to step over all the
// accounts in a slow, but easily verifiable way. Note this function is used for
// initialization, use `newBinaryAccountIterator` as the API.
func (dl *diffLayer) initBinaryAccountIterator() Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// initBinaryStorageIterator creates a simplistic iterator to step over all the
// storage slots in a slow, but easily verifiable way. Note this function is used
// for initialization, use `newBinaryStorageIterator` as the API.
func (dl *diffLayer) initBinaryStorageIterator(account common.Hash) Iterator {
	_ = "STUB: not implemented"
	return *new(Iterator)
}

// If the storage in this layer is already destructed, discard all
// deeper layers but still return an valid single-branch iterator.

// The parent is disk layer, don't need to take care "destructed"
// anymore.

// If the storage in this layer is already destructed, discard all
// deeper layers but still return an valid single-branch iterator.

// Next steps the iterator forward one element, returning false if exhausted,
// or an error if iteration failed for some reason (e.g. root being iterated
// becomes stale and garbage collected).
func (it *binaryIterator) Next() bool { _ = "STUB: not implemented"; return false }

// Now we need to advance one of them

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
func (it *binaryIterator) Error() error {
	_ = "STUB: not implemented"

	// Hash returns the hash of the account the iterator is currently at.
	return nil
}

func (it *binaryIterator) Hash() common.Hash {
	_ = "STUB: not implemented"

	// Account returns the RLP encoded slim account the iterator is currently at, or
	// nil if the iterated snapshot stack became stale (you can check Error after
	// to see if it failed or not).
	//
	// Note the returned account is not a copy, please don't modify it.
	return *new(common.Hash)
}

func (it *binaryIterator) Account() []byte { _ = "STUB: not implemented"; return nil }

// The topmost iterator must be `diffAccountIterator`

// Slot returns the raw storage slot data the iterator is currently at, or
// nil if the iterated snapshot stack became stale (you can check Error after
// to see if it failed or not).
//
// Note the returned slot is not a copy, please don't modify it.
func (it *binaryIterator) Slot() []byte { _ = "STUB: not implemented"; return nil }

// Release recursively releases all the iterators in the stack.
func (it *binaryIterator) Release() { _ = "STUB: not implemented"; return }

// newBinaryAccountIterator creates a simplistic account iterator to step over
// all the accounts in a slow, but easily verifiable way.
func (dl *diffLayer) newBinaryAccountIterator() AccountIterator {
	_ = "STUB: not implemented"
	return *new(AccountIterator)
}

// newBinaryStorageIterator creates a simplistic account iterator to step over
// all the storage slots in a slow, but easily verifiable way.
func (dl *diffLayer) newBinaryStorageIterator(account common.Hash) StorageIterator {
	_ = "STUB: not implemented"
	return *new(StorageIterator)
}
