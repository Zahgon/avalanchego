// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package txgossiptest provides test helpers for mempool operations.
package txgossiptest

import (
	"context"
	"testing"

	"github.com/ava-labs/libevm/core/txpool"
	"github.com/ava-labs/libevm/core/types"
)

// WaitUntilPending waits until all transactions provided are marked as pending in `pool`.
func WaitUntilPending(tb testing.TB, ctx context.Context, pool *txpool.TxPool, txs ...*types.Transaction) {
	_ = "STUB: not implemented"
	return
}

// size arbitrary
/*reorgs but ignored by legacypool*/

// Optimistically check current mempool - any reorgs after this will
// certainly be caught by the subscription.

// already found all txs
