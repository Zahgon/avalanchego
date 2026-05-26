// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tracers

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/internal/ethapi"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/eth/tracers/logger"
	"github.com/ava-labs/libevm/ethdb"
)

const (
	// defaultTraceTimeout is the amount of time a single transaction can execute
	// by default before being forcefully aborted.
	defaultTraceTimeout = 5 * time.Second

	// defaultTraceReexec is the number of blocks the tracer is willing to go back
	// and reexecute to produce missing historical state necessary to run a specific
	// trace.
	defaultTraceReexec = uint64(128)

	// defaultTracechainMemLimit is the size of the triedb, at which traceChain
	// switches over and tries to use a disk-backed database instead of building
	// on top of memory.
	// For non-archive nodes, this limit _will_ be overblown, as disk-backed tries
	// will only be found every ~15K blocks or so.
	defaultTracechainMemLimit = common.StorageSize(500 * 1024 * 1024)

	// maximumPendingTraceStates is the maximum number of states allowed waiting
	// for tracing. The creation of trace state will be paused if the unused
	// trace states exceed this limit.
	maximumPendingTraceStates = 128
)

var errTxNotFound = errors.New("transaction not found")

// StateReleaseFunc is used to deallocate resources held by constructing a
// historical state for tracing purposes.
type StateReleaseFunc func()

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error)
	BadBlocks() ([]*types.Block, []*core.BadBlockReason)
	GetTransaction(ctx context.Context, txHash common.Hash) (bool, *types.Transaction, common.Hash, uint64, uint64, error)
	RPCGasCap() uint64
	ChainConfig() *params.ChainConfig
	Engine() consensus.Engine
	ChainDb() ethdb.Database
	StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, StateReleaseFunc, error)
	StateAtNextBlock(ctx context.Context, parent, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, StateReleaseFunc, error)
	StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, StateReleaseFunc, error)
}

// baseAPI holds the collection of common methods for API and FileTracerAPI.
type baseAPI struct {
	backend Backend
}

// API is the collection of tracing APIs exposed over the private debugging endpoint.
type API struct {
	baseAPI
}

// NewAPI creates a new API definition for the tracing methods of the Ethereum service.
func NewAPI(backend Backend) *API { _ = "STUB: not implemented"; return nil }

// FileTracerAPI is the collection of additional tracing APIs exposed over the private
// debugging endpoint that log their output to a file.
type FileTracerAPI struct {
	baseAPI
}

// NewFileTracerAPI creates a new API definition for the tracing methods of the Ethererum
// service that log their output to a file.
func NewFileTracerAPI(backend Backend) *FileTracerAPI { _ = "STUB: not implemented"; return nil }

// chainContext constructs the context reader which is used by the evm for reading
// the necessary chain context.
func (api *baseAPI) chainContext(ctx context.Context) core.ChainContext {
	_ = "STUB: not implemented"
	return *new(core.ChainContext)
}

// blockByNumber is the wrapper of the chain access function offered by the backend.
// It will return an error if the block is not found.
func (api *baseAPI) blockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// blockByHash is the wrapper of the chain access function offered by the backend.
// It will return an error if the block is not found.
func (api *baseAPI) blockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// blockByNumberAndHash is the wrapper of the chain access function offered by
// the backend. It will return an error if the block is not found.
//
// Note this function is friendly for the light client which can only retrieve the
// historical(before the CHT) header/block by number.
func (api *baseAPI) blockByNumberAndHash(ctx context.Context, number rpc.BlockNumber, hash common.Hash) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceConfig holds extra parameters to trace functions.
type TraceConfig struct {
	*logger.Config
	Tracer  *string
	Timeout *string
	Reexec  *uint64
	// Config specific to given tracer. Note struct logger
	// config are historically embedded in main object.
	TracerConfig json.RawMessage
}

// TraceCallConfig is the config for traceCall API. It holds one more
// field to override the state for tracing.
type TraceCallConfig struct {
	TraceConfig
	StateOverrides *ethapi.StateOverride
	BlockOverrides *ethapi.BlockOverrides
	TxIndex        *hexutil.Uint
}

// StdTraceConfig holds extra parameters to standard-json trace functions.
type StdTraceConfig struct {
	logger.Config
	Reexec *uint64
	TxHash common.Hash
}

// txTraceResult is the result of a single transaction trace.
type txTraceResult struct {
	TxHash common.Hash `json:"txHash"`           // transaction hash
	Result interface{} `json:"result,omitempty"` // Trace results produced by the tracer
	Error  string      `json:"error,omitempty"`  // Trace failure produced by the tracer
}

func (t *txTraceResult) String() string { _ = "STUB: not implemented"; return "" }

// blockTraceTask represents a single block trace task when an entire chain is
// being traced.
type blockTraceTask struct {
	statedb *state.StateDB   // Intermediate state prepped for tracing
	block   *types.Block     // Block to trace the transactions from
	release StateReleaseFunc // The function to release the held resource for this task
	results []*txTraceResult // Trace results procudes by the task
}

// blockTraceResult represets the results of tracing a single block when an entire
// chain is being traced.
type blockTraceResult struct {
	Block  hexutil.Uint64   `json:"block"`  // Block number corresponding to this trace
	Hash   common.Hash      `json:"hash"`   // Block hash corresponding to this trace
	Traces []*txTraceResult `json:"traces"` // Trace results produced by the task
}

// txTraceTask represents a single transaction trace task when an entire block
// is being traced.
type txTraceTask struct {
	statedb *state.StateDB // Intermediate state prepped for tracing
	index   int            // Transaction offset in the block
}

// TraceChain returns the structured logs created during the execution of EVM
// between two blocks (excluding start) and returns them as a JSON object.
func (api *API) TraceChain(ctx context.Context, start, end rpc.BlockNumber, config *TraceConfig) (*rpc.Subscription, error) {
	_ = "STUB: not implemented" // Fetch the block interval that we want to trace
	return nil, nil
}

// Tracing a chain is a **long** operation, only do with subscriptions

// traceChain configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The tracing chain range includes
// the end block but excludes the start one. The return value will be one item per
// transaction, dependent on the requested tracer.
// The tracing procedure should be aborted in case the closed signal is received.
func (api *API) traceChain(start, end *types.Block, config *TraceConfig, closed <-chan interface{}) chan *blockTraceResult {
	_ = "STUB: not implemented"
	return nil
}

// Fetch and execute the block trace taskCh

// Trace all the transactions contained within

// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect

// Tracing state is used up, queue it for de-referencing. Note the
// state is the parent state of trace block, use block.number-1 as
// the state number.

// Stream the result back to the result catcher or abort on teardown

// Start a goroutine to feed all the blocks into the tracers

// Ensure everything is properly cleaned up on any exit path

// Clean out any pending release functions of trace states.

// Log the chain result

// Feed all the blocks both into the tracer, as well as fast process concurrently

// Stop tracing if interruption was requested

// Print progress logs if long enough time elapsed

// Retrieve the parent block and target block for tracing.

// Make sure the state creator doesn't go too far. Too many unprocessed
// trace state may cause the oldest state to become stale(e.g. in
// path-based scheme).

// Prepare the statedb for tracing. Don't use the live database for
// tracing to avoid persisting state junks into the database. Switch
// over to `preferDisk` mode only if the memory usage exceeds the
// limit, the trie database will be reconstructed from scratch only
// if the relevant state is available in disk.

// Clean out any pending release functions of trace state. Note this
// step must be done after constructing tracing state, because the
// tracing state of block next depends on the parent state and construction
// may fail if we release too early.

// Send the block over to the concurrent tracers (if not in the fast-forward phase)

// Keep reading the trace results and stream them to result channel.

// Queue up next received result

// Stream completed traces to the result channel

// It will be blocked in case the channel consumer doesn't take the
// tracing result in time(e.g. the websocket connect is not stable)
// which will eventually block the entire chain tracer. It's the
// expected behavior to not waste node resources for a non-active user.

// TraceBlockByNumber returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) TraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceBlockByHash returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) TraceBlockByHash(ctx context.Context, hash common.Hash, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceBlock returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *baseAPI) TraceBlock(ctx context.Context, blob hexutil.Bytes, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceBlockFromFile returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *FileTracerAPI) TraceBlockFromFile(ctx context.Context, file string, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceBadBlock returns the structured logs created during the execution of
// EVM against a block pulled from the pool of bad ones and returns them as a JSON
// object.
func (api *API) TraceBadBlock(ctx context.Context, hash common.Hash, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	// Search for the bad block corresponding to [hash].
	return nil, nil
}

// StandardTraceBlockToFile dumps the structured logs created during the
// execution of EVM to the local file system and returns a list of files
// to the caller.
func (api *FileTracerAPI) StandardTraceBlockToFile(ctx context.Context, hash common.Hash, config *StdTraceConfig) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IntermediateRoots executes a block (bad- or canon- or side-), and returns a list
// of intermediate roots: the stateroot after each transaction.
func (api *API) IntermediateRoots(ctx context.Context, hash common.Hash, config *TraceConfig) ([]common.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We intentionally don't return the error here: if we do, then the RPC server will not
// return the roots. Most likely, the caller already knows that a certain transaction fails to
// be included, but still want the intermediate roots that led to that point.
// It may happen the tx_N causes an erroneous state, which in turn causes tx_N+M to not be
// executable.
// N.B: This should never happen while tracing canon blocks, only when tracing bad blocks.

// calling IntermediateRoot will internally call Finalize on the state
// so any modifications are written to the trie

// StandardTraceBadBlockToFile dumps the structured logs created during the
// execution of EVM against a block pulled from the pool of bad ones to the
// local file system and returns a list of files to the caller.
func (api *FileTracerAPI) StandardTraceBadBlockToFile(ctx context.Context, hash common.Hash, config *StdTraceConfig) ([]string, error) {
	_ = "STUB: not implemented"
	// Search for the bad block corresponding to [hash].
	return nil, nil
}

// traceBlock configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requested tracer.
func (api *baseAPI) traceBlock(ctx context.Context, block *types.Block, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare base state

// JS tracers have high overhead. In this case run a parallel
// process that generates states in one thread and traces txes
// in separate worker threads.

// Native tracers have low overhead

// Generate the next state snapshot fast without tracing

// Finalize the state so any modifications are written to the trie
// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect

// traceBlockParallel is for tracers that have a high overhead (read JS tracers). One thread
// runs along and executes txes without tracing enabled to generate their prestate.
// Worker threads take the tasks and the prestate and trace them.
func (api *baseAPI) traceBlockParallel(ctx context.Context, block *types.Block, statedb *state.StateDB, config *TraceConfig) ([]*txTraceResult, error) {
	_ = "STUB: not implemented"
	// Execute all the transaction contained within the block concurrently
	return nil, nil
}

// Fetch and execute the next transaction trace tasks

// Feed the transactions into the tracers and return

// Send the trace task over for execution

// Generate the next state snapshot fast without tracing

// Finalize the state so any modifications are written to the trie
// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect

// If execution failed in between, abort

// standardTraceBlockToFile configures a new tracer which uses standard JSON output,
// and traces either a full block or an individual transaction. The return value will
// be one filename per transaction traced.
func (api *FileTracerAPI) standardTraceBlockToFile(ctx context.Context, block *types.Block, config *StdTraceConfig) ([]string, error) {
	_ = "STUB: not implemented"
	// If we're tracing a single transaction, make sure it's present
	return nil, nil
}

// Retrieve the tracing configurations, or use default values

// Execute transaction, either tracing all or just the requested one

// Check if there are any overrides: the caller may wish to enable a future
// fork when executing this block. Note, such overrides are only applicable to the
// actual specified block, not any preceding blocks that we have to go through
// in order to obtain the state.
// Therefore, it's perfectly valid to specify `"futureForkBlock": 0`, to enable `futureFork`

// Note: This copies the config, to not screw up the main config

// Prepare the transaction for un-traced execution

// If the transaction needs tracing, swap out the configs

// Generate a unique temporary file to dump it into

// Swap out the noop logger to the standard tracer

// Execute the transaction and flush any traces to disk

// Finalize the state so any modifications are written to the trie
// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect

// If we've traced the transaction we were looking for, abort

// containsTx reports whether the transaction with a certain hash
// is contained within the specified block.
func containsTx(block *types.Block, hash common.Hash) bool { _ = "STUB: not implemented"; return false }

// TraceTransaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *API) TraceTransaction(ctx context.Context, hash common.Hash, config *TraceConfig) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only mined txes are supported

// It shouldn't happen in practice.

// TraceCall lets you trace a given eth_call. It collects the structured logs
// created during the execution of EVM if the given transaction was added on
// top of the provided block and returns them as a JSON object.
// If no transaction index is specified, the trace will be conducted on the state
// after executing the specified block. However, if a transaction index is provided,
// the trace will be conducted on the state after executing the specified transaction
// within the specified block.
func (api *API) TraceCall(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, config *TraceCallConfig) (interface{}, error) {
	_ = "STUB: not implemented"
	// Try to retrieve the specified block
	return nil, nil
}

// We don't have access to the miner here. For tracing 'future' transactions,
// it can be done with block- and state-overrides instead, which offers
// more flexibility and stability than trying to trace on 'pending', since
// the contents of 'pending' is unstable and probably not a true representation
// of what the next actual block is likely to contain.

// try to recompute the state

// Apply the customization rules if required.

// Overriding the block number to n+1 is a common way for wallets to
// simulate transactions, however without the following fix, a contract
// can assert it is being simulated by checking if blockhash(n) == 0x0 and
// can behave differently during the simulation. (#32175 for more info)
// --
// Modify the parent hash and number so that downstream, blockContext's
// GetHash function can correctly return n.

// Apply all relevant upgrades from [originalTime] to the block time set in the override.
// Should be applied before the state overrides.

// Execute the trace

// traceTx configures a new tracer according to the provided configuration, and
// executes the given message in the provided environment. The return value will
// be tracer dependent.
func (api *baseAPI) traceTx(ctx context.Context, message *core.Message, txctx *Context, vmctx vm.BlockContext, statedb *state.StateDB, config *TraceConfig) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default tracer is the struct logger

// Define a meaningful timeout of a single transaction trace

// Stop evm execution. Note cancellation is not necessarily immediate.

// Call Prepare to clear out the statedb access list

// APIs return the collection of RPC services the tracer package offers.
func APIs(backend Backend) []rpc.API {
	_ = "STUB: not implemented"
	// Append all the local APIs and return
	return nil
}

// overrideConfig returns a copy of [original] with network upgrades enabled by [override] enabled,
// along with a boolean that indicates whether the copy is canonical (equivalent to the original).
func overrideConfig(original *params.ChainConfig, override *params.ChainConfig) (*params.ChainConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
