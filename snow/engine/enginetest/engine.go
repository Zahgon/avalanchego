// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var (
	errGossip                        = errors.New("unexpectedly called Gossip")
	errNotify                        = errors.New("unexpectedly called Notify")
	errGetStateSummaryFrontier       = errors.New("unexpectedly called GetStateSummaryFrontier")
	errGetStateSummaryFrontierFailed = errors.New("unexpectedly called GetStateSummaryFrontierFailed")
	errStateSummaryFrontier          = errors.New("unexpectedly called StateSummaryFrontier")
	errGetAcceptedStateSummary       = errors.New("unexpectedly called GetAcceptedStateSummary")
	errGetAcceptedStateSummaryFailed = errors.New("unexpectedly called GetAcceptedStateSummaryFailed")
	errAcceptedStateSummary          = errors.New("unexpectedly called AcceptedStateSummary")
	errGetAcceptedFrontier           = errors.New("unexpectedly called GetAcceptedFrontier")
	errGetAcceptedFrontierFailed     = errors.New("unexpectedly called GetAcceptedFrontierFailed")
	errAcceptedFrontier              = errors.New("unexpectedly called AcceptedFrontier")
	errGetAccepted                   = errors.New("unexpectedly called GetAccepted")
	errGetAcceptedFailed             = errors.New("unexpectedly called GetAcceptedFailed")
	errAccepted                      = errors.New("unexpectedly called Accepted")
	errGet                           = errors.New("unexpectedly called Get")
	errGetAncestors                  = errors.New("unexpectedly called GetAncestors")
	errGetFailed                     = errors.New("unexpectedly called GetFailed")
	errGetAncestorsFailed            = errors.New("unexpectedly called GetAncestorsFailed")
	errPut                           = errors.New("unexpectedly called Put")
	errAncestors                     = errors.New("unexpectedly called Ancestors")
	errPushQuery                     = errors.New("unexpectedly called PushQuery")
	errPullQuery                     = errors.New("unexpectedly called PullQuery")
	errQueryFailed                   = errors.New("unexpectedly called QueryFailed")
	errChits                         = errors.New("unexpectedly called Chits")
	errSimplex                       = errors.New("unexpectedly called Simplex")
	errStart                         = errors.New("unexpectedly called Start")

	_ common.Engine = (*Engine)(nil)
)

// Engine is a test engine
type Engine struct {
	T *testing.T

	CantStart,

	CantIsBootstrapped,
	CantTimeout,
	CantGossip,
	CantHalt,
	CantShutdown,

	CantContext,

	CantNotify,

	CantGetStateSummaryFrontier,
	CantGetStateSummaryFrontierFailed,
	CantStateSummaryFrontier,

	CantGetAcceptedStateSummary,
	CantGetAcceptedStateSummaryFailed,
	CantAcceptedStateSummary,

	CantGetAcceptedFrontier,
	CantGetAcceptedFrontierFailed,
	CantAcceptedFrontier,

	CantGetAccepted,
	CantGetAcceptedFailed,
	CantAccepted,

	CantGet,
	CantGetAncestors,
	CantGetFailed,
	CantGetAncestorsFailed,
	CantPut,
	CantAncestors,

	CantPushQuery,
	CantPullQuery,
	CantQueryFailed,
	CantChits,
	CantSimplex,

	CantConnected,
	CantDisconnected,

	CantHealth,

	CantAppRequest,
	CantAppResponse,
	CantAppGossip,
	CantAppRequestFailed,

	CantGetVM bool

	StartF                       func(ctx context.Context, startReqID uint32) error
	IsBootstrappedF              func() bool
	ContextF                     func() *snow.ConsensusContext
	HaltF                        func(context.Context)
	TimeoutF, GossipF, ShutdownF func(context.Context) error
	NotifyF                      func(context.Context, common.Message) error
	GetF, GetAncestorsF          func(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error
	PullQueryF                   func(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID, requestedHeight uint64) error
	PutF                         func(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte) error
	PushQueryF                   func(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte, requestedHeight uint64) error
	AncestorsF                   func(ctx context.Context, nodeID ids.NodeID, requestID uint32, containers [][]byte) error
	AcceptedFrontierF            func(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error
	GetAcceptedF, AcceptedF      func(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredIDs set.Set[ids.ID]) error
	ChitsF                       func(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) error
	SimplexF                     func(ctx context.Context, nodeID ids.NodeID, msg *p2p.Simplex) error
	GetStateSummaryFrontierF, GetStateSummaryFrontierFailedF, GetAcceptedStateSummaryFailedF,
	GetAcceptedFrontierF, GetFailedF, GetAncestorsFailedF,
	QueryFailedF, GetAcceptedFrontierFailedF, GetAcceptedFailedF func(ctx context.Context, nodeID ids.NodeID, requestID uint32) error
	AppRequestFailedF        func(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error
	StateSummaryFrontierF    func(ctx context.Context, nodeID ids.NodeID, requestID uint32, summary []byte) error
	GetAcceptedStateSummaryF func(ctx context.Context, nodeID ids.NodeID, requestID uint32, keys set.Set[uint64]) error
	AcceptedStateSummaryF    func(ctx context.Context, nodeID ids.NodeID, requestID uint32, summaryIDs set.Set[ids.ID]) error
	ConnectedF               func(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error
	DisconnectedF            func(ctx context.Context, nodeID ids.NodeID) error
	HealthF                  func(context.Context) (interface{}, error)
	GetVMF                   func() common.VM
	AppRequestF              func(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, msg []byte) error
	AppResponseF             func(ctx context.Context, nodeID ids.NodeID, requestID uint32, msg []byte) error
	AppGossipF               func(ctx context.Context, nodeID ids.NodeID, msg []byte) error
}

func (e *Engine) Default(cant bool) { _ = "STUB: not implemented"; return }

func (e *Engine) Start(ctx context.Context, startReqID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) Notify(ctx context.Context, msg common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetStateSummaryFrontier(ctx context.Context, validatorID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) StateSummaryFrontier(ctx context.Context, validatorID ids.NodeID, requestID uint32, summary []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetStateSummaryFrontierFailed(ctx context.Context, validatorID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAcceptedStateSummary(ctx context.Context, validatorID ids.NodeID, requestID uint32, keys set.Set[uint64]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AcceptedStateSummary(ctx context.Context, validatorID ids.NodeID, requestID uint32, summaryIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAcceptedStateSummaryFailed(ctx context.Context, validatorID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAcceptedFrontierFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AcceptedFrontier(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAccepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAcceptedFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Accepted(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerIDs set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Get(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAncestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) GetAncestorsFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Put(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Ancestors(ctx context.Context, nodeID ids.NodeID, requestID uint32, containers [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) PushQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, container []byte, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) PullQuery(ctx context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) QueryFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) AppGossip(ctx context.Context, nodeID ids.NodeID, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Chits(ctx context.Context, nodeID ids.NodeID, requestID uint32, preferredID ids.ID, preferredIDAtHeight ids.ID, acceptedID ids.ID, acceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Simplex(ctx context.Context, nodeID ids.NodeID, msg *p2p.Simplex) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
