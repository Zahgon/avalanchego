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

package eth

import (
	"context"

	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/eth/tracers"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
)

// noopReleaser is returned in case there is no operation expected
// for releasing state.
var noopReleaser = tracers.StateReleaseFunc(func() {})

func (eth *Ethereum) hashState(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (statedb *state.StateDB, release tracers.StateReleaseFunc, err error) {
	_ = "STUB: not implemented"
	// Do not support re-executing historical blocks to grab state
	return nil, *new(tracers.StateReleaseFunc), nil
}

// The state is only for reading purposes, check the state presence in
// live database.

// The state is available in live database, create a reference
// on top to prevent garbage collection and return a release
// function to deref it.

// The state is both for reading and writing, or it's unavailable in disk,
// try to construct/recover the state over an ephemeral trie.Database for
// isolating the live one.

// Create an ephemeral trie.Database for isolating the live one. Otherwise
// the internal junks created by tracing will be persisted into the disk.
// TODO(rjl493456442), clean cache is disabled to prevent memory leak,
// please re-enable it for better performance.

// The optional base statedb is given, mark the start point as parent block

// Otherwise, try to reexec blocks until we find a state or reach our limit

// Create an ephemeral trie.Database for isolating the live one. Otherwise
// the internal junks created by tracing will be persisted into the disk.
// TODO(rjl493456442), clean cache is disabled to prevent memory leak,
// please re-enable it for better performance.

// If we didn't check the live database, do check state over ephemeral database,
// otherwise we would rewind past a persisted block (specific corner case is
// chain tracing from the genesis).

// Database does not have the state for the given block, try to regenerate

// State is available at historical point, re-execute the blocks on top for
// the desired state.

// Print progress logs if long enough time elapsed

// Retrieve the next block to regenerate and process it

// Finalize the state so any modifications are written to the trie

// Hold the state reference and also drop the parent state
// to prevent accumulating too many nodes in memory.

// all memory is contained within the nodes return in hashdb

// This is compatible with both PathDB and FirewoodDB schemes.
func (eth *Ethereum) pathState(block *types.Block) (*state.StateDB, func(), error) {
	_ = "STUB: not implemented"
	// Check if the requested state is available in the live chain.
	return nil, nil, nil
}

// TODO historic state is not supported in path-based scheme.
// Fully archive node in pbss will be implemented by relying
// on state history, but needs more work on top.

// firewoodState reconstructs the state at the requested block (`header`) by
// walking back to a persisted revision or genesis, then re-executing blocks
// forward.
//
// The walk-back is bounded by `reexec`. If no persisted revision or genesis is
// found within `reexec` blocks of the requested block, this returns an error.
func (eth *Ethereum) firewoodState(ctx context.Context, header *types.Header, reexec uint64) (_ *state.StateDB, _ tracers.StateReleaseFunc, finalErr error) {
	_ = "STUB: not implemented"
	// Fast path: state is available directly.
	return nil, *new(tracers.StateReleaseFunc), nil
}

// Get the Firewood TrieDB.

// Genesis state is not in Firewood; reconstruct it from the genesis
// spec using an in-memory hash-based trie.

// If the target block is genesis, return the state directly.

// The target block is past genesis, so we need the genesis root
// and header as the starting point for re-execution.

// Get the base revision.

// Create initial Reconstructed from the base revision.

// Create a single accessor for the entire re-execution; the underlying
// Reconstructed is mutated in place so the accessor remains valid.
// Root hashing is deferred until after replay, when the target root is
// validated once against the requested header.

/* computeRootOnHash */

// Re-execute blocks forward from current+1 to the target block.

// If using a reconstructed revision, compute the root hash here as root computation
// was deferred during block reexecution.

// Before returning, reopen a clean StateDB against the same reconstructed view,
// now with normal root computation enabled.

/* computeRootOnHash */

// inMemoryGenesisDB creates an in-memory hash-based trie database populated
// with the committed genesis state.
func (eth *Ethereum) inMemoryGenesisDB() (state.Database, error) {
	_ = "STUB: not implemented"
	return *new(state.Database), nil
}

// stateAtBlock retrieves the state database associated with a certain block.
// If no state is locally available for the given block, a number of blocks
// are attempted to be reexecuted to generate the desired state. The optional
// base layer statedb can be provided which is regarded as the statedb of the
// parent block.
//
// An additional release function will be returned if the requested state is
// available. Release is expected to be invoked when the returned state is no
// longer needed. Its purpose is to prevent resource leaking. Though it can be
// noop in some cases.
//
// Parameters:
//   - block:      The block for which we want the state(state = block.Root)
//   - reexec:     The maximum number of blocks to reprocess trying to obtain the desired state
//   - base:       If the caller is tracing multiple blocks, the caller can provide the parent
//     state continuously from the callsite.
//   - readOnly:   If true, then the live 'blockchain' state database is used. No mutation should
//     be made from caller, e.g. perform Commit or other 'save-to-disk' changes.
//     Otherwise, the trash generated by caller may be persisted permanently.
//   - preferDisk: This arg can be used by the caller to signal that even though the 'base' is
//     provided, it would be preferable to start from a fresh state, if we have it
//     on disk.
func (eth *Ethereum) stateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (statedb *state.StateDB, release tracers.StateReleaseFunc, err error) {
	_ = "STUB: not implemented"
	return nil, *new(tracers.StateReleaseFunc), nil
}

// stateAtTransaction returns the execution environment of a certain transaction.
func (eth *Ethereum) stateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	// Short circuit if it's genesis block.
	return nil, *new(vm.BlockContext), nil, *new(tracers.StateReleaseFunc), nil
}

// Create the parent state database

// Lookup the statedb of parent block from the live database,
// otherwise regenerate it on the flight.

// Recompute transactions up to the target index.

// Assemble the transaction call message and return if the requested offset

// Not yet the searched for transaction, execute on top of the current state

// Ensure any modifications are committed to the state
// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect

// StateAtNextBlock is a helper function that returns the state at the next block.
// It wraps StateAtBlock and handles the case where Upgrades are applied to the
// next block.
// This is different than using StateAtBlock with [nextBlock] because it will
// apply the upgrades to the [parent] state before returning it.
func (eth *Ethereum) StateAtNextBlock(ctx context.Context, parent *types.Block, nextBlock *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, tracers.StateReleaseFunc, error) {
	_ = "STUB: not implemented"
	// Get state for [parent]
	return nil, *new(tracers.StateReleaseFunc), nil
}

// Apply upgrades here for the [nextBlock]
