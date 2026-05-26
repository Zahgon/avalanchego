// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/version"
)

var (
	errInitialize       = errors.New("unexpectedly called Initialize")
	errSetState         = errors.New("unexpectedly called SetState")
	errShutdown         = errors.New("unexpectedly called Shutdown")
	errCreateHandlers   = errors.New("unexpectedly called CreateHandlers")
	errNewHTTPHandler   = errors.New("unexpectedly called NewHTTPHandler")
	errHealthCheck      = errors.New("unexpectedly called HealthCheck")
	errConnected        = errors.New("unexpectedly called Connected")
	errDisconnected     = errors.New("unexpectedly called Disconnected")
	errVersion          = errors.New("unexpectedly called Version")
	errAppRequest       = errors.New("unexpectedly called AppRequest")
	errAppResponse      = errors.New("unexpectedly called AppResponse")
	errAppRequestFailed = errors.New("unexpectedly called AppRequestFailed")
	errAppGossip        = errors.New("unexpectedly called AppGossip")

	_ common.VM = (*VM)(nil)
)

// VM is a test vm
type VM struct {
	T *testing.T

	CantInitialize, CantSetState,
	CantShutdown, CantCreateHandlers, CantNewHTTPHandler,
	CantHealthCheck, CantConnected, CantDisconnected, CantVersion,
	CantAppRequest, CantAppResponse, CantAppGossip, CantAppRequestFailed bool

	InitializeF       func(ctx context.Context, chainCtx *snow.Context, db database.Database, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, fxs []*common.Fx, appSender common.AppSender) error
	SetStateF         func(ctx context.Context, state snow.State) error
	ShutdownF         func(context.Context) error
	CreateHandlersF   func(context.Context) (map[string]http.Handler, error)
	NewHTTPHandlerF   func(context.Context) (http.Handler, error)
	ConnectedF        func(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error
	DisconnectedF     func(ctx context.Context, nodeID ids.NodeID) error
	HealthCheckF      func(context.Context) (interface{}, error)
	AppRequestF       func(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, msg []byte) error
	AppResponseF      func(ctx context.Context, nodeID ids.NodeID, requestID uint32, msg []byte) error
	AppGossipF        func(ctx context.Context, nodeID ids.NodeID, msg []byte) error
	AppRequestFailedF func(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error
	VersionF          func(context.Context) (string, error)
	WaitForEventF     common.Subscription
}

func (vm *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

func (vm *VM) Default(cant bool) { _ = "STUB: not implemented"; return }

func (vm *VM) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes,
	upgradeBytes,
	configBytes []byte,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) SetState(ctx context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VM) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) NewHTTPHandler(ctx context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (vm *VM) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) AppGossip(ctx context.Context, nodeID ids.NodeID, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Connected(ctx context.Context, id ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Disconnected(ctx context.Context, id ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Version(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
