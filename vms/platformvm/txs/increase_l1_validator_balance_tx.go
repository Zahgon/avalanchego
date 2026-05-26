// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var (
	_ UnsignedTx = (*IncreaseL1ValidatorBalanceTx)(nil)

	ErrZeroBalance = errors.New("balance must be greater than 0")
)

type IncreaseL1ValidatorBalanceTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// ID corresponding to the validator
	ValidationID ids.ID `serialize:"true" json:"validationID"`
	// Balance <= sum($AVAX inputs) - sum($AVAX outputs) - TxFee
	Balance uint64 `serialize:"true" json:"balance"`
}

func (tx *IncreaseL1ValidatorBalanceTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

func (tx *IncreaseL1ValidatorBalanceTx) Visit(visitor Visitor) error {
	_ = "STUB: not implemented"
	return nil
}
