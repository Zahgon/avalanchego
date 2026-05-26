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
	"github.com/ava-labs/avalanchego/graft/coreth/consensus"
	"github.com/ava-labs/avalanchego/graft/coreth/core/extstate"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/libevm"
	"github.com/ava-labs/libevm/libevm/stateconf"
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
	// which is used in coreth to process historical pre-AP1 blocks with the
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

// StateDBAP0 implements the GetCommittedState behavior that existed prior to
// the AP1 upgrade.
//
// Since launch, state keys have been normalized to allow for multicoin
// balances. However, at launch GetCommittedState was not updated. This meant
// that gas refunds were not calculated as expected for SSTORE opcodes.
//
// This oversight was fixed in AP1, but in order to execute blocks prior to AP1
// and generate the same merkle root, this behavior must be maintained.
//
// See the [extstate] package for details around state key normalization.
type StateDBAP0 struct {
	*extstate.StateDB
}

func (s *StateDBAP0) GetCommittedState(addr common.Address, key common.Hash, _ ...stateconf.StateDBStateOption) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
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

// If we don't have an explicit author (i.e. not mining), extract from the header

// Ignore error, we're past header validation

// NewEVMBlockContextWithPredicateResults creates a new context for use in the
// EVM with an override for the predicate results. The miner uses this to pass
// predicate results to the EVM when header.Extra is not fully formed yet.
func NewEVMBlockContextWithPredicateResults(rules extras.AvalancheRules, header *types.Header, chain ChainContext, author *common.Address, predicateBytes []byte) vm.BlockContext {
	_ = "STUB: not implemented"
	return *new(vm.BlockContext)
}

// Note this only sets the block context, which is the hand-off point for
// the EVM. The actual header is not modified.

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
