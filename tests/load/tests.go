// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"sync"

	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/load/contracts"
	"github.com/ava-labs/avalanchego/utils/sampler"
)

var maxFeeCap = big.NewInt(300000000000)

// NewRandomTest creates a RandomWeightedTest containing a collection of EVM
// load testing scenarios.
//
// This function handles the setup of the tests and also assigns each test
// a weight based on its C-Chain frequency and computational intensity.
func NewRandomTest(
	ctx context.Context,
	chainID *big.Int,
	worker *Worker,
	source rand.Source,
	tokenContract *contracts.ERC20,
) (*RandomWeightedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// value specifies the amount to send in a transfer test

// Random values are written to slots to ensure that the same value isn't
// being written to a slot, as the odds of a pRNG choosing the same value
// twice is practically zero. Using random values simplifies gas calculations
// as it removes the need to use an SLOAD operation to verify a different
// value is being written.
//#nosec G404
//#nosec G404

// minimum gas used: 21_000

// minimum gas used: 84_000

// minimum gas used: 242_000

// minimum gas used: 61_000

// minimum gas used: 302_100

// minimum gas used: 290_000

// minimum gas used: 23_000

// minimum gas used: 155_900

// minimum gas used: 52_300

type RandomWeightedTest struct {
	tests       []Test
	weighted    sampler.Weighted
	totalWeight int64

	mu   sync.Mutex
	rand *rand.Rand
}

func NewRandomWeightedTest(
	weightedTests []WeightedTest,
	source rand.Source,
) (*RandomWeightedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize weighted set

//#nosec G404

func (r *RandomWeightedTest) Run(tc tests.TestContext, wallet *Wallet) {
	_ = "STUB: not implemented"
	return
}

type WeightedTest struct {
	Test   Test
	Weight uint64
}

type TransferTest struct {
	Value *big.Int
}

func (t TransferTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

// Generate non-existent account address

type ReadTest struct {
	Contract *contracts.LoadSimulator
	Offset   *big.Int
	NumSlots *big.Int
}

func (r ReadTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

type WriteTest struct {
	Contract *contracts.LoadSimulator
	NumSlots *big.Int

	mu   sync.Mutex
	Rand *rand.Rand
}

func (w *WriteTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

type ModifyTest struct {
	Contract *contracts.LoadSimulator
	NumSlots *big.Int

	mu   sync.Mutex
	Rand *rand.Rand
}

func (m *ModifyTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

type HashTest struct {
	Contract      *contracts.LoadSimulator
	Value         *big.Int
	NumIterations *big.Int
}

func (h HashTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

type DeployTest struct {
	Contract *contracts.LoadSimulator
}

func (d DeployTest) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

type LargeCalldataTest struct {
	Contract *contracts.LoadSimulator
	Calldata []byte
}

func (l LargeCalldataTest) Run(tc tests.TestContext, wallet *Wallet) {
	_ = "STUB: not implemented"
	return
}

type TrieStressTest struct {
	Contract  *contracts.TrieStressTest
	NumValues *big.Int
}

func (t TrieStressTest) Run(tc tests.TestContext, wallet *Wallet) {
	_ = "STUB: not implemented"
	return
}

type ERC20Test struct {
	Contract *contracts.ERC20
	Value    *big.Int
}

func (e ERC20Test) Run(tc tests.TestContext, wallet *Wallet) { _ = "STUB: not implemented"; return }

// Generate non-existent account address

func executeContractTx(
	tc tests.TestContext,
	wallet *Wallet,
	txFunc func(*bind.TransactOpts) (*types.Transaction, error),
) {
	_ = "STUB: not implemented"
	return
}

// newTxOpts returns transactions options for contract calls, with sending disabled
func newTxOpts(
	key *ecdsa.PrivateKey,
	chainID *big.Int,
	maxFeeCap *big.Int,
	nonce uint64,
) (*bind.TransactOpts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
