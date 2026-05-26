// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
)

var (
	_ syncTask = (*mainTrieTask)(nil)
	_ syncTask = (*storageTrieTask)(nil)
)

type syncTask interface {
	// IterateLeafs should return an iterator over
	// trie leafs already persisted to disk for this
	// trie. Used for restoring progress in case of an
	// interrupted sync and for hashing segments.
	IterateLeafs(seek common.Hash) ethdb.Iterator

	// callbacks used to form a LeafSyncTask
	OnStart() (bool, error)
	OnLeafs(ctx context.Context, db ethdb.KeyValueWriter, keys, vals [][]byte) error
	OnFinish() error
}

type mainTrieTask struct {
	sync *stateSync
}

func NewMainTrieTask(sync *stateSync) syncTask { _ = "STUB: not implemented"; return *new(syncTask) }

func (m *mainTrieTask) IterateLeafs(seek common.Hash) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

// OnStart always returns false since the main trie task cannot be skipped.
func (*mainTrieTask) OnStart() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (m *mainTrieTask) OnFinish() error { _ = "STUB: not implemented"; return nil }

func (m *mainTrieTask) OnLeafs(ctx context.Context, db ethdb.KeyValueWriter, keys, vals [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// loop over the keys, decode them as accounts, then check for any
// storage or code we need to sync as well.

// persist the account data

// check if this account has storage root that we need to fetch

// check if this account has code and add it to codeHashes to fetch
// at the end of this loop.

// Add collected code hashes to the code fetcher.

type storageTrieTask struct {
	sync     *stateSync
	root     common.Hash
	accounts []common.Hash
}

func NewStorageTrieTask(sync *stateSync, root common.Hash, accounts []common.Hash) syncTask {
	_ = "STUB: not implemented"
	return *new(syncTask)
}

func (s *storageTrieTask) IterateLeafs(seek common.Hash) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

func (s *storageTrieTask) OnStart() (bool, error) {
	_ = "STUB: not implemented"
	// check if this storage root is on disk
	return false, nil
}

//nolint:nilerr // the storage trie does not exist, so it should be rerequested

// If the storage trie is already on disk, we only need to populate the storage snapshot for [accountHash]
// with the trie contents. There is no need to re-sync the trie, since it is already present.

// If the storage trie cannot be iterated (due to an incomplete trie from pruning this storage trie in the past)
// then we re-sync it here. Therefore, this error is not fatal and we can safely continue here.

// Populating the snapshot from the existing storage trie succeeded,
// return true to skip this task.

func (s *storageTrieTask) OnFinish() error { _ = "STUB: not implemented"; return nil }

func (s *storageTrieTask) OnLeafs(ctx context.Context, db ethdb.KeyValueWriter, keys, vals [][]byte) error {
	_ = "STUB: not implemented"
	// persists the trie leafs to the snapshot for all accounts associated with this root
	return nil
}

// Check context cancellation before processing each account to allow early exit during shutdown.
