// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	_ UnsignedTx = (*BaseTx)(nil)

	ErrNilTx = errors.New("tx is nil")

	errOutputsNotSorted      = errors.New("outputs not sorted")
	errInputsNotSortedUnique = errors.New("inputs not sorted and unique")
)

// BaseTx contains fields common to many transaction types. It should be
// embedded in transaction implementations.
type BaseTx struct {
	avax.BaseTx `serialize:"true"`

	// true iff this transaction has already passed syntactic verification
	SyntacticallyVerified bool `json:"-"`

	unsignedBytes []byte // Unsigned byte representation of this data
}

func (tx *BaseTx) SetBytes(unsignedBytes []byte) { _ = "STUB: not implemented"; return }

func (tx *BaseTx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (tx *BaseTx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (tx *BaseTx) Outputs() []*avax.TransferableOutput {
	_ = "STUB: not implemented"

	// InitCtx sets the FxID fields in the inputs and outputs of this [BaseTx]. Also
	// sets the [ctx] to the given [vm.ctx] so that the addresses can be json
	// marshalled into human readable format
	return nil
}

func (tx *BaseTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// SyntacticVerify returns nil iff this tx is well formed
func (tx *BaseTx) SyntacticVerify(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

// already passed syntactic verification

func (tx *BaseTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
