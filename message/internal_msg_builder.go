// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//nolint:staticcheck // proto generates interfaces that fail linting
package message

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/version"
)

var (
	disconnected  = &Disconnected{}
	gossipRequest = &GossipRequest{}

	_ fmt.Stringer    = (*GetStateSummaryFrontierFailed)(nil)
	_ chainIDGetter   = (*GetStateSummaryFrontierFailed)(nil)
	_ requestIDGetter = (*GetStateSummaryFrontierFailed)(nil)

	_ fmt.Stringer    = (*GetAcceptedStateSummaryFailed)(nil)
	_ chainIDGetter   = (*GetAcceptedStateSummaryFailed)(nil)
	_ requestIDGetter = (*GetAcceptedStateSummaryFailed)(nil)

	_ fmt.Stringer    = (*GetAcceptedFrontierFailed)(nil)
	_ chainIDGetter   = (*GetAcceptedFrontierFailed)(nil)
	_ requestIDGetter = (*GetAcceptedFrontierFailed)(nil)

	_ fmt.Stringer    = (*GetAcceptedFailed)(nil)
	_ chainIDGetter   = (*GetAcceptedFailed)(nil)
	_ requestIDGetter = (*GetAcceptedFailed)(nil)

	_ fmt.Stringer     = (*GetAncestorsFailed)(nil)
	_ chainIDGetter    = (*GetAncestorsFailed)(nil)
	_ requestIDGetter  = (*GetAncestorsFailed)(nil)
	_ engineTypeGetter = (*GetAncestorsFailed)(nil)

	_ fmt.Stringer    = (*GetFailed)(nil)
	_ chainIDGetter   = (*GetFailed)(nil)
	_ requestIDGetter = (*GetFailed)(nil)

	_ fmt.Stringer    = (*QueryFailed)(nil)
	_ chainIDGetter   = (*QueryFailed)(nil)
	_ requestIDGetter = (*QueryFailed)(nil)

	_ fmt.Stringer = (*Disconnected)(nil)

	_ fmt.Stringer = (*GossipRequest)(nil)
)

type GetStateSummaryFrontierFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *GetStateSummaryFrontierFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetStateSummaryFrontierFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetStateSummaryFrontierFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalGetStateSummaryFrontierFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GetAcceptedStateSummaryFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *GetAcceptedStateSummaryFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetAcceptedStateSummaryFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetAcceptedStateSummaryFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalGetAcceptedStateSummaryFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GetAcceptedFrontierFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *GetAcceptedFrontierFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetAcceptedFrontierFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetAcceptedFrontierFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalGetAcceptedFrontierFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GetAcceptedFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *GetAcceptedFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetAcceptedFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetAcceptedFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalGetAcceptedFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GetAncestorsFailed struct {
	ChainID    ids.ID         `json:"chain_id,omitempty"`
	RequestID  uint32         `json:"request_id,omitempty"`
	EngineType p2p.EngineType `json:"engine_type,omitempty"`
}

func (m *GetAncestorsFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetAncestorsFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetAncestorsFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *GetAncestorsFailed) GetEngineType() p2p.EngineType {
	_ = "STUB: not implemented"
	return *new(p2p.EngineType)
}

func InternalGetAncestorsFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
	engineType p2p.EngineType,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GetFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *GetFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *GetFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalGetFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type QueryFailed struct {
	ChainID   ids.ID `json:"chain_id,omitempty"`
	RequestID uint32 `json:"request_id,omitempty"`
}

func (m *QueryFailed) String() string { _ = "STUB: not implemented"; return "" }

func (m *QueryFailed) GetChainId() []byte { _ = "STUB: not implemented"; return nil }

func (m *QueryFailed) GetRequestId() uint32 { _ = "STUB: not implemented"; return 0 }

func InternalQueryFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type Connected struct {
	NodeVersion *version.Application `json:"node_version,omitempty"`
}

func (m *Connected) String() string { _ = "STUB: not implemented"; return "" }

func InternalConnected(nodeID ids.NodeID, nodeVersion *version.Application) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type Disconnected struct{}

func (Disconnected) String() string { _ = "STUB: not implemented"; return "" }

func InternalDisconnected(nodeID ids.NodeID) *InboundMessage { _ = "STUB: not implemented"; return nil }

type VMMessage struct {
	Notification uint32 `json:"notification,omitempty"`
}

func (m *VMMessage) String() string { _ = "STUB: not implemented"; return "" }

func InternalVMMessage(
	nodeID ids.NodeID,
	notification uint32,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}

type GossipRequest struct{}

func (GossipRequest) String() string { _ = "STUB: not implemented"; return "" }

func InternalGossipRequest(
	nodeID ids.NodeID,
) *InboundMessage {
	_ = "STUB: not implemented"
	return nil
}
