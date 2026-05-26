// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"context"
	"time"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ simplex.BlockBuilder = (*BlockBuilder)(nil)

type BlockBuilder struct {
	log          logging.Logger
	vm           block.ChainVM
	blockTracker *blockTracker
}

const (
	maxBackoff  = 5 * time.Second
	initBackoff = 10 * time.Millisecond
)

// BuildBlock continuously tries to build a block until the context is cancelled. If there are no blocks to be built, it will wait for an event from the VM.
// It returns false if the context was cancelled, otherwise it returns the built block and true.
func (b *BlockBuilder) BuildBlock(ctx context.Context, metadata simplex.ProtocolMetadata, blacklist simplex.Blacklist) (simplex.VerifiedBlock, bool) {
	_ = "STUB: not implemented"
	return *new(simplex.VerifiedBlock), false
}

// Reset backoff after a successful block build

// WaitForPendingBlock blocks until a new block is ready to be built from the VM, or until the
// context is cancelled.
func (b *BlockBuilder) WaitForPendingBlock(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *BlockBuilder) waitForPendingBlock(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// backoff waits for `backoff` duration before returning the next backoff duration.
// It doubles the backoff duration each time it is called, up to a maximum of `maxBackoff`.
func backoff(ctx context.Context, backoff time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// exponential backoff
