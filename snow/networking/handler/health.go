// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

var ErrNotConnectedEnoughStake = errors.New("not connected to enough stake")

func (h *handler) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *handler) networkHealthCheck() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SnowParameters can be nil if the subnet is configured for Simplex consensus

func (h *handler) getDisconnectedValidators() set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	return nil
}

// vdrs - connectedVdrs is equal to the disconnectedVdrs
