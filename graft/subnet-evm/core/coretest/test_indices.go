// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package coretest

import (
	"testing"

	"github.com/ava-labs/libevm/ethdb"
)

// checkTxIndicesHelper checks that the transaction indices are correctly stored in the database.
// [expectedTail] is the expected value of the tail index.
// [indexedFrom] is the block number from which the transactions should be indexed.
// [indexedTo] is the block number to which the transactions should be indexed.
// [head] is the block number of the head block.
func CheckTxIndices(t *testing.T, expectedTail *uint64, indexedFrom uint64, indexedTo uint64, head uint64, db ethdb.Database, allowNilBlocks bool) {
	_ = "STUB: not implemented"
	return
}
