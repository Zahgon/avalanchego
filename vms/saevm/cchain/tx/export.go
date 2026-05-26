// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	// Imported for [atomic.UnsignedExportTx.Burned] comment resolution.
	_ "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"

	"github.com/ava-labs/avalanchego/graft/coreth/core/extstate"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"

	chainsatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

var _ Unsigned = (*Export)(nil)

// Export is the unsigned component of a transaction that transfers assets from
// the C-Chain to either the P-Chain or the X-Chain. It modifies the C-Chain
// state and produces UTXOs in the shared memory between the C-Chain and the
// destination chain.
type Export struct {
	NetworkID        uint32                     `serialize:"true" json:"networkID"`
	BlockchainID     ids.ID                     `serialize:"true" json:"blockchainID"`
	DestinationChain ids.ID                     `serialize:"true" json:"destinationChain"`
	Ins              []Input                    `serialize:"true" json:"inputs"`
	ExportedOutputs  []*avax.TransferableOutput `serialize:"true" json:"exportedOutputs"`
}

// Input identifies an account + nonce pair on the C-Chain that authorizes the
// asset and quantity to deduct.
//
// If the AssetID is AVAX, the amount will be scaled up to account for the EVM's
// higher denomination.
type Input struct {
	Address common.Address `serialize:"true" json:"address"`
	Amount  uint64         `serialize:"true" json:"amount"`
	AssetID ids.ID         `serialize:"true" json:"assetID"`
	Nonce   uint64         `serialize:"true" json:"nonce"`
}

// Compare orders [Input] values by [Input.Address] and [Input.AssetID].
func (i Input) Compare(other Input) int { _ = "STUB: not implemented"; return 0 }

func (e *Export) inputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// AccountInputID returns the Account+Nonce pair as a unique [ids.ID].
//
// It is safe to assume that the returned ID never conflicts with a UTXO ID.
func AccountInputID(address common.Address, nonce uint64) ids.ID {
	_ = "STUB: not implemented"
	return *new(ids.ID)
}

// 32 bytes long
// add 8 bytes
// add 24 bytes

// Like [atomic.UnsignedExportTx.Burned], burned will error if the sum of the
// inputs exceeds MaxUint64, even if the total amount burned could be
// represented as a uint64.
//
// Because the total supply of AVAX fits in a uint64, this doesn't matter in
// practice and allows for easier fuzzing.
func (e *Export) burned(assetID ids.ID) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

var errOutputsNotSorted = errors.New("outputs not sorted")

func (e *Export) sanityCheck(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

// Like [atomic.UnsignedExportTx.Verify], outputs aren't enforced to be
// unique. This is safe because each output's UTXO is keyed by txID and
// outputIndex, so duplicate outputs still produce distinct UTXOs.

var (
	sigCache = secp256k1.NewRecoverCache(1024)

	errIncorrectNumSignatures = errors.New("incorrect number of signatures")
	errRecoveringPublicKey    = errors.New("recovering public key")
	errAddressMismatch        = errors.New("signature does not match address")
)

func (e *Export) verifyCredentials(_ chainsatomic.SharedMemory, creds []Credential) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(StephenButtolph): Parallelize signature verification. This is
// non-trivial, because transactions frequently contain duplicate
// signatures, which are currently being cached.

func (e *Export) numSigs() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

var errMultipleNonces = errors.New("multiple nonces for address")

func (e *Export) asOp(avaxAssetID ids.ID) (op, error) {
	_ = "STUB: not implemented"
	return *new(op), nil
}

// Even if no AVAX is debited, non-AVAX inputs MUST increment the nonce.

func (e *Export) atomicRequests(txID ids.ID) (ids.ID, *chainsatomic.Requests, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

//#nosec G115 -- Won't overflow

var errInsufficientFunds = errors.New("insufficient funds")

// transferNonAVAX subtracts the non-AVAX balances from the statedb.
func (e *Export) transferNonAVAX(avaxAssetID ids.ID, statedb *extstate.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}
