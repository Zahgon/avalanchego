// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txgossip

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/txpool"
	"github.com/ava-labs/libevm/core/types"
	"github.com/holiman/uint256"
)

// A LazyTransaction couples a [txpool.LazyTransaction] with its sender.
type LazyTransaction struct {
	*txpool.LazyTransaction
	Sender common.Address
}

// TransactionsByPriority calls [txpool.TxPool.Pending] with the given filter,
// collapses the results into a slice, and sorts said slice by decreasing gas
// tip then chronologically. Transactions from the same sender are merely sorted
// by increasing nonce.
func (s *Set) TransactionsByPriority(filter txpool.PendingFilter) []*LazyTransaction {
	_ = "STUB: not implemented"
	// TODO(arr4n) investigate optimisations; e.g. skipping entire accounts once
	// the block builder has found that a lower-nonced tx is invalid.
	return nil
}

// [txpool.TxPool.Pending] already returns each slice in nonce order
// and we're performing a stable sort. A direct comparison of nonces
// would require resolving the lazy transaction.

// Higher tips first

// effectiveGasTip is equivalent to [types.Transaction.EffectiveGasTip] but
// assumes that `baseFee` is either nil or <= the transaction's fee cap. This
// assumption avoids the need for [types.ErrGasFeeCapTooLow]. If this invariant
// is broken, effectiveGasTip returns zero.
func (ltx *LazyTransaction) effectiveGasTip(baseFee *uint256.Int) *uint256.Int {
	_ = "STUB: not implemented"
	return nil
}

// Resolve shadows the equivalent method on the [txpool.LazyTransaction],
// extending its return signature to include a boolean that is true i.f.f. the
// transaction is non-nil. This avoids the foot-gun of not knowing that a nil
// check needs to be performed.
func (ltx *LazyTransaction) Resolve() (*types.Transaction, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
