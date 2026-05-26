// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/utils/compression"
	"github.com/ava-labs/avalanchego/utils/ips"
)

var _ OutboundMsgBuilder = (*outMsgBuilder)(nil)

// OutboundMsgBuilder builds outbound messages. Outbound messages are returned
// with a reference count of 1. Once the reference count hits 0, the message
// bytes should no longer be accessed.
type OutboundMsgBuilder interface {
	Handshake(
		networkID uint32,
		myTime uint64,
		ip netip.AddrPort,
		client string,
		major uint32,
		minor uint32,
		patch uint32,
		upgradeTime uint64,
		ipSigningTime uint64,
		ipNodeIDSig []byte,
		ipBLSSig []byte,
		trackedSubnets []ids.ID,
		supportedACPs []uint32,
		objectedACPs []uint32,
		knownPeersFilter []byte,
		knownPeersSalt []byte,
		requestAllSubnetIPs bool,
	) (*OutboundMessage, error)

	GetPeerList(
		knownPeersFilter []byte,
		knownPeersSalt []byte,
		requestAllSubnetIPs bool,
	) (*OutboundMessage, error)

	PeerList(
		peers []*ips.ClaimedIPPort,
		bypassThrottling bool,
	) (*OutboundMessage, error)

	Ping(
		primaryUptime uint32,
	) (*OutboundMessage, error)

	Pong() (*OutboundMessage, error)

	GetStateSummaryFrontier(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
	) (*OutboundMessage, error)

	StateSummaryFrontier(
		chainID ids.ID,
		requestID uint32,
		summary []byte,
	) (*OutboundMessage, error)

	GetAcceptedStateSummary(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		heights []uint64,
	) (*OutboundMessage, error)

	AcceptedStateSummary(
		chainID ids.ID,
		requestID uint32,
		summaryIDs []ids.ID,
	) (*OutboundMessage, error)

	GetAcceptedFrontier(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
	) (*OutboundMessage, error)

	AcceptedFrontier(
		chainID ids.ID,
		requestID uint32,
		containerID ids.ID,
	) (*OutboundMessage, error)

	GetAccepted(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		containerIDs []ids.ID,
	) (*OutboundMessage, error)

	Accepted(
		chainID ids.ID,
		requestID uint32,
		containerIDs []ids.ID,
	) (*OutboundMessage, error)

	GetAncestors(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		containerID ids.ID,
		engineType p2p.EngineType,
	) (*OutboundMessage, error)

	Ancestors(
		chainID ids.ID,
		requestID uint32,
		containers [][]byte,
	) (*OutboundMessage, error)

	Get(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		containerID ids.ID,
	) (*OutboundMessage, error)

	Put(
		chainID ids.ID,
		requestID uint32,
		container []byte,
	) (*OutboundMessage, error)

	PushQuery(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		container []byte,
		requestedHeight uint64,
	) (*OutboundMessage, error)

	PullQuery(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		containerID ids.ID,
		requestedHeight uint64,
	) (*OutboundMessage, error)

	Chits(
		chainID ids.ID,
		requestID uint32,
		preferredID ids.ID,
		preferredIDAtHeight ids.ID,
		acceptedID ids.ID,
		acceptedHeight uint64,
	) (*OutboundMessage, error)

	AppRequest(
		chainID ids.ID,
		requestID uint32,
		deadline time.Duration,
		msg []byte,
	) (*OutboundMessage, error)

	AppResponse(
		chainID ids.ID,
		requestID uint32,
		msg []byte,
	) (*OutboundMessage, error)

	AppError(
		chainID ids.ID,
		requestID uint32,
		errorCode int32,
		errorMessage string,
	) (*OutboundMessage, error)

	AppGossip(
		chainID ids.ID,
		msg []byte,
	) (*OutboundMessage, error)

	SimplexMessage(
		msg *p2p.Simplex,
	) (*OutboundMessage, error)
}

type outMsgBuilder struct {
	compressionType compression.Type

	builder *msgBuilder
}

// Use "message.NewCreator" to import this function
// since we do not expose "msgBuilder" yet
func newOutboundBuilder(compressionType compression.Type, builder *msgBuilder) OutboundMsgBuilder {
	_ = "STUB: not implemented"
	return *new(OutboundMsgBuilder)
}

func (b *outMsgBuilder) Ping(
	primaryUptime uint32,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Pong() (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Handshake(
	networkID uint32,
	myTime uint64,
	ip netip.AddrPort,
	client string,
	major uint32,
	minor uint32,
	patch uint32,
	upgradeTime uint64,
	ipSigningTime uint64,
	ipNodeIDSig []byte,
	ipBLSSig []byte,
	trackedSubnets []ids.ID,
	supportedACPs []uint32,
	objectedACPs []uint32,
	knownPeersFilter []byte,
	knownPeersSalt []byte,
	requestAllSubnetIPs bool,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetPeerList(
	knownPeersFilter []byte,
	knownPeersSalt []byte,
	requestAllSubnetIPs bool,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) PeerList(peers []*ips.ClaimedIPPort, bypassThrottling bool) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetStateSummaryFrontier(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) StateSummaryFrontier(
	chainID ids.ID,
	requestID uint32,
	summary []byte,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetAcceptedStateSummary(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	heights []uint64,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AcceptedStateSummary(
	chainID ids.ID,
	requestID uint32,
	summaryIDs []ids.ID,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetAcceptedFrontier(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AcceptedFrontier(
	chainID ids.ID,
	requestID uint32,
	containerID ids.ID,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetAccepted(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerIDs []ids.ID,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Accepted(
	chainID ids.ID,
	requestID uint32,
	containerIDs []ids.ID,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) GetAncestors(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerID ids.ID,
	engineType p2p.EngineType,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Ancestors(
	chainID ids.ID,
	requestID uint32,
	containers [][]byte,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Get(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerID ids.ID,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Put(
	chainID ids.ID,
	requestID uint32,
	container []byte,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) PushQuery(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	container []byte,
	requestedHeight uint64,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) PullQuery(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	containerID ids.ID,
	requestedHeight uint64,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) Chits(
	chainID ids.ID,
	requestID uint32,
	preferredID ids.ID,
	preferredIDAtHeight ids.ID,
	acceptedID ids.ID,
	acceptedHeight uint64,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AppRequest(
	chainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	msg []byte,
) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AppResponse(chainID ids.ID, requestID uint32, msg []byte) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AppError(chainID ids.ID, requestID uint32, errorCode int32, errorMessage string) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) AppGossip(chainID ids.ID, msg []byte) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *outMsgBuilder) SimplexMessage(msg *p2p.Simplex) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
