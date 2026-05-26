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
	_ BanffBlock = (*BanffCommitBlock)(nil)
	_ Block      = (*ApricotCommitBlock)(nil)
)

type BanffCommitBlock struct {
	Time               uint64 `serialize:"true" json:"time"`
	ApricotCommitBlock `serialize:"true"`
}

func (b *BanffCommitBlock) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *BanffCommitBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewBanffCommitBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
) (*BanffCommitBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ApricotCommitBlock struct {
	CommonBlock `serialize:"true"`
}

func (b *ApricotCommitBlock) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (*ApricotCommitBlock) InitCtx(*snow.Context) { _ = "STUB: not implemented"; return }

func (*ApricotCommitBlock) Txs() []*txs.Tx { _ = "STUB: not implemented"; return nil }

func (b *ApricotCommitBlock) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

func NewApricotCommitBlock(
	parentID ids.ID,
	height uint64,
) (*ApricotCommitBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
