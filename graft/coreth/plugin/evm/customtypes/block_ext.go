// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customtypes

import (
	"math/big"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/rlp"

	"github.com/ava-labs/avalanchego/vms/evm/acp226"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

// SetBlockExtra sets the [BlockBodyExtra] `extra` in the [Block] `b`.
func SetBlockExtra(b *ethtypes.Block, extra *BlockBodyExtra) { _ = "STUB: not implemented"; return }

// BlockBodyExtra is a struct containing extra fields used by Avalanche
// in the [Block] and [Body].
type BlockBodyExtra struct {
	Version uint32
	ExtData *[]byte
}

// Copy deep copies the [BlockBodyExtra] `b` and returns it.
// It is notably used in the following functions:
// - [ethtypes.Block.Body]
// - [ethtypes.Block.WithSeal]
// - [ethtypes.Block.WithBody]
// - [ethtypes.Block.WithWithdrawals]
func (b *BlockBodyExtra) Copy() *BlockBodyExtra { _ = "STUB: not implemented"; return nil }

// BodyRLPFieldPointersForEncoding returns the fields that should be encoded
// for the [Body] and [BlockBodyExtra].
// Note the following fields are added (+) and removed (-) compared to geth:
// - (-) [ethtypes.Body] `Withdrawals` field
// - (+) [BlockBodyExtra] `Version` field
// - (+) [BlockBodyExtra] `ExtData` field
func (b *BlockBodyExtra) BodyRLPFieldsForEncoding(body *ethtypes.Body) *rlp.Fields {
	_ = "STUB: not implemented"
	return nil
}

// BodyRLPFieldPointersForDecoding returns the fields that should be decoded to
// for the [Body] and [BlockBodyExtra].
func (b *BlockBodyExtra) BodyRLPFieldPointersForDecoding(body *ethtypes.Body) *rlp.Fields {
	_ = "STUB: not implemented"
	return nil
}

// BlockRLPFieldPointersForEncoding returns the fields that should be encoded
// for the [Block] and [BlockBodyExtra].
// Note the following fields are added (+) and removed (-) compared to geth:
// - (-) [ethtypes.Block] `Withdrawals` field
// - (+) [BlockBodyExtra] `Version` field
// - (+) [BlockBodyExtra] `ExtData` field
func (b *BlockBodyExtra) BlockRLPFieldsForEncoding(block *ethtypes.BlockRLPProxy) *rlp.Fields {
	_ = "STUB: not implemented"
	return nil
}

// BlockRLPFieldPointersForDecoding returns the fields that should be decoded to
// for the [Block] and [BlockBodyExtra].
func (b *BlockBodyExtra) BlockRLPFieldPointersForDecoding(block *ethtypes.BlockRLPProxy) *rlp.Fields {
	_ = "STUB: not implemented"
	return nil
}

func BlockExtData(b *ethtypes.Block) []byte { _ = "STUB: not implemented"; return nil }

func BlockVersion(b *ethtypes.Block) uint32 { _ = "STUB: not implemented"; return 0 }

func BlockExtDataGasUsed(b *ethtypes.Block) *big.Int { _ = "STUB: not implemented"; return nil }

func BlockGasCost(b *ethtypes.Block) *big.Int { _ = "STUB: not implemented"; return nil }

func BlockTimeMilliseconds(b *ethtypes.Block) *uint64 { _ = "STUB: not implemented"; return nil }

func BlockMinDelayExcess(b *ethtypes.Block) *acp226.DelayExcess {
	_ = "STUB: not implemented"
	return nil
}

func CalcExtDataHash(extdata []byte) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func BlockTime(eth *ethtypes.Header) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func NewBlockWithExtData(
	header *ethtypes.Header, txs []*ethtypes.Transaction, uncles []*ethtypes.Header, receipts []*ethtypes.Receipt,
	hasher ethtypes.TrieHasher, extdata []byte, recalc bool,
) *ethtypes.Block {
	_ = "STUB: not implemented"
	return nil
}
