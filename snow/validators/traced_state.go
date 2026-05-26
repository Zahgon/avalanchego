// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
)

var _ State = (*tracedState)(nil)

type tracedState struct {
	s                         State
	getMinimumHeightTag       string
	getCurrentHeightTag       string
	getSubnetIDTag            string
	getWarpValidatorSetsTag   string
	getValidatorSetTag        string
	getCurrentValidatorSetTag string
	tracer                    trace.Tracer
}

func Trace(s State, name string, tracer trace.Tracer) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (s *tracedState) GetMinimumHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *tracedState) GetCurrentHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *tracedState) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (s *tracedState) GetWarpValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]WarpSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *tracedState) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*GetValidatorOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *tracedState) GetCurrentValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) (map[ids.ID]*GetCurrentValidatorOutput, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
