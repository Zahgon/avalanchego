// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
)

var _ InboundMsgBuilder = (*inMsgBuilder)(nil)

type InboundMsgBuilder interface {
	// Parse reads given bytes as InboundMessage
	Parse(
		bytes []byte,
		nodeID ids.NodeID,
		onFinishedHandling func(),
	) (*InboundMessage, error)
}

type inMsgBuilder struct {
	builder *msgBuilder
}

func newInboundBuilder(builder *msgBuilder) InboundMsgBuilder {
	_ = "STUB: not implemented"
	return *new(InboundMsgBuilder)
}

func (b *inMsgBuilder) Parse(bytes []byte, nodeID ids.NodeID, onFinishedHandling func()) (*InboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InboundGetStateSummaryFrontier(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundStateSummaryFrontier(
	chainID ids.ID,
	requestID uint32,
	summary []byte,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundGetAcceptedStateSummary(
	chainID ids.ID,
	requestID uint32,
	heights []uint64,
	deadline time.Duration,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAcceptedStateSummary(
	chainID ids.ID,
	requestID uint32,
	summaryIDs []ids.ID,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundGetAcceptedFrontier(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAcceptedFrontier(
	chainID ids.ID,
	requestID uint32,
	containerID ids.ID,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundGetAccepted(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerIDs []ids.ID,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAccepted(
	chainID ids.ID,
	requestID uint32,
	containerIDs []ids.ID,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundPushQuery(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	container []byte,
	requestedHeight uint64,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundPullQuery(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerID ids.ID,
	requestedHeight uint64,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundChits(
	chainID ids.ID,
	requestID uint32,
	preferredID ids.ID,
	preferredIDAtHeight ids.ID,
	acceptedID ids.ID,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAppRequest(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	msg []byte,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAppError(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
	errorCode int32,
	errorMessage string,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func InboundAppResponse(
	chainID ids.ID,
	requestID uint32,
	msg []byte,
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

// InboundSimplexMessage creates a new InboundMessage for simplex messages.
func InboundSimplexMessage(
	nodeID ids.NodeID,
	msg *p2p.Simplex,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func encodeIDs(ids []ids.ID, result [][]byte) { _ = "STUB: not implemented"; return }
