// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/chains/atomic/gsharedmemory"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/ids/galiasreader"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/common/appsender"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators/gvalidators"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/resource"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/avalanchego/vms/components/chain"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/gwarp"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"

	vmpb "github.com/ava-labs/avalanchego/proto/pb/vm"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	dto "github.com/prometheus/client_model/go"
)

// TODO: Enable these to be configured by the user
const (
	decidedCacheSize    = 64 * units.MiB
	missingCacheSize    = 2048
	unverifiedCacheSize = 64 * units.MiB
	bytesToIDCacheSize  = 64 * units.MiB
)

var (
	errUnsupportedFXs                       = errors.New("unsupported feature extensions")
	errBatchedParseBlockWrongNumberOfBlocks = errors.New("BatchedParseBlock returned different number of blocks than expected")

	_ block.ChainVM                      = (*VMClient)(nil)
	_ block.BuildBlockWithContextChainVM = (*VMClient)(nil)
	_ block.BatchedChainVM               = (*VMClient)(nil)
	_ block.StateSyncableVM              = (*VMClient)(nil)
	_ prometheus.Gatherer                = (*VMClient)(nil)

	_ snowman.Block           = (*blockClient)(nil)
	_ block.WithVerifyContext = (*blockClient)(nil)

	_ block.StateSummary = (*summaryClient)(nil)
)

// VMClient is an implementation of a VM that talks over RPC.
type VMClient struct {
	*chain.State
	logger          logging.Logger
	client          vmpb.VMClient
	runtime         runtime.Stopper
	pid             int
	processTracker  resource.ProcessTracker
	metricsGatherer metrics.MultiGatherer

	sharedMemory         *gsharedmemory.Server
	bcLookup             *galiasreader.Server
	appSender            *appsender.Server
	validatorStateServer *gvalidators.Server
	warpSignerServer     *gwarp.Server

	serverCloser grpcutils.ServerCloser
	conns        []*grpc.ClientConn

	grpcServerMetrics *grpc_prometheus.ServerMetrics
}

// NewClient returns a VM connected to a remote VM
func NewClient(
	clientConn *grpc.ClientConn,
	runtime runtime.Stopper,
	pid int,
	processTracker resource.ProcessTracker,
	metricsGatherer metrics.MultiGatherer,
	logger logging.Logger,
) *VMClient {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	upgradeBytes []byte,
	configBytes []byte,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If fetching the alias fails, we default to the chain's ID

// Register metrics

// Initialize the database

// We don't need to check whether this is a block.WithVerifyContext because
// we'll never Verify this block.

func getNetworkUpgrades(u upgrade.Config) *vmpb.NetworkUpgrades {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) newDBServer(db database.Database) *grpc.Server {
	_ = "STUB: not implemented"
	return nil
}

// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md

// Register services

// Ensure metric counters are zeroed on restart

func (vm *VMClient) newInitServer() *grpc.Server { _ = "STUB: not implemented"; return nil }

// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md

// Register services

// Ensure metric counters are zeroed on restart

func (vm *VMClient) SetState(ctx context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't need to check whether this is a block.WithVerifyContext because
// we'll never Verify this block.

func (vm *VMClient) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VMClient) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMClient) NewHTTPHandler(ctx context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (vm *VMClient) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

func (vm *VMClient) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

// If the underlying VM doesn't actually implement this method, its [BuildBlock]
// method will be called instead.
func (vm *VMClient) buildBlockWithContext(ctx context.Context, blockCtx *block.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VMClient) buildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VMClient) parseBlock(ctx context.Context, bytes []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VMClient) getBlock(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VMClient) SetPreference(ctx context.Context, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	// HealthCheck is a special case, where we want to fail fast instead of block.
	return nil, nil
}

func (vm *VMClient) Version(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (vm *VMClient) AppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) AppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) AppRequestFailed(ctx context.Context, nodeID ids.NodeID, requestID uint32, appErr *common.AppError) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) AppGossip(ctx context.Context, nodeID ids.NodeID, msg []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VMClient) Gather() ([]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMClient) GetAncestors(
	ctx context.Context,
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrivalTime time.Duration,
) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMClient) batchedParseBlock(ctx context.Context, blksBytes [][]byte) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VMClient) GetBlockIDAtHeight(ctx context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VMClient) StateSyncEnabled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (vm *VMClient) GetOngoingSyncStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *VMClient) GetLastStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *VMClient) ParseStateSummary(ctx context.Context, summaryBytes []byte) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *VMClient) GetStateSummary(ctx context.Context, summaryHeight uint64) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *VMClient) newBlockFromBuildBlock(resp *vmpb.BuildBlockResponse) (*blockClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type blockClient struct {
	vm *VMClient

	id                  ids.ID
	parentID            ids.ID
	bytes               []byte
	height              uint64
	time                time.Time
	shouldVerifyWithCtx bool
}

func (b *blockClient) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *blockClient) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *blockClient) Reject(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *blockClient) Parent() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *blockClient) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *blockClient) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *blockClient) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *blockClient) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *blockClient) ShouldVerifyWithContext(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *blockClient) VerifyWithContext(ctx context.Context, blockCtx *block.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type summaryClient struct {
	vm *VMClient

	id     ids.ID
	height uint64
	bytes  []byte
}

func (s *summaryClient) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *summaryClient) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *summaryClient) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *summaryClient) Accept(ctx context.Context) (block.StateSyncMode, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSyncMode), nil
}
