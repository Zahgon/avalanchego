// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlist

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"

	_ "embed"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
)

// AllowList is an abstraction that allows other precompiles to manage
// which addresses can call the precompile by maintaining an allowlist
// in the storage trie.

const (
	ModifyAllowListGasCost = contract.WriteGasCostPerSlot
	ReadAllowListGasCost   = contract.ReadGasCostPerSlot

	allowListInputLen = common.HashLength
)

var (
	// Error returned when an invalid write is attempted
	ErrCannotModifyAllowList = errors.New("cannot modify allow list")

	// AllowListRawABI contains the raw ABI of AllowList library interface.
	//go:embed IAllowList.abi
	AllowListRawABI string

	AllowListABI = contract.ParseABI(AllowListRawABI)
)

// GetAllowListStatus returns the allow list role of [address] for the precompile
// at [precompileAddr]
func GetAllowListStatus(state contract.StateReader, precompileAddr common.Address, address common.Address) Role {
	_ = "STUB: not implemented"
	// Generate the state key for [address]
	return *new(Role)
}

// SetAllowListRole sets the permissions of [address] to [role] for the precompile
// at [precompileAddr].
// assumes [role] has already been verified as valid.
func SetAllowListRole(stateDB contract.StateDB, precompileAddr, address common.Address, role Role) {
	_ = "STUB: not implemented"
	// Generate the state key for [address]
	return
}

// Assign [role] to the address
// This stores the [role] in the contract storage with address [precompileAddr]
// and [addressKey] hash. It means that any reusage of the [addressKey] for different value
// conflicts with the same slot [role] is stored.
// Precompile implementations must use a different key than [addressKey]

func PackModifyAllowList(address common.Address, role Role) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnpackModifyAllowListInput(input []byte, r Role, useStrictMode bool) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// createAllowListRoleSetter returns an execution function for setting the allow list status of the input address argument to [role].
// This execution function is specific to [precompileAddr].
func createAllowListRoleSetter(precompileAddr common.Address, role Role) contract.RunStatefulPrecompileFunc {
	_ = "STUB: not implemented"
	//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
	return *new(contract.RunStatefulPrecompileFunc)
}

// do not use strict mode after Durango

// Verify that the caller is an admin with permission to modify the allow list

// Verify that the address we are trying to modify has a status that allows it to be modified

// PackReadAllowList packs [address] into the input data to the read allow list function
func PackReadAllowList(address common.Address) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnpackReadAllowListInput(input []byte, useStrictMode bool) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func PackReadAllowListOutput(roleNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createReadAllowList returns an execution function that reads the allow list for the given [precompileAddr].
// The execution function parses the input into a single address and returns the 32 byte hash that specifies the
// designated role of that address
func createReadAllowList(precompileAddr common.Address) contract.RunStatefulPrecompileFunc {
	_ = "STUB: not implemented"
	//nolint:revive // General-purpose types lose the meaning of args if unused ones are removed
	return *new(contract.RunStatefulPrecompileFunc)
}

// We skip the fixed length check with Durango

// CreateAllowListPrecompile returns a StatefulPrecompiledContract with R/W control of an allow list at [precompileAddr]
func CreateAllowListPrecompile(precompileAddr common.Address) contract.StatefulPrecompiledContract {
	_ = "STUB: not implemented"
	// Construct the contract with no fallback function.
	return *new(contract.StatefulPrecompiledContract)
}

func CreateAllowListFunctions(precompileAddr common.Address) []*contract.StatefulPrecompileFunction {
	_ = "STUB: not implemented"
	return nil
}
