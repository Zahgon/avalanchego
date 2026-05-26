// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ common.Sender = (*tracedSender)(nil)

type tracedSender struct {
	sender common.Sender
	tracer trace.Tracer
}

func Trace(sender common.Sender, tracer trace.Tracer) common.Sender {
	_ = "STUB: not implemented"
	return *new(common.Sender)
}

func (s *tracedSender) SendGetStateSummaryFrontier(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendStateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, summary []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendGetAcceptedStateSummary(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, heights []uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendAcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendGetAcceptedFrontier(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendAcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendGetAccepted(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendAccepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendGetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containers [][]byte) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendGet(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendPut(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendPushQuery(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, container []byte, requestedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendPullQuery(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, containerID ids.ID, requestedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendChits(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *tracedSender) SendAppRequest(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, appRequestBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tracedSender) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, appResponseBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tracedSender) SendAppError(ctx context.Context, nodeID ids.NodeID, requestID uint32, errorCode int32, errorMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *tracedSender) SendAppGossip(
	ctx context.Context,
	config common.SendConfig,
	appGossipBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}
