// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"io"

	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/bloombits"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/eth"
	"github.com/ava-labs/libevm/eth/filters"
	"github.com/ava-labs/libevm/ethdb"
)

// chainIndexer implements the subset of [ethapi.Backend] required to back a
// [core.ChainIndexer].
type chainIndexer struct {
	Chain
}

var _ core.ChainIndexerChain = chainIndexer{}

func (c chainIndexer) CurrentHeader() *types.Header { _ = "STUB: not implemented"; return nil }

// A bloomOverrider constructs Bloom filters from persisted receipts instead of
// relying on the [types.Header] field.
type bloomOverrider struct {
	chain Chain
}

var _ filters.BloomOverrider = bloomOverrider{}

// OverrideHeaderBloom returns the Bloom filter of the receipts generated when
// executing the respective block, whereas the [types.Header] carries those
// settled by the block.
func (b bloomOverrider) OverrideHeaderBloom(header *types.Header) types.Bloom {
	_ = "STUB: not implemented"
	return *new(types.Bloom)
}

// bloomIndexer provides the [bloomIndexer.BloomStatus] and
// [bloomIndexer.ServiceFilter] methods of an [ethapi.Backend] implementation.
type bloomIndexer struct {
	indexer  *core.ChainIndexer
	size     uint64
	handlers *eth.BloomHandlers
}

// newBloomIndexer creates a [bloomIndexer] and starts the indexer to run with
// events from `chain`.
//
// The consumer must call [bloomIndexer.Close] to release allocated resources.
func newBloomIndexer(db ethdb.Database, chain core.ChainIndexerChain, override filters.BloomOverrider, size uint64) *bloomIndexer {
	_ = "STUB: not implemented"
	return nil
}

func (b *bloomIndexer) BloomStatus() (size uint64, sections uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (b *bloomIndexer) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	_ = "STUB: not implemented"
	return
}

var _ io.Closer = (*bloomIndexer)(nil)

func (b *bloomIndexer) Close() error { _ = "STUB: not implemented"; return nil }

var _ core.ChainIndexerBackend = (*bloomBackend)(nil)

// bloomBackend is a wrapper around a [core.BloomIndexer] that overrides
// Process() to allow for custom bloom-filter generation.
type bloomBackend struct {
	*core.BloomIndexer
	filters.BloomOverrider
}

func (b *bloomBackend) Process(ctx context.Context, hdr *types.Header) error {
	_ = "STUB: not implemented"
	return nil
}
