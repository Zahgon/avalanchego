// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

const blockCacheSize = 64 * units.MiB

var (
	errBlockWrongVersion = errors.New("wrong version")

	_ BlockState = (*blockState)(nil)
)

type BlockState interface {
	GetBlock(blkID ids.ID) (block.Block, error)
	PutBlock(blk block.Block) error
	DeleteBlock(blkID ids.ID) error
}

type blockState struct {
	// Caches BlockID -> Block. If the Block is nil, that means the block is not
	// in storage.
	blkCache cache.Cacher[ids.ID, *blockWrapper]

	db database.Database
}

type blockWrapper struct {
	Block  []byte         `serialize:"true"`
	Status choices.Status `serialize:"true"`

	block block.Block
}

func cachedBlockSize(_ ids.ID, bw *blockWrapper) int { _ = "STUB: not implemented"; return 0 }

func NewBlockState(db database.Database) BlockState {
	_ = "STUB: not implemented"
	return *new(BlockState)
}

func NewMeteredBlockState(db database.Database, namespace string, metrics prometheus.Registerer) (BlockState, error) {
	_ = "STUB: not implemented"
	return *new(BlockState), nil
}

func (s *blockState) GetBlock(blkID ids.ID) (block.Block, error) {
	_ = "STUB: not implemented"
	return *new(block.Block), nil
}

// The key was in the database

func (s *blockState) PutBlock(blk block.Block) error { _ = "STUB: not implemented"; return nil }

func (s *blockState) DeleteBlock(blkID ids.ID) error { _ = "STUB: not implemented"; return nil }
