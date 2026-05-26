// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

type addValidatorRules struct {
	assetID           ids.ID
	minValidatorStake uint64
	maxValidatorStake uint64
	minStakeDuration  time.Duration
	maxStakeDuration  time.Duration
	minDelegationFee  uint32
}

func getValidatorRules(
	backend *Backend,
	chainState state.Chain,
	subnetID ids.ID,
) (*addValidatorRules, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type addDelegatorRules struct {
	assetID                  ids.ID
	minDelegatorStake        uint64
	maxValidatorStake        uint64
	minStakeDuration         time.Duration
	maxStakeDuration         time.Duration
	maxValidatorWeightFactor byte
}

func getDelegatorRules(
	backend *Backend,
	chainState state.Chain,
	subnetID ids.ID,
) (*addDelegatorRules, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetValidator returns information about the given validator, which may be a
// current validator or pending validator.
func GetValidator(state state.Chain, subnetID ids.ID, nodeID ids.NodeID) (*state.Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This node is currently validating the subnet.

// Unexpected error occurred.

// overDelegated returns true if [validator] will be overdelegated when adding [delegator].
//
// A [validator] would become overdelegated if:
// - the maximum total weight on [validator] exceeds [weightLimit]
func overDelegated(
	state state.Chain,
	validator *state.Staker,
	weightLimit uint64,
	delegatorWeight uint64,
	delegatorStartTime time.Time,
	delegatorEndTime time.Time,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetMaxWeight returns the maximum total weight of the [validator], including
// its own weight, between [startTime] and [endTime].
// The weight changes are applied in the order they will be applied as chain
// time advances.
// Invariant:
// - [validator.StartTime] <= [startTime] < [endTime] <= [validator.EndTime]
func GetMaxWeight(
	chainState state.Chain,
	validator *state.Staker,
	startTime time.Time,
	endTime time.Time,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: We can optimize this by moving the current total weight to be
//       stored in the validator state.
//
// Calculate the current total weight on this validator, including the
// weight of the actual validator and the sum of the weights of all of the
// currently active delegators.

// Iterate over the future stake weight changes and calculate the maximum
// total weight on the validator, only including the points in the time
// range [startTime, endTime].

// [delegator.NextTime] > [endTime]

// This delegation change (and all following changes) occurs after
// [endTime]. Since we're calculating the max amount staked in
// [startTime, endTime], we can stop.

// [delegator.NextTime] >= [startTime]

// We have advanced time to be at the inside of the delegation
// window. Make sure that the max weight is updated accordingly.

// Because we assume [startTime] < [endTime], we have advanced time to
// be at the end of the delegation window. Make sure that the max weight is
// updated accordingly.

func GetTransformSubnetTx(chain state.Chain, subnetID ids.ID) (*txs.TransformSubnetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
