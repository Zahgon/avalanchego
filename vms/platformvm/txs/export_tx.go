// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var (
	_ UnsignedTx = (*ExportTx)(nil)

	ErrWrongLocktime   = errors.New("wrong locktime reported")
	errNoExportOutputs = errors.New("no export outputs")
)

// ExportTx is an unsigned exportTx
type ExportTx struct {
	BaseTx `serialize:"true"`

	// Which chain to send the funds to
	DestinationChain ids.ID `serialize:"true" json:"destinationChain"`

	// Outputs that are exported to the chain
	ExportedOutputs []*avax.TransferableOutput `serialize:"true" json:"exportedOutputs"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [UnsignedExportTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *ExportTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// SyntacticVerify this transaction is well-formed
func (tx *ExportTx) SyntacticVerify(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

// already passed syntactic verification

func (tx *ExportTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
