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
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/options"
	"github.com/holiman/uint256"
)

// BlockGen creates blocks for testing.
// See GenerateChain for a detailed explanation.
type BlockGen struct {
	i       int
	cm      *chainMaker
	parent  *types.Block
	header  *types.Header
	statedb *state.StateDB

	gasPool  *GasPool
	txs      []*types.Transaction
	receipts []*types.Receipt
	uncles   []*types.Header

	engine           consensus.Engine
	onBlockGenerated func(*types.Block)
}

// SetCoinbase sets the coinbase of the generated block.
// It can be called at most once.
func (b *BlockGen) SetCoinbase(addr common.Address) { _ = "STUB: not implemented"; return }

// SetExtra sets the extra data field of the generated block.
func (b *BlockGen) SetExtra(data []byte) { _ = "STUB: not implemented"; return }

// AppendExtra appends data to the extra data field of the generated block.
func (b *BlockGen) AppendExtra(data []byte) { _ = "STUB: not implemented"; return }

// SetNonce sets the nonce field of the generated block.
func (b *BlockGen) SetNonce(nonce types.BlockNonce) { _ = "STUB: not implemented"; return }

// SetDifficulty sets the difficulty field of the generated block. This method is
// useful for Clique tests where the difficulty does not depend on time. For the
// ethash tests, please use OffsetTime, which implicitly recalculates the diff.
func (b *BlockGen) SetDifficulty(diff *big.Int) { _ = "STUB: not implemented"; return }

// Difficulty returns the currently calculated difficulty of the block.
func (b *BlockGen) Difficulty() *big.Int { _ = "STUB: not implemented"; return nil }

// SetParentBeaconRoot sets the parent beacon root field of the generated
// block.
func (b *BlockGen) SetParentBeaconRoot(root common.Hash) { _ = "STUB: not implemented"; return }

// addTx adds a transaction to the generated block. If no coinbase has
// been set, the block's coinbase is set to the zero address.
//
// There are a few options can be passed as well in order to run some
// customized rules.
// - bc:       enables the ability to query historical block hashes for BLOCKHASH
// - vmConfig: extends the flexibility for customizing evm rules, e.g. enable extra EIPs
func (b *BlockGen) addTx(bc *BlockChain, vmConfig vm.Config, tx *types.Transaction) {
	_ = "STUB: not implemented"
	return
}

// AddTx adds a transaction to the generated block. If no coinbase has
// been set, the block's coinbase is set to the zero address.
//
// AddTx panics if the transaction cannot be executed. In addition to the protocol-imposed
// limitations (gas limit, etc.), there are some further limitations on the content of
// transactions that can be added. Notably, contract code relying on the BLOCKHASH
// instruction will panic during execution if it attempts to access a block number outside
// of the range created by GenerateChain.
func (b *BlockGen) AddTx(tx *types.Transaction) { _ = "STUB: not implemented"; return }

// AddTxWithChain adds a transaction to the generated block. If no coinbase has
// been set, the block's coinbase is set to the zero address.
//
// AddTxWithChain panics if the transaction cannot be executed. In addition to the
// protocol-imposed limitations (gas limit, etc.), there are some further limitations on
// the content of transactions that can be added. If contract code relies on the BLOCKHASH
// instruction, the block in chain will be returned.
func (b *BlockGen) AddTxWithChain(bc *BlockChain, tx *types.Transaction) {
	_ = "STUB: not implemented"
	return
}

// AddTxWithVMConfig adds a transaction to the generated block. If no coinbase has
// been set, the block's coinbase is set to the zero address.
// The evm interpreter can be customized with the provided vm config.
func (b *BlockGen) AddTxWithVMConfig(tx *types.Transaction, config vm.Config) {
	_ = "STUB: not implemented"
	return
}

// GetBalance returns the balance of the given address at the generated block.
func (b *BlockGen) GetBalance(addr common.Address) *uint256.Int {
	_ = "STUB: not implemented"
	return nil
}

// AddUncheckedTx forcefully adds a transaction to the block without any validation.
//
// AddUncheckedTx will cause consensus failures when used during real
// chain processing. This is best used in conjunction with raw block insertion.
func (b *BlockGen) AddUncheckedTx(tx *types.Transaction) { _ = "STUB: not implemented"; return }

// Number returns the block number of the block being generated.
func (b *BlockGen) Number() *big.Int { _ = "STUB: not implemented"; return nil }

// Timestamp returns the timestamp of the block being generated.
func (b *BlockGen) Timestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// BaseFee returns the EIP-1559 base fee of the block being generated.
func (b *BlockGen) BaseFee() *big.Int { _ = "STUB: not implemented"; return nil }

// Gas returns the amount of gas left in the current block.
func (b *BlockGen) Gas() uint64 { _ = "STUB: not implemented"; return 0 }

// Signer returns a valid signer instance for the current block.
func (b *BlockGen) Signer() types.Signer { _ = "STUB: not implemented"; return *new(types.Signer) }

// AddUncheckedReceipt forcefully adds a receipts to the block without a
// backing transaction.
//
// AddUncheckedReceipt will cause consensus failures when used during real
// chain processing. This is best used in conjunction with raw block insertion.
func (b *BlockGen) AddUncheckedReceipt(receipt *types.Receipt) { _ = "STUB: not implemented"; return }

// TxNonce returns the next valid transaction nonce for the
// account at addr. It panics if the account does not exist.
func (b *BlockGen) TxNonce(addr common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

// AddUncle adds an uncle header to the generated block.
func (b *BlockGen) AddUncle(h *types.Header) { _ = "STUB: not implemented"; return }

// PrevBlock returns a previously generated block by number. It panics if
// num is greater or equal to the number of the block being generated.
// For index -1, PrevBlock returns the parent block given to GenerateChain.
func (b *BlockGen) PrevBlock(index int) *types.Block { _ = "STUB: not implemented"; return nil }

// OffsetTime modifies the time instance of a block, implicitly changing its
// associated difficulty. It's useful to test scenarios where forking is not
// tied to chain length directly.
func (b *BlockGen) OffsetTime(seconds int64) { _ = "STUB: not implemented"; return }

// SetOnBlockGenerated sets a callback function to be invoked after each block is generated
func (b *BlockGen) SetOnBlockGenerated(onBlockGenerated func(*types.Block)) {
	_ = "STUB: not implemented"
	return
}

// GenerateChain creates a chain of n blocks. The first block's
// parent will be the provided parent. db is used to store
// intermediate states and should contain the parent's state trie.
//
// The generator function is called with a new block generator for
// every block. Any transactions and uncles added to the generator
// become part of the block. If gen is nil, the blocks will be empty
// and their coinbase will be the zero address.
//
// Blocks created by GenerateChain do not contain valid proof of work
// values. Inserting them into BlockChain requires use of FakePow or
// a similar non-validating proof of work implementation.
func GenerateChain(config *params.ChainConfig, parent *types.Block, engine consensus.Engine, db ethdb.Database, n int, gap uint64, gen func(int, *BlockGen)) ([]*types.Block, []types.Receipts, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type generateChainConfig struct {
	commitToDisk bool
}

// GenerateChainOption configures [GenerateChainFromStateCache].
type GenerateChainOption = options.Option[generateChainConfig]

// WithoutDiskCommit skips persisting trie state to disk, allowing blocks
// to be separately accepted by a VM on the same database.
func WithoutDiskCommit() GenerateChainOption {
	_ = "STUB: not implemented"
	return *new(GenerateChainOption)
}

// GenerateChainFromStateCache is exactly like [GenerateChain], except allows other [triedb.Database] implementations.
func GenerateChainFromStateCache(config *params.ChainConfig, parent *types.Block, engine consensus.Engine, stateCache state.Database, n int, gap uint64, gen func(int, *BlockGen), opts ...GenerateChainOption) ([]*types.Block, []types.Receipts, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Execute any user modifications to the block

// Finalize and seal the block

// Write state changes to db

// Post-process the receipts.
// Here we assign the final block hash and other info into the receipt.
// In order for DeriveFields to work, the transaction and receipt lists need to be
// of equal length. If AddUncheckedTx or AddUncheckedReceipt are used, there will be
// extra ones, so we just trim the lists here.

// Re-expand to ensure all receipts are returned.

// Advance the chain.

// GenerateChainWithGenesis is a wrapper of GenerateChain which will initialize
// genesis block to database first according to the provided genesis specification
// then generate chain on top.
func GenerateChainWithGenesis(genesis *Genesis, engine consensus.Engine, n int, gap uint64, gen func(int, *BlockGen)) (ethdb.Database, []*types.Block, []types.Receipts, error) {
	_ = "STUB: not implemented"
	return *new(ethdb.Database), nil, nil, nil
}

func (cm *chainMaker) makeHeader(parent *types.Block, gap uint64, state *state.StateDB, engine consensus.Engine) *types.Header {
	_ = "STUB: not implemented"
	return nil
	// block time is fixed at [gap] seconds
}

// chainMaker contains the state of chain generation.
type chainMaker struct {
	bottom      *types.Block
	engine      consensus.Engine
	config      *params.ChainConfig
	chain       []*types.Block
	chainByHash map[common.Hash]*types.Block
	receipts    []types.Receipts
}

func newChainMaker(bottom *types.Block, config *params.ChainConfig, engine consensus.Engine) *chainMaker {
	_ = "STUB: not implemented"
	return nil
}

func (cm *chainMaker) add(b *types.Block, r []*types.Receipt) { _ = "STUB: not implemented"; return }

func (cm *chainMaker) blockByNumber(number uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// ChainReader/ChainContext implementation

// Config returns the chain configuration (for consensus.ChainReader).
func (cm *chainMaker) Config() *params.ChainConfig {
	_ = "STUB: not implemented"

	// Engine returns the consensus engine (for ChainContext).
	return nil
}

func (cm *chainMaker) Engine() consensus.Engine {
	_ = "STUB: not implemented"
	return *new(consensus.Engine)
}

func (cm *chainMaker) CurrentHeader() *types.Header { _ = "STUB: not implemented"; return nil }

func (cm *chainMaker) GetHeaderByNumber(number uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

func (cm *chainMaker) GetHeaderByHash(hash common.Hash) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

func (cm *chainMaker) GetHeader(hash common.Hash, number uint64) *types.Header {
	_ = "STUB: not implemented"
	return nil
}

func (cm *chainMaker) GetBlock(hash common.Hash, number uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}
