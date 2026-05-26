// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorstest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
)

var (
	errMinimumHeight          = errors.New("unexpectedly called GetMinimumHeight")
	errCurrentHeight          = errors.New("unexpectedly called GetCurrentHeight")
	errSubnetID               = errors.New("unexpectedly called GetSubnetID")
	errGetWarpValidatorSets   = errors.New("unexpectedly called GetWarpValidatorSets")
	errGetValidatorSet        = errors.New("unexpectedly called GetValidatorSet")
	errGetCurrentValidatorSet = errors.New("unexpectedly called GetCurrentValidatorSet")
)

var _ validators.State = (*State)(nil)

type State struct {
	T testing.TB

	CantGetMinimumHeight,
	CantGetCurrentHeight,
	CantGetSubnetID,
	CantGetWarpValidatorSets,
	CantGetValidatorSet,
	CantGetCurrentValidatorSet bool

	GetMinimumHeightF       func(ctx context.Context) (uint64, error)
	GetCurrentHeightF       func(ctx context.Context) (uint64, error)
	GetSubnetIDF            func(ctx context.Context, chainID ids.ID) (ids.ID, error)
	GetWarpValidatorSetsF   func(ctx context.Context, height uint64) (map[ids.ID]validators.WarpSet, error)
	GetValidatorSetF        func(ctx context.Context, height uint64, subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error)
	GetCurrentValidatorSetF func(ctx context.Context, subnetID ids.ID) (map[ids.ID]*validators.GetCurrentValidatorOutput, uint64, error)
}

func (vm *State) GetMinimumHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (vm *State) GetCurrentHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (vm *State) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *State) GetWarpValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]validators.WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *State) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *State) GetCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) (map[ids.ID]*validators.GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
