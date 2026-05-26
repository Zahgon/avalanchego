// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package firewood

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/trie/trienode"
)

var _ state.Trie = (*accountTrie)(nil)

// accountTrie implements [state.Trie] for managing account states.
// Although it fulfills the [state.Trie] interface, it has some important differences:
//  1. [accountTrie.Commit] is not used as expected in the state package. The `StorageTrie` doesn't return
//     values, and we thus rely on the `accountTrie`. Additionally, no [trienode.NodeSet] is
//     actually constructed, since Firewood manages nodes internally and the list of changes
//     is not needed externally.
//  2. The [accountTrie.Hash] method actually creates the [ffi.Proposal], since Firewood cannot calculate
//     the hash of the trie without committing it.
//
// Note this is not concurrent safe.
type accountTrie struct {
	*baseTrie
	fw         *TrieDB
	parentRoot common.Hash
}

func newAccountTrie(root common.Hash, db *TrieDB) (*accountTrie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start with hasChanges true to allow computing the proposal hash

// Hash returns the current hash of the state trie.
// This will create the necessary proposals to guarantee that the changes can
// later be committed. All new proposals will be tracked by the [TrieDB].
// If there are no changes since the last call, the cached root is returned.
// On error, the zero hash is returned.
func (a *accountTrie) Hash() common.Hash { _ = "STUB: not implemented"; return *new(common.Hash) }

func (a *accountTrie) hash() (common.Hash, error) {
	_ = "STUB: not implemented"
	// If we haven't already hashed, we need to do so.
	return *new(common.Hash), nil
}

// Avoid re-hashing until next update

// Commit returns the new root hash of the trie and an empty [trienode.NodeSet].
// The boolean input is ignored, as it is a relic of the StateTrie implementation.
// If the changes are not yet already tracked by the [TrieDB], they are created.
func (a *accountTrie) Commit(bool) (common.Hash, *trienode.NodeSet, error) {
	_ = "STUB: not implemented"
	// Get the hash of the trie.
	// Ensures all changes are tracked by the Database.
	return *new(common.Hash), nil, nil
}

// Copy creates a deep copy of the [accountTrie].
// The [database.Reader] is shared, since it is read-only.
func (a *accountTrie) Copy() *accountTrie { _ = "STUB: not implemented"; return nil }
