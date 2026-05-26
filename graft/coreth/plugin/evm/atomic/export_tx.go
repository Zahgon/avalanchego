// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_                           UnsignedAtomicTx       = (*UnsignedExportTx)(nil)
	_                           secp256k1fx.UnsignedTx = (*UnsignedExportTx)(nil)
	ErrExportNonAVAXInputBanff                         = errors.New("export input cannot contain non-AVAX in Banff")
	ErrExportNonAVAXOutputBanff                        = errors.New("export output cannot contain non-AVAX in Banff")
	ErrInsufficientFunds                               = errors.New("insufficient funds")
	ErrInvalidNonce                                    = errors.New("invalid nonce")
	ErrNoExportOutputs                                 = errors.New("tx has no export outputs")
	errOverflowExport                                  = errors.New("overflow when computing export amount + txFee")
)

// UnsignedExportTx is an unsigned ExportTx
type UnsignedExportTx struct {
	Metadata
	// ID of the network on which this tx was issued
	NetworkID uint32 `serialize:"true" json:"networkID"`
	// ID of this blockchain.
	BlockchainID ids.ID `serialize:"true" json:"blockchainID"`
	// Which chain to send the funds to
	DestinationChain ids.ID `serialize:"true" json:"destinationChain"`
	// Inputs
	Ins []EVMInput `serialize:"true" json:"inputs"`
	// Outputs that are exported to the chain
	ExportedOutputs []*avax.TransferableOutput `serialize:"true" json:"exportedOutputs"`
}

// InputUTXOs returns a set of all the hash(address:nonce) exporting funds.
func (utx *UnsignedExportTx) InputUTXOs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// Total populated bytes is exactly 32 bytes.
// 8 (Nonce) + 4 (Address Length) + 20 (Address)

// Verify this transaction is well-formed
func (utx *UnsignedExportTx) Verify(
	ctx *snow.Context,
	rules extras.Rules,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure that the tx has a valid peer chain ID

// Note that SameSubnet verifies that [tx.DestinationChain] isn't this
// chain's ID

func (utx *UnsignedExportTx) GasUsed(fixedFee bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Amount of [assetID] burned by this transaction
func (utx *UnsignedExportTx) Burned(assetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (utx *UnsignedExportTx) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }

// AtomicOps returns the atomic operations for this transaction.
func (utx *UnsignedExportTx) AtomicOps() (ids.ID, *atomic.Requests, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

// NewExportTx returns a new ExportTx
func NewExportTx(
	ctx *snow.Context,
	rules extras.Rules,
	state StateDB,
	assetID ids.ID, // AssetID of the tokens to export
	amount uint64, // Amount of tokens to export
	chainID ids.ID, // Chain to send the UTXOs to
	to ids.ShortID, // Address of chain recipient
	baseFee *big.Int, // fee to use post-AP3
	keys []*secp256k1.PrivateKey, // Pay the fee and provide the tokens
) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:prealloc // sizes depend on runtime values not known at declaration
//nolint:prealloc // sizes depend on runtime values not known at declaration

// consume non-AVAX

// Create the transaction

// EVMStateTransfer executes the state update from the atomic export transaction
func (utx *UnsignedExportTx) EVMStateTransfer(ctx *snow.Context, state StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

// We multiply the input amount by x2cRate to convert AVAX back to the appropriate
// denomination before export.

// getSpendableFunds returns a list of EVMInputs and keys (in corresponding
// order) to total [amount] of [assetID] owned by [keys].
// Note: we return [][]*secp256k1.PrivateKey even though each input
// corresponds to a single key, so that the signers can be passed in to
// [tx.Sign] which supports multiple keys on a single input.
func getSpendableFunds(
	ctx *snow.Context,
	state StateDB,
	keys []*secp256k1.PrivateKey,
	assetID ids.ID,
	amount uint64,
) ([]EVMInput, [][]*secp256k1.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Note: we assume that each key in [keys] is unique, so that iterating over
// the keys will not produce duplicated nonces in the returned EVMInput slice.

// If the asset is AVAX, we divide by the x2cRate to convert back to the correct
// denomination of AVAX that can be exported.

// getSpendableAVAXWithFee returns a list of EVMInputs and keys (in corresponding
// order) to total [amount] + [fee] of [AVAX] owned by [keys].
// This function accounts for the added cost of the additional inputs needed to
// create the transaction and makes sure to skip any keys with a balance that is
// insufficient to cover the additional fee.
// Note: we return [][]*secp256k1.PrivateKey even though each input
// corresponds to a single key, so that the signers can be passed in to
// [tx.Sign] which supports multiple keys on a single input.
func getSpendableAVAXWithFee(
	ctx *snow.Context,
	state StateDB,
	keys []*secp256k1.PrivateKey,
	amount uint64,
	cost uint64,
	baseFee *big.Int,
) ([]EVMInput, [][]*secp256k1.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Note: we assume that each key in [keys] is unique, so that iterating over
// the keys will not produce duplicated nonces in the returned EVMInput slice.

// Since the asset is AVAX, we divide by the x2cRate to convert back to
// the correct denomination of AVAX that can be exported.

// If the balance for [addr] is insufficient to cover the additional cost
// of adding an input to the transaction, skip adding the input altogether

// Update the cost for the next iteration

// Use the entire [balance] as an input, but if the required [amount]
// is less than the balance, update the [inputAmount] to spend the
// minimum amount to finish the transaction.
