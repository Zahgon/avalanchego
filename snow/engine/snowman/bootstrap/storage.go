// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/bootstrap/interval"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

const (
	batchWritePeriod      = 64
	iteratorReleasePeriod = 1024
	logPeriod             = 5 * time.Second
	minBlocksToCompact    = 5000
)

// getMissingBlockIDs returns the ID of the blocks that should be fetched to
// attempt to make a single continuous range from
// (lastAcceptedHeight, highestTrackedHeight].
//
// For example, if the tree currently contains heights [1, 4, 6, 7] and the
// lastAcceptedHeight is 2, this function will return the IDs corresponding to
// blocks [3, 5].
func getMissingBlockIDs(
	ctx context.Context,
	db database.KeyValueReader,
	nonVerifyingParser block.Parser,
	tree *interval.Tree,
	lastAcceptedHeight uint64,
) (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// process a series of consecutive blocks starting at [blk].
//
//   - blk is a block that is assumed to have been marked as acceptable by the
//     bootstrapping engine.
//   - ancestors is a set of blocks that can be used to lookup blocks.
//
// If [blk]'s height is <= the last accepted height, then it will be removed
// from the missingIDs set.
//
// Returns a newly discovered blockID that should be fetched.
func process(
	db database.KeyValueWriterDeleter,
	tree *interval.Tree,
	missingBlockIDs set.Set[ids.ID],
	lastAcceptedHeight uint64,
	blk snowman.Block,
	ancestors map[ids.ID]snowman.Block,
) (ids.ID, bool, error) {
	_ = "STUB: not implemented"

	// It's possible that missingBlockIDs contain values contained inside of
	// ancestors. So, it's important to remove IDs from the set for each
	// iteration, not just the first block's ID.
	return *new(ids.ID), false, nil
}

// If the parent was provided in the ancestors set, we can immediately
// process it.

// execute all the blocks tracked by the tree. If a block is in the tree but is
// already accepted based on the lastAcceptedHeight, it will be removed from the
// tree but not executed.
//
// execute assumes that getMissingBlockIDs would return an empty set.
//
// TODO: Replace usage of haltable with context cancellation.
func execute(
	ctx context.Context,
	shouldHalt func() bool,
	log logging.Func,
	db database.Database,
	nonVerifyingParser block.Parser,
	tree *interval.Tree,
	lastAcceptedHeight uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Not a fatal error, log and move on.

// Not a fatal error, log and move on.

// Add the first sample to the EtaTracker to establish an accurate baseline

// Periodically write the batch to disk to avoid memory pressure.

// Periodically release and re-grab the database iterator to avoid
// keeping a reference to an old database revision.

// The batch must be written here to avoid re-processing a block.

// We specify the starting key of the iterator so that the
// underlying database doesn't need to scan over the, potentially
// not yet compacted, blocks we just deleted.

// Use the tracked previous progress for accurate ETA calculation

// Only log if we have a valid ETA estimate
