// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"time"

	// Imported to allow IDE resolution of comments like [types.Block]. The
	// package is imported in other files so this is a no-op beyond devex.
	_ "github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/saevm/adaptor"
)

var _ adaptor.BlockProperties = (*Block)(nil)

// ID returns [types.Block.Hash] from the wrapped [types.Block].
func (b *Block) ID() ids.ID {
	_ = "STUB: not implemented"
	return *

	// Parent returns [types.Block.ParentHash] from the wrapped [types.Block].
	new(ids.ID)
}

func (b *Block) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// Bytes returns the RLP encoding of the wrapped [types.Block]. If encoding
// returns an error, it is logged at the ERROR level and a nil slice is
// returned.
func (b *Block) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Height returns [types.Block.NumberU64] from the wrapped [types.Block].
func (b *Block) Height() uint64 { _ = "STUB: not implemented"; return 0 }

// Timestamp returns the timestamp of the wrapped [types.Block], at
// [time.Second] resolution.
func (b *Block) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

//#nosec G115 -- Won't overflow for a few millennia
