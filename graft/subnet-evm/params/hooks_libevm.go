// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package params

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/libevm"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

// invalidateDelegateTime is the Unix timestamp for August 2nd, 2025, midnight Eastern Time
// (August 2nd, 2025, 04:00 UTC)
const InvalidateDelegateUnix = 1754107200

// P256VerifyAddress is the address of the p256 signature verification precompile
var P256VerifyAddress = common.BytesToAddress([]byte{0x1, 0x00})

type RulesExtra extras.Rules

func GetRulesExtra(r Rules) *extras.Rules { _ = "STUB: not implemented"; return nil }

func (r RulesExtra) CanCreateContract(ac *libevm.AddressContext, gas uint64, state libevm.StateReader) (uint64, error) {
	_ = "STUB: not implemented"
	// If the allow list is enabled, check that [ac.Origin] has permission to deploy a contract.
	return 0, nil
}

func (RulesExtra) CanExecuteTransaction(common.Address, *common.Address, libevm.StateReader) error {
	_ = "STUB: not implemented"
	// TODO: Migrate call for txallowlist precompile to here from core/state_transition.go
	// when that is used from libevm.
	return nil
}

// MinimumGasConsumption is a no-op.
func (RulesExtra) MinimumGasConsumption(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// AccessListGas computes the intrinsic gas for an access list.
// When predicaters exist, it calculates gas per-tuple, delegating to predicate
// contracts for addresses that have them. Otherwise, it returns override=false
// to use the default calculation.
func (r RulesExtra) AccessListGas(accessList libevm.AccessList) (uint64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

var PrecompiledContractsGranite = map[common.Address]vm.PrecompiledContract{
	P256VerifyAddress: &vm.P256Verify{},
}

func (r RulesExtra) ActivePrecompiles(existing []common.Address) []common.Address {
	_ = "STUB: not implemented"
	return nil
}

func (r RulesExtra) currentPrecompiles() map[common.Address]vm.PrecompiledContract {
	_ = "STUB: not implemented"
	return nil
}

// precompileOverrideBuiltin specifies precompiles that were activated prior to the
// dynamic precompile activation registry.
// These were only active historically and are not active in the current network.
func (r RulesExtra) precompileOverrideBuiltin(addr common.Address) (libevm.PrecompiledContract, bool) {
	_ = "STUB: not implemented"
	return *new(libevm.PrecompiledContract), false
}

func makePrecompile(contract contract.StatefulPrecompiledContract) libevm.PrecompiledContract {
	_ = "STUB: not implemented"
	return *new(libevm.PrecompiledContract)
}

// Should never happen

// Should never happen, as results are already validated in block validation

// Others always allowed

// Otherwise, we allow the precompile to be called

func (r RulesExtra) PrecompileOverride(addr common.Address) (libevm.PrecompiledContract, bool) {
	_ = "STUB: not implemented"
	return *new(libevm.PrecompiledContract), false
}

type accessibleState struct {
	env          vm.PrecompileEnvironment
	blockContext *precompileBlockContext
}

func (a accessibleState) GetStateDB() contract.StateDB {
	_ = "STUB: not implemented"
	// TODO the contracts should be refactored to call `env.ReadOnlyState`
	// or `env.StateDB` based on the env.ReadOnly() flag
	return *new(contract.StateDB)
}

func (a accessibleState) GetBlockContext() contract.BlockContext {
	_ = "STUB: not implemented"
	return *new(contract.BlockContext)
}

func (a accessibleState) GetRules() precompileconfig.Rules {
	_ = "STUB: not implemented"
	return *new(precompileconfig.Rules)
}

func (a accessibleState) GetSnowContext() *snow.Context { _ = "STUB: not implemented"; return nil }

func (a accessibleState) GetPrecompileEnv() vm.PrecompileEnvironment {
	_ = "STUB: not implemented"
	return *new(vm.PrecompileEnvironment)
}

type precompileBlockContext struct {
	number           *big.Int
	time             uint64
	predicateResults predicate.BlockResults
}

func (p *precompileBlockContext) Number() *big.Int { _ = "STUB: not implemented"; return nil }

func (p *precompileBlockContext) Timestamp() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *precompileBlockContext) GetPredicateResults(txHash common.Hash, precompileAddress common.Address) set.Bits {
	_ = "STUB: not implemented"
	return *new(set.Bits)
}
