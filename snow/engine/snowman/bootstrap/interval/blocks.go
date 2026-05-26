// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package interval

import "github.com/ava-labs/avalanchego/database"

// Add the block to the tree and return if the parent block should be fetched,
// but wasn't desired before.
func Add(
	db database.KeyValueWriterDeleter,
	tree *Tree,
	lastAcceptedHeight uint64,
	height uint64,
	blkBytes []byte,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We know that height is greater than lastAcceptedHeight here, so height-1
// is guaranteed not to underflow.

// Remove the block from the tree.
func Remove(
	db database.KeyValueWriterDeleter,
	tree *Tree,
	height uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}
