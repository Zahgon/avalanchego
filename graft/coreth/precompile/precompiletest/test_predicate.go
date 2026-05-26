// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompiletest

import (
	"testing"

	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

// PredicateTest defines a unit test/benchmark for verifying a precompile predicate.
type PredicateTest struct {
	Name string

	Config precompileconfig.Config

	PredicateContext *precompileconfig.PredicateContext

	Predicate   predicate.Predicate
	Rules       precompileconfig.Rules
	Gas         uint64
	GasErr      error
	ExpectedErr error
}

func (test PredicateTest) Run(t testing.TB) { _ = "STUB: not implemented"; return }

func RunPredicateTests(t *testing.T, tests []PredicateTest) { _ = "STUB: not implemented"; return }

func (test PredicateTest) RunBenchmark(b *testing.B) { _ = "STUB: not implemented"; return }

// Keep it as uint64, multiply 100 to get two digit float later

func RunPredicateBenchmarks(b *testing.B, tests []PredicateTest) { _ = "STUB: not implemented"; return }
