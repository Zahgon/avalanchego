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
	_ BanffBlock = (*BanffProposalBlock)(nil)
	_ Block      = (*ApricotProposalBlock)(nil)
)

type BanffProposalBlock struct {
	Time                 uint64    `serialize:"true" json:"time"`
	Transactions         []*txs.Tx `serialize:"true" json:"txs"`
	ApricotProposalBlock `serialize:"true"`
}

func (b *BanffProposalBlock) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (b *BanffProposalBlock) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (b *BanffProposalBlock) Timestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (b *BanffProposalBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *BanffProposalBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewBanffProposalBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
	proposalTx *txs.Tx,
	decisionTxs []*txs.Tx,
) (*BanffProposalBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ApricotProposalBlock struct {
	CommonBlock `serialize:"true"`
	Tx          *txs.Tx `serialize:"true" json:"tx"`
}

func (b *ApricotProposalBlock) initialize(bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *ApricotProposalBlock) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (b *ApricotProposalBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *ApricotProposalBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

// NewApricotProposalBlock is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Proposal blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotProposalBlock(
	parentID ids.ID,
	height uint64,
	tx *txs.Tx,
) (*ApricotProposalBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
