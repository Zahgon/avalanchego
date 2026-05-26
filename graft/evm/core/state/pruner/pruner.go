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
// Copyright 2021 The go-ethereum Authors
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

package pruner

import (
	"time"

	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
)

const (
	// stateBloomFilePrefix is the filename prefix of state bloom filter.
	stateBloomFilePrefix = "statebloom"

	// stateBloomFileSuffix is the filename suffix of state bloom filter.
	stateBloomFileSuffix = "bf.gz"

	// stateBloomFileTempSuffix is the filename suffix of state bloom filter
	// while it is being written out to detect write aborts.
	stateBloomFileTempSuffix = ".tmp"

	// rangeCompactionThreshold is the minimal deleted entry number for
	// triggering range compaction. It's a quite arbitrary number but just
	// to avoid triggering range compaction because of small deletion.
	rangeCompactionThreshold = 100000
)

// Config includes all the configurations for pruning.
type Config struct {
	Datadir   string // The directory of the state database
	BloomSize uint64 // The Megabytes of memory allocated to bloom-filter
}

// Pruner is an offline tool to prune the stale state with the
// help of the snapshot. The workflow of pruner is very simple:
//
//   - iterate the snapshot, reconstruct the relevant state
//   - iterate the database, delete all other state entries which
//     don't belong to the target state and the genesis state
//
// It can take several hours(around 2 hours for mainnet) to finish
// the whole pruning work. It's recommended to run this offline tool
// periodically in order to release the disk usage and improve the
// disk read performance to some extent.
type Pruner struct {
	config      Config
	chainHeader *types.Header
	db          ethdb.Database
	stateBloom  *stateBloom
	snaptree    *snapshot.Tree
}

// NewPruner creates the pruner instance.
func NewPruner(db ethdb.Database, config Config) (*Pruner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Offline pruning is only supported in legacy hash based scheme.

// Note: we refuse to start a pruning session unless the snapshot disk layer exists, which should prevent
// us from ever needing to enter RecoverPruning in an invalid pruning session (a session where we do not have
// the protected trie in the triedb and in the snapshot disk layer).

// The relevant snapshot(s) might not exist

// Sanitize the bloom filter size if it's too small.

func prune(maindb ethdb.Database, stateBloom *stateBloom, bloomPath string, start time.Time) error {
	_ = "STUB: not implemented"
	// Delete all stale trie nodes in the disk. With the help of state bloom
	// the trie nodes(and codes) belong to the active state will be filtered
	// out. A very small part of stale tries will also be filtered because of
	// the false-positive rate of bloom filter. But the assumption is held here
	// that the false-positive is low enough(~0.05%). The probability of the
	// dangling node is the state root is super low. So the dangling nodes in
	// theory will never ever be visited again.
	return nil
}

// We wrap iter.Release() in an anonymous function so that the [iter]
// value captured is the value of [iter] at the end of the function as opposed
// to incorrectly capturing the first iterator immediately.

// All state entries don't belong to specific state and genesis are deleted here
// - trie node
// - legacy contract code
// - new-scheme contract code

// Realistically will never remain uninited

// +1s to avoid division by zero

// Recreate the iterator after every batch commit in order
// to allow the underlying compactor to delete the entries.

// Write marker to DB to indicate offline pruning finished successfully. We write before calling os.RemoveAll
// to guarantee that if the node dies midway through pruning, then this will run during RecoverPruning.

// Delete the state bloom, it marks the entire pruning procedure is
// finished. If any crashes or manual exit happens before this,
// `RecoverPruning` will pick it up in the next restarts to redo all
// the things.

// Start compactions, will remove the deleted data from the disk immediately.
// Note for small pruning, the compaction is skipped.

// Prune deletes all historical state nodes except the nodes belong to the
// specified state version. If user doesn't specify the state version, use
// the bottom-most snapshot diff layer as the target.
func (p *Pruner) Prune(root common.Hash) error {
	_ = "STUB: not implemented"
	// If the state bloom filter is already committed previously,
	// reuse it for pruning instead of generating a new one. It's
	// mandatory because a part of state may already be deleted,
	// the recovery procedure is necessary.
	return nil
}

// If the target state root is not specified, return a fatal error.

// Ensure the root is really present. The weak assumption
// is the presence of root can indicate the presence of the
// entire trie.

// Traverse the target state, re-construct the whole state trie and
// commit to the given bloom filter.

// Traverse the genesis, put all genesis state entries into the
// bloom filter too.

// RecoverPruning will resume the pruning procedure during the system restart.
// This function is used in this case: user tries to prune state data, but the
// system was interrupted midway because of crash or manual-kill. In this case
// if the bloom filter for filtering active state is already constructed, the
// pruning can be resumed. What's more if the bloom filter is constructed, the
// pruning **has to be resumed**. Otherwise a lot of dangling nodes may be left
// in the disk.
func RecoverPruning(datadir string, db ethdb.Database) error { _ = "STUB: not implemented"; return nil }

// nothing to recover

// All the state roots of the middle layers should be forcibly pruned,
// otherwise the dangling state will be left.

// extractGenesis loads the genesis state and commits all the state entries
// into the given bloomfilter.
func extractGenesis(db ethdb.Database, stateBloom *stateBloom) error {
	_ = "STUB: not implemented"
	return nil
}

// Embedded nodes don't have hash.

// If it's a leaf node, yes we are touching an account,
// dig into the storage trie further.

func bloomFilterName(datadir string, hash common.Hash) string { _ = "STUB: not implemented"; return "" }

func isBloomFilter(filename string) (bool, common.Hash) {
	_ = "STUB: not implemented"
	return false, *new(common.Hash)
}

func findBloomFilter(datadir string) (string, common.Hash, error) {
	_ = "STUB: not implemented"
	return "", *new(common.Hash), nil
}
