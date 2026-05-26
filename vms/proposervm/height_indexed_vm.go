// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
)

const pruneCommitPeriod = 1024

// vm.ctx.Lock should be held
func (vm *VM) GetBlockIDAtHeight(ctx context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// fork not reached yet. Block must be pre-fork

func (vm *VM) updateHeightIndex(height uint64, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// The fork was already reached. Just update the index.

// This is the first post fork block, store the fork height.

// Note: The last accepted block is not considered a historical block. Which
// is why <= is used rather than <. This prevents the user from only storing
// the last accepted block, which can never be safe due to the non-atomic
// commits between the proposervm database and the innerVM's database.

// Note: heightToDelete is >= forkHeight, so it is guaranteed not to
// underflow.

// Block may have already been deleted. This can happen due to a
// proposervm rollback, the node having recently state-synced, or the
// user reconfiguring the node to store more historical blocks than a
// prior run.

// TODO: Support async deletion of old blocks.
func (vm *VM) pruneOldBlocks() error { _ = "STUB: not implemented"; return nil }

// Chain hasn't forked yet

// TODO: Refactor to use DB iterators.
//
// Note: vm.lastAcceptedHeight is guaranteed to be >= height, so the
// subtraction can never underflow.

// Note: height is < vm.lastAcceptedHeight, so it is guaranteed not to
// overflow.
