// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracker

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ Targeter = (*targeter)(nil)

type Targeter interface {
	// Returns the target usage of the given node.
	TargetUsage(nodeID ids.NodeID) float64
}

type TargeterConfig struct {
	// VdrAlloc is the amount of the resource to split over validators, weighted
	// by stake.
	VdrAlloc float64 `json:"vdrAlloc"`

	// MaxNonVdrUsage is the amount of the resource which, if utilized, will
	// result in allocations being based only on the stake weighted allocation.
	MaxNonVdrUsage float64 `json:"maxNonVdrUsage"`

	// MaxNonVdrNodeUsage is the amount of the resource to allocate to a node
	// before adding the stake weighted allocation.
	MaxNonVdrNodeUsage float64 `json:"maxNonVdrNodeUsage"`
}

func NewTargeter(
	logger logging.Logger,
	config *TargeterConfig,
	vdrs validators.Manager,
	tracker Tracker,
) Targeter {
	_ = "STUB: not implemented"
	return *new(Targeter)
}

type targeter struct {
	vdrs               validators.Manager
	log                logging.Logger
	tracker            Tracker
	vdrAlloc           float64
	maxNonVdrUsage     float64
	maxNonVdrNodeUsage float64
}

func (t *targeter) TargetUsage(nodeID ids.NodeID) float64 {
	_ = "STUB: not implemented"
	// This node's at-large allocation is min([remaining at large], [max at large for a given peer])
	return 0
}

// This node gets a stake-weighted portion of the validator allocation.
