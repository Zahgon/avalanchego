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
// Copyright 2023 The go-ethereum Authors
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

package gasestimator

import (
	"context"

	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
)

// Options are the contextual parameters to execute the requested call.
//
// Whilst it would be possible to pass a blockchain object that aggregates all
// these together, it would be excessively hard to test. Splitting the parts out
// allows testing without needing a proper live chain.
type Options struct {
	Config *params.ChainConfig // Chain configuration for hard fork selection
	Chain  core.ChainContext   // Chain context to access past block hashes
	Header *types.Header       // Header defining the block context to execute in
	State  *state.StateDB      // Pre-state on top of which to estimate the gas

	ErrorRatio float64 // Allowed overestimation ratio for faster estimation termination
}

// Estimate returns the lowest possible gas limit that allows the transaction to
// run successfully with the provided context options. It returns an error if the
// transaction would always revert, or if there are unexpected failures.
func Estimate(ctx context.Context, call *core.Message, opts *Options, gasCap uint64) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	// Binary search the gas limit, as it may need to be higher than the amount used
	return 0, nil, nil
}

// lowest-known gas limit where tx execution fails
// lowest-known gas limit where tx execution succeeds

// Determine the highest gas limit can be used during the estimation.

// Normalize the max fee per gas the call is willing to spend.

// Recap the highest gas limit with account's available balance.

// If the allowance is larger than maximum uint64, skip checking

// Recap the highest gas allowance with specified gascap.

// If the transaction is a plain value transfer, short circuit estimation and
// directly try 21000. Returning 21000 without any execution is dangerous as
// some tx field combos might bump the price up even for plain transfers (e.g.
// unused access list items). Ever so slightly wasteful, but safer overall.

// We first execute the transaction at the highest allowable gas limit, since if this fails we
// can return error immediately.

// For almost any transaction, the gas consumed by the unconstrained execution
// above lower-bounds the gas limit required for it to succeed. One exception
// is those that explicitly check gas remaining in order to execute within a
// given limit, but we probably don't want to return the lowest possible gas
// limit for these cases anyway.

// There's a fairly high chance for the transaction to execute successfully
// with gasLimit set to the first execution's usedGas + gasRefund. Explicitly
// check that gas amount and use as a limit for the binary search.

// This should not happen under normal conditions since if we make it this far the
// transaction had run without error at least once before.

// Binary search for the smallest gas limit that allows the tx to execute successfully.

// It is a bit pointless to return a perfect estimation, as changing
// network conditions require the caller to bump it up anyway. Since
// wallets tend to use 20-25% bump, allowing a small approximation
// error is fine (as long as it's upwards).

// Most txs don't need much higher gas limit than their gas used, and most txs don't
// require near the full block limit of gas, so the selection of where to bisect the
// range here is skewed to favor the low side.

// This should not happen under normal conditions since if we make it this far the
// transaction had run without error at least once before.

// execute is a helper that executes the transaction under a given gas limit and
// returns true if the transaction fails for a reason that might be related to
// not enough gas. A non-nil error means execution failed due to reasons unrelated
// to the gas limit.
func execute(ctx context.Context, call *core.Message, opts *Options, gasLimit uint64) (bool, *core.ExecutionResult, error) {
	_ = "STUB: not implemented"
	// Configure the call for this specific execution (and revert the change after)
	return false, nil, nil
}

// Execute the call and separate execution faults caused by a lack of gas or
// other non-fixable conditions

// Special case, raise gas limit

// Bail out

// run assembles the EVM as defined by the consensus rules and runs the requested
// call invocation.
func run(ctx context.Context, call *core.Message, opts *Options) (*core.ExecutionResult, error) {
	_ = "STUB: not implemented"
	// Assemble the call and the call context
	return nil, nil
}

// Monitor the outer context and interrupt the EVM upon cancellation. To avoid
// a dangling goroutine until the outer estimation finishes, create an internal
// context for the lifetime of this method call.

// Execute the call, returning a wrapped error or the result
