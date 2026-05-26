// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_                           UnsignedAtomicTx       = (*UnsignedImportTx)(nil)
	_                           secp256k1fx.UnsignedTx = (*UnsignedImportTx)(nil)
	ErrImportNonAVAXInputBanff                         = errors.New("import input cannot contain non-AVAX in Banff")
	ErrImportNonAVAXOutputBanff                        = errors.New("import output cannot contain non-AVAX in Banff")
	ErrNoImportInputs                                  = errors.New("tx has no imported inputs")
	ErrWrongChainID                                    = errors.New("tx has wrong chain ID")
	ErrNoEVMOutputs                                    = errors.New("tx has no EVM outputs")
	ErrInputsNotSortedUnique                           = errors.New("inputs not sorted and unique")
	ErrOutputsNotSortedUnique                          = errors.New("outputs not sorted and unique")
	ErrOutputsNotSorted                                = errors.New("tx outputs not sorted")
	errNilBaseFeeApricotPhase3                         = errors.New("nil base fee is invalid after apricotPhase3")
	errInsufficientFundsForFee                         = errors.New("insufficient AVAX funds to pay transaction fee")
)

// UnsignedImportTx is an unsigned ImportTx
type UnsignedImportTx struct {
	Metadata
	// ID of the network on which this tx was issued
	NetworkID uint32 `serialize:"true" json:"networkID"`
	// ID of this blockchain.
	BlockchainID ids.ID `serialize:"true" json:"blockchainID"`
	// Which chain to consume the funds from
	SourceChain ids.ID `serialize:"true" json:"sourceChain"`
	// Inputs that consume UTXOs produced on the chain
	ImportedInputs []*avax.TransferableInput `serialize:"true" json:"importedInputs"`
	// Outputs
	Outs []EVMOutput `serialize:"true" json:"outputs"`
}

// InputUTXOs returns the UTXOIDs of the imported funds
func (utx *UnsignedImportTx) InputUTXOs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// Verify this transaction is well-formed
func (utx *UnsignedImportTx) Verify(
	ctx *snow.Context,
	rules extras.Rules,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure that the tx has a valid peer chain ID

// Note that SameSubnet verifies that [tx.SourceChain] isn't this
// chain's ID

func (utx *UnsignedImportTx) GasUsed(fixedFee bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Amount of [assetID] burned by this transaction
func (utx *UnsignedImportTx) Burned(assetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AtomicOps returns imported inputs spent on this transaction
// We spend imported UTXOs here rather than in verification because
// we don't want to remove an imported UTXO in verification
// only to have the transaction not be Accepted. This would be inconsistent.
// Recall that imported UTXOs are not kept in a versionDB.
func (utx *UnsignedImportTx) AtomicOps() (ids.ID, *atomic.Requests, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

// NewImportTx returns a new ImportTx
func NewImportTx(
	ctx *snow.Context,
	rules extras.Rules,
	time uint64,
	chainID ids.ID, // chain to import from
	to common.Address, // Address of recipient
	baseFee *big.Int, // fee to use post-AP3
	kc *secp256k1fx.Keychain, // Keychain to use for signing the atomic UTXOs
	atomicUTXOs []*avax.UTXO, // UTXOs to spend
) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This will create unique outputs (in the context of sorting)
// since each output will have a unique assetID

// Skip the AVAX amount since it is included separately to account for
// the fee

// AVAX output
// imported amount goes toward paying tx fee

// If no outputs are produced, return an error.
// Note: this can happen if there is exactly enough AVAX to pay the
// transaction fee, but no other funds to be imported.

// Create the transaction

// EVMStateTransfer performs the state transfer to increase the balances of
// accounts accordingly with the imported EVMOutputs
func (utx *UnsignedImportTx) EVMStateTransfer(ctx *snow.Context, state StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

// If the asset is AVAX, convert the input amount in nAVAX to gWei by
// multiplying by the x2c rate.

func (utx *UnsignedImportTx) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }
