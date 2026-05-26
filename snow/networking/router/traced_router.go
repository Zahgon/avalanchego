// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow/networking/handler"
	"github.com/ava-labs/avalanchego/snow/networking/timeout"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var _ Router = (*tracedRouter)(nil)

type tracedRouter struct {
	router Router
	tracer trace.Tracer
}

func Trace(router Router, tracer trace.Tracer) Router {
	_ = "STUB: not implemented"
	return *new(Router)
}

func (r *tracedRouter) Initialize(
	nodeID ids.NodeID,
	log logging.Logger,
	timeoutManager *timeout.Manager,
	closeTimeout time.Duration,
	criticalChains set.Set[ids.ID],
	sybilProtectionEnabled bool,
	trackedSubnets set.Set[ids.ID],
	onFatal func(exitCode int),
	healthConfig HealthConfig,
	reg prometheus.Registerer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *tracedRouter) RegisterRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
	op message.Op,
	failedMsg *message.InboundMessage,
	engineType p2p.EngineType,
) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) HandleInbound(ctx context.Context, msg *message.InboundMessage) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) HandleInternal(ctx context.Context, msg *message.InboundMessage) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) Shutdown(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *tracedRouter) AddChain(ctx context.Context, chain handler.Handler) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

func (r *tracedRouter) Benched(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) Unbenched(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (r *tracedRouter) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
