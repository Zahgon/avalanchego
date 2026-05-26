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
	_ BanffBlock = (*BanffStandardBlock)(nil)
	_ Block      = (*ApricotStandardBlock)(nil)
)

type BanffStandardBlock struct {
	Time                 uint64 `serialize:"true" json:"time"`
	ApricotStandardBlock `serialize:"true"`
}

func (b *BanffStandardBlock) Timestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (b *BanffStandardBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewBanffStandardBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
	txs []*txs.Tx,
) (*BanffStandardBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ApricotStandardBlock struct {
	CommonBlock  `serialize:"true"`
	Transactions []*txs.Tx `serialize:"true" json:"txs"`
}

func (b *ApricotStandardBlock) initialize(bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *ApricotStandardBlock) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (b *ApricotStandardBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *ApricotStandardBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

// NewApricotStandardBlock is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Standard blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotStandardBlock(
	parentID ids.ID,
	height uint64,
	txs []*txs.Tx,
) (*ApricotStandardBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
