// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snapshot

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
)

func (t *Tree) DiskAccountIterator(seek common.Hash) AccountIterator {
	_ = "STUB: not implemented"
	return *new(AccountIterator)
}

func (t *Tree) DiskStorageIterator(account common.Hash, seek common.Hash) StorageIterator {
	_ = "STUB: not implemented"
	return *new(StorageIterator)
}

type SnapshotIterable interface {
	Snapshot

	// AccountIterator creates an account iterator over an arbitrary layer.
	AccountIterator(seek common.Hash) AccountIterator

	// StorageIterator creates a storage iterator over an arbitrary layer.
	StorageIterator(account common.Hash, seek common.Hash) (StorageIterator, bool)
}

// NewDiskLayer creates a diskLayer for direct access to the contents of the on-disk
// snapshot. Does not perform any validation.
func NewDiskLayer(diskdb ethdb.KeyValueStore) SnapshotIterable {
	_ = "STUB: not implemented"
	return *new(SnapshotIterable)
}

// state sync uses iterators to access data, so this cache is not used.
// initializing it out of caution.

// NewTestTree creates a *Tree with a pre-populated diskLayer
func NewTestTree(diskdb ethdb.KeyValueStore, blockHash, root common.Hash) *Tree {
	_ = "STUB: not implemented"
	return nil
}
