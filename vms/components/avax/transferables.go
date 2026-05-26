// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	ErrNilTransferableOutput   = errors.New("nil transferable output is not valid")
	ErrNilTransferableFxOutput = errors.New("nil transferable feature extension output is not valid")
	ErrOutputsNotSorted        = errors.New("outputs not sorted")

	ErrNilTransferableInput   = errors.New("nil transferable input is not valid")
	ErrNilTransferableFxInput = errors.New("nil transferable feature extension input is not valid")
	ErrInputsNotSortedUnique  = errors.New("inputs not sorted and unique")

	_ verify.Verifiable                  = (*TransferableOutput)(nil)
	_ verify.Verifiable                  = (*TransferableInput)(nil)
	_ utils.Sortable[*TransferableInput] = (*TransferableInput)(nil)
)

// Amounter is a data structure that has an amount of something associated with it
type Amounter interface {
	snow.ContextInitializable
	// Amount returns how much value this element represents of the asset in its
	// transaction.
	Amount() uint64
}

// Coster is a data structure that has a cost associated with it
type Coster interface {
	// Cost returns how much this element costs to be included in its
	// transaction.
	Cost() (uint64, error)
}

// TransferableIn is the interface a feature extension must provide to transfer
// value between features extensions.
type TransferableIn interface {
	verify.Verifiable
	Amounter
	Coster
}

// TransferableOut is the interface a feature extension must provide to transfer
// value between features extensions.
type TransferableOut interface {
	snow.ContextInitializable
	verify.State
	Amounter
}

type TransferableOutput struct {
	Asset `serialize:"true"`
	// FxID has serialize false because we don't want this to be encoded in bytes
	FxID ids.ID          `serialize:"false" json:"fxID"`
	Out  TransferableOut `serialize:"true"  json:"output"`
}

func (out *TransferableOutput) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// Output returns the feature extension output that this Output is using.
func (out *TransferableOutput) Output() TransferableOut {
	_ = "STUB: not implemented"
	return *new(TransferableOut)
}

func (out *TransferableOutput) Verify() error { _ = "STUB: not implemented"; return nil }

type innerSortTransferableOutputs struct {
	outs  []*TransferableOutput
	codec codec.Manager
}

func (outs *innerSortTransferableOutputs) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

func (outs *innerSortTransferableOutputs) Len() int { _ = "STUB: not implemented"; return 0 }

func (outs *innerSortTransferableOutputs) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SortTransferableOutputs sorts output objects
func SortTransferableOutputs(outs []*TransferableOutput, c codec.Manager) {
	_ = "STUB: not implemented"
	return
}

// IsSortedTransferableOutputs returns true if output objects are sorted
func IsSortedTransferableOutputs(outs []*TransferableOutput, c codec.Manager) bool {
	_ = "STUB: not implemented"
	return false
}

type TransferableInput struct {
	UTXOID `serialize:"true"`
	Asset  `serialize:"true"`
	// FxID has serialize false because we don't want this to be encoded in bytes
	FxID ids.ID         `serialize:"false" json:"fxID"`
	In   TransferableIn `serialize:"true"  json:"input"`
}

// Input returns the feature extension input that this Input is using.
func (in *TransferableInput) Input() TransferableIn {
	_ = "STUB: not implemented"
	return *new(TransferableIn)
}

func (in *TransferableInput) Verify() error { _ = "STUB: not implemented"; return nil }

func (in *TransferableInput) Compare(other *TransferableInput) int {
	_ = "STUB: not implemented"
	return 0
}

type innerSortTransferableInputsWithSigners struct {
	ins     []*TransferableInput
	signers [][]*secp256k1.PrivateKey
}

func (ins *innerSortTransferableInputsWithSigners) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

func (ins *innerSortTransferableInputsWithSigners) Len() int { _ = "STUB: not implemented"; return 0 }

func (ins *innerSortTransferableInputsWithSigners) Swap(i, j int) {
	_ = "STUB: not implemented"
	return
}

// SortTransferableInputsWithSigners sorts the inputs and signers based on the
// input's utxo ID
func SortTransferableInputsWithSigners(ins []*TransferableInput, signers [][]*secp256k1.PrivateKey) {
	_ = "STUB: not implemented"
	return
}

// VerifyTx verifies that the inputs and outputs flowcheck, including a fee.
// Additionally, this verifies that the inputs and outputs are sorted.
func VerifyTx(
	feeAmount uint64,
	feeAssetID ids.ID,
	allIns [][]*TransferableInput,
	allOuts [][]*TransferableOutput,
	c codec.Manager,
) error {
	_ = "STUB: not implemented"
	return nil
}

// The txFee must be burned

// Add all the outputs to the flow checker and make sure they are sorted

// Add all the inputs to the flow checker and make sure they are sorted
