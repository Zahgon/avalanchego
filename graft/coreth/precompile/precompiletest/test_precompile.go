// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompiletest

import (
	"testing"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/coreth/core/extstate"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/contract"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/modules"
	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

// PrecompileTest is a test case for a precompile
type PrecompileTest struct {
	Name string
	// Caller is the address of the precompile caller
	Caller common.Address
	// Input the raw input bytes to the precompile
	Input []byte
	// InputFn is a function that returns the raw input bytes to the precompile
	// If specified, Input will be ignored.
	InputFn func(t testing.TB) []byte
	// SuppliedGas is the amount of gas supplied to the precompile
	SuppliedGas uint64
	// ReadOnly is whether the precompile should be called in read only
	// mode. If true, the precompile should not modify the state.
	ReadOnly bool
	// Config is the config to use for the precompile
	// It should be the same precompile config that is used in the
	// precompile's configurator.
	// If nil, Configure will not be called.
	Config precompileconfig.Config
	// Predicates that the precompile should have access to.
	Predicates []predicate.Predicate
	// SetupBlockContext sets the expected calls on MockBlockContext for the test execution.
	SetupBlockContext func(*contract.MockBlockContext)
	// AfterHook is called after the precompile is called.
	AfterHook func(t testing.TB, state contract.StateDB)
	// ExpectedRes is the expected raw byte result returned by the precompile
	ExpectedRes []byte
	// ExpectedErr is the expected error returned by the precompile
	ExpectedErr error
	// ChainConfig is the chain config to use for the precompile's block context
	// If nil, the default chain config will be used.
	ChainConfig precompileconfig.ChainConfig
	// Rules is the rules to use for the precompile's block context
	Rules extras.AvalancheRules
}

type PrecompileRunparams struct {
	AccessibleState contract.AccessibleState
	Caller          common.Address
	ContractAddress common.Address
	Input           []byte
	SuppliedGas     uint64
	ReadOnly        bool
}

func (test PrecompileTest) Run(t *testing.T, module modules.Module) {
	_ = "STUB: not implemented"
	return
}

func (test PrecompileTest) Bench(b *testing.B, module modules.Module) {
	_ = "STUB: not implemented"
	return
}

// Revert to the previous snapshot and take a new snapshot, so we can reset the state after execution

// Ignore return values for benchmark

// Keep it as uint64, multiply 100 to get two digit float later

// Execute the test one final time to ensure that if our RevertToSnapshot logic breaks such that each run is actually failing or resulting in unexpected behavior
// the benchmark should catch the error here.

func (test PrecompileTest) setup(t testing.TB, module modules.Module, state *testStateDB) PrecompileRunparams {
	_ = "STUB: not implemented"
	return *new(PrecompileRunparams)
}

func RunPrecompileTests(t *testing.T, module modules.Module, tests []PrecompileTest) {
	_ = "STUB: not implemented"
	return
}

func RunPrecompileBenchmarks(b *testing.B, module modules.Module, tests []PrecompileTest) {
	_ = "STUB: not implemented"
	return
}

// testStateDB allows for mocking the predicate storage slots without calling
// Prepare on the statedb.
type testStateDB struct {
	*extstate.StateDB

	predicates map[common.Address][]predicate.Predicate
}

func newTestStateDB(t testing.TB, predicates map[common.Address][]predicate.Predicate) *testStateDB {
	_ = "STUB: not implemented"
	return nil
}

func (s *testStateDB) GetPredicate(address common.Address, index int) (predicate.Predicate, bool) {
	_ = "STUB: not implemented"
	return *new(predicate.Predicate), false
}
