// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package worstcase provides the worst-case balance and nonce tracking needed
// to safely include transactions that are guaranteed to be valid during
// execution.
package worstcase

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/params"
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/gastime"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/saedb"

	saeparams "github.com/ava-labs/avalanchego/vms/saevm/params"
)

// State tracks the worst-case gas price and account state as operations are
// executed.
//
// Usage of [State] must follow the pattern:
//  1. [State.StartBlock] for each block to be included.
//  2. [State.GasLimit] and [State.BaseFee] to query the block's parameters.
//  3. [State.ApplyTx] or [State.Apply] for each [types.Transaction] or
//     [hook.Op] to include in the block, respectively.
//  4. [State.GasUsed] to query the total gas used in the block.
//  5. [State.FinishBlock] to finalize the block's gas time.
//  6. Repeat from step 1 for the next block.
type State struct {
	hooks  hook.Points
	config *params.ChainConfig

	db    *state.StateDB
	clock *gastime.Time
	// expectedParentHash is used to sanity check that blocks are provided in
	// order. The [types.Header] in the `curr` field is modified to reflect
	// worst-case bounds (which will almost certainly differ from actual values
	// when replaying historical blocks) so its hash can't be used.
	expectedParentHash common.Hash

	qSize, blockSize, maxBlockSize gas.Gas

	baseFee             *uint256.Int
	curr                *types.Header
	signer              types.Signer
	minOpBurnerBalances []map[common.Address]*uint256.Int
}

var errSettledBlockNotExecuted = errors.New("block marked for settling has not finished execution yet")

// NewState constructs a new worst-case state on top of the settled block.
func NewState(
	hooks hook.Points,
	config *params.ChainConfig,
	settled *blocks.Block,
	opener saedb.StateDBOpener,
) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	maxGasSecondsPerBlock = saeparams.TauSeconds * saeparams.Lambda
)

var (
	errNonConsecutiveBlocks = errors.New("non-consecutive blocks")
	// ErrQueueFull is returned by [State.StartBlock] if the queue is not able
	// to accept new blocks. This can be rectified by building a block at a
	// later time such that additional blocks are settled and the queue is
	// sufficiently drained.
	ErrQueueFull = errors.New("queue exceeds gas threshold for new block")
)

// StartBlock updates the worst-case state to the beginning of the provided
// block.
//
// It is not necessary for [types.Header.GasLimit] nor [types.Header.BaseFee] to
// be set. However, all other fields should be populated and
// [types.Header.ParentHash] must match the previous block's hash.
//
// If the queue is too full to accept another block, [ErrQueueFull] is returned.
func (s *State) StartBlock(h *types.Header) error { _ = "STUB: not implemented"; return nil }

// [State.FinishBlock] returns a clone so we can reuse the alloc here

// expectedParentHash is updated prior to modifying the GasLimit and BaseFee
// to ensure that historical block hashes are not modified.

// We MUST use the block's timestamp, not the execution clock's, otherwise
// we might enable an upgrade too early.

// safeMaxBlockSize returns the maximum block size for the clock's rate,
// possibly capping it so a full closed queue still fits in [gas.Gas]. At the
// time of writing, the cap is ~6e17, so capping is exceedingly unlikely.
func safeMaxBlockSize(clock *gastime.Time) gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

// GasLimit returns the available gas limit for the current block.
func (s *State) GasLimit() uint64 { _ = "STUB: not implemented"; return 0 }

// BaseFee returns the worst-case base fee for the current block.
func (s *State) BaseFee() *uint256.Int { _ = "STUB: not implemented"; return nil }

var errCostOverflow = errors.New("Cost() overflows uint256")

// ApplyTx validates the transaction both intrinsically and in the context of
// worst-case gas assumptions of all previous operations. This provides an upper
// bound on the total cost of the transaction such that a nil error returned by
// ApplyTx guarantees that the sender of the transaction will have sufficient
// balance to cover its costs if consensus accepts the same operation set
// (and order) as was applied.
//
// If the transaction can not be applied, an error is returned and the state is
// not modified.
//
// TODO: Consider exporting txToOp and expecting users to call Apply directly.
func (s *State) ApplyTx(tx *types.Transaction) error { _ = "STUB: not implemented"; return nil }

// No byte-size limit needed as gas validation (intrinsic gas ≤
// tx gas ≤ block gas limit) already enforces an implicit size
// limit (2MB) on transactions.

// While EOA enforcement is not possible to guarantee in worst-case
// execution, we can prevent most cases here.
//
// TODO: We must still handle non-EOA issuance later during actual
// execution.

func bigToUint256(v *big.Int) (_ uint256.Int, overflow bool) {
	_ = "STUB: not implemented"
	return *new(uint256.Int), false
}

// mulAdd returns a*b + c and reports whether overflow occurred.
func mulAdd(a uint64, b, c *uint256.Int) (_ uint256.Int, overflow bool) {
	_ = "STUB: not implemented"
	return *new(uint256.Int), false
}

func txToOp(from common.Address, tx *types.Transaction, baseFee *uint256.Int) (hook.Op, error) {
	_ = "STUB: not implemented"
	// for convenience when returning zero value
	return *new(hook.Op), nil
}

// effectiveGasPrice = min(gasFeeCap, baseFee + gasTipCap)

// Mint MUST NOT be populated here because this transaction may revert.

// Apply attempts to apply the operation to this state.
//
// If the operation can not be applied, an error is returned and the state is
// not modified.
//
// Operations are invalid if any of the following are true:
//
//   - The operation consumes more gas than the block has available.
//   - The operation specifies too low of a gas price.
//   - The operation is from an account with an incorrect or invalid nonce.
//   - The operation is from an account with an insufficient balance.
func (s *State) Apply(o hook.Op) error { _ = "STUB: not implemented"; return nil }

// MUST be before `o.ApplyTo()` to mirror [saexec.Executor] check

// GasUsed returns the gas used for the current block.
func (s *State) GasUsed() uint64 { _ = "STUB: not implemented"; return 0 }

// FinishBlock advances the [gastime.Time] in preparation for the next block.
//
// The returned bounds assume that every non-nil error from [State.ApplyTx]
// resulted in said transaction being included, which is reflected in the
// indexing of tx-sender balances.
func (s *State) FinishBlock() (*blocks.WorstCaseBounds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
