// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package blockstest provides test helpers for constructing [Streaming
// Asynchronous Execution] (SAE) blocks.
//
// [Streaming Asynchronous Execution]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
package blockstest

import (
	"testing"

	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/options"
	"github.com/ava-labs/libevm/params"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/hook/hookstest"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// An EthBlockOption configures the default block properties created by
// [NewEthBlock].
type EthBlockOption = options.Option[ethBlockProperties]

// NewEthBlock constructs a raw Ethereum block with the given arguments.
func NewEthBlock(tb testing.TB, parent *types.Block, txs types.Transactions, opts ...EthBlockOption) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// synchronoous

type ethBlockProperties struct {
	header        *types.Header
	receipts      types.Receipts
	ops           []hookstest.Op
	settledHeight uint64
}

// ModifyHeader returns an option to modify the [types.Header] constructed by
// [NewEthBlock]. It SHOULD NOT modify the `Number` and `ParentHash`, but MAY
// modify any other field.
func ModifyHeader(fn func(*types.Header)) EthBlockOption {
	_ = "STUB: not implemented"
	return *new(EthBlockOption)
}

// WithReceipts returns an option to set the receipts of a block constructed by
// [NewEthBlock].
func WithReceipts(rs types.Receipts) EthBlockOption {
	_ = "STUB: not implemented"
	return *new(EthBlockOption)
}

// WithOps returns an option to set the ops of a block constructed by
// [NewEthBlock].
func WithOps(ops []hookstest.Op) EthBlockOption {
	_ = "STUB: not implemented"
	return *new(EthBlockOption)
}

// A BlockOption configures the default block properties created by [NewBlock].
type BlockOption = options.Option[blockProperties]

// NewBlock constructs an SAE block, wrapping the raw Ethereum block.
func NewBlock(tb testing.TB, eth *types.Block, parent, lastSettled *blocks.Block, opts ...BlockOption) *blocks.Block {
	_ = "STUB: not implemented"
	return nil
}

type blockProperties struct {
	logger logging.Logger
}

// WithLogger overrides the logger passed to [blocks.New] by [NewBlock].
func WithLogger(l logging.Logger) BlockOption { _ = "STUB: not implemented"; return *new(BlockOption) }

// NewGenesis constructs a new [core.Genesis], writes it to the database, and
// returns wraps [core.Genesis.ToBlock] with [NewBlock]. It assumes a nil
// [triedb.Config] unless overridden by a [WithTrieDBConfig]. The block is
// marked as both executed and synchronous.
func NewGenesis(tb testing.TB, db ethdb.Database, xdb saetypes.ExecutionResults, config *params.ChainConfig, alloc types.GenesisAlloc, opts ...GenesisOption) *blocks.Block {
	_ = "STUB: not implemented"
	return nil
}

type genesisConfig struct {
	tdbConfig *triedb.Config
	timestamp uint64
	gasTarget gas.Gas
	gasExcess gas.Gas
}

// A GenesisOption configures [NewGenesis].
type GenesisOption = options.Option[genesisConfig]

// WithTrieDBConfig override the [triedb.Config] used by [NewGenesis].
func WithTrieDBConfig(tc *triedb.Config) GenesisOption {
	_ = "STUB: not implemented"
	return *new(GenesisOption)
}

// WithTimestamp overrides the timestamp used by [NewGenesis].
func WithTimestamp(timestamp uint64) GenesisOption {
	_ = "STUB: not implemented"
	return *new(GenesisOption)
}

// WithGasTarget overrides the gas target used by [NewGenesis].
func WithGasTarget(target gas.Gas) GenesisOption {
	_ = "STUB: not implemented"
	return *new(GenesisOption)
}

// WithGasExcess overrides the gas excess used by [NewGenesis].
func WithGasExcess(excess gas.Gas) GenesisOption {
	_ = "STUB: not implemented"
	return *new(GenesisOption)
}
