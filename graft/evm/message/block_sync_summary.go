// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"context"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var _ Syncable = (*BlockSyncSummary)(nil)

// BlockSyncSummary provides the information necessary to sync a node starting
// at the given block.
type BlockSyncSummary struct {
	BlockNumber uint64      `serialize:"true"`
	BlockHash   common.Hash `serialize:"true"`
	BlockRoot   common.Hash `serialize:"true"`

	summaryID  ids.ID
	bytes      []byte
	acceptImpl AcceptImplFn
}

func NewBlockSyncSummary(c codec.Manager, blockHash common.Hash, blockNumber uint64, blockRoot common.Hash) (*BlockSyncSummary, error) {
	_ = "STUB: not implemented"
	// We intentionally do not use the acceptImpl here and leave it for the parser to set.
	return nil, nil
}

func (s *BlockSyncSummary) GetBlockHash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *BlockSyncSummary) GetBlockRoot() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func (s *BlockSyncSummary) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *BlockSyncSummary) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *BlockSyncSummary) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *BlockSyncSummary) String() string { _ = "STUB: not implemented"; return "" }

func (s *BlockSyncSummary) Accept(context.Context) (block.StateSyncMode, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSyncMode), nil
}
