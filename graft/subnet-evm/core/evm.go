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
// Copyright 2016 The go-ethereum Authors
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

package core

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/consensus"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/libevm"

	"github.com/holiman/uint256"
)

// RegisterExtras registers hooks with libevm to achieve Avalanche behaviour of
// the EVM. It MUST NOT be called more than once and therefore is only allowed
// to be used in tests and `package main`, to avoid polluting other packages
// that transitively depend on this one but don't need registration.
func RegisterExtras() {
	_ = "STUB: not implemented"
	// Although the registration function refers to just Hooks (not Extras) this
	// will be changed in the future to standardise across libevm, hence the
	// name of the function we're in.
	return
}

// WithTempRegisteredExtras runs `fn` with temporary registration otherwise
// equivalent to a call to [RegisterExtras], but limited to the life of `fn`.
//
// This function is not intended for direct use. Use
// `evm.WithTempRegisteredLibEVMExtras()` instead as it calls this along with
// all other temporary-registration functions.
func WithTempRegisteredExtras(lock libevm.ExtrasLock, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

type hooks struct{}

// PreprocessingGasCharge is not necessary.
// It is required to implement the vm.Hooks interface, but is only needed post-SAE.
func (hooks) PreprocessingGasCharge(common.Hash) (uint64, error) {
	_ = "STUB: not implemented"

	// OverrideNewEVMArgs is a hook that is called in [vm.NewEVM].
	// It allows for the modification of the EVM arguments before the EVM is created.
	// Specifically, we set Random to be the same as Difficulty since Shanghai.
	// This allows using the same jump table as upstream.
	// Then we set Difficulty to 0 as it is post Merge in upstream.
	// Additionally we wrap the StateDB with the appropriate StateDB wrapper,
	// which is used in subnet-evm to process historical pre-AP1 blocks with the
	// [StateDbAP1.GetCommittedState] method as it was historically.
	return 0, nil
}

func (hooks) OverrideNewEVMArgs(args *vm.NewEVMArgs) *vm.NewEVMArgs {
	_ = "STUB: not implemented"
	return nil
}

func (hooks) OverrideEVMResetArgs(rules params.Rules, args *vm.EVMResetArgs) *vm.EVMResetArgs {
	_ = "STUB: not implemented"
	return nil
}

func wrapStateDB(rules params.Rules, statedb vm.StateDB) vm.StateDB {
	_ = "STUB: not implemented"
	return *new(vm.StateDB)
}

// ChainContext supports retrieving headers and consensus parameters from the
// current blockchain to be used during transaction processing.
type ChainContext interface {
	// Engine retrieves the chain's consensus engine.
	Engine() consensus.Engine

	// GetHeader returns the header corresponding to the hash/number argument pair.
	GetHeader(common.Hash, uint64) *types.Header
}

// NewEVMBlockContext creates a new context for use in the EVM.
func NewEVMBlockContext(header *types.Header, chain ChainContext, author *common.Address) vm.BlockContext {
	_ = "STUB: not implemented"
	return *new(vm.BlockContext)
}

// Prior to Durango, the VM enforces the extra data is smaller than or
// equal to this size. After Durango, the VM pre-verifies the extra
// data past the dynamic fee rollup window is valid.

// As mentioned above, we pre-verify the extra data to ensure this never happens.
// If we hit an error, construct a new block context rather than use a potentially half initialized value
// as defense in depth.

// NewEVMBlockContextWithPredicateResults creates a new context for use in the EVM with an override for the predicate results that is not present
// in header.Extra.
// This function is used to create a BlockContext when the header Extra data is not fully formed yet and it's more efficient to pass in predicateResults
// directly rather than re-encode the latest results when executing each individual transaction.
func NewEVMBlockContextWithPredicateResults(header *types.Header, chain ChainContext, author *common.Address, predicateBytes []byte) vm.BlockContext {
	_ = "STUB: not implemented"
	return *new(vm.BlockContext)
}

// Note this only sets the block context, which is the hand-off point for
// the EVM. The actual header is not modified.

func newEVMBlockContext(header *types.Header, chain ChainContext, author *common.Address, extra []byte) vm.BlockContext {
	_ = "STUB: not implemented"
	return *new(vm.BlockContext)
}

// If we don't have an explicit author (i.e. not mining), extract from the header

// Ignore error, we're past header validation

// NewEVMTxContext creates a new transaction context for a single transaction.
func NewEVMTxContext(msg *Message) vm.TxContext {
	_ = "STUB: not implemented"
	return *new(vm.TxContext)
}

// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn(ref *types.Header, chain ChainContext) func(n uint64) common.Hash {
	_ = "STUB: not implemented"
	// Cache will initially contain [refHash.parent],
	// Then fill up with [refHash.p, refHash.pp, refHash.ppp, ...]
	return nil
}

// This situation can happen if we're doing tracing and using
// block overrides.

// If there's no hash cache yet, make one

// No luck in the cache, but we can start iterating from the last element we already know

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, amount *uint256.Int) bool {
	_ = "STUB: not implemented"
	return false
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
	_ = "STUB: not implemented"
	return
}
