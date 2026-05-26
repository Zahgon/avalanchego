// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"context"
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
)

const (
	validatorSetsCacheSize = 8
	subnetIDsCacheSize     = 4096 // At 64 bytes per entry, this is ~256 KB
)

var (
	_ State = (*lockedState)(nil)
	_ State = (*cachedState)(nil)

	ErrUnfinalizedHeight = errors.New("validator set height not yet finalized")
)

// State allows the lookup of validator sets on specified subnets at the
// requested P-chain height.
type State interface {
	// GetMinimumHeight returns the minimum height of the block still in the
	// proposal window.
	GetMinimumHeight(context.Context) (uint64, error)
	// GetCurrentHeight returns the current height of the P-chain.
	GetCurrentHeight(context.Context) (uint64, error)

	// GetSubnetID returns the subnetID of the provided chain.
	GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error)

	// GetWarpValidatorSets returns the canonical warp validator set for all
	// subnets at the requested P-chain height.
	//
	// If a subnet is not present in the returned map, that indicates that the
	// subnet is not currently able to produce valid warp message signatures.
	//
	// The returned map should not be modified.
	GetWarpValidatorSets(ctx context.Context, height uint64) (map[ids.ID]WarpSet, error)

	// GetValidatorSet returns the validators of the provided subnet at the
	// requested P-chain height.
	// The returned map should not be modified.
	GetValidatorSet(
		ctx context.Context,
		height uint64,
		subnetID ids.ID,
	) (map[ids.NodeID]*GetValidatorOutput, error)

	// GetCurrentValidatorSet returns the current validators of the provided subnet
	// and the current P-Chain height.
	// Map is keyed by ValidationID.
	GetCurrentValidatorSet(
		ctx context.Context,
		subnetID ids.ID,
	) (map[ids.ID]*GetCurrentValidatorOutput, uint64, error)
}

type lockedState struct {
	lock sync.Locker
	s    State
}

func NewLockedState(lock sync.Locker, s State) State { _ = "STUB: not implemented"; return *new(State) }

func (s *lockedState) GetMinimumHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *lockedState) GetCurrentHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *lockedState) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (s *lockedState) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *lockedState) GetWarpValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *lockedState) GetCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) (map[ids.ID]*GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type noValidators struct {
	State
}

func NewNoValidatorsState(state State) State { _ = "STUB: not implemented"; return *new(State) }

func (*noValidators) GetWarpValidatorSets(context.Context, uint64) (map[ids.ID]WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*noValidators) GetValidatorSet(context.Context, uint64, ids.ID) (map[ids.NodeID]*GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *noValidators) GetCurrentValidatorSet(ctx context.Context, _ ids.ID) (map[ids.ID]*GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type cachedState struct {
	State

	// Caches the subnet ID for given blockchain IDs.
	// Key: blockchain ID
	// Value: subnet ID
	subnetIDsCache cache.Cacher[ids.ID, ids.ID]

	// Caches validators for all subnets at various heights.
	// Key: height
	// Value: mapping subnet ID -> validator set
	validatorSetsCache cache.Cacher[uint64, map[ids.ID]WarpSet]
}

func NewCachedState(state State) State { _ = "STUB: not implemented"; return *new(State) }

func (c *cachedState) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (c *cachedState) GetWarpValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
