// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
)

// trieQueue persists storage trie roots with their associated
// accounts when [RegisterStorageTrie] is called. These are
// later returned from the [getNextTrie] method.
type trieQueue struct {
	db              ethdb.Database
	nextStorageRoot []byte
}

func NewTrieQueue(db ethdb.Database) *trieQueue { _ = "STUB: not implemented"; return nil }

// clearIfRootDoesNotMatch clears progress and segment markers if
// the persisted root does not match the root we are syncing to.
func (t *trieQueue) clearIfRootDoesNotMatch(root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// If no sync root exists, treat it as empty hash (no previous sync).

// if not resuming, clear all progress markers

// RegisterStorageTrie is called by the main trie's leaf handling callbacks
// It adds a key built as [syncProgressPrefix+root+account] to the database.
// getNextTrie iterates this prefix to find storage tries and accounts
// associated with them.
func (t *trieQueue) RegisterStorageTrie(root common.Hash, account common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// StorageTrieDone is called when a storage trie has completed syncing.
// This removes any progress markers for the trie.
func (t *trieQueue) StorageTrieDone(root common.Hash) error { _ = "STUB: not implemented"; return nil }

// getNextTrie returns the next storage trie to sync, along with a slice
// of accounts that point to the returned storage trie.
// Returns true if there are more storage tries to sync and false otherwise.
// Note: if a non-nil root is returned, getNextTrie guarantees that there will be at least
// one account hash in the returned slice.
func (t *trieQueue) getNextTrie() (common.Hash, []common.Hash, bool, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, false, nil
}

// Iterate over the keys to find the next storage trie root and all of the account hashes that contain the same storage root.

// Unpack the state root and account hash from the current key

// Set the root for the first pass

// If the next root is different than the originally set root, then we've iterated over all of the account hashes that
// have the same storage trie root. Set more to be true, since there is at least one more storage trie.

// If we found another account with the same root, add the accountHash.

func (t *trieQueue) countTries() (int, error) { _ = "STUB: not implemented"; return 0, nil }
