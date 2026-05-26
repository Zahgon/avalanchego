// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package synctest

import (
	"math/rand"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/graft/evm/utils/utilstest"
)

// FillAccountsWithOverlappingStorage adds [numAccounts] randomly generated accounts to the secure trie at [root]
// and commits it to [trieDB]. For each 3 accounts created:
// - One does not have a storage trie,
// - One has a storage trie shared with other accounts (total number of shared storage tries [numOverlappingStorageRoots]),
// - One has a uniquely generated storage trie,
// returns the new trie root and a map of funded keys to StateAccount structs.
// This is only safe for HashDB, as path-based DBs do not share storage tries.
func FillAccountsWithOverlappingStorage(
	t *testing.T, r *rand.Rand, s state.Database, root common.Hash, numAccounts int, numOverlappingStorageRoots int,
) (common.Hash, map[*utilstest.Key]*types.StateAccount) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// unmodified account
// account with overlapping storage root

// account with unique storage root

// GenerateIndependentTrie creates a trie with [numKeys] random key-value pairs inside of [trieDB].
// Returns the root of the generated trie, the slice of keys inserted into the trie in lexicographical
// order, and the slice of corresponding values.
//
// This is safe for use with HashDB, intended for use creating a storage trie independent of an account,
// or for an atomic trie.
func GenerateIndependentTrie(t *testing.T, r *rand.Rand, trieDB *triedb.Database, numKeys int, keySize int) (common.Hash, [][]byte, [][]byte) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

// FillIndependentTrie fills a given trie with [numKeys] random keys, each of size [keySize]
// returns inserted keys and values
//
// This is safe for use with HashDB.
func FillIndependentTrie(t *testing.T, r *rand.Rand, start, numKeys int, keySize int, trieDB *triedb.Database, root common.Hash) (common.Hash, [][]byte, [][]byte) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

// Generate key-value pairs

// min 128 bytes, max 255 bytes

// Commit the root to [trieDB]

// AssertTrieConsistency ensures given trieDB [a] and [b] both have the same
// non-empty trie at [root]. (all key/value pairs must be equal)
//
// This is only safe for HashDB or PathDB, since Firewood doesn't store trie nodes individually.
func AssertTrieConsistency(t testing.TB, root common.Hash, a, b *triedb.Database, onLeaf func(key, val []byte) error) {
	_ = "STUB: not implemented"
	return
}

// CorruptTrie deletes every [n]th trie node from the trie given by [tr] from the underlying [db].
// Assumes [tr] can be iterated without issue.
//
// This is only safe for HashDB or PathDB, since Firewood doesn't store trie nodes individually.
func CorruptTrie(t *testing.T, diskdb ethdb.Batcher, tr *trie.Trie, n int) {
	_ = "STUB: not implemented"
	// Delete some trie nodes
	return
}

// FillAccounts adds [numAccounts] randomly generated accounts to the secure trie at [root] and commits it to [trieDB].
// [onAccount] is called if non-nil so the caller can modify the account before it is stored in the trie.
// If the trie in the callback is used (i.e. tr.Hash() doesn't return the empty root), the account's storage root will be updated to match.
// Returns the new trie root and a map of funded keys to StateAccount structs.
func FillAccounts(
	t *testing.T, r *rand.Rand, s state.Database, root common.Hash, numAccounts int,
	onAccount func(*testing.T, int, common.Address, types.StateAccount, state.Trie) types.StateAccount,
) (common.Hash, map[*utilstest.Key]*types.StateAccount) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// If the storage trie was used, update the account's storage root and pass nodes to TrieDB.

// block hashes required for Firewood

// FillAccountsWithStorageAndCode is a helper function that calls [FillAccounts] with an [onAccount] callback that randomly assigns accounts to have code and storage.
// Approximately half of accounts created will have unrunnable contracts and non-empty storage tries, and the other half will be EOAs.
func FillAccountsWithStorageAndCode(t *testing.T, r *rand.Rand, serverDB state.Database, root common.Hash, numAccounts int) (common.Hash, map[*utilstest.Key]*types.StateAccount) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// FillStorageForAccount adds [numStorageKeys] random key-value pairs to the storage trie for [addr] in [storageTr].
func FillStorageForAccount(
	t *testing.T, r *rand.Rand, numStorageKeys int,
	addr common.Address, storageTr state.Trie,
) {
	_ = "STUB: not implemented"
	return
}

func makeKeyValues(t *testing.T, r *rand.Rand, numKeys, keySize int) ([][]byte, [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate key-value pairs

// min 128 bytes, max 255 bytes
