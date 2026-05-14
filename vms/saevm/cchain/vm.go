// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/triedb"
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
	"github.com/ava-labs/avalanchego/graft/evm/utils/rpc"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/evm/acp226"
	"github.com/ava-labs/avalanchego/vms/evm/database"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/api"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/state"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/txpool"
	"github.com/ava-labs/avalanchego/vms/saevm/sae"

	avadb "github.com/ava-labs/avalanchego/database"
	corethparams "github.com/ava-labs/avalanchego/graft/coreth/params"
)

// VM is a harness around an [sae.VM], providing an `Initialize`
// method that supports being asynchronous since genesis or after a previously
// accepted synchronous block.
type VM struct {
	*sae.VM // created by [VM.Initialize]

	ctx    *snow.Context
	state  *state.State
	txpool *txpool.Txpool

	// onClose are executed in reverse order during [SinceGenesis.Shutdown].
	// If a resource depends on another resource, it MUST be added AFTER the
	// resource it depends on.
	onClose []func() error
}

var ethDBPrefix = []byte("ethdb")

// Initialize initializes the VM.
func (v *VM) Initialize(
	ctx context.Context,
	snowCtx *snow.Context,
	avaDB avadb.Database,
	genesisBytes []byte,
	_ []byte,
	configBytes []byte,
	_ []*common.Fx,
	appSender common.AppSender,
) error {
	// [prefixdb.NewNested] is used because coreth used to be run as a plugin.
	// This meant that the database's prefix was not compacted, because the
	// provided database was wrapped by the rpcchainvm.
	db := rawdb.NewDatabase(database.New(prefixdb.NewNested(ethDBPrefix, avaDB)))
	tdb := triedb.NewDatabase(db, v.config.DBConfig.TrieDBConfig)

	snowCtx.Log.Info("parsing genesis")

	genesis, err := parseGenesis(snowCtx, genesisBytes)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(%T): %w", genesis, err)
	}

	snowCtx.Log.Info("establishing last synchronous block")

	var lastSync *types.Block
	lastSyncBytes, err := state.ReadLastSync(avaDB)
	switch {
	case err == nil:
		lastSync = new(types.Block)
		if err := rlp.DecodeBytes(lastSyncBytes, lastSync); err != nil {
			return fmt.Errorf("rlp.DecodeBytes(..., %T): %w", lastSync, err)
		}
	case errors.Is(err, avadb.ErrNotFound):
		lastSync = genesis.ToBlock()
	default:
		return err
	}

	snowCtx.Log.Info("setting up the genesis",
		zap.Stringer("lastID", ids.ID(lastSync.Hash())),
		zap.Uint64("lastHeight", lastSync.NumberU64()),
	)

	// TODO: Are these reasonable?
	config, _, err := core.SetupGenesisBlock(db, tdb, genesis, lastSync.Hash(), false)
	if err != nil {
		return fmt.Errorf("core.SetupGenesisBlock(...): %w", err)
	}

	snowCtx.Log.Info("constructing cross-chain state")

	cchainState, err := state.New(snowCtx, avaDB)
	if err != nil {
		return fmt.Errorf("creating cchain state: %w", err)
	}
	v.onClose = append(v.onClose, cchainState.Close)

	snowCtx.Log.Info("parsing user config")

	userConfig, err := ParseConfig(configBytes)
	if err != nil {
		return err
	}

	snowCtx.Log.Info("parsing warp message overrides")

	warpMessages, err := userConfig.WarpMessages()
	if err != nil {
		return err
	}

	var desiredDelayExcess *acp226.DelayExcess
	if userConfig.MinDelayTarget != nil {
		desiredDelayExcess = new(acp226.DelayExcess)
		*desiredDelayExcess = acp226.DesiredDelayExcess(*userConfig.MinDelayTarget)
	}
	var desiredTargetExcess *acp176.TargetExcess
	if userConfig.GasTarget != nil {
		desiredTargetExcess = new(acp176.TargetExcess)
		*desiredTargetExcess = acp176.DesiredTargetExcess(*userConfig.GasTarget)
	}

	pendingTxs := txpool.NewPending()
	warpStorage := saewarp.NewStorage(avaDB, warpMessages...)

	hooks := hook.NewPoints(
		snowCtx,
		cchainState,
		config,
		desiredDelayExcess,
		desiredTargetExcess,
		pendingTxs,
		warpStorage,
	)

	snowCtx.Log.Info("constructing the sae VM")

	inner, err := sae.NewVM(ctx, hooks, v.config, snowCtx, config, db, lastSync, appSender)
	if err != nil {
		return err
	}
	v.VM = inner
	v.ctx = snowCtx
	v.state = cchainState

	v.mempool, err = txpool.New(snowCtx, config, pendingTxs, inner, 1024)
	if err != nil {
		return fmt.Errorf("creating txpool: %w", err)
	}
	v.onClose = append(v.onClose, func() error {
		v.mempool.Close()
		return nil
	})

	snowCtx.Log.Info("initialized saevm")

	return nil
}

// TODO: copied from coreth
func parseGenesis(ctx *snow.Context, bytes []byte) (*core.Genesis, error) {
	g := new(core.Genesis)
	if err := json.Unmarshal(bytes, g); err != nil {
		return nil, fmt.Errorf("parsing genesis: %w", err)
	}

	// Populate the Avalanche config extras.
	configExtra := corethparams.GetExtra(g.Config)
	configExtra.AvalancheContext = extras.AvalancheContext{
		SnowCtx: ctx,
	}
	configExtra.NetworkUpgrades = extras.GetNetworkUpgrades(ctx.NetworkUpgrades)

	// If Durango is scheduled, schedule the Warp Precompile at the same time.
	if configExtra.DurangoBlockTimestamp != nil {
		configExtra.PrecompileUpgrades = append(configExtra.PrecompileUpgrades, extras.PrecompileUpgrade{
			Config: warpcontract.NewDefaultConfig(configExtra.DurangoBlockTimestamp),
		})
	}
	if err := configExtra.Verify(); err != nil {
		return nil, fmt.Errorf("invalid chain config: %w", err)
	}

	// Align all the Ethereum upgrades to the Avalanche upgrades
	if err := corethparams.SetEthUpgrades(g.Config); err != nil {
		return nil, fmt.Errorf("setting eth upgrades: %w", err)
	}
	return g, nil
}

const (
	avaxServiceName       = "avax"
	avaxHTTPExtensionPath = "/" + avaxServiceName
)

func (v *VM) CreateHandlers(ctx context.Context) (map[string]http.Handler, error) {
	m, err := v.VM.CreateHandlers(ctx)
	if err != nil {
		return nil, err
	}

	service := api.NewService(v.ctx, v.GethRPCBackends(), v.mempool, v.pushGossiper, v.state)
	handler, err := rpc.NewHandler(avaxServiceName, service)
	if err != nil {
		return nil, fmt.Errorf("rpc.NewHandler(%s, ...): %w", avaxServiceName, err)
	}

	m[avaxHTTPExtensionPath] = handler
	return m, nil
}

// Prevent busy looping when the chain is more advanced than the mempool.
const waitForEventDelay = 100 * time.Millisecond

var errNoPreference = errors.New("no preferred block")

// WaitForEvent waits for a transaction to be in the txpool or for the SAE VM to
// produce an event.
func (v *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	// TODO(StephenButtolph): Do not busy loop with [common.PendingTxs]. The
	// mempools are cleared after block execution, so we may still have
	// transactions in the mempool while blocks containing those transactions
	// are processing.
	// TODO(StephenButtolph): Wait until we are allowed to build a block.

	ctx, cancel := context.WithCancel(ctx)
	type result struct {
		msg common.Message
		err error
	}
	results := make(chan result, 2)
	go func() {
		defer cancel()
		msg, err := v.VM.WaitForEvent(ctx)
		results <- result{msg, err}
	}()
	go func() {
		defer cancel()
		err := v.mempool.AwaitTxs(ctx)
		results <- result{common.PendingTxs, err}
	}()

	r := <-results
	return r.msg, r.err
}

func (v *VM) RejectBlock(ctx context.Context, b *blocks.Block) error {
	// If the block is rejected, the transactions might get dropped from the
	// network. If the transactions are still valid, it is a better UX to add
	// them into our mempool.
	txs, err := tx.ParseSlice(customtypes.BlockExtData(b.EthBlock()))
	if err != nil {
		return fmt.Errorf("parsing txs: %w", err)
	}
	for _, tx := range txs {
		_ = v.mempool.Add(tx)
	}
	return v.VM.RejectBlock(ctx, b)
}

func (v *VM) Shutdown(ctx context.Context) error {
	errs := make([]error, len(v.onClose))
	for i, f := range slices.Backward(v.onClose) {
		errs[i] = f()
	}
	if err := errors.Join(errs...); err != nil {
		return fmt.Errorf("closing resources: %w", err)
	}

	if v.VM == nil {
		return nil
	}
	return v.VM.Shutdown(ctx)
}
