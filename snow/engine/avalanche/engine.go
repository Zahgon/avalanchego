// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
)

var (
	_ common.Engine = (*engine)(nil)

	errUnexpectedStart = errors.New("unexpectedly started engine")
)

type engine struct {
	common.AllGetsServer

	// list of NoOpsHandler for messages dropped by engine
	common.StateSummaryFrontierHandler
	common.AcceptedStateSummaryHandler
	common.AcceptedFrontierHandler
	common.AcceptedHandler
	common.AncestorsHandler
	common.PutHandler
	common.QueryHandler
	common.ChitsHandler
	common.AppHandler
	common.InternalHandler
	common.SimplexHandler

	ctx *snow.ConsensusContext
}

func New(
	ctx *snow.ConsensusContext,
	gets common.AllGetsServer,
) common.Engine {
	_ = "STUB: not implemented"
	return *new(common.Engine)
}

func (*engine) Start(context.Context, uint32) error { _ = "STUB: not implemented"; return nil }

func (e *engine) Context() *snow.ConsensusContext { _ = "STUB: not implemented"; return nil }

func (*engine) HealthCheck(context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
