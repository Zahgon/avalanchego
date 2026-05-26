// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// While an argument can be made for embedding the [types.Block] in [Block],
// instead of aliasing methods, that risks incorrect usage of subtle differences
// under SAE. These methods are direct aliases i.f.f. the interpretation is
// unambiguous, otherwise their names are clarified (e.g.
// [Block.SettledStateRoot]).

// EthBlock returns the raw EVM block wrapped by b. Prefer accessing its
// properties via the methods aliased on [Block] as some (e.g.
// [types.Block.Root]) have ambiguous interpretation under SAE.
func (b *Block) EthBlock() *types.Block {
	_ = "STUB: not implemented"

	// Body is equivalent to calling [types.Block.Body] on the block returned by
	// [Block.EthBlock].
	return nil
}

func (b *Block) Body() *types.Body {
	_ = "STUB: not implemented"

	// SettledStateRoot returns the state root after execution of the last block
	// settled by b. It is a convenience wrapper for calling [types.Block.Root] on
	// the wrapped [types.Block].
	return nil
}

func (b *Block) SettledStateRoot() common.Hash {
	_ = "STUB: not implemented"

	// BuildTime returns the Unix timestamp of the block, which is the canonical
	// inclusion time of its transactions; see [Block.ExecutedByGasTime] for their
	// execution timestamp. BuildTime is a convenience wrapper for calling
	// [types.Block.Time] on the wrapped [types.Block].
	return *new(common.Hash)
}

func (b *Block) BuildTime() uint64 {
	_ = "STUB: not implemented"

	// Hash returns [types.Block.Hash] from the wrapped [types.Block].
	return 0
}

func (b *Block) Hash() common.Hash {
	_ = "STUB: not implemented"

	// Header returns [types.Block.Header] from the wrapped [types.Block].
	return *new(common.Hash)
}

func (b *Block) Header() *types.Header {
	_ = "STUB: not implemented"

	// ParentHash returns [types.Block.ParentHash] from the wrapped [types.Block].
	return nil
}

func (b *Block) ParentHash() common.Hash {
	_ = "STUB: not implemented"
	return *

	// NumberU64 returns [types.Block.NumberU64] from the wrapped [types.Block].
	new(common.Hash)
}

func (b *Block) NumberU64() uint64 { _ = "STUB: not implemented"; return 0 }

// Number returns [types.Block.Number] from the wrapped [types.Block].
func (b *Block) Number() *big.Int {
	_ = "STUB: not implemented"

	// Transactions returns [types.Block.Transactions] from the wrapped [types.Block].
	return nil
}

func (b *Block) Transactions() types.Transactions {
	_ = "STUB: not implemented"
	return *new(types.Transactions)
}
