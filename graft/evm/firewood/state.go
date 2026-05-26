// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package firewood

import (
	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/log"
)

func init() {
	state.RegisterDatabaseInterceptor(wrapIfFirewood)
	if err := ffi.StartMetrics(); err != nil {
		log.Crit("starting firewood metrics", "error", err)
	}
}

var _ state.Database = (*stateAccessor)(nil)

type stateAccessor struct {
	state.Database
	triedb *TrieDB
}

func wrapIfFirewood(db state.Database) state.Database {
	_ = "STUB: not implemented"
	return *new(state.Database)
}

// OpenTrie opens the main account trie.
func (s *stateAccessor) OpenTrie(root common.Hash) (state.Trie, error) {
	_ = "STUB: not implemented"
	return *new(state.Trie), nil
}

// OpenStorageTrie opens a wrapped version of the account trie.
//
//nolint:revive // removing names loses context.
func (*stateAccessor) OpenStorageTrie(stateRoot common.Hash, addr common.Address, accountRoot common.Hash, self state.Trie) (state.Trie, error) {
	_ = "STUB: not implemented"
	return *new(state.Trie), nil
}

// CopyTrie returns a deep copy of the given trie.
// It can be altered by the caller.
func (*stateAccessor) CopyTrie(t state.Trie) state.Trie {
	_ = "STUB: not implemented"
	return *new(state.Trie)
}

// The storage trie just wraps the account trie, so we must re-open it separately.
