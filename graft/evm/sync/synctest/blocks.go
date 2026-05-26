// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package synctest

import (
	"testing"

	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
)

var (
	// testKey is a pre-generated private key for test transactions.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	// testAddr is the address derived from testKey.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)
)

// BlockGeneratorConfig configures block generation.
type BlockGeneratorConfig struct {
	// TxDataSize sets a fixed size for transaction data in each block.
	// If TxDataSizeFunc is set, this field is ignored.
	TxDataSize int

	// TxDataSizeFunc returns the transaction data size for a given block index.
	// Block index 0 is the genesis block (which has no transactions).
	// If nil, TxDataSize is used for all blocks.
	TxDataSizeFunc func(blockIndex int) int

	// GasLimit sets the gas limit for generated blocks.
	// Defaults to params.GenesisGasLimit if zero.
	GasLimit uint64
}

// GenerateTestBlocks creates a chain of test blocks for sync testing.
// Returns a slice of blocks where blocks[0] is the genesis block.
//
// Each non-genesis block contains a single signed transaction to ensure
// unique block hashes. The blocks are minimal but structurally valid
// for testing block sync functionality.
func GenerateTestBlocks(t *testing.T, numBlocks int, cfg *BlockGeneratorConfig) []*types.Block {
	_ = "STUB: not implemented"
	return nil
}

func newGenesisBlock(gasLimit uint64) *types.Block { _ = "STUB: not implemented"; return nil }

func newBlock(t *testing.T, parent *types.Block, blockNum, txDataSize int, gasLimit uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// 10 seconds between blocks

func makeTxData(blockNum, size int) []byte { _ = "STUB: not implemented"; return nil }

func newSignedTx(t *testing.T, blockNum int, txData []byte) *types.Transaction {
	_ = "STUB: not implemented"

	// Use non-zero gas price to ensure gas cost is calculated.
	// Calldata gas: 16 per non-zero byte, 4 per zero byte.
	// We use 16 as upper bound since makeTxData produces non-zero bytes.
	return nil
}

// nonce
