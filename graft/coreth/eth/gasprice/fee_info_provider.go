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
// Copyright 2015 The go-ethereum Authors
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
	"math/big"

	"github.com/ava-labs/libevm/core/types"
	lru "github.com/hashicorp/golang-lru"
)

// additional slots in the header cache to allow processing queries
// for previous blocks (with full number of blocks desired) if new
// blocks are added concurrently.
const feeCacheExtraSlots = 5

type feeInfoProvider struct {
	cache          *lru.Cache
	backend        OracleBackend
	newHeaderAdded func() // callback used in tests
}

// feeInfo is the type of data stored in feeInfoProvider's cache.
type feeInfo struct {
	baseFee   *big.Int   // base fee of the block
	tips      []*big.Int // tips for txs to be included in the block
	timestamp uint64     // timestamp of the block header
}

// newFeeInfoProvider returns a bounded buffer with [size] slots to
// store [*feeInfo] for the most recently accepted blocks.
func newFeeInfoProvider(backend OracleBackend, size int) (*feeInfoProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if size is zero, we return early as there is no
// reason for a goroutine to subscribe to the chain's
// accepted event.

// subscribe to the chain accepted event

// addHeader processes header into a feeInfo struct and caches the result.
func (f *feeInfoProvider) addHeader(ctx context.Context, header *types.Header, txs []*types.Transaction) (*feeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get returns the feeInfo for block with [number] if present in the cache
// and a boolean representing if it was found.
func (f *feeInfoProvider) get(number uint64) (*feeInfo, bool) {
	_ = "STUB: not implemented"
	// Note: use Peek on LRU to use it as a bounded buffer.
	return nil, false
}

// populateCache populates [f] with [size] blocks up to last accepted.
// Note: assumes [size] is greater than zero.
func (f *feeInfoProvider) populateCache(size int) error { _ = "STUB: not implemented"; return nil }

// Note: "size-1" because we need a total of size blocks.
