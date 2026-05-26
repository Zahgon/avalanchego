// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ common.Sender    = (*Sender)(nil)
	_ common.AppSender = (*SenderStub)(nil)

	errSendAppRequest  = errors.New("unexpectedly called SendAppRequest")
	errSendAppResponse = errors.New("unexpectedly called SendAppResponse")
	errSendAppError    = errors.New("unexpectedly called SendAppError")
	errSendAppGossip   = errors.New("unexpectedly called SendAppGossip")
)

// Sender is a test sender
type Sender struct {
	T *testing.T

	CantSendGetStateSummaryFrontier, CantSendStateSummaryFrontier,
	CantSendGetAcceptedStateSummary, CantSendAcceptedStateSummary,
	CantSendGetAcceptedFrontier, CantSendAcceptedFrontier,
	CantSendGetAccepted, CantSendAccepted,
	CantSendGet, CantSendGetAncestors, CantSendPut, CantSendAncestors,
	CantSendPullQuery, CantSendPushQuery, CantSendChits,
	CantSendAppRequest, CantSendAppResponse, CantSendAppError,
	CantSendAppGossip bool

	SendGetStateSummaryFrontierF func(context.Context, set.Set[ids.NodeID], uint32)
	SendStateSummaryFrontierF    func(context.Context, ids.NodeID, uint32, []byte)
	SendGetAcceptedStateSummaryF func(context.Context, set.Set[ids.NodeID], uint32, []uint64)
	SendAcceptedStateSummaryF    func(context.Context, ids.NodeID, uint32, []ids.ID)
	SendGetAcceptedFrontierF     func(context.Context, set.Set[ids.NodeID], uint32)
	SendAcceptedFrontierF        func(context.Context, ids.NodeID, uint32, ids.ID)
	SendGetAcceptedF             func(context.Context, set.Set[ids.NodeID], uint32, []ids.ID)
	SendAcceptedF                func(context.Context, ids.NodeID, uint32, []ids.ID)
	SendGetF                     func(context.Context, ids.NodeID, uint32, ids.ID)
	SendGetAncestorsF            func(context.Context, ids.NodeID, uint32, ids.ID)
	SendPutF                     func(context.Context, ids.NodeID, uint32, []byte)
	SendAncestorsF               func(context.Context, ids.NodeID, uint32, [][]byte)
	SendPushQueryF               func(context.Context, set.Set[ids.NodeID], uint32, []byte, uint64)
	SendPullQueryF               func(context.Context, set.Set[ids.NodeID], uint32, ids.ID, uint64)
	SendChitsF                   func(context.Context, ids.NodeID, uint32, ids.ID, ids.ID, ids.ID, uint64)
	SendAppRequestF              func(context.Context, set.Set[ids.NodeID], uint32, []byte) error
	SendAppResponseF             func(context.Context, ids.NodeID, uint32, []byte) error
	SendAppErrorF                func(context.Context, ids.NodeID, uint32, int32, string) error
	SendAppGossipF               func(context.Context, common.SendConfig, []byte) error
}

// Default set the default callable value to [cant]
func (s *Sender) Default(cant bool) { _ = "STUB: not implemented"; return }

// SendGetStateSummaryFrontier calls SendGetStateSummaryFrontierF if it was
// initialized. If it wasn't initialized and this function shouldn't be called
// and testing was initialized, then testing will fail.
func (s *Sender) SendGetStateSummaryFrontier(ctx context.Context, validatorIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

// SendStateSummaryFrontier calls SendStateSummaryFrontierF if it was
// initialized. If it wasn't initialized and this function shouldn't be called
// and testing was initialized, then testing will fail.
func (s *Sender) SendStateSummaryFrontier(ctx context.Context, validatorID ids.NodeID, requestID uint32, summary []byte) {
	_ = "STUB: not implemented"
	return
}

// SendGetAcceptedStateSummary calls SendGetAcceptedStateSummaryF if it was
// initialized. If it wasn't initialized and this function shouldn't be called
// and testing was initialized, then testing will fail.
func (s *Sender) SendGetAcceptedStateSummary(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, heights []uint64) {
	_ = "STUB: not implemented"
	return
}

// SendAcceptedStateSummary calls SendAcceptedStateSummaryF if it was
// initialized. If it wasn't initialized and this function shouldn't be called
// and testing was initialized, then testing will fail.
func (s *Sender) SendAcceptedStateSummary(ctx context.Context, validatorID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendGetAcceptedFrontier calls SendGetAcceptedFrontierF if it was initialized.
// If it wasn't initialized and this function shouldn't be called and testing
// was initialized, then testing will fail.
func (s *Sender) SendGetAcceptedFrontier(ctx context.Context, validatorIDs set.Set[ids.NodeID], requestID uint32) {
	_ = "STUB: not implemented"
	return
}

// SendAcceptedFrontier calls SendAcceptedFrontierF if it was initialized. If it
// wasn't initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAcceptedFrontier(ctx context.Context, validatorID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendGetAccepted calls SendGetAcceptedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendGetAccepted(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendAccepted calls SendAcceptedF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAccepted(ctx context.Context, validatorID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendGet calls SendGetF if it was initialized. If it wasn't initialized and
// this function shouldn't be called and testing was initialized, then testing
// will fail.
func (s *Sender) SendGet(ctx context.Context, vdr ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendGetAncestors calls SendGetAncestorsF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendGetAncestors(ctx context.Context, validatorID ids.NodeID, requestID uint32, containerID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// SendPut calls SendPutF if it was initialized. If it wasn't initialized and
// this function shouldn't be called and testing was initialized, then testing
// will fail.
func (s *Sender) SendPut(ctx context.Context, vdr ids.NodeID, requestID uint32, container []byte) {
	_ = "STUB: not implemented"
	return
}

// SendAncestors calls SendAncestorsF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAncestors(ctx context.Context, vdr ids.NodeID, requestID uint32, containers [][]byte) {
	_ = "STUB: not implemented"
	return
}

// SendPushQuery calls SendPushQueryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendPushQuery(ctx context.Context, vdrs set.Set[ids.NodeID], requestID uint32, container []byte, requestedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

// SendPullQuery calls SendPullQueryF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendPullQuery(ctx context.Context, vdrs set.Set[ids.NodeID], requestID uint32, containerID ids.ID, requestedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

// SendChits calls SendChitsF if it was initialized. If it wasn't initialized
// and this function shouldn't be called and testing was initialized, then
// testing will fail.
func (s *Sender) SendChits(ctx context.Context, vdr ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) {
	_ = "STUB: not implemented"
	return
}

// SendAppRequest calls SendAppRequestF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAppRequest(ctx context.Context, nodeIDs set.Set[ids.NodeID], requestID uint32, appRequestBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SendAppResponse calls SendAppResponseF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, appResponseBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SendAppError calls SendAppErrorF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAppError(ctx context.Context, nodeID ids.NodeID, requestID uint32, code int32, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// SendAppGossip calls SendAppGossipF if it was initialized. If it wasn't
// initialized and this function shouldn't be called and testing was
// initialized, then testing will fail.
func (s *Sender) SendAppGossip(
	ctx context.Context,
	config common.SendConfig,
	appGossipBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SenderStub is a stub sender that returns values received on method-specific channels.
type SenderStub struct {
	SentAppRequest, SentAppResponse,
	SentAppGossip chan []byte

	SentAppError chan *common.AppError
}

func (f SenderStub) SendAppRequest(_ context.Context, _ set.Set[ids.NodeID], _ uint32, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f SenderStub) SendAppResponse(_ context.Context, _ ids.NodeID, _ uint32, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f SenderStub) SendAppError(_ context.Context, _ ids.NodeID, _ uint32, errorCode int32, errorMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f SenderStub) SendAppGossip(_ context.Context, _ common.SendConfig, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}
