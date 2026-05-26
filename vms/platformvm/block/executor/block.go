// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ snowman.Block             = (*Block)(nil)
	_ snowman.OracleBlock       = (*Block)(nil)
	_ smblock.WithVerifyContext = (*Block)(nil)
)

// Exported for testing in platformvm package.
type Block struct {
	block.Block
	manager *manager
}

func (*Block) ShouldVerifyWithContext(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *Block) VerifyWithContext(ctx context.Context, blockContext *smblock.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If the chain is bootstrapped and the warp messages haven't been verified,
// we must verify them.

// If the block was previously executed, we don't need to execute it again,
// we can just mark that the warp messages are valid at this height.

// Since this is the first time we are verifying this block, we must execute
// the state transitions to generate the state diffs.

func (b *Block) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Block) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Block) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Block) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *Block) Options(context.Context) ([2]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
