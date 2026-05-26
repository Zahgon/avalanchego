// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gasprice

import (
	"math/big"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/utils/logging"
)

type transaction struct {
	gas uint64   // [types.Transaction.Gas]
	tip *big.Int // [types.Transaction.EffectiveGasTip]
}

func newTx(tx *types.Transaction, baseFee *big.Int) transaction {
	_ = "STUB: not implemented"
	return *new(transaction)
}

func (t transaction) Compare(o transaction) int { _ = "STUB: not implemented"; return 0 }

type block struct {
	timestamp uint64        // [types.Header.Time]
	gasUsed   uint64        // [types.Header.GasUsed]
	gasLimit  uint64        // [types.Header.GasLimit]
	baseFee   *big.Int      // [types.Header.BaseFee]
	txs       []transaction // sorted ascending by tip
}

func newBlock(blk *types.Block) *block { _ = "STUB: not implemented"; return nil }

// tipPercentiles computes the gas-weighted tip at each requested percentile.
// all of which MUST be sorted in ascending order.
//
// Because block builders sequence transactions without executing them in SAE,
// we accumulate gas limits, not the gas charged.
func (b *block) tipPercentiles(percentiles []float64) []*big.Int {
	_ = "STUB: not implemented"
	return nil
}

// TODO:(StephenButtolph): Improve from `O(txs + percentiles)` to
// `O(percentiles * log(txs))` by binary searching for each threshold if
// networks with large blocks encounter performance degradation.

type blockCache struct {
	log     logging.Logger
	backend Backend
	// TODO(StephenButtolph): Use a ring-buffer rather than an LRU cache if we
	// observe cache contention in production.
	cache *lru.Cache[uint64, *block]
}

func newBlockCache(log logging.Logger, backend Backend, size int) *blockCache {
	_ = "STUB: not implemented"
	return nil
}

// getBlock returns the block at height n. If the block does not exist, it will
// return nil.
func (b *blockCache) getBlock(n uint64) *block { _ = "STUB: not implemented"; return nil }

//#nosec G115 -- Block numbers were previously resolved

// Don't cache a nil block. It may be populated in the future.

func (b *blockCache) cacheBlock(blk *types.Block) *block { _ = "STUB: not implemented"; return nil }
