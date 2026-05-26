// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var (
	_ StateSummaryFrontierHandler = (*noOpStateSummaryFrontierHandler)(nil)
	_ AcceptedStateSummaryHandler = (*noOpAcceptedStateSummaryHandler)(nil)
	_ AcceptedFrontierHandler     = (*noOpAcceptedFrontierHandler)(nil)
	_ AcceptedHandler             = (*noOpAcceptedHandler)(nil)
	_ AncestorsHandler            = (*noOpAncestorsHandler)(nil)
	_ PutHandler                  = (*noOpPutHandler)(nil)
	_ QueryHandler                = (*noOpQueryHandler)(nil)
	_ ChitsHandler                = (*noOpChitsHandler)(nil)
	_ AppHandler                  = (*noOpAppHandler)(nil)
	_ InternalHandler             = (*noOpInternalHandler)(nil)
	_ SimplexHandler              = (*noOpSimplexHandler)(nil)
)

type noOpStateSummaryFrontierHandler struct {
	log logging.Logger
}

func NewNoOpStateSummaryFrontierHandler(log logging.Logger) StateSummaryFrontierHandler {
	_ = "STUB: not implemented"
	return *new(StateSummaryFrontierHandler)
}

func (nop *noOpStateSummaryFrontierHandler) StateSummaryFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpStateSummaryFrontierHandler) GetStateSummaryFrontierFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpAcceptedStateSummaryHandler struct {
	log logging.Logger
}

func NewNoOpAcceptedStateSummaryHandler(log logging.Logger) AcceptedStateSummaryHandler {
	_ = "STUB: not implemented"
	return *new(AcceptedStateSummaryHandler)
}

func (nop *noOpAcceptedStateSummaryHandler) AcceptedStateSummary(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAcceptedStateSummaryHandler) GetAcceptedStateSummaryFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpAcceptedFrontierHandler struct {
	log logging.Logger
}

func NewNoOpAcceptedFrontierHandler(log logging.Logger) AcceptedFrontierHandler {
	_ = "STUB: not implemented"
	return *new(AcceptedFrontierHandler)
}

func (nop *noOpAcceptedFrontierHandler) AcceptedFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAcceptedFrontierHandler) GetAcceptedFrontierFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpAcceptedHandler struct {
	log logging.Logger
}

func NewNoOpAcceptedHandler(log logging.Logger) AcceptedHandler {
	_ = "STUB: not implemented"
	return *new(AcceptedHandler)
}

func (nop *noOpAcceptedHandler) Accepted(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAcceptedHandler) GetAcceptedFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpAncestorsHandler struct {
	log logging.Logger
}

func NewNoOpAncestorsHandler(log logging.Logger) AncestorsHandler {
	_ = "STUB: not implemented"
	return *new(AncestorsHandler)
}

func (nop *noOpAncestorsHandler) Ancestors(_ context.Context, nodeID ids.NodeID, requestID uint32, _ [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAncestorsHandler) GetAncestorsFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpPutHandler struct {
	log logging.Logger
}

func NewNoOpPutHandler(log logging.Logger) PutHandler {
	_ = "STUB: not implemented"
	return *new(PutHandler)
}

func (nop *noOpPutHandler) Put(_ context.Context, nodeID ids.NodeID, requestID uint32, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpPutHandler) GetFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpQueryHandler struct {
	log logging.Logger
}

func NewNoOpQueryHandler(log logging.Logger) QueryHandler {
	_ = "STUB: not implemented"
	return *new(QueryHandler)
}

func (nop *noOpQueryHandler) PullQuery(_ context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpQueryHandler) PushQuery(_ context.Context, nodeID ids.NodeID, requestID uint32, _ []byte, requestedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpChitsHandler struct {
	log logging.Logger
}

func NewNoOpChitsHandler(log logging.Logger) ChitsHandler {
	_ = "STUB: not implemented"
	return *new(ChitsHandler)
}

func (nop *noOpChitsHandler) Chits(_ context.Context, nodeID ids.NodeID, requestID uint32, preferredID, preferredIDAtHeight, acceptedID ids.ID, acceptedHeight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpChitsHandler) QueryFailed(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpAppHandler struct {
	log logging.Logger
}

func NewNoOpAppHandler(log logging.Logger) AppHandler {
	_ = "STUB: not implemented"
	return *new(AppHandler)
}

func (nop *noOpAppHandler) AppRequest(_ context.Context, nodeID ids.NodeID, requestID uint32, _ time.Time, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAppHandler) AppRequestFailed(_ context.Context, nodeID ids.NodeID, requestID uint32, appErr *AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAppHandler) AppResponse(_ context.Context, nodeID ids.NodeID, requestID uint32, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAppHandler) AppGossip(_ context.Context, nodeID ids.NodeID, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpInternalHandler struct {
	log logging.Logger
}

func NewNoOpInternalHandler(log logging.Logger) InternalHandler {
	_ = "STUB: not implemented"
	return *new(InternalHandler)
}

func (nop *noOpInternalHandler) Connected(
	_ context.Context,
	nodeID ids.NodeID,
	nodeVersion *version.Application,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpInternalHandler) Disconnected(_ context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpInternalHandler) Gossip(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpInternalHandler) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpInternalHandler) Notify(_ context.Context, msg Message) error {
	_ = "STUB: not implemented"
	return nil
}

type noOpSimplexHandler struct {
	log logging.Logger
}

func NewNoOpSimplexHandler(log logging.Logger) SimplexHandler {
	_ = "STUB: not implemented"
	return *new(SimplexHandler)
}

func (nop *noOpSimplexHandler) Simplex(_ context.Context, nodeID ids.NodeID, _ *p2p.Simplex) error {
	_ = "STUB: not implemented"
	return nil
}

var _ AllGetsServer = (*noOpAllGetsServer)(nil)

type noOpAllGetsServer struct {
	log logging.Logger
}

func NewNoOpAllGetsServer(log logging.Logger) AllGetsServer {
	_ = "STUB: not implemented"
	return *new(AllGetsServer)
}

func (nop *noOpAllGetsServer) GetStateSummaryFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAllGetsServer) GetAcceptedStateSummary(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[uint64]) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAllGetsServer) GetAcceptedFrontier(_ context.Context, nodeID ids.NodeID, requestID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAllGetsServer) GetAccepted(_ context.Context, nodeID ids.NodeID, requestID uint32, _ set.Set[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAllGetsServer) GetAncestors(_ context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (nop *noOpAllGetsServer) Get(_ context.Context, nodeID ids.NodeID, requestID uint32, containerID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}
