// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utxo

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Verifier = (*verifier)(nil)

	ErrUnsupportedTxType            = errors.New("unsupported tx type")
	ErrInsufficientUnlockedFunds    = errors.New("insufficient unlocked funds")
	ErrInsufficientLockedFunds      = errors.New("insufficient locked funds")
	errWrongNumberCredentials       = errors.New("wrong number of credentials")
	errWrongNumberUTXOs             = errors.New("wrong number of UTXOs")
	errAssetIDMismatch              = errors.New("input asset ID does not match UTXO asset ID")
	errLocktimeMismatch             = errors.New("input locktime does not match UTXO locktime")
	errLockedFundsNotMarkedAsLocked = errors.New("locked funds not marked as locked")
)

type Verifier interface {
	// Verify that [tx] is semantically valid.
	// [ins] and [outs] are the inputs and outputs of [tx].
	// [creds] are the credentials of [tx], which allow [ins] to be spent.
	// [unlockedProduced] is the map of assets that were produced and their
	// amounts.
	// The [ins] must have at least [unlockedProduced] than the [outs].
	//
	// Precondition: [tx] has already been syntactically verified.
	//
	// Note: [unlockedProduced] is modified by this method.
	VerifySpend(
		tx txs.UnsignedTx,
		utxoDB avax.UTXOGetter,
		ins []*avax.TransferableInput,
		outs []*avax.TransferableOutput,
		creds []verify.Verifiable,
		unlockedProduced map[ids.ID]uint64,
	) error

	// Verify that [tx] is semantically valid.
	// [utxos[i]] is the UTXO being consumed by [ins[i]].
	// [ins] and [outs] are the inputs and outputs of [tx].
	// [creds] are the credentials of [tx], which allow [ins] to be spent.
	// [unlockedProduced] is the map of assets that were produced and their
	// amounts.
	// The [ins] must have at least [unlockedProduced] more than the [outs].
	//
	// Precondition: [tx] has already been syntactically verified.
	//
	// Note: [unlockedProduced] is modified by this method.
	VerifySpendUTXOs(
		tx txs.UnsignedTx,
		utxos []*avax.UTXO,
		ins []*avax.TransferableInput,
		outs []*avax.TransferableOutput,
		creds []verify.Verifiable,
		unlockedProduced map[ids.ID]uint64,
	) error
}

func NewVerifier(
	ctx *snow.Context,
	clk *mockable.Clock,
	fx fx.Fx,
) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

type verifier struct {
	ctx *snow.Context
	clk *mockable.Clock
	fx  fx.Fx
}

func (h *verifier) VerifySpend(
	tx txs.UnsignedTx,
	utxoDB avax.UTXOGetter,
	ins []*avax.TransferableInput,
	outs []*avax.TransferableOutput,
	creds []verify.Verifiable,
	unlockedProduced map[ids.ID]uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *verifier) VerifySpendUTXOs(
	tx txs.UnsignedTx,
	utxos []*avax.UTXO,
	ins []*avax.TransferableInput,
	outs []*avax.TransferableOutput,
	creds []verify.Verifiable,
	unlockedProduced map[ids.ID]uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify credentials are well-formed.

// Time this transaction is being verified

// Track the amount of unlocked transfers
// assetID -> amount

// Track the amount of locked transfers and their owners
// assetID -> locktime -> ownerID -> amount

// The UTXO consumed by [input]

// Set [locktime] to this UTXO's locktime, if applicable

// The UTXO says it's locked until [locktime], but this input, which
// consumes it, is not locked even though [locktime] hasn't passed. This
// is invalid.

// This input is locked, but its locktime is wrong

// Verify that this tx's credentials allow [in] to be spent

// Set [locktime] to this output's locktime, if applicable

// Make sure that for each assetID and locktime, tokens produced <= tokens consumed

// More unlocked tokens produced than consumed. Invalid.

// GetInputOutputs returns the input/output utxos and any AVAX that is produced
// as part of the execution of the tx
func GetInputOutputs(tx txs.UnsignedTx) (
	[]*avax.TransferableInput,
	[]*avax.TransferableOutput,
	uint64,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// inputOutputGetter gets the utxos and AVAX produced for each tx type
type inputOutputGetter struct {
	// InputUTXOs is the utxos consumed by the tx
	InputUTXOs []*avax.TransferableInput
	// OutputUTXOs is the utxos produced by the tx
	OutputUTXOs []*avax.TransferableOutput
	// ProducedAVAX is produced by the execution of this tx that does not have a
	// corresponding UTXO
	ProducedAVAX uint64
}

func (i *inputOutputGetter) AddValidatorTx(tx *txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) AddSubnetValidatorTx(tx *txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) AddDelegatorTx(tx *txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) CreateChainTx(tx *txs.CreateChainTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) CreateSubnetTx(tx *txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (i *inputOutputGetter) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (*inputOutputGetter) AdvanceTimeTx(*txs.AdvanceTimeTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*inputOutputGetter) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) RemoveSubnetValidatorTx(tx *txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) TransformSubnetTx(tx *txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) AddPermissionlessValidatorTx(tx *txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) AddPermissionlessDelegatorTx(tx *txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) TransferSubnetOwnershipTx(tx *txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

// ConvertSubnetToL1Tx treats validator balances like produced AVAX because
// the fee payer must have enough input AVAX to cover the initial state of the
// L1 validators
func (i *inputOutputGetter) ConvertSubnetToL1Tx(tx *txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterL1ValidatorTx treats the validator balance like produced AVAX because
// the fee payer must have enough input AVAX to cover the initial state of the
// validator
func (i *inputOutputGetter) RegisterL1ValidatorTx(tx *txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) SetL1ValidatorWeightTx(tx *txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterL1ValidatorTx treats the validator balance like produced AVAX because
// the fee payer must have enough input AVAX to cover the increase in balance
func (i *inputOutputGetter) IncreaseL1ValidatorBalanceTx(tx *txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) DisableL1ValidatorTx(tx *txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *inputOutputGetter) getUTXOs(tx txs.BaseTx) { _ = "STUB: not implemented"; return }
