// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"errors"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"

	vmpb "github.com/ava-labs/avalanchego/proto/pb/vm"
)

var (
	_ vmpb.VMServer = (*VMServer)(nil)

	originalStderr = os.Stderr

	errExpectedBlockWithVerifyContext = errors.New("expected block.WithVerifyContext")
	errNilNetworkUpgradesPB           = errors.New("network upgrades protobuf is nil")
)

// VMServer is a VM that is managed over RPC.
type VMServer struct {
	vmpb.UnsafeVMServer

	vm block.ChainVM
	// If nil, the underlying VM doesn't implement the interface.
	bVM block.BuildBlockWithContextChainVM
	// If nil, the underlying VM doesn't implement the interface.
	ssVM block.StateSyncableVM

	allowShutdown *utils.Atomic[bool]

	metrics metrics.MultiGatherer
	db      database.Database
	log     logging.Logger

	serverCloser grpcutils.ServerCloser
	connCloser   wrappers.Closer

	ctx    *snow.Context
	closed chan struct{}
}

// NewServer returns a vm instance connected to a remote vm instance
func NewServer(vm block.ChainVM, allowShutdown *utils.Atomic[bool]) *VMServer {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMServer) Initialize(ctx context.Context, req *vmpb.InitializeRequest) (*vmpb.InitializeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Current state of process metrics

// Go process metrics using debug.GCStats

// gRPC client metrics

// Dial the database

// TODO: Allow the logger to be configured by the client

// Ignore closing errors to return the original error

// Signs warp messages

// The validator state is cached to avoid the gRPC overhead.

// TODO: support remaining snowman++ fields

// Ignore errors closing resources to return the original error

// Ignore errors closing resources to return the original error

// Ignore errors closing resources to return the original error

func (vm *VMServer) SetState(ctx context.Context, stateReq *vmpb.SetStateRequest) (*vmpb.SetStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Shutdown(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) CreateHandlers(ctx context.Context, _ *emptypb.Empty) (*vmpb.CreateHandlersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start HTTP service

func (vm *VMServer) NewHTTPHandler(ctx context.Context, _ *emptypb.Empty) (*vmpb.NewHTTPHandlerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start HTTP service

func (vm *VMServer) WaitForEvent(ctx context.Context, _ *emptypb.Empty) (*vmpb.WaitForEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Connected(ctx context.Context, req *vmpb.ConnectedRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Disconnected(ctx context.Context, req *vmpb.DisconnectedRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the underlying VM doesn't actually implement this method, its [BuildBlock]
// method will be called instead.
func (vm *VMServer) BuildBlock(ctx context.Context, req *vmpb.BuildBlockRequest) (*vmpb.BuildBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) ParseBlock(ctx context.Context, req *vmpb.ParseBlockRequest) (*vmpb.ParseBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetBlock(ctx context.Context, req *vmpb.GetBlockRequest) (*vmpb.GetBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) SetPreference(ctx context.Context, req *vmpb.SetPreferenceRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Health(ctx context.Context, _ *emptypb.Empty) (*vmpb.HealthResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Version(ctx context.Context, _ *emptypb.Empty) (*vmpb.VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) AppRequest(ctx context.Context, req *vmpb.AppRequestMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) AppRequestFailed(ctx context.Context, req *vmpb.AppRequestFailedMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) AppResponse(ctx context.Context, req *vmpb.AppResponseMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) AppGossip(ctx context.Context, req *vmpb.AppGossipMsg) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) Gather(context.Context, *emptypb.Empty) (*vmpb.GatherResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetAncestors(ctx context.Context, req *vmpb.GetAncestorsRequest) (*vmpb.GetAncestorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) BatchedParseBlock(
	ctx context.Context,
	req *vmpb.BatchedParseBlockRequest,
) (*vmpb.BatchedParseBlockResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetBlockIDAtHeight(
	ctx context.Context,
	req *vmpb.GetBlockIDAtHeightRequest,
) (*vmpb.GetBlockIDAtHeightResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) StateSyncEnabled(ctx context.Context, _ *emptypb.Empty) (*vmpb.StateSyncEnabledResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetOngoingSyncStateSummary(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vmpb.GetOngoingSyncStateSummaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetLastStateSummary(ctx context.Context, _ *emptypb.Empty) (*vmpb.GetLastStateSummaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) ParseStateSummary(
	ctx context.Context,
	req *vmpb.ParseStateSummaryRequest,
) (*vmpb.ParseStateSummaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) GetStateSummary(
	ctx context.Context,
	req *vmpb.GetStateSummaryRequest,
) (*vmpb.GetStateSummaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) BlockVerify(ctx context.Context, req *vmpb.BlockVerifyRequest) (*vmpb.BlockVerifyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) BlockAccept(ctx context.Context, req *vmpb.BlockAcceptRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) BlockReject(ctx context.Context, req *vmpb.BlockRejectRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMServer) StateSummaryAccept(
	ctx context.Context,
	req *vmpb.StateSummaryAcceptRequest,
) (*vmpb.StateSummaryAcceptResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertNetworkUpgrades(pbUpgrades *vmpb.NetworkUpgrades) (upgrade.Config, error) {
	_ = "STUB: not implemented"
	return *new(upgrade.Config), nil
}
