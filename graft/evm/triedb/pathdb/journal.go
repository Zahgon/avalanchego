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
// Copyright 2022 The go-ethereum Authors
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
	"errors"
	"io"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/rlp"
)

var (
	errMissJournal       = errors.New("journal not found")
	errMissVersion       = errors.New("version not found")
	errUnexpectedVersion = errors.New("unexpected journal version")
	errMissDiskRoot      = errors.New("disk layer root not found")
	errUnmatchedJournal  = errors.New("unmatched journal")
)

const journalVersion uint64 = 0

// journalNode represents a trie node persisted in the journal.
type journalNode struct {
	Path []byte // Path of the node in the trie
	Blob []byte // RLP-encoded trie node blob, nil means the node is deleted
}

// journalNodes represents a list trie nodes belong to a single account
// or the main account trie.
type journalNodes struct {
	Owner common.Hash
	Nodes []journalNode
}

// journalAccounts represents a list accounts belong to the layer.
type journalAccounts struct {
	Addresses []common.Address
	Accounts  [][]byte
}

// journalStorage represents a list of storage slots belong to an account.
type journalStorage struct {
	Incomplete bool
	Account    common.Address
	Hashes     []common.Hash
	Slots      [][]byte
}

// loadJournal tries to parse the layer journal from the disk.
func (db *Database) loadJournal(diskRoot common.Hash) (layer, error) {
	_ = "STUB: not implemented"
	return *new(layer), nil
}

// Firstly, resolve the first element as the journal version

// Secondly, resolve the disk layer root, ensure it's continuous
// with disk layer. Note now we can ensure it's the layer journal
// correct version, so we expect everything can be resolved properly.

// The journal is not matched with persistent state, discard them.
// It can happen that geth crashes without persisting the journal.

// Load the disk layer from the journal

// Load all the diff layers from the journal

// loadLayers loads a pre-existing state layer backed by a key-value store.
func (db *Database) loadLayers() layer {
	_ = "STUB: not implemented"
	// Retrieve the root node of persistent state.
	return *new(layer)
}

// Load the layers by resolving the journal

// journal is not matched(or missing) with the persistent state, discard
// it. Display log for discarding journal, but try to avoid showing
// useless information when the db is created from scratch.

// Return single layer with persistent state.

// loadDiskLayer reads the binary blob from the layer journal, reconstructing
// a new disk layer on it.
func (db *Database) loadDiskLayer(r *rlp.Stream) (layer, error) {
	_ = "STUB: not implemented"
	// Resolve disk layer root
	return *new(layer), nil
}

// Resolve the state id of disk layer, it can be different
// with the persistent id tracked in disk, the id distance
// is the number of transitions aggregated in disk layer.

// Resolve nodes cached in node buffer

// Calculate the internal state transitions by id difference.

// loadDiffLayer reads the next sections of a layer journal, reconstructing a new
// diff and verifying that it can be linked to the requested parent.
func (db *Database) loadDiffLayer(parent layer, r *rlp.Stream) (layer, error) {
	_ = "STUB: not implemented"
	// Read the next diff journal entry
	return *new(layer), nil
}

// The first read may fail with EOF, marking the end of the journal

// Read in-memory trie nodes from journal

// Read state changes from journal

// journal implements the layer interface, marshaling the un-flushed trie nodes
// along with layer meta data into provided byte buffer.
func (dl *diskLayer) journal(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Ensure the layer didn't get stale

// Step one, write the disk root into the journal.

// Step two, write the corresponding state id into the journal

// Step three, write all unwritten nodes into the journal

// journal implements the layer interface, writing the memory layer contents
// into a buffer to be stored in the database as the layer journal.
func (dl *diffLayer) journal(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// journal the parent first

// Everything below was journaled, persist this layer too

// Write the accumulated trie nodes into buffer

// Write the accumulated state changes into buffer

// Journal commits an entire diff hierarchy to disk into a single journal entry.
// This is meant to be used during shutdown to persist the layer without
// flattening everything down (bad for reorgs). And this function will mark the
// database as read-only to prevent all following mutation to disk.
func (db *Database) Journal(root common.Hash) error {
	_ = "STUB: not implemented"
	// Retrieve the head layer to journal from.
	return nil
}

// disk layer only on noop runs (likely) or deep reorgs (unlikely)

// Run the journaling

// Short circuit if the database is in read only mode.

// Firstly write out the metadata of journal

// The stored state in disk might be empty, convert the
// root to emptyRoot in this case.

// Secondly write out the state root in disk, ensure all layers
// on top are continuous with disk.

// Finally write out the journal of each layer in reverse order.

// Store the journal into the database and return

// Set the db in read only mode to reject all following mutations
