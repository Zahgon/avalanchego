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
// Copyright 2014 The go-ethereum Authors
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

// Package core implements the Ethereum consensus protocol.
package core

import (
	"context"
	"sync"

	"github.com/ava-labs/libevm/core/types"
)

type blockAndState struct {
	block    *types.Block
	hasState bool
	err      error
}

type blockChainIterator struct {
	bc *BlockChain

	nextReadBlockHeight   uint64
	nextBlockHeightToRead uint64
	blocks                []*blockAndState
	blocksRead            chan *blockAndState
	heightsToRead         chan uint64

	wg        sync.WaitGroup
	closeOnce sync.Once
	onClose   chan struct{}
}

func newBlockChainIterator(bc *BlockChain, start uint64, parallelism int) *blockChainIterator {
	_ = "STUB: not implemented"
	return nil
}

// Start [parallelism] worker threads to read block information

// Start a goroutine to read incoming heights from [heightsToRead]
// fetch the corresponding block information and place it on the
// [blocksRead] channel.

// Read heights in from [heightsToRead]

// populateReaders ie. adds task for [parallelism] threads

// populateReaders adds the heights for the next [parallelism] blocks to
// [blocksToRead]. This is called piecewise to ensure that each of the blocks
// is read within Next and set in [blocks] before moving on to the next tranche
// of blocks.
func (i *blockChainIterator) populateReaders(lastAccepted uint64) {
	_ = "STUB: not implemented"
	return
}

// Next retrieves the next consecutive block in the iteration
func (i *blockChainIterator) Next(ctx context.Context) (*types.Block, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// If the nextBlock in the iteration has already been populated
// return the block immediately.

// Otherwise, keep reading in block info from [blocksRead]
// and populate the [blocks] buffer until we hit the actual
// next block in the iteration.

// Stop closes the [onClose] channel signalling all worker threads to exit
// and waits for all of the worker threads to finish.
func (i *blockChainIterator) Stop() { _ = "STUB: not implemented"; return }
