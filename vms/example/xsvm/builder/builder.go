// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/chain"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/tx"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

const MaxTxsPerBlock = 10

var _ Builder = (*builder)(nil)

type Builder interface {
	SetPreference(preferred ids.ID)
	AddTx(ctx context.Context, tx *tx.Tx) error
	WaitForEvent(ctx context.Context) (common.Message, error)
	BuildBlock(ctx context.Context, blockContext *smblock.Context) (chain.Block, error)
}

type builder struct {
	chainContext *snow.Context
	chain        chain.Chain

	preference ids.ID
	// pendingTxsCond is awoken once there is at least one pending transaction.
	pendingTxsCond *lock.Cond
	pendingTxs     *linked.Hashmap[ids.ID, *tx.Tx]
}

func New(chainContext *snow.Context, chain chain.Chain) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) SetPreference(preferred ids.ID) { _ = "STUB: not implemented"; return }

func (b *builder) AddTx(_ context.Context, newTx *tx.Tx) error {
	_ = "STUB: not implemented"
	// TODO: verify [tx] against the currently preferred state
	return nil
}

func (b *builder) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

func (b *builder) BuildBlock(ctx context.Context, blockContext *smblock.Context) (chain.Block, error) {
	_ = "STUB: not implemented"
	return *new(chain.Block), nil
}

// This tx was invalid, drop it and continue block building

// TODO: populate fees

// This tx was invalid, drop it and continue block building
