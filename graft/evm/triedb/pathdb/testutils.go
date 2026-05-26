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
// Copyright 2023 The go-ethereum Authors
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

package pathdb

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/trie/triestate"
)

// testHasher is a test utility for computing root hash of a batch of state
// elements. The hash algorithm is to sort all the elements in lexicographical
// order, concat the key and value in turn, and perform hash calculation on
// the concatenated bytes. Except the root hash, a nodeset will be returned
// once Commit is called, which contains all the changes made to hasher.
type testHasher struct {
	owner   common.Hash            // owner identifier
	root    common.Hash            // original root
	dirties map[common.Hash][]byte // dirty states
	cleans  map[common.Hash][]byte // clean states
}

// newTestHasher constructs a hasher object with provided states.
func newTestHasher(owner common.Hash, root common.Hash, cleans map[common.Hash][]byte) (*testHasher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get returns the value for key stored in the trie.
func (h *testHasher) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Update associates key with value in the trie.
func (h *testHasher) Update(key, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete removes any existing value for key from the trie.
func (h *testHasher) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// Commit computes the new hash of the states and returns the set with all
// state changes.
func (h *testHasher) Commit(collectLeaf bool) (common.Hash, *trienode.NodeSet, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

// Include the dirty root node as well.

// hash performs the hash computation upon the provided states.
func hash(states map[common.Hash][]byte) (common.Hash, []byte) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

type hashLoader struct {
	accounts map[common.Hash][]byte
	storages map[common.Hash]map[common.Hash][]byte
}

func newHashLoader(accounts map[common.Hash][]byte, storages map[common.Hash]map[common.Hash][]byte) *hashLoader {
	_ = "STUB: not implemented"
	return nil
}

// OpenTrie opens the main account trie.
func (l *hashLoader) OpenTrie(root common.Hash) (triestate.Trie, error) {
	_ = "STUB: not implemented"
	return *new(triestate.Trie), nil
}

// OpenStorageTrie opens the storage trie of an account.
func (l *hashLoader) OpenStorageTrie(stateRoot common.Hash, addrHash, root common.Hash) (triestate.Trie, error) {
	_ = "STUB: not implemented"
	return *new(triestate.Trie), nil
}
