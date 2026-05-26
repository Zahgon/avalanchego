// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package syncutils

import (
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
)

var (
	_ ethdb.Iterator = (*AccountIterator)(nil)
	_ ethdb.Iterator = (*StorageIterator)(nil)
)

// AccountIterator wraps a [snapshot.AccountIterator] to conform to [ethdb.Iterator]
// accounts will be returned in consensus (FullRLP) format for compatibility with trie data.
type AccountIterator struct {
	snapshot.AccountIterator
	err error
	val []byte
}

func (it *AccountIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (it *AccountIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *AccountIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *AccountIterator) Error() error { _ = "STUB: not implemented"; return nil }

// StorageIterator wraps a [snapshot.StorageIterator] to conform to [ethdb.Iterator]
type StorageIterator struct {
	snapshot.StorageIterator
}

func (it *StorageIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *StorageIterator) Value() []byte { _ = "STUB: not implemented"; return nil }
