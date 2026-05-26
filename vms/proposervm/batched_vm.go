// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"

	statelessblock "github.com/ava-labs/avalanchego/vms/proposervm/block"
)

var _ block.BatchedChainVM = (*VM)(nil)

func (vm *VM) GetAncestors(
	ctx context.Context,
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrievalTime time.Duration,
) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// hereinafter loop over proposerVM cache and DB, possibly till snowman++
// fork is hit

// maybe we have hit the proposerVM fork here?

// Ensure response size isn't too large. Include wrappers.IntLen because
// the size of the message is included with each container, and the size
// is repr. by an int.

// reached maximum size or ran out of time

// snowman++ fork may have been hit.

// return what we have

func (vm *VM) BatchedParseBlock(ctx context.Context, blks [][]byte) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parse all inner blocks at once

func (vm *VM) getStatelessBlk(blkID ids.ID) (statelessblock.Block, error) {
	_ = "STUB: not implemented"
	return *new(statelessblock.Block), nil
}
