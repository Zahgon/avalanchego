// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"context"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var _ message.Syncable = (*Summary)(nil)

// Summary provides the information necessary to sync a node starting
// at the given block.
type Summary struct {
	*message.BlockSyncSummary `serialize:"true"`
	AtomicRoot                common.Hash `serialize:"true"`

	summaryID  ids.ID
	bytes      []byte
	acceptImpl message.AcceptImplFn
}

func NewSummary(blockHash common.Hash, blockNumber uint64, blockRoot common.Hash, atomicRoot common.Hash) (*Summary, error) {
	_ = "STUB: not implemented"
	// We intentionally do not use the acceptImpl here and leave it for the parser to set.
	return nil, nil
}

func (a *Summary) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (a *Summary) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (a *Summary) String() string { _ = "STUB: not implemented"; return "" }

func (a *Summary) Accept(context.Context) (block.StateSyncMode, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSyncMode), nil
}
