// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
)

func (vm *blockVM) GetAncestors(
	ctx context.Context,
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrivalTime time.Duration,
) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *blockVM) BatchedParseBlock(ctx context.Context, blks [][]byte) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
