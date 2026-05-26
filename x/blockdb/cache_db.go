// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

import (
	"sync/atomic"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/database"
)

var _ database.HeightIndex = (*cacheDB)(nil)

// cacheDB caches data from the underlying [Database].
//
// Operations (Get, Has, Put) are not atomic with the underlying database.
// Concurrent writes to the same height can result in cache inconsistencies where
// the cache and database contain different values. This limitation is acceptable
// because concurrent writes to the same height are not an intended use case.
type cacheDB struct {
	db     *Database
	cache  *lru.Cache[BlockHeight, BlockData]
	closed atomic.Bool
}

func newCacheDB(db *Database, size uint16) *cacheDB { _ = "STUB: not implemented"; return nil }

func (c *cacheDB) Get(height BlockHeight) (BlockData, error) {
	_ = "STUB: not implemented"
	return *new(BlockData), nil
}

func (c *cacheDB) Put(height BlockHeight, data BlockData) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheDB) Has(height BlockHeight) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *cacheDB) Sync(start, end uint64) error { _ = "STUB: not implemented"; return nil }

func (c *cacheDB) Close() error { _ = "STUB: not implemented"; return nil }
