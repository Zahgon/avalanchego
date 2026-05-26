// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/window"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
)

const (
	// MaxRecentlyAcceptedWindowSize is the maximum number of blocks that the
	// recommended minimum height will lag behind the last accepted block.
	MaxRecentlyAcceptedWindowSize = 64
	// MinRecentlyAcceptedWindowSize is the minimum number of blocks that the
	// recommended minimum height will lag behind the last accepted block.
	MinRecentlyAcceptedWindowSize = 0
	// RecentlyAcceptedWindowTTL is the amount of time after a block is accepted
	// to avoid recommending it as the minimum height. The size constraints take
	// precedence over this time constraint.
	RecentlyAcceptedWindowTTL = 30 * time.Second

	validatorSetsCacheSize = 64
)

var (
	_ validators.State = (*Manager)(nil)

	errUnfinalizedHeight = errors.New("failed to fetch validator set at unfinalized height")
)

func NewManager(
	cfg config.Internal,
	state *state.State,
	metrics metrics.Metrics,
	clk *mockable.Clock,
) *Manager {
	_ = "STUB: not implemented"
	return nil
}

// Manager implements [validators.State] and additionally tracks recently accepted block IDs via OnAcceptedBlockID.
// TODO: Remove requirement for the P-chain's context lock to be held when
// calling exported functions.
type Manager struct {
	cfg     config.Internal
	state   *state.State
	metrics metrics.Metrics
	clk     *mockable.Clock

	// Maps caches for each subnet that is currently tracked.
	// Key: Subnet ID
	// Value: cache mapping height -> validator set map
	caches map[ids.ID]cache.Cacher[uint64, map[ids.NodeID]*validators.GetValidatorOutput]

	// sliding window of blocks that were recently accepted
	recentlyAccepted window.Window[ids.ID]
}

// GetMinimumHeight returns the height of the most recent block beyond the
// horizon of our recentlyAccepted window.
//
// Because the time between blocks is arbitrary, we're only guaranteed that
// the window's configured TTL amount of time has passed once an element
// expires from the window.
//
// To try to always return a block older than the window's TTL, we return the
// parent of the oldest element in the window (as an expired element is always
// guaranteed to be sufficiently stale). If we haven't expired an element yet
// in the case of a process restart, we default to the lastAccepted block's
// height which is likely (but not guaranteed) to also be older than the
// window's configured TTL.
//
// If [UseCurrentHeight] is true, we override the block selection policy
// described above and we will always return the last accepted block height
// as the minimum.
func (m *Manager) GetMinimumHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We subtract 1 from the height of [oldest] because we want the height of
// the last block accepted before the [recentlyAccepted] window.
//
// There is guaranteed to be a block accepted before this window because the
// first block added to [recentlyAccepted] window is >= height 1.

func (m *Manager) GetCurrentHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: Pass the context into the state.
func (m *Manager) getCurrentHeight(context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Manager) GetWarpValidatorSets(
	ctx context.Context,
	targetHeight uint64,
) (map[ids.ID]validators.WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we can't flatten the validator set, skip it and disallow warp
// message verification from this subnet.

func (m *Manager) GetValidatorSet(
	ctx context.Context,
	targetHeight uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get the start time to track metrics

// cache the validator set

func (m *Manager) getValidatorSetCache(subnetID ids.ID) cache.Cacher[uint64, map[ids.NodeID]*validators.GetValidatorOutput] {
	_ = "STUB: not implemented"
	// Only cache tracked subnets
	return nil
}

func (m *Manager) makeAllValidatorSets(
	ctx context.Context,
	targetHeight uint64,
) (map[ids.ID]map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rebuild subnet validators at [targetHeight]
//
// Note: Since we are attempting to generate the validator set at
// [targetHeight], we want to apply the diffs from
// (targetHeight, currentHeight]. Because the state interface is implemented
// to be inclusive, we apply diffs in [targetHeight + 1, currentHeight].

func (m *Manager) makeValidatorSet(
	ctx context.Context,
	targetHeight uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Rebuild subnet validators at [targetHeight]
//
// Note: Since we are attempting to generate the validator set at
// [targetHeight], we want to apply the diffs from
// (targetHeight, currentHeight]. Because the state interface is implemented
// to be inclusive, we apply diffs in [targetHeight + 1, currentHeight].

func (m *Manager) getAllCurrentValidatorSets(
	ctx context.Context,
) (map[ids.ID]map[ids.NodeID]*validators.GetValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *Manager) getCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *Manager) GetSubnetID(_ context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// OnAcceptedBlockID registers the ID of the latest accepted block.
// It is used to update the recentlyAccepted sliding window.
func (m *Manager) OnAcceptedBlockID(blkID ids.ID) { _ = "STUB: not implemented"; return }

func (m *Manager) GetCurrentValidatorSet(ctx context.Context, subnetID ids.ID) (map[ids.ID]*validators.GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
