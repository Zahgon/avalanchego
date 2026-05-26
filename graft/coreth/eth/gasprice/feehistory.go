// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package gasprice

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/libevm/core/types"
)

var (
	errInvalidPercentile     = errors.New("invalid reward percentile")
	errRequestBeyondHead     = errors.New("request beyond head block")
	errBeyondHistoricalLimit = errors.New("request beyond historical limit")
)

const (
	maxQueryLimit = 100
)

// txGasAndReward is sorted in ascending order based on reward
type txGasAndReward struct {
	gasUsed uint64
	reward  *big.Int
}

type slimBlock struct {
	GasUsed  uint64
	GasLimit uint64
	BaseFee  *big.Int
	Txs      []txGasAndReward
}

// processBlock prepares a [slimBlock] from a retrieved block and list of
// receipts. This slimmed block can be cached and used for future calls.
func processBlock(block *types.Block, receipts types.Receipts) *slimBlock {
	_ = "STUB: not implemented"
	return nil
}

// processPercentiles returns baseFee, gasUsedRatio, and optionally reward percentiles (if any are
// requested)
func (sb *slimBlock) processPercentiles(percentiles []float64) ([]*big.Int, *big.Int, float64) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

// rewards were not requested

// return an all zero row if there are no transactions to gather data from

// sb transactions are already sorted by tip, so we don't need to re-sort

// resolveBlockRange resolves the specified block range to absolute block numbers while also
// enforcing backend specific limitations.
// Note: an error is only returned if retrieving the head header has failed. If there are no
// retrievable blocks in the specified range then zero block count is returned with no error.
func (oracle *Oracle) resolveBlockRange(ctx context.Context, lastBlock rpc.BlockNumber, blocks uint64) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	// Query either pending block or head header and set headBlock
	return 0, 0, nil
}

// Pending block not supported by backend, process until latest block

// If the requested last block reaches further back than [oracle.maxBlockHistory] past the last accepted block return an error
// Note: this allows some blocks past this point to be fetched since it will start fetching [blocks] from this point.

// If the requested block is above the accepted block return an error

// Ensure not trying to retrieve before genesis

// Truncate blocks range if extending past [oracle.maxBlockHistory]

// It is not possible that [blocks] could be <= 0 after
// truncation as the [lastBlock] requested will at least by fetchable.
// Otherwise, we would've returned an error earlier.

// FeeHistory returns data relevant for fee estimation based on the specified range of blocks.
// The range can be specified either with absolute block numbers or ending with the latest
// or pending block. Backends may or may not support gathering data from the pending block
// or blocks older than a certain age (specified in maxHistory). The first block of the
// actually processed range is returned to avoid ambiguity when parts of the requested range
// are not available or when the head has changed during processing this request.
// Three arrays are returned based on the processed blocks:
//   - reward: the requested percentiles of effective priority fees per gas of transactions in each
//     block, sorted in ascending order and weighted by gas used.
//   - baseFee: base fee per gas in the given block
//   - gasUsedRatio: gasUsed/gasLimit in the given block
//
// Note: baseFee includes the next block after the newest of the returned range, because this
// value can be derived from the newest block.
func (oracle *Oracle) FeeHistory(ctx context.Context, blocks uint64, unresolvedLastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// returning with no data and no error means there are no retrievable blocks

// Check if the context has errored

// getting no block and no error means we are requesting into the future (might happen because of a reorg)
