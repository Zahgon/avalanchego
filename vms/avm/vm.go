// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/avalanchego/vms/avm/block"
	"github.com/ava-labs/avalanchego/vms/avm/config"
	"github.com/ava-labs/avalanchego/vms/avm/network"
	"github.com/ava-labs/avalanchego/vms/avm/state"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/avm/utxo"
	"github.com/ava-labs/avalanchego/vms/components/avax"

	blockbuilder "github.com/ava-labs/avalanchego/vms/avm/block/builder"
	blockexecutor "github.com/ava-labs/avalanchego/vms/avm/block/executor"
	extensions "github.com/ava-labs/avalanchego/vms/avm/fxs"
	avmmetrics "github.com/ava-labs/avalanchego/vms/avm/metrics"
	txexecutor "github.com/ava-labs/avalanchego/vms/avm/txs/executor"
)

var (
	errIncompatibleFx            = errors.New("incompatible feature extension")
	errUnknownFx                 = errors.New("unknown feature extension")
	errGenesisAssetMustHaveState = errors.New("genesis asset must have non-empty state")

	_ vertex.LinearizableVMWithEngine = (*VM)(nil)
)

type VM struct {
	network.Atomic

	config.Config

	metrics avmmetrics.Metrics

	avax.AddressManager
	ids.Aliaser
	utxo.Spender

	// Contains information of where this VM is executing
	ctx *snow.Context

	// Used to check local time
	clock mockable.Clock

	registerer prometheus.Registerer

	connectedPeers map[ids.NodeID]*version.Application

	parser block.Parser

	appSender common.AppSender

	// State management
	state state.State

	// asset id that will be used for fees
	feeAssetID ids.ID

	baseDB database.Database
	db     *versiondb.Database

	typeToFxIndex map[reflect.Type]int
	fxs           []*extensions.ParsedFx

	walletService WalletService

	txBackend *txexecutor.Backend

	// Cancelled on shutdown
	onShutdownCtx context.Context
	// Call [onShutdownCtxCancel] to cancel [onShutdownCtx] during Shutdown()
	onShutdownCtxCancel context.CancelFunc
	awaitShutdown       sync.WaitGroup

	networkConfig network.Config
	// These values are only initialized after the chain has been linearized.
	blockbuilder.Builder
	chainManager blockexecutor.Manager
	network      *network.Network
}

func (vm *VM) Connected(ctx context.Context, nodeID ids.NodeID, version *version.Application) error {
	_ = "STUB: not implemented"
	// If the chain isn't linearized yet, we must track the peers externally
	// until the network is initialized.
	return nil
}

func (vm *VM) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	// If the chain isn't linearized yet, we must track the peers externally
	// until the network is initialized.
	return nil
}

/*
 ******************************************************************************
 ********************************* Common VM **********************************
 ******************************************************************************
 */

func (vm *VM) Initialize(
	_ context.Context,
	ctx *snow.Context,
	db database.Database,
	genesisBytes []byte,
	_ []byte,
	configBytes []byte,
	fxs []*common.Fx,
	appSender common.AppSender,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize metrics as soon as possible

// onBootstrapStarted is called by the consensus engine when it starts bootstrapping this chain
func (vm *VM) onBootstrapStarted() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) onNormalOperationsStarted() error { _ = "STUB: not implemented"; return nil }

func (vm *VM) SetState(_ context.Context, state snow.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*VM) Version(context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (vm *VM) CreateHandlers(context.Context) (map[string]http.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// name this service "avm"

// name this service "wallet"

func (*VM) NewHTTPHandler(context.Context) (http.Handler, error) {
	_ = "STUB: not implemented"

	/*
	 ******************************************************************************
	 ********************************** Chain VM **********************************
	 ******************************************************************************
	 */return *new(http.Handler), nil
}

func (vm *VM) GetBlock(_ context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) ParseBlock(_ context.Context, blkBytes []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (vm *VM) SetPreference(_ context.Context, blkID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (vm *VM) LastAccepted(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (vm *VM) GetBlockIDAtHeight(_ context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

/*
 ******************************************************************************
 *********************************** DAG VM ***********************************
 ******************************************************************************
 */

func (vm *VM) Linearize(ctx context.Context, stopVertexID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: The context lock is not held when calling network.IssueTx.

// Notify the network of our current peers

// Note: It's important only to switch the networking stack after the full
// chainVM has been initialized. Traffic will immediately start being
// handled asynchronously.

// Invariant: PushGossip must never grab the context lock.

// Invariant: PullGossip must never grab the context lock.

func (vm *VM) ParseTx(_ context.Context, bytes []byte) (snowstorm.Tx, error) {
	_ = "STUB: not implemented"
	return *new(snowstorm.Tx), nil
}

/*
 ******************************************************************************
 ********************************** JSON API **********************************
 ******************************************************************************
 */

// issueTxFromRPC attempts to send a transaction to consensus.
//
// Invariant: The context lock is not held
// Invariant: This function is only called after Linearize has been called.
func (vm *VM) issueTxFromRPC(tx *txs.Tx) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

/*
 ******************************************************************************
 ********************************** Helpers ***********************************
 ******************************************************************************
 */

func (vm *VM) initGenesis(genesisBytes []byte) error { _ = "STUB: not implemented"; return nil }

// secure this by defaulting to avaxAsset

func (vm *VM) initState(tx *txs.Tx) { _ = "STUB: not implemented"; return }

// lookupAssetID looks for an ID aliased by [asset] and if it fails
// attempts to parse [asset] into an ID
func (vm *VM) lookupAssetID(asset string) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// Invariant: onAccept is called when [tx] is being marked as accepted, but
// before its state changes are applied.
// TODO: Remove [onAccept] once the deprecated APIs this powers are removed.
func (vm *VM) onAccept(tx *txs.Tx) { _ = "STUB: not implemented"; return }
