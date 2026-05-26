// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import "github.com/ava-labs/avalanchego/database"

// WriteAll writes all of the batches to the underlying database of baseBatch.
// Assumes all batches have the same underlying database.
func WriteAll(baseBatch database.Batch, batches ...database.Batch) error {
	_ = "STUB: not implemented"
	return nil
}

// Replay the inner batches onto [baseBatch] so that it includes all DB
// operations as they would be applied to the base database.

// Write all of the combined operations in one atomic batch.
