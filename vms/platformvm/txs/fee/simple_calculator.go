// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package fee

import "github.com/ava-labs/avalanchego/vms/platformvm/txs"

var _ Calculator = (*SimpleCalculator)(nil)

type SimpleCalculator struct {
	txFee uint64
}

func NewSimpleCalculator(fee uint64) *SimpleCalculator { _ = "STUB: not implemented"; return nil }

func (c *SimpleCalculator) CalculateFee(txs.UnsignedTx) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
