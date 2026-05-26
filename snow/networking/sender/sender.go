// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sender

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/networking/timeout"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

const opLabel = "op"

var (
	_ common.Sender = (*sender)(nil)

	opLabels = []string{opLabel}
)

// sender is a wrapper around an ExternalSender.
// Messages to this node are put directly into [router] rather than
// being sent over the network via the wrapped ExternalSender.
// sender registers outbound requests with [router] so that [router]
// fires a timeout if we don't get a response to the request.
type sender struct {
	ctx        *snow.ConsensusContext
	msgCreator message.OutboundMsgBuilder

	sender   ExternalSender // Actually does the sending over the network
	router   router.InternalHandler
	timeouts *timeout.Manager

	// Counts how many request have failed because the node was benched
	failedDueToBench *prometheus.CounterVec // op

	engineType p2p.EngineType
	subnet     subnets.Subnet
}

func New(
	ctx *snow.ConsensusContext,
	msgCreator message.OutboundMsgBuilder,
	externalSender ExternalSender,
	router router.InternalHandler,
	timeouts *timeout.Manager,
	engineType p2p.EngineType,
	subnet subnets.Subnet,
	reg prometheus.Registerer,
) (common.Sender, error) {
	_ = "STUB: not implemented"
	return *new(common.Sender), nil
}

func (s *sender) SendGetStateSummaryFrontier(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendStateSummaryFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, summary []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendGetAcceptedStateSummary(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, heights []uint64) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendAcceptedStateSummary(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendGetAcceptedFrontier(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendAcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendGetAccepted(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendAccepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendGetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// Sending a GetAncestors to myself will fail. To avoid constantly sending
// myself requests when not connected to any peers, we rely on the timeout
// firing to deliver the GetAncestorsFailed message.

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendAncestors(_ context.Context, nodeID ids.NodeID, requestID uint32, containers [][]byte) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendGet(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// Sending a Get to myself always fails.

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendPut(_ context.Context, nodeID ids.NodeID, requestID uint32, container []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendPushQuery(
	ctx context.Context,
	nodeIDs set.Set[ids.NodeID],
	requestID uint32,
	container []byte,
	requestedHeight uint64,
) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendPullQuery(
	ctx context.Context,
	nodeIDs set.Set[ids.NodeID],
	requestID uint32,
	containerID ids.ID,
	requestedHeight uint64,
) {
	_ = "STUB: not implemented"
	return
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendChits(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	preferredID ids.ID,
	preferredIDAtHeight ids.ID,
	acceptedID ids.ID,
	acceptedHeight uint64,
) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) SendAppRequest(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// The deadline is used as a best-effort communication to the peer for when
// we expect the message by. It's not guaranteed to exactly match our
// registered timeout.

func (s *sender) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) SendAppError(ctx context.Context, nodeID ids.NodeID, requestID uint32, errorCode int32, errorMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) SendAppGossip(
	_ context.Context,
	to common.SendConfig,
	bytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) sendUnlessError(
	log logging.Logger,
	to common.SendConfig,
	msg *message.OutboundMessage,
	err error,
) set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	return nil
}
