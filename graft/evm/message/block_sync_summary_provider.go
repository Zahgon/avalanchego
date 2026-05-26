// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var _ SyncSummaryProvider = (*BlockSyncSummaryProvider)(nil)

type BlockSyncSummaryProvider struct {
	codec codec.Manager
}

func NewBlockSyncSummaryProvider(c codec.Manager) *BlockSyncSummaryProvider {
	_ = "STUB: not implemented"
	return nil
}

// StateSummaryAtBlock returns the block state summary at [block] if valid.
func (c *BlockSyncSummaryProvider) StateSummaryAtBlock(blk *types.Block) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (c *BlockSyncSummaryProvider) Parse(summaryBytes []byte, acceptImpl AcceptImplFn) (Syncable, error) {
	_ = "STUB: not implemented"
	return *new(Syncable), nil
}
