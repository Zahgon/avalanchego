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
	ethsnapshot "github.com/ava-labs/libevm/core/state/snapshot"
	"github.com/ava-labs/libevm/ethdb"
)

// Iterator is an iterator to step over all the accounts or the specific
// storage in a snapshot which may or may not be composed of multiple layers.
type Iterator = ethsnapshot.Iterator

// AccountIterator is an iterator to step over all the accounts in a snapshot,
// which may or may not be composed of multiple layers.
type AccountIterator = ethsnapshot.AccountIterator

// StorageIterator is an iterator to step over the specific storage in a snapshot,
// which may or may not be composed of multiple layers.
type StorageIterator = ethsnapshot.StorageIterator

// diffAccountIterator is an account iterator that steps over the accounts (both
// live and deleted) contained within a single diff layer. Higher order iterators
// will use the deleted accounts to skip deeper iterators.
type diffAccountIterator struct {
	// curHash is the current hash the iterator is positioned on. The field is
	// explicitly tracked since the referenced diff layer might go stale after
	// the iterator was positioned and we don't want to fail accessing the old
	// hash as long as the iterator is not touched any more.
	curHash common.Hash

	layer *diffLayer    // Live layer to retrieve values from
	keys  []common.Hash // Keys left in the layer to iterate
	fail  error         // Any failures encountered (stale)
}

// AccountIterator creates an account iterator over a single diff layer.
func (dl *diffLayer) AccountIterator(seek common.Hash) AccountIterator {
	_ = "STUB: not implemented"
	// Seek out the requested starting account
	return *new(AccountIterator)
}

// Assemble and returned the already seeked iterator

// Next steps the iterator forward one element, returning false if exhausted.
func (it *diffAccountIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator was already stale, consider it a programmer error. Although
	// we could just return false here, triggering this path would probably mean
	// somebody forgot to check for Error, so lets blow up instead of undefined
	// behavior that's hard to debug.
	return false
}

// Stop iterating if all keys were exhausted

// Iterator seems to be still alive, retrieve and cache the live hash

// key cached, shift the iterator and notify the user of success

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
func (it *diffAccountIterator) Error() error {
	_ = "STUB: not implemented"

	// Hash returns the hash of the account the iterator is currently at.
	return nil
}

func (it *diffAccountIterator) Hash() common.Hash {
	_ = "STUB: not implemented"

	// Account returns the RLP encoded slim account the iterator is currently at.
	// This method may _fail_, if the underlying layer has been flattened between
	// the call to Next and Account. That type of error will set it.Err.
	// This method assumes that flattening does not delete elements from
	// the accountdata mapping (writing nil into it is fine though), and will panic
	// if elements have been deleted.
	//
	// Note the returned account is not a copy, please don't modify it.
	return *new(common.Hash)
}

func (it *diffAccountIterator) Account() []byte { _ = "STUB: not implemented"; return nil }

// Release is a noop for diff account iterators as there are no held resources.
func (it *diffAccountIterator) Release() {
	_ = "STUB: not implemented"

	// diskAccountIterator is an account iterator that steps over the live accounts
	// contained within a disk layer.
	return
}

type diskAccountIterator struct {
	layer *diskLayer
	it    ethdb.Iterator
}

// AccountIterator creates an account iterator over a disk layer.
func (dl *diskLayer) AccountIterator(seek common.Hash) AccountIterator {
	_ = "STUB: not implemented"
	return *new(AccountIterator)
}

// Next steps the iterator forward one element, returning false if exhausted.
func (it *diskAccountIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator was already exhausted, don't bother
	return false
}

// Try to advance the iterator and release it if we reached the end

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
//
// A diff layer is immutable after creation content wise and can always be fully
// iterated without error, so this method always returns nil.
func (it *diskAccountIterator) Error() error { _ = "STUB: not implemented"; return nil }

// Iterator is exhausted and released

// Hash returns the hash of the account the iterator is currently at.
func (it *diskAccountIterator) Hash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// The prefix will be truncated

// Account returns the RLP encoded slim account the iterator is currently at.
func (it *diskAccountIterator) Account() []byte { _ = "STUB: not implemented"; return nil }

// Release releases the database snapshot held during iteration.
func (it *diskAccountIterator) Release() {
	_ = "STUB: not implemented"
	// The iterator is auto-released on exhaustion, so make sure it's still alive
	return
}

// diffStorageIterator is a storage iterator that steps over the specific storage
// (both live and deleted) contained within a single diff layer. Higher order
// iterators will use the deleted slot to skip deeper iterators.
type diffStorageIterator struct {
	// curHash is the current hash the iterator is positioned on. The field is
	// explicitly tracked since the referenced diff layer might go stale after
	// the iterator was positioned and we don't want to fail accessing the old
	// hash as long as the iterator is not touched any more.
	curHash common.Hash
	account common.Hash

	layer *diffLayer    // Live layer to retrieve values from
	keys  []common.Hash // Keys left in the layer to iterate
	fail  error         // Any failures encountered (stale)
}

// StorageIterator creates a storage iterator over a single diff layer.
// Except the storage iterator is returned, there is an additional flag
// "destructed" returned. If it's true then it means the whole storage is
// destructed in this layer(maybe recreated too), don't bother deeper layer
// for storage retrieval.
func (dl *diffLayer) StorageIterator(account common.Hash, seek common.Hash) (StorageIterator, bool) {
	_ = "STUB: not implemented"
	// Create the storage for this account even it's marked
	// as destructed. The iterator is for the new one which
	// just has the same address as the deleted one.
	return *new(StorageIterator), false
}

// Assemble and returned the already seeked iterator

// Next steps the iterator forward one element, returning false if exhausted.
func (it *diffStorageIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator was already stale, consider it a programmer error. Although
	// we could just return false here, triggering this path would probably mean
	// somebody forgot to check for Error, so lets blow up instead of undefined
	// behavior that's hard to debug.
	return false
}

// Stop iterating if all keys were exhausted

// Iterator seems to be still alive, retrieve and cache the live hash

// key cached, shift the iterator and notify the user of success

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
func (it *diffStorageIterator) Error() error {
	_ = "STUB: not implemented"

	// Hash returns the hash of the storage slot the iterator is currently at.
	return nil
}

func (it *diffStorageIterator) Hash() common.Hash {
	_ = "STUB: not implemented"

	// Slot returns the raw storage slot value the iterator is currently at.
	// This method may _fail_, if the underlying layer has been flattened between
	// the call to Next and Value. That type of error will set it.Err.
	// This method assumes that flattening does not delete elements from
	// the storage mapping (writing nil into it is fine though), and will panic
	// if elements have been deleted.
	//
	// Note the returned slot is not a copy, please don't modify it.
	return *new(common.Hash)
}

func (it *diffStorageIterator) Slot() []byte { _ = "STUB: not implemented"; return nil }

// Storage slot might be nil(deleted), but it must exist

// Release is a noop for diff account iterators as there are no held resources.
func (it *diffStorageIterator) Release() {
	_ = "STUB: not implemented"

	// diskStorageIterator is a storage iterator that steps over the live storage
	// contained within a disk layer.
	return
}

type diskStorageIterator struct {
	layer   *diskLayer
	account common.Hash
	it      ethdb.Iterator
}

// StorageIterator creates a storage iterator over a disk layer.
// If the whole storage is destructed, then all entries in the disk
// layer are deleted already. So the "destructed" flag returned here
// is always false.
func (dl *diskLayer) StorageIterator(account common.Hash, seek common.Hash) (StorageIterator, bool) {
	_ = "STUB: not implemented"
	return *new(StorageIterator), false
}

// create prefix to be rawdb.SnapshotStoragePrefix + account[:]

// Next steps the iterator forward one element, returning false if exhausted.
func (it *diskStorageIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator was already exhausted, don't bother
	return false
}

// Try to advance the iterator and release it if we reached the end

// Error returns any failure that occurred during iteration, which might have
// caused a premature iteration exit (e.g. snapshot stack becoming stale).
//
// A diff layer is immutable after creation content wise and can always be fully
// iterated without error, so this method always returns nil.
func (it *diskStorageIterator) Error() error { _ = "STUB: not implemented"; return nil }

// Iterator is exhausted and released

// Hash returns the hash of the storage slot the iterator is currently at.
func (it *diskStorageIterator) Hash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

// The prefix will be truncated

// Slot returns the raw storage slot content the iterator is currently at.
func (it *diskStorageIterator) Slot() []byte { _ = "STUB: not implemented"; return nil }

// Release releases the database snapshot held during iteration.
func (it *diskStorageIterator) Release() {
	_ = "STUB: not implemented"
	// The iterator is auto-released on exhaustion, so make sure it's still alive
	return
}
