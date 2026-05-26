// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var _ Engine = (*tracedEngine)(nil)

type tracedEngine struct {
	engine Engine
	tracer trace.Tracer
}

func TraceEngine(engine Engine, tracer trace.Tracer) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

func (e *tracedEngine) GetStateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) StateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, summary []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetStateSummaryFrontierFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, heights set.Set[uint64]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAcceptedStateSummaryFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAcceptedFrontierFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAccepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Accepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAcceptedFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Ancestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containers [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetAncestorsFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Get(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Put(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) GetFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) PullQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) PushQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Chits(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) QueryFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) AppGossip(ctx context.Context, nodeID ids.NodeID, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *tracedEngine) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *tracedEngine) Notify(ctx context.Context, msg Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tracedEngine) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *tracedEngine) Simplex(ctx context.Context, nodeID ids.NodeID, msg *p2p.Simplex) error {
	_ = "STUB: not implemented"
	return nil
}
