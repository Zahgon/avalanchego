// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/network"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	snowmanblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	blockbuilder "github.com/ava-labs/avalanchego/vms/platformvm/block/builder"
	blockexecutor "github.com/ava-labs/avalanchego/vms/platformvm/block/executor"
	platformvmmetrics "github.com/ava-labs/avalanchego/vms/platformvm/metrics"
)

var (
	_ snowmanblock.ChainVM                         = (*VM)(nil)
	_ snowmanblock.BuildBlockWithContextChainVM    = (*VM)(nil)
	_ snowmanblock.SetPreferenceWithContextChainVM = (*VM)(nil)
	_ secp256k1fx.VM                               = (*VM)(nil)
	_ validators.State                             = (*VM)(nil)
)

type VM struct {
	config.Internal
	blockbuilder.Builder
	*network.Network
	validators.State

	metrics platformvmmetrics.Metrics

	// Used to get time. Useful for faking time during tests.
	clock mockable.Clock

	uptimeManager uptime.Manager

	// The context of this vm
	ctx *snow.Context
	db  database.Database

	state *state.State

	fx            fx.Fx
	codecRegistry codec.Registry

	// Bootstrapped remembers if this chain has finished bootstrapping or not
	bootstrapped utils.Atomic[bool]

	manager blockexecutor.Manager

	// Cancelled on shutdown
	onShutdownCtx context.Context
	// Call [onShutdownCtxCancel] to cancel [onShutdownCtx] during Shutdown()
	onShutdownCtxCancel context.CancelFunc
}

// Initialize this blockchain.
// [vm.ChainManager] and [vm.vdrMgr] must be set before this function is called.
func (vm *VM) Initialize(
	ctx context.Context,
	chainCtx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	_ []byte,
	configBytes []byte,
	_ []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize metrics as soon as possible

// Note: this codec is never used to serialize anything

// TODO: Wait for this goroutine to exit during Shutdown once the platformvm
// has better control of the context lock.

// Create all of the chains that the database says exist

// Incrementing [awaitShutdown] would cause a deadlock since
// [periodicallyPruneMempool] grabs the context lock.

func (vm *VM) periodicallyPruneMempool(frequency time.Duration) { _ = "STUB: not implemented"; return }

func (vm *VM) pruneMempool() error { _ = "STUB: not implemented"; return nil }

// Packing all of the transactions in order performs additional checks that
// the MempoolTxVerifier doesn't include. So, evicting transactions from
// here is expected to happen occasionally.

// Create all chains that exist that this node validates.
func (vm *VM) initBlockchains() error { _ = "STUB: not implemented"; return nil }

// Create the subnet with ID [subnetID]
func (vm *VM) createSubnet(subnetID ids.ID) error { _ = "STUB: not implemented"; return nil }

// onBootstrapStarted marks this VM as bootstrapping
func (vm *VM) onBootstrapStarted() error { _ = "STUB: not implemented"; return nil }

// onNormalOperationsStarted marks this VM as bootstrapped
func (vm *VM) onNormalOperationsStarted() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) SetState(_ context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown this blockchain
func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (vm *VM) ParseBlock(_ context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	// Note: blocks to be parsed are not verified, so we must used blocks.Codec
	// rather than blocks.GenesisCodec
	return *new(snowman.Block), nil
}

func (vm *VM) GetBlock(_ context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// LastAccepted returns the block most recently accepted
func (vm *VM) LastAccepted(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// SetPreference sets the preferred block to be the one with ID [blkID]
func (vm *VM) SetPreference(_ context.Context, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) SetPreferenceWithContext(_ context.Context, blkID ids.ID, blockCtx *snowmanblock.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*VM) Version(context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CreateHandlers returns a map where:
// * keys are API endpoint extensions
// * values are API handlers
func (vm *VM) CreateHandlers(context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*VM) NewHTTPHandler(context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (vm *VM) Connected(ctx context.Context, nodeID ids.NodeID, version *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) CodecRegistry() codec.Registry {
	_ = "STUB: not implemented"
	return *new(codec.Registry)
}

func (vm *VM) Clock() *mockable.Clock { _ = "STUB: not implemented"; return nil }

func (vm *VM) Logger() logging.Logger { _ = "STUB: not implemented"; return *new(logging.Logger) }

func (vm *VM) GetBlockIDAtHeight(_ context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VM) issueTxFromRPC(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }
