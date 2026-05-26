// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var _ UnsignedTx = (*DisableL1ValidatorTx)(nil)

type DisableL1ValidatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// ID corresponding to the validator
	ValidationID ids.ID `serialize:"true" json:"validationID"`
	// Authorizes this validator to be disabled
	DisableAuth verify.Verifiable `serialize:"true" json:"disableAuthorization"`
}

func (tx *DisableL1ValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

func (tx *DisableL1ValidatorTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
