// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extstate

import (
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/libevm"

	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
)

// RegisterExtras registers hooks with libevm to achieve Avalanche state
// management. It MUST NOT be called more than once and therefore is only
// allowed to be used in tests and `package main`, to avoid polluting other
// packages that transitively depend on this one but don't need registration.
//
// Of note, a call to RegisterExtras will result in state-key normalization
// unless [stateconf.SkipStateKeyTransformation] is used.
func RegisterExtras() { _ = "STUB: not implemented"; return }

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

type normalizeStateKeysHook struct{}

// TransformStateKey transforms all keys with [normalizeStateKey].
func (normalizeStateKeysHook) TransformStateKey(_ common.Address, key common.Hash) common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

type StateDB struct {
	*state.StateDB

	// Ordered storage slots to be used in predicate verification as set in the tx access list.
	// Only set in [StateDB.Prepare], and un-modified through execution.
	predicates map[common.Address][]predicate.Predicate
}

// New creates a new [StateDB] with the given [state.StateDB], wrapping it with
// additional functionality.
func New(vm *state.StateDB) *StateDB { _ = "STUB: not implemented"; return nil }

func (s *StateDB) Prepare(rules params.Rules, sender, coinbase common.Address, dst *common.Address, precompiles []common.Address, list types.AccessList) {
	_ = "STUB: not implemented"
	return
}

// GetPredicate returns the storage slots associated with the address, index pair.
// A list of access tuples can be included within transaction types post EIP-2930. The address
// is declared directly on the access tuple and the index is the i'th occurrence of an access
// tuple with the specified address.
//
// Ex. AccessList[[AddrA, Predicate1], [AddrB, Predicate2], [AddrA, Predicate3]]
// In this case, the caller could retrieve predicates 1-3 with the following calls:
// GetPredicate(AddrA, 0) -> Predicate1
// GetPredicate(AddrB, 0) -> Predicate2
// GetPredicate(AddrA, 1) -> Predicate3
func (s *StateDB) GetPredicate(address common.Address, index int) (predicate.Predicate, bool) {
	_ = "STUB: not implemented"
	return *new(predicate.Predicate), false
}

// Retrieve the balance from the given address or 0 if object not found
func (s *StateDB) GetBalanceMultiCoin(addr common.Address, coinID common.Hash) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// AddBalance adds amount to the account associated with addr.
func (s *StateDB) AddBalanceMultiCoin(addr common.Address, coinID common.Hash, amount *big.Int) {
	_ = "STUB: not implemented"
	return
}

// used to cause touch

// SubBalance subtracts amount from the account associated with addr.
func (s *StateDB) SubBalanceMultiCoin(addr common.Address, coinID common.Hash, amount *big.Int) {
	_ = "STUB: not implemented"
	return
}

// Note: It's not needed to set the IsMultiCoin (extras) flag here, as this
// call would always be preceded by a call to AddBalanceMultiCoin, which would
// set the extra flag. Seems we should remove the redundant code.

// normalizeStateKey sets the 0th bit of the first byte in `key` to 0.
// This partitions normal state storage from multicoin storage.
func normalizeStateKey(key *common.Hash) {
	_ = "STUB: not implemented"

	// normalizeCoinID sets the 0th bit of the first byte in `coinID` to 1.
	// This partitions multicoin storage from normal state storage.
	return
}

func normalizeCoinID(coinID *common.Hash) { _ = "STUB: not implemented"; return }
