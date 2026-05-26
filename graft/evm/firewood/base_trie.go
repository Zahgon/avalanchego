// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package firewood

import (
	"errors"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb/database"
)

var (
	errNodeIteratorNotImplemented = errors.New("NodeIterator not implemented for Firewood")
	errProveNotImplemented        = errors.New("Prove not implemented for Firewood")
)

// baseTrie contains the shared state and methods for all Firewood
// trie implementations. It provides the read/write operations that are
// identical between [accountTrie] and [reconstructedAccountTrie].
//
// Not concurrent-safe.
type baseTrie struct {
	reader     database.Reader
	root       common.Hash
	dirtyKeys  map[string][]byte
	updateOps  []ffi.BatchOp
	hasChanges bool
}

// GetAccount returns the state account associated with an address.
// - If the account has been updated, the new value is returned.
// - If the account has been deleted, (nil, nil) is returned.
// - If the account does not exist, (nil, nil) is returned.
func (a *baseTrie) GetAccount(addr common.Address) (*types.StateAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First check if there's a pending update for this account

// If the value is empty, it indicates deletion
// Invariant: All encoded values have length > 0

// Decode and return the updated account

// No pending update found, read from the underlying reader

// Decode the account node

// GetStorage returns the value associated with a storage key for a given account address.
// - If the storage slot has been updated, the new value is returned.
// - If the storage slot has been deleted, (nil, nil) is returned.
// - If the storage slot does not exist, (nil, nil) is returned.
func (a *baseTrie) GetStorage(addr common.Address, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// If the account has been deleted, we should return nil
	return nil, nil
}

// Check if there's a pending update for this storage slot

// If the value is empty, it indicates deletion

// Decode and return the updated storage value

// No pending update found, read from the underlying reader

// Decode the storage value

// UpdateAccount replaces or creates the state account associated with an address.
// This new value will be returned for subsequent `GetAccount` calls.
func (a *baseTrie) UpdateAccount(addr common.Address, account *types.StateAccount) error {
	_ = "STUB: not implemented"
	// Queue the keys and values for later commit
	return nil
}

// Mark that there are changes to commit

// UpdateStorage replaces or creates the value associated with a storage key for a given account address.
// This new value will be returned for subsequent `GetStorage` calls.
func (a *baseTrie) UpdateStorage(addr common.Address, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Queue the keys and values for later commit

// Mark that there are changes to commit

// DeleteAccount removes the state account associated with an address.
func (a *baseTrie) DeleteAccount(addr common.Address) error { _ = "STUB: not implemented"; return nil }

// Queue the key for deletion

// Remove all storage
// Mark that there are changes to commit

// DeleteStorage removes the value associated with a storage key for a given account address.
func (a *baseTrie) DeleteStorage(addr common.Address, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Queue the key for deletion

// Mark that there are changes to commit

// UpdateContractCode implements state.Trie.
// Contract code is controlled by `rawdb`, so we don't need to do anything here.
// This always returns nil.
func (*baseTrie) UpdateContractCode(common.Address, common.Hash, []byte) error {
	_ = "STUB: not implemented"

	// GetKey implements state.Trie.
	// Preimages are not yet supported in Firewood.
	// It always returns nil.
	return nil
}

func (*baseTrie) GetKey([]byte) []byte {
	_ = "STUB: not implemented"

	// NodeIterator implements state.Trie.
	// Firewood does not support iterating over internal nodes.
	// This always returns an error.
	return nil
}

func (*baseTrie) NodeIterator([]byte) (trie.NodeIterator, error) {
	_ = "STUB: not implemented"
	return *new(trie.NodeIterator), nil
}

// Prove implements state.Trie.
// Firewood does not support providing key proofs.
// This always returns an error.
func (*baseTrie) Prove([]byte, ethdb.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

// copy creates a deep copy of the baseTrie fields.
// The [database.Reader] is shared, since it is read-only.
func (a *baseTrie) copy() *baseTrie { _ = "STUB: not implemented"; return nil }

// each ffi.BatchOp is read-only, safe to shallow copy
