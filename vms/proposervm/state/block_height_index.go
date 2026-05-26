// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
)

const cacheSize = 8192 // max cache entries

var (
	_ HeightIndex = (*heightIndex)(nil)

	heightPrefix   = []byte("height")
	metadataPrefix = []byte("metadata")

	forkKey = []byte("fork")
)

type HeightIndexGetter interface {
	// GetMinimumHeight return the smallest height of an indexed blockID. If
	// there are no indexed blockIDs, ErrNotFound will be returned.
	GetMinimumHeight() (uint64, error)
	GetBlockIDAtHeight(height uint64) (ids.ID, error)

	// Fork height is stored when the first post-fork block/option is accepted.
	// Before that, fork height won't be found.
	GetForkHeight() (uint64, error)
}

type HeightIndexWriter interface {
	SetForkHeight(height uint64) error
	SetBlockIDAtHeight(height uint64, blkID ids.ID) error
	DeleteBlockIDAtHeight(height uint64) error
}

// HeightIndex contains mapping of blockHeights to accepted proposer block IDs
// along with some metadata (fork height and checkpoint).
type HeightIndex interface {
	HeightIndexWriter
	HeightIndexGetter
}

type heightIndex struct {
	versiondb.Commitable

	// Caches block height -> proposerVMBlockID.
	heightsCache cache.Cacher[uint64, ids.ID]

	heightDB   database.Database
	metadataDB database.Database
}

func NewHeightIndex(db database.Database, commitable versiondb.Commitable) HeightIndex {
	_ = "STUB: not implemented"
	return *new(HeightIndex)
}

func (hi *heightIndex) GetMinimumHeight() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (hi *heightIndex) GetBlockIDAtHeight(height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (hi *heightIndex) SetBlockIDAtHeight(height uint64, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (hi *heightIndex) DeleteBlockIDAtHeight(height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (hi *heightIndex) GetForkHeight() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (hi *heightIndex) SetForkHeight(height uint64) error { _ = "STUB: not implemented"; return nil }
