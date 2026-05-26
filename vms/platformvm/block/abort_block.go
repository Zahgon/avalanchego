// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ BanffBlock = (*BanffAbortBlock)(nil)
	_ Block      = (*ApricotAbortBlock)(nil)
)

type BanffAbortBlock struct {
	Time              uint64 `serialize:"true" json:"time"`
	ApricotAbortBlock `serialize:"true"`
}

func (b *BanffAbortBlock) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *BanffAbortBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewBanffAbortBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
) (*BanffAbortBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ApricotAbortBlock struct {
	CommonBlock `serialize:"true"`
}

func (b *ApricotAbortBlock) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (*ApricotAbortBlock) InitCtx(*snow.Context) { _ = "STUB: not implemented"; return }

func (*ApricotAbortBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *ApricotAbortBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

// NewApricotAbortBlock is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Abort blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotAbortBlock(
	parentID ids.ID,
	height uint64,
) (*ApricotAbortBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
