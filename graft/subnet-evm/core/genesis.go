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

package core

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/common/math"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/triedb"
)

//go:generate go tool gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go

var errGenesisNoConfig = errors.New("genesis has no chain configuration")

// Deprecated: use types.Account instead.
type GenesisAccount = types.Account

// Deprecated: use types.GenesisAlloc instead.
type GenesisAlloc = types.GenesisAlloc

type Airdrop struct {
	// Address strings are hex-formatted common.Address
	Address common.Address `json:"address"`
}

// Genesis specifies the header fields, state of a genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type Genesis struct {
	Config        *params.ChainConfig `json:"config"`
	Nonce         uint64              `json:"nonce"`
	Timestamp     uint64              `json:"timestamp"`
	ExtraData     []byte              `json:"extraData"`
	GasLimit      uint64              `json:"gasLimit"   gencodec:"required"`
	Difficulty    *big.Int            `json:"difficulty" gencodec:"required"`
	Mixhash       common.Hash         `json:"mixHash"`
	Coinbase      common.Address      `json:"coinbase"`
	Alloc         types.GenesisAlloc  `json:"alloc"      gencodec:"required"`
	AirdropHash   common.Hash         `json:"airdropHash"`
	AirdropAmount *big.Int            `json:"airdropAmount"`
	AirdropData   []byte              `json:"-"` // provided in a separate file, not serialized in this struct.

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number        uint64      `json:"number"`
	GasUsed       uint64      `json:"gasUsed"`
	ParentHash    common.Hash `json:"parentHash"`
	BaseFee       *big.Int    `json:"baseFeePerGas"` // EIP-1559
	ExcessBlobGas *uint64     `json:"excessBlobGas"` // EIP-4844
	BlobGasUsed   *uint64     `json:"blobGasUsed"`   // EIP-4844
}

// field type overrides for gencodec
type genesisSpecMarshaling struct {
	Nonce         math.HexOrDecimal64
	Timestamp     math.HexOrDecimal64
	ExtraData     hexutil.Bytes
	GasLimit      math.HexOrDecimal64
	GasUsed       math.HexOrDecimal64
	Number        math.HexOrDecimal64
	Difficulty    *math.HexOrDecimal256
	Alloc         map[common.UnprefixedAddress]types.Account
	BaseFee       *math.HexOrDecimal256
	AirdropAmount *math.HexOrDecimal256
	ExcessBlobGas *math.HexOrDecimal64
	BlobGasUsed   *math.HexOrDecimal64
}

// GenesisMismatchError is raised when trying to overwrite an existing
// genesis block with an incompatible one.
type GenesisMismatchError struct {
	Stored, New common.Hash
}

func (e *GenesisMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

// SetupGenesisBlock writes or updates the genesis block in db.
// The block that will be used is:
//
//	                     genesis == nil       genesis != nil
//	                  +------------------------------------------
//	db has no genesis |  main-net default  |  genesis
//	db has genesis    |  from DB           |  genesis (if compatible)

// The argument [genesis] must be specified and must contain a valid chain config.
// If the genesis block has already been set up, then we verify the hash matches the genesis passed in
// and that the chain config contained in genesis is backwards compatible with what is stored in the database.
//
// The stored chain configuration will be updated if it is compatible (i.e. does not
// specify a fork block below the local head block). In case of a conflict, the
// error is a *params.ConfigCompatError and the new, unwritten config is returned.
func SetupGenesisBlock(
	db ethdb.Database, triedb *triedb.Database, genesis *Genesis, lastAcceptedHash common.Hash, skipChainConfigCheckCompatible bool,
) (*params.ChainConfig, common.Hash, error) {
	_ = "STUB: not implemented"
	return nil, *new(common.Hash), nil
}

// Just commit the new block if there is no stored genesis block.

// The genesis block is present(perhaps in ancient database) while the
// state database is not initialized yet. It can happen that the node
// is initialized with an external ancient store. Commit genesis state
// in this case.

// Ensure the stored genesis matches with the given one.

// With Firewood's deferred persistence, a crash before the
// first persist leaves the trie database empty while LevelDB
// already has blocks beyond genesis. In this case, only recommit
// genesis state to the trieDB.

// Check whether the genesis block is already written.

// Get the existing chain configuration.

// If there is no previously stored chain config, write the chain config to disk.

// Note: this can happen since we did not previously write the genesis block and chain config in the same batch.

// Notes on the following line:
// - this is needed in coreth to handle the case where existing nodes do not
//   have the Berlin or London forks initialized by block number on disk.
//   See https://github.com/ava-labs/coreth/pull/667/files
// - this is not needed in subnet-evm but it does not impact it either

// Check config compatibility and write the config. Compatibility errors
// are returned to the caller unless we're already at block zero.
// we use last accepted block for cfg compatibility check. Note this allows
// the node to continue if it previously halted due to attempting to process blocks with
// an incorrect chain config.

// this should never happen, but we check anyway
// when we start syncing from scratch, the last accepted block
// will be genesis block

// Required to write the chain config to disk to ensure both the chain config and upgrade bytes are persisted to disk.
// Note: this intentionally removes an extra check from upstream.

// IsVerkle indicates whether the state is already stored in a verkle
// tree at genesis time.
func (g *Genesis) IsVerkle() bool { _ = "STUB: not implemented"; return false }

// ToBlock returns the genesis block according to genesis specification.
func (g *Genesis) ToBlock() *types.Block { _ = "STUB: not implemented"; return nil }

func (g *Genesis) trieConfig() *triedb.Config { _ = "STUB: not implemented"; return nil }

// TODO: migrate this function to "flush" for more similarity with upstream.
func (g *Genesis) toBlock(db ethdb.Database, triedb *triedb.Database) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configure any stateful precompiles that should be enabled in the genesis.

// Do custom allocation after airdrop in case an address shows up in standard
// allocation

// When Etna/Cancun is active, `BlockGasCost` are decoded to 0 if it's nil.
// This is because these fields come before the other optional Cancun fields in RLP order.
// This only occurs with a serialized and written genesis block, and then reading it back.
// While this does not affect anything (because we don't use `ToBlock` to retrieve the genesis block),
// it's still confusing and breaking few tests. So we set it here to 0 to make it consistent.

// EIP-4788: The parentBeaconBlockRoot of the genesis block is always
// the zero hash. This is because the genesis block does not have a parent
// by definition.

// EIP-4844 fields

// Granite: set TimeMilliseconds

// Create the genesis block to use the block hash

// Commit newly generated states into disk if it's not empty.

// Commit writes the block and state of a genesis specification to the database.
// The block is committed as the canonical head block.
func (g *Genesis) Commit(db ethdb.Database, triedb *triedb.Database) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustCommit writes the genesis block and state to db, panicking on error.
// The block is committed as the canonical head block.
func (g *Genesis) MustCommit(db ethdb.Database, triedb *triedb.Database) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

func (g *Genesis) Verify() error {
	_ = "STUB: not implemented"
	// Make sure genesis gas limit is consistent
	return nil
}

// Verify config

// GenesisBlockForTesting creates and writes a block in which addr has the given wei balance.
func GenesisBlockForTesting(db ethdb.Database, addr common.Address, balance *big.Int) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// ReadBlockByHash reads the block with the given hash from the database.
func ReadBlockByHash(db ethdb.Reader, hash common.Hash) *types.Block {
	_ = "STUB: not implemented"
	return nil
}
