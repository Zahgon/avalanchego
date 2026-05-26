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
	_ UnsignedTx = (*ImportTx)(nil)

	errNoImportInputs = errors.New("tx has no imported inputs")
)

// ImportTx is an unsigned importTx
type ImportTx struct {
	BaseTx `serialize:"true"`

	// Which chain to consume the funds from
	SourceChain ids.ID `serialize:"true" json:"sourceChain"`

	// Inputs that consume UTXOs produced on the chain
	ImportedInputs []*avax.TransferableInput `serialize:"true" json:"importedInputs"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [ImportTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *ImportTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// InputUTXOs returns the UTXOIDs of the imported funds
func (tx *ImportTx) InputUTXOs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (tx *ImportTx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// SyntacticVerify this transaction is well-formed
func (tx *ImportTx) SyntacticVerify(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

// already passed syntactic verification

func (tx *ImportTx) Visit(visitor Visitor) error { _ = "STUB: not implemented"; return nil }
