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
// Copyright 2023 The go-ethereum Authors
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

package eth

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/libevm/core/types"
)

// AdminAPI is the collection of Ethereum full node related APIs for node
// administration.
type AdminAPI struct {
	eth *Ethereum
}

// NewAdminAPI creates a new instance of AdminAPI.
func NewAdminAPI(eth *Ethereum) *AdminAPI { _ = "STUB: not implemented"; return nil }

// ExportChain exports the current blockchain into a local file,
// or a range of blocks if first and last are non-nil.
func (api *AdminAPI) ExportChain(file string, first *uint64, last *uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// File already exists. Allowing overwrite could be a DoS vector,
// since the 'file' may point to arbitrary paths on the drive.

// Make sure we can create the file to export into

// Export the blockchain

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	_ = "STUB: not implemented"
	return false
}

// ImportChain imports a blockchain from a local file.
func (api *AdminAPI) ImportChain(file string) (bool, error) {
	_ = "STUB: not implemented"
	// Make sure the can access the file to import
	return false, nil
}

// Run actual the import in pre-configured batches

// Load a batch of blocks from the input file

// ignore the genesis block when importing blocks

// Import the batch and reset the buffer
