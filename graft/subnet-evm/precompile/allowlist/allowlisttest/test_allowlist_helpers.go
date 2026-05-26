// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlisttest

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"

	sim "github.com/ava-labs/avalanchego/graft/subnet-evm/ethclient/simulated"
)

// All precompiles that use the AllowList pattern extend IAllowList in Solidity, so their Go bindings
// will automatically have these 5 methods and implement the AllowListContract interface.
// Defining this interface allows the helper functions to work with any precompile that uses the AllowList pattern.
type AllowListContract interface {
	ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error)
	SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)
	SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)
	SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)
	SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)
}

// VerifyRole checks that the given address has the expected role in the allow list.
func VerifyRole(t *testing.T, contract AllowListContract, address common.Address, expectedRole allowlist.Role) {
	_ = "STUB: not implemented"
	return
}

// SetAsEnabled sets the given address as Enabled in the allow list.
func SetAsEnabled(t *testing.T, b *sim.Backend, contract AllowListContract, auth *bind.TransactOpts, address common.Address) {
	_ = "STUB: not implemented"
	return
}

// SetAsAdmin sets the given address as Admin in the allow list.
func SetAsAdmin(t *testing.T, b *sim.Backend, contract AllowListContract, auth *bind.TransactOpts, address common.Address) {
	_ = "STUB: not implemented"
	return
}

// SetAsManager sets the given address as Manager in the allow list.
func SetAsManager(t *testing.T, b *sim.Backend, contract AllowListContract, auth *bind.TransactOpts, address common.Address) {
	_ = "STUB: not implemented"
	return
}

// SetAsNone revokes the role of the given address in the allow list.
func SetAsNone(t *testing.T, b *sim.Backend, contract AllowListContract, auth *bind.TransactOpts, address common.Address) {
	_ = "STUB: not implemented"
	return
}
