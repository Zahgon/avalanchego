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
// Copyright 2015 The go-ethereum Authors
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

package tests

import (
	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/triedb"
)

// StateTestState groups all the state database objects together for use in tests.
type StateTestState struct {
	StateDB   *state.StateDB
	TrieDB    *triedb.Database
	Snapshots *snapshot.Tree
	TempDir   string
}

// MakePreState creates a state containing the given allocation.
func MakePreState(db ethdb.Database, accounts types.GenesisAlloc, snapshotter bool, scheme string) StateTestState {
	_ = "STUB: not implemented"
	// Set database path
	return *new(StateTestState)
}

// Commit and re-open to start with a clean state.

// If snapshot is requested, initialize the snapshotter and use it in state.

// Close should be called when the state is no longer needed, ie. after running the test.
func (st *StateTestState) Close() { _ = "STUB: not implemented"; return }

// Need to call Disable here to quit the snapshot generator goroutine.
