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

package core

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
)

// BlockValidator is responsible for validating block headers, uncles and
// processed state.
//
// BlockValidator implements Validator.
type BlockValidator struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for validating
}

// NewBlockValidator returns a new block validator which is safe for re-use
func NewBlockValidator(config *params.ChainConfig, blockchain *BlockChain, engine consensus.Engine) *BlockValidator {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBody validates the given block's uncles and verifies the block
// header's transaction and uncle roots. The headers are assumed to be already
// validated at this point.
func (v *BlockValidator) ValidateBody(block *types.Block) error {
	_ = "STUB: not implemented"
	// Check whether the block is already imported.
	return nil
}

// Header validity is known at this point. Here we verify that uncle and transactions
// given in the block body match the header.

// Blob transactions may be present after the Cancun fork.

// Count the number of blobs to validate against the header's blobGasUsed

// If the tx is a blob tx, it must NOT have a sidecar attached to be valid in a block.

// The individual checks for blob validity (version-check + not empty)
// happens in StateTransition.

// Check blob gas usage.

// div because the header is surely good vs the body might be bloated

// Ancestor block must be known.

// ValidateState validates the various changes that happen after a state transition,
// such as amount of used gas, the receipt roots and the state root itself.
func (v *BlockValidator) ValidateState(block *types.Block, statedb *state.StateDB, receipts types.Receipts, usedGas uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate the received block's bloom with the one derived from the generated receipts.
// For valid blocks this should always validate to true.

// Tre receipt Trie's root (R = (Tr [[H1, R1], ... [Hn, Rn]]))

// Validate the state root against the received state root and throw
// an error if they don't match.

// CalcGasLimit computes the gas limit of the next block after parent. It aims
// to keep the baseline gas above the provided floor, and increase it towards the
// ceil if the blocks are full. If the ceil is exceeded, it will always decrease
// the gas allowance.
func CalcGasLimit(parentGasUsed, parentGasLimit, gasFloor, gasCeil uint64) uint64 {
	_ = "STUB: not implemented"
	// contrib = (parentGasUsed * 3 / 2) / 1024
	return 0
}

// decay = parentGasLimit / 1024 -1

/*
	strategy: gasLimit of block-to-mine is set based on parent's
	gasUsed value.  if parentGasUsed > parentGasLimit * (2/3) then we
	increase it, otherwise lower it (or leave it unchanged if it's right
	at that usage) the amount increased/decreased depends on how far away
	from parentGasLimit * (2/3) parentGasUsed is.
*/

// If we're outside our allowed gas range, we try to hone towards them
