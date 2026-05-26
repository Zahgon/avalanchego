// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"errors"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

var ErrMissingPredicateContext = errors.New("missing predicate context")

// CheckBlockPredicates verifies the predicates of a block of transactions and
// returns the results.
//
// Returning an error invalidates the block.
func CheckBlockPredicates(
	rules params.Rules,
	predicateContext *precompileconfig.PredicateContext,
	txs []*types.Transaction,
) (predicate.BlockResults, error) {
	_ = "STUB: not implemented"
	return *new(predicate.BlockResults), nil
}

// TODO: Calculate tx predicates concurrently.

// CheckTxPredicates verifies the predicates of a transaction and returns the
// results.
//
// Returning an error invalidates the transaction.
func CheckTxPredicates(
	rules params.Rules,
	predicateContext *precompileconfig.PredicateContext,
	tx *types.Transaction,
) (predicate.PrecompileResults, error) {
	_ = "STUB: not implemented"
	// Check that the transaction can cover its IntrinsicGas, including the gas
	// required by the predicate, before verifying the predicate.
	return *new(predicate.PrecompileResults), nil
}

// Short circuit early if there are no precompile predicates to verify

// Prepare the predicate storage slots from the transaction's access list

// If there are no predicates to verify, return early and skip requiring the
// proposervm block context to be populated.

// Since address is only added to predicateArguments when there's a
// valid predicate in the ruleset there's no need to check if the
// predicate exists here.
