// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customrawdb

import (
	"errors"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/params"
)

var (
	// errInvalidData indicates the stored value exists but is malformed or undecodable.
	errInvalidData                  = errors.New("invalid data")
	errFailedToGetUpgradeConfig     = errors.New("failed to get upgrade config")
	errFailedToMarshalUpgradeConfig = errors.New("failed to marshal upgrade config")

	upgradeConfigPrefix = []byte("upgrade-config-")
	// offlinePruningKey tracks runs of offline pruning.
	offlinePruningKey = []byte("OfflinePruning")
	// populateMissingTriesKey tracks runs of trie backfills.
	populateMissingTriesKey = []byte("PopulateMissingTries")
	// pruningDisabledKey tracks whether the node has ever run in archival mode
	// to ensure that a user does not accidentally corrupt an archival node.
	pruningDisabledKey = []byte("PruningDisabled")
	// acceptorTipKey tracks the tip of the last accepted block that has been fully processed.
	acceptorTipKey = []byte("AcceptorTipKey")
	// snapshotBlockHashKey tracks the block hash of the last snapshot.
	snapshotBlockHashKey = []byte("SnapshotBlockHash")
)

// WriteOfflinePruning writes a time marker of the last attempt to run offline pruning.
// The marker is written when offline pruning completes and is deleted when the node
// is started successfully with offline pruning disabled. This ensures users must
// disable offline pruning and start their node successfully between runs of offline
// pruning.
func WriteOfflinePruning(db ethdb.KeyValueWriter, ts time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadOfflinePruning reads the most recent timestamp of an attempt to run offline
// pruning if present.
func ReadOfflinePruning(db ethdb.KeyValueReader) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// DeleteOfflinePruning deletes any marker of the last attempt to run offline pruning.
func DeleteOfflinePruning(db ethdb.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

// WritePopulateMissingTries writes a marker for the current attempt to populate
// missing tries.
func WritePopulateMissingTries(db ethdb.KeyValueWriter, ts time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadPopulateMissingTries reads the most recent timestamp of an attempt to
// re-populate missing trie nodes.
func ReadPopulateMissingTries(db ethdb.KeyValueReader) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// DeletePopulateMissingTries deletes any marker of the last attempt to
// re-populate missing trie nodes.
func DeletePopulateMissingTries(db ethdb.KeyValueWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// WritePruningDisabled writes a marker to track whether the node has ever run
// with pruning disabled.
func WritePruningDisabled(db ethdb.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

// HasPruningDisabled returns true if there is a marker present indicating that
// the node has run with pruning disabled at some point.
func HasPruningDisabled(db ethdb.KeyValueReader) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// WriteAcceptorTip writes `hash` as the last accepted block that has been fully processed.
func WriteAcceptorTip(db ethdb.KeyValueWriter, hash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadAcceptorTip reads the hash of the last accepted block that was fully processed.
// If there is no value present (the index is being initialized for the first time), then the
// empty hash is returned.
func ReadAcceptorTip(db ethdb.KeyValueReader) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// ReadChainConfig retrieves the consensus settings based on the given genesis hash.
// The provided `upgradeConfig` (any JSON-unmarshalable type) will be populated if present on disk.
func ReadChainConfig[T any](db ethdb.KeyValueReader, hash common.Hash, upgradeConfig *T) (*params.ChainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteChainConfig writes the chain config settings to the database.
// The provided `upgradeConfig` (any JSON-marshalable type) will be stored alongside the chain config.
func WriteChainConfig[T any](db ethdb.KeyValueWriter, hash common.Hash, config *params.ChainConfig, upgradeConfig T) error {
	_ = "STUB: not implemented"
	return nil
}

// NewAccountSnapshotsIterator returns an iterator for walking all of the accounts in the snapshot
func NewAccountSnapshotsIterator(db ethdb.Iteratee) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

// ReadSnapshotBlockHash retrieves the hash of the block whose state is contained in
// the persisted snapshot.
func ReadSnapshotBlockHash(db ethdb.KeyValueReader) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// WriteSnapshotBlockHash stores the root of the block whose state is contained in
// the persisted snapshot.
func WriteSnapshotBlockHash(db ethdb.KeyValueWriter, blockHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteSnapshotBlockHash deletes the hash of the block whose state is contained in
// the persisted snapshot. Since snapshots are not immutable, this method can
// be used during updates, so a crash or failure will mark the entire snapshot
// invalid.
func DeleteSnapshotBlockHash(db ethdb.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }

func writeTimeMarker(db ethdb.KeyValueWriter, key []byte, ts time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func readTimeMarker(db ethdb.KeyValueReader, key []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	// Check existence first to map missing marker to a stable sentinel error.
	return *new(time.Time), nil
}

func upgradeConfigKey(hash common.Hash) []byte { _ = "STUB: not implemented"; return nil }
