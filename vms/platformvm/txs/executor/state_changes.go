// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/validators/fee"
)

var (
	ErrChildBlockEarlierThanParent     = errors.New("proposed timestamp before current chain time")
	ErrChildBlockAfterStakerChangeTime = errors.New("proposed timestamp later than next staker change time")
	ErrChildBlockBeyondSyncBound       = errors.New("proposed timestamp is too far in the future relative to local time")
)

// VerifyNewChainTime returns nil if the [newChainTime] is a valid chain time.
// Requires:
//   - [newChainTime] >= [currentChainTime]: to ensure chain time advances
//     monotonically.
//   - [newChainTime] <= [now] + [SyncBound]: to ensure chain time approximates
//     "real" time.
//   - [newChainTime] <= [nextStakerChangeTime]: so that no staking set changes
//     are skipped.
func VerifyNewChainTime(
	config fee.Config,
	newChainTime time.Time,
	now time.Time,
	currentState state.Chain,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Only allow timestamp to be reasonably far forward

// nextStakerChangeTime is calculated last to ensure that the function is
// able to be calculated efficiently.

// Only allow timestamp to move as far forward as the time of the next
// staker set change

// AdvanceTimeTo applies all state changes to [parentState] resulting from
// advancing the chain time to [newChainTime].
//
// Returns true iff the validator set changed.
func AdvanceTimeTo(
	backend *Backend,
	parentState state.Chain,
	newChainTime time.Time,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// advanceTimeTo returns the state diff on top of parentState resulting from
// advancing the chain time to newChainTime. It also returns a boolean
// indicating if the validator set changed.
//
// parentState is not modified.
func advanceTimeTo(
	backend *Backend,
	parentState state.Chain,
	newChainTime time.Time,
) (*state.Diff, bool, error) {
	_ = "STUB: not implemented"
	// We promote pending stakers to current stakers first and remove
	// completed stakers from the current staker set. We assume that any
	// promoted staker will not immediately be removed from the current staker
	// set. This is guaranteed by the following invariants.
	//
	// Invariant: MinStakeDuration > 0 => guarantees [StartTime] != [EndTime]
	// Invariant: [newChainTime] <= nextStakerChangeTime.
	return nil, false, nil
}

// Promote any pending stakers to current if [StartTime] <= [newChainTime].
//
// Pending validators are promoted before pending delegators so that when we attempt to promote a current delegator,
// the current validator exists to satisfy a defensive check that a current delegator being added must have an
// existing current validator.
//
// Pending stakers are ordered such that ties in Staker.NextTime are broken by Staker.Priority. It is possible for a
// tie to result in an Apricot pending delegator being ordered before a pending validator, so we must do two passes
// to promote pending validators before pending delegators to respect the expected ordering. Rewards however must be
// performed in the original staker iterator ordering to respect the historical ordering.
//
// Invariant: It is not safe to modify the state while iterating over it,
// so we use the parentState's iterator rather than the changes iterator.
// ParentState must not be modified before this iterator is released.

// Only permissionless networks (including the primary network) are eligible for rewards

// Invariant: reward.Calculator.Calculate can never return a potentialReward
//            such that supply + potentialReward > maximumSupply.

// Buffer state changes so that we can perform validator updates before delegator updates to respect state's
// expected order of operations.

// Remove any current stakers whose [EndTime] <= [newChainTime].
//
// Invariant: It is not safe to modify the state while iterating over it,
// so we use the parentState's iterator rather than the changes iterator.
// ParentState must not be modified before this iterator is released.

// Invariant: Permissioned stakers are encountered first for a given
//            timestamp because their priority is the smallest.

// Permissionless stakers are removed by the RewardValidatorTx, not
// an AdvanceTimeTx.

// Calculate number of seconds the time is advancing

// Remove all expiries whose timestamp now implies they can never be re-issued.
//
// The expiry timestamp is the time at which it is no longer valid, so any
// expiry with a timestamp less than or equal to the new chain time can be
// removed.
//
// Ref: https://github.com/avalanche-foundation/ACPs/tree/e333b335c34c8692d84259d21bd07b2bb849dc2c/ACPs/77-reinventing-subnets#registerl1validatortx
func removeStaleExpiries(
	parentState state.Chain,
	changes *state.Diff,
	newChainTimeUnix uint64,
) error {
	_ = "STUB: not implemented"
	// Invariant: It is not safe to modify the state while iterating over it, so
	// we use the parentState's iterator rather than the changes iterator.
	// ParentState must not be modified before this iterator is released.
	return nil
}

// The expiry iterator is sorted in order of increasing timestamp. Once
// we find a non-expired expiry, we can break.

func advanceDynamicFeeState(
	config gas.Config,
	changes *state.Diff,
	seconds uint64,
) {
	_ = "STUB: not implemented"
	return
}

// advanceValidatorFeeState advances the validator fee state by [seconds]. L1
// validators are read from [parentState] and written to [changes] to avoid
// modifying state while an iterator is held.
func advanceValidatorFeeState(
	config fee.Config,
	parentState state.Chain,
	changes *state.Diff,
	seconds uint64,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Invariant: It is not safe to modify the state while iterating over it,
// so we use the parentState's iterator rather than the changes iterator.
// ParentState must not be modified before this iterator is released.

// GetActiveL1ValidatorsIterator iterates in order of increasing
// EndAccumulatedFee, so we can break early.

// Deactivate the validator

func GetRewardsCalculator(
	backend *Backend,
	parentState state.Chain,
	subnetID ids.ID,
) (reward.Calculator, error) {
	_ = "STUB: not implemented"
	return *new(reward.Calculator), nil
}
