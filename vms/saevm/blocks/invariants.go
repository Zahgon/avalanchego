// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/holiman/uint256"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/gastime"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
)

// WorstCaseBounds define the limits of certain values, predicted by the block
// builder, that a [Block] will encounter when eventually executed.
type WorstCaseBounds struct {
	MaxBaseFee *uint256.Int
	// LatestEndTime is the worst-case [gastime.Time] after this block's gas has been
	// consumed and the target updated. Its [gastime.Time.BaseFee] is an upper
	// bound on the next block's base fee because the next block's
	// [gastime.Time.BeforeBlock] can only reduce the excess.
	LatestEndTime *gastime.Time
	// Invariant: keys of individual maps MUST be identical to those of the
	// respective [hook.Op.Burn] map. For transaction-derived Ops, there is
	// always 1 entry.
	MinOpBurnerBalances []map[common.Address]*uint256.Int
}

// SetWorstCaseBounds sets the bounds, which MUST be done before execution.
func (b *Block) SetWorstCaseBounds(lim *WorstCaseBounds) {
	_ = "STUB: not implemented"

	// WorstCaseBounds returns the argument passed to [Block.SetWorstCaseBounds].
	return
}

func (b *Block) WorstCaseBounds() *WorstCaseBounds {
	_ = "STUB: not implemented"

	// CheckBaseFeeBound logs at ERROR if the `actual` base fee is greater than the
	// predicted upper bound passed to [Block.SetWorstCaseBounds].
	//
	// Such a violation, while potentially critical, might not result in failed
	// execution so no error is returned and execution MUST continue optimistically.
	// Any such log in development will cause tests to fail.
	return nil
}

func (b *Block) CheckBaseFeeBound(actual *uint256.Int) { _ = "STUB: not implemented"; return }

// Coverage visualisation

// CheckSenderBalanceBound logs at ERROR if the balance of the `tx` sender is
// less than the predicted lower bound passed to [Block.SetWorstCaseBounds].
// [state.StateDB.SetTxContext] MUST have already been called.
//
// Such a violation, while potentially critical, might not result in failed
// execution so no error is returned and execution MUST continue optimistically.
// Any such log in development will cause tests to fail.
func (b *Block) CheckSenderBalanceBound(stateDB *state.StateDB, signer types.Signer, tx *types.Transaction) {
	_ = "STUB: not implemented"
	return
}

// CheckOpBurnerBalanceBounds is equivalent to [Block.CheckSenderBalanceBound],
// performed for every address in [hook.Op.Burn] instead of only for a single
// transaction sender.
//
// For the purposes of calculating the Op's index in the block, a
// [types.Transaction] is also considered to be an Op.
func (b *Block) CheckOpBurnerBalanceBounds(stateDB *state.StateDB, opIndexInBlock int, op hook.Op) {
	_ = "STUB: not implemented"
	return
}

func (b *Block) checkBalanceBounds(log logging.Logger, stateDB *state.StateDB, opIndexInBlock int, accounts ...common.Address) {
	_ = "STUB: not implemented"
	return
}

// Coverage visualisation

// A LifeCycleStage defines the progression of a block from acceptance through
// to settlement.
type LifeCycleStage int

// Valid [LifeCycleStage] values. Blocks proceed in increasing stage numbers,
// but specific values MUST NOT be relied upon to be stable.
const (
	NotExecuted LifeCycleStage = iota
	Executed
	Settled

	Accepted = NotExecuted
)

func (b *Block) brokenInvariantErr(msg string) error { _ = "STUB: not implemented"; return nil }

// CheckInvariants checks internal invariants against expected stage, typically
// only used during database recovery.
func (b *Block) CheckInvariants(expect LifeCycleStage) error { _ = "STUB: not implemented"; return nil }

// not executed

// executed

// settled

// not settled
