// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"context"
	"errors"
	"time"

	"gonum.org/v1/gonum/mathext/prng"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/sampler"
)

// Proposer list constants
const (
	WindowDuration = 5 * time.Second

	MaxVerifyWindows = 6
	MaxVerifyDelay   = MaxVerifyWindows * WindowDuration // 30 seconds

	MaxBuildWindows = 60
	MaxBuildDelay   = MaxBuildWindows * WindowDuration // 5 minutes

	MaxLookAheadSlots  = 720
	MaxLookAheadWindow = MaxLookAheadSlots * WindowDuration // 1 hour
)

var (
	_ Windower = (*windower)(nil)

	ErrAnyoneCanPropose         = errors.New("anyone can propose")
	ErrUnexpectedSamplerFailure = errors.New("unexpected sampler failure")
)

type Windower interface {
	// Proposers returns the proposer list for building a block at [blockHeight]
	// when the validator set is defined at [pChainHeight]. The list is returned
	// in order. The minimum delay of a validator is the index they appear times
	// [WindowDuration].
	Proposers(
		ctx context.Context,
		blockHeight,
		pChainHeight uint64,
		maxWindows int,
	) ([]ids.NodeID, error)

	// Delay returns the amount of time that [validatorID] must wait before
	// building a block at [blockHeight] when the validator set is defined at
	// [pChainHeight].
	Delay(
		ctx context.Context,
		blockHeight,
		pChainHeight uint64,
		validatorID ids.NodeID,
		maxWindows int,
	) (time.Duration, error)

	// In the Post-Durango windowing scheme, every validator active at
	// [pChainHeight] gets specific slots it can propose in (instead of being
	// able to propose from a given time on as it happens Pre-Durango).
	// [ExpectedProposer] calculates which nodeID is scheduled to propose a
	// block of height [blockHeight] at [slot].
	// If no validators are currently available, [ErrAnyoneCanPropose] is
	// returned.
	ExpectedProposer(
		ctx context.Context,
		blockHeight,
		pChainHeight,
		slot uint64,
	) (ids.NodeID, error)

	// In the Post-Durango windowing scheme, every validator active at
	// [pChainHeight] gets specific slots it can propose in (instead of being
	// able to propose from a given time on as it happens Pre-Durango).
	// [MinDelayForProposer] specifies how long [nodeID] needs to wait for its
	// slot to start. Delay is specified as starting from slot zero start.
	// (which is parent timestamp). For efficiency reasons, we cap the slot
	// search to [MaxLookAheadSlots].
	// If no validators are currently available, [ErrAnyoneCanPropose] is
	// returned.
	MinDelayForProposer(
		ctx context.Context,
		blockHeight,
		pChainHeight uint64,
		nodeID ids.NodeID,
		startSlot uint64,
	) (time.Duration, error)
}

// windower interfaces with P-Chain and it is responsible for calculating the
// delay for the block submission window of a given validator
type windower struct {
	state       validators.State
	subnetID    ids.ID
	chainSource uint64
	logger      logging.Logger
}

func New(state validators.State, subnetID, chainID ids.ID, logger logging.Logger) Windower {
	_ = "STUB: not implemented"
	return *new(Windower)
}

func (w *windower) Proposers(ctx context.Context, blockHeight, pChainHeight uint64, maxWindows int) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	// Note: The 32-bit prng is used here for legacy reasons. All other usages
	// of a prng in this file should use the 64-bit version.
	return nil, nil
}

func (w *windower) Delay(ctx context.Context, blockHeight, pChainHeight uint64, validatorID ids.NodeID, maxWindows int) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (w *windower) ExpectedProposer(
	ctx context.Context,
	blockHeight,
	pChainHeight,
	slot uint64,
) (ids.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), nil
}

func (w *windower) MinDelayForProposer(
	ctx context.Context,
	blockHeight,
	pChainHeight uint64,
	nodeID ids.NodeID,
	startSlot uint64,
) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// no slots scheduled for the max window we inspect. Return max delay

func (w *windower) makeSampler(
	ctx context.Context,
	pChainHeight uint64,
	source sampler.Source,
) (sampler.WeightedWithoutReplacement, []validatorData, error) {
	_ = "STUB: not implemented"
	// Get the canonical representation of the validator set at the provided
	// p-chain height.
	return *new(sampler.WeightedWithoutReplacement), nil, nil
}

// Ignore inactive ACP-77 validators.

// Note: validators are sorted by ID. Sorting by weight would not create a
// canonically sorted list.

func (w *windower) expectedProposer(
	validators []validatorData,
	source *prng.MT19937_64,
	sampler sampler.WeightedWithoutReplacement,
	blockHeight,
	slot uint64,
) (ids.NodeID, error) {
	_ = "STUB: not implemented"
	// Slot is reversed to utilize a different state space in the seed than the
	// height. If the slot was not reversed the state space would collide;
	// biasing the seed generation. For example, without reversing the slot
	// height=0 and slot=1 would equal height=1 and slot=0.
	return *new(ids.NodeID), nil
}

func TimeToSlot(start, now time.Time) uint64 { _ = "STUB: not implemented"; return 0 }
