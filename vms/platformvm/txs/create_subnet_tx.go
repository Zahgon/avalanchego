// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
)

var _ UnsignedTx = (*CreateSubnetTx)(nil)

// CreateSubnetTx is an unsigned proposal to create a new subnet
type CreateSubnetTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// Who is authorized to manage this subnet
	Owner fx.Owner `serialize:"true" json:"owner"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [CreateSubnetTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *CreateSubnetTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// SyntacticVerify verifies that this transaction is well-formed
func (tx *CreateSubnetTx) SyntacticVerify(ctx *snow.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already passed syntactic verification

func (tx *CreateSubnetTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
