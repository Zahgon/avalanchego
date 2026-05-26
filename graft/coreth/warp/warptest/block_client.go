// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// warptest exposes common functionality for testing the warp package.
package warptest

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
)

// EmptyBlockClient returns an error if a block is requested
var EmptyBlockClient BlockClient = MakeBlockClient()

type BlockClient func(ctx context.Context, blockID ids.ID) (snowman.Block, error)

func (f BlockClient) GetAcceptedBlock(ctx context.Context, blockID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *

	// MakeBlockClient returns a new BlockClient that returns the provided blocks.
	// If a block is requested that isn't part of the provided blocks, an error is
	// returned.
	new(snowman.Block), nil
}

func MakeBlockClient(blkIDs ...ids.ID) BlockClient {
	_ = "STUB: not implemented"
	return *new(BlockClient)
}
