// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowmantest

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils"
)

const (
	GenesisHeight uint64 = 0
)

var (
	_ utils.Sortable[*Block] = (*Block)(nil)

	GenesisID        = ids.GenerateTestID()
	GenesisTimestamp = upgrade.InitiallyActiveTime
	GenesisBytes     = GenesisID[:]
	Genesis          = BuildChain(1)[0]
)

func BuildChild(parent *Block) *Block { _ = "STUB: not implemented"; return nil }

func BuildChain(length int) []*Block { _ = "STUB: not implemented"; return nil }

func BuildDescendants(parent *Block, length int) []*Block { _ = "STUB: not implemented"; return nil }

type Block struct {
	snowtest.Decidable

	ParentV    ids.ID
	HeightV    uint64
	TimestampV time.Time
	VerifyV    error
	BytesV     []byte
}

func (b *Block) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *Block) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *Block) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *Block) Verify(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *Block) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Block) Compare(other *Block) int { _ = "STUB: not implemented"; return 0 }
