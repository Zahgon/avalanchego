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

	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/graft/evm/utils/rpc"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/evm/database"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/state"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/txpool"
	"github.com/ava-labs/avalanchego/vms/saevm/sae"

	avadb "github.com/ava-labs/avalanchego/database"
)

// VM is a harness around an [sae.VM], providing an `Initialize`
// method that supports being asynchronous since genesis or after a previously
// accepted synchronous block.
type VM struct {
	*sae.VM // created by [VM.Initialize]

	config sae.Config
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

	// TODO(StephenButtolph): Replace this with Coreth's genesis format.
	genesis := new(core.Genesis)
	if err := json.Unmarshal(genesisBytes, genesis); err != nil {
		return fmt.Errorf("json.Unmarshal(%T): %v", genesis, err)
	}
	config, _, err := core.SetupGenesisBlock(db, tdb, genesis)
	if err != nil {
		return fmt.Errorf("core.SetupGenesisBlock(...): %v", err)
	}

	cchainState, err := state.New(snowCtx, avaDB)
	if err != nil {
		return fmt.Errorf("creating cchain state: %w", err)
	}
	v.onClose = append(v.onClose, cchainState.Close)

	pendingTxs := txpool.NewPending()
	hooks := newHooks(
		snowCtx,
		cchainState,
		pendingTxs,
	)
	inner, err := sae.NewVM(ctx, hooks, v.config, snowCtx, config, db, genesis.ToBlock(), appSender)
	if err != nil {
		return err
	}
	v.VM = inner
	v.ctx = snowCtx
	v.state = cchainState

	const maxTxPoolSize = 1024
	v.txpool, err = txpool.New(snowCtx, config, pendingTxs, inner, maxTxPoolSize)
	if err != nil {
		return fmt.Errorf("creating txpool: %w", err)
	}
	v.onClose = append(v.onClose, func() error {
		v.txpool.Close()
		return nil
	})
	return nil
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

	service, err := newService(v.ctx, v.txpool, v.state)
	if err != nil {
		return nil, fmt.Errorf("creating avax service: %w", err)
	}
	handler, err := rpc.NewHandler(avaxServiceName, service)
	if err != nil {
		return nil, fmt.Errorf("rpc.NewHandler(%s, ...): %w", avaxServiceName, err)
	}

	m[avaxHTTPExtensionPath] = handler
	return m, nil
}

// WaitForEvent waits for a transaction to be in the txpool or for the SAE VM to
// produce an event.
func (v *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	// TODO(StephenButtolph): Do not busy loop with [common.PendingTxs]. The
	// txpools are cleared after block execution, so we may still have
	// transactions in the txpool while blocks containing those transactions are
	// processing.

	// TODO(StephenButtolph): Wait until the minimum block delay has passed.

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
		err := v.txpool.AwaitTxs(ctx)
		results <- result{common.PendingTxs, err}
	}()

	r := <-results
	return r.msg, r.err
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
