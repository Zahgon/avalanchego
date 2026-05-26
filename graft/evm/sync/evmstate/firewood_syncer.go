// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"context"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/database/merkle/firewood/syncer"
	"github.com/ava-labs/avalanchego/graft/evm/sync/code"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
	"github.com/ava-labs/avalanchego/network/p2p"

	merklesync "github.com/ava-labs/avalanchego/database/merkle/sync"
)

var (
	_ types.Syncer    = (*FirewoodSyncer)(nil)
	_ types.Finalizer = (*FirewoodSyncer)(nil)
)

type FirewoodSyncer struct {
	s         *merklesync.Syncer[*syncer.RangeProof, struct{}]
	cancel    context.CancelFunc
	codeQueue *code.Queue
	// finalizeOnce is initialized in the constructor to make Finalize idempotent.
	finalizeOnce func() error
}

func NewFirewoodSyncer(config syncer.Config, db *ffi.Database, target common.Hash, codeQueue *code.Queue, client *p2p.Client) (*FirewoodSyncer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overwritten in Sync

func (f *FirewoodSyncer) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (f *FirewoodSyncer) Finalize() error { _ = "STUB: not implemented"; return nil }

// finish performs the finalization logic for the FirewoodSyncer inside a [sync.Once].
// This is linked to the [sync.Once] in the constructor, and should not be called directly.
func (f *FirewoodSyncer) finish() error {
	_ = "STUB: not implemented"
	// Ensure the syncer stops work and the code queue closes on exit.
	return nil
}

func (*FirewoodSyncer) ID() string { _ = "STUB: not implemented"; return "" }

func (*FirewoodSyncer) Name() string { _ = "STUB: not implemented"; return "" }
