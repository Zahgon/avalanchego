// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/graft/evm/sync/client"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

const blocksPerRequest = 32

var (
	_                        types.Syncer = (*BlockSyncer)(nil)
	errBlocksToFetchRequired              = errors.New("blocksToFetch must be > 0")
	errFromHashRequired                   = errors.New("fromHash must be non-zero when fromHeight > 0")
)

type BlockSyncer struct {
	db            ethdb.Database
	client        client.Client
	fromHash      common.Hash
	fromHeight    uint64
	blocksToFetch uint64
}

func NewSyncer(client client.Client, db ethdb.Database, fromHash common.Hash, fromHeight uint64, blocksToFetch uint64) (*BlockSyncer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Name returns the human-readable name for this sync task.
func (*BlockSyncer) Name() string { _ = "STUB: not implemented"; return "" }

// ID returns the stable identifier for this sync task.
func (*BlockSyncer) ID() string { _ = "STUB: not implemented"; return "" }

// Sync fetches (up to) BlocksToFetch blocks from peers
// using Client and writes them to disk.
// the process begins with FromHash and it fetches parents recursively.
// fetching starts from the first ancestor not found on disk
//
// TODO: We could inspect the database more accurately to ensure we never fetch
// any blocks that are locally available.
// We could also prevent overrequesting blocks, if the number of blocks needed
// to be fetched isn't a multiple of blocksPerRequest.
func (s *BlockSyncer) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// first, check for blocks already available on disk so we don't
// request them from peers.

// block was not found

// block exists

// get any blocks we couldn't find on disk from peers and write
// them to disk.
