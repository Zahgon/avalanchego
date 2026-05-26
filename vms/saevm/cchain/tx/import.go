// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	// Imported for [atomic.UnsignedImportTx.Burned] comment resolution.
	_ "github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"

	"github.com/ava-labs/avalanchego/graft/coreth/core/extstate"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"

	chainsatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

var _ Unsigned = (*Import)(nil)

// Import is the unsigned component of a transaction that transfers assets from
// either the P-Chain or the X-Chain to the C-Chain. It consumes UTXOs in the
// shared memory between the C-Chain and the source chain and increases balances
// in the C-Chain state.
type Import struct {
	NetworkID      uint32                    `serialize:"true" json:"networkID"`
	BlockchainID   ids.ID                    `serialize:"true" json:"blockchainID"`
	SourceChain    ids.ID                    `serialize:"true" json:"sourceChain"`
	ImportedInputs []*avax.TransferableInput `serialize:"true" json:"importedInputs"`
	Outs           []Output                  `serialize:"true" json:"outputs"`
}

// Output specifies an account on the C-Chain whose balance of the specified
// asset should be increased.
//
// If the AssetID is AVAX, the amount will be scaled up to account for the EVM's
// higher denomination.
type Output struct {
	Address common.Address `serialize:"true" json:"address"`
	Amount  uint64         `serialize:"true" json:"amount"`
	AssetID ids.ID         `serialize:"true" json:"assetID"`
}

// Compare orders [Output] values by [Output.Address] and [Output.AssetID].
func (o Output) Compare(other Output) int { _ = "STUB: not implemented"; return 0 }

func (i *Import) inputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// Like [atomic.UnsignedImportTx.Burned], burned will error if the sum of the
// inputs exceeds MaxUint64, even if the total amount burned could be
// represented as a uint64.
//
// Because the total supply of AVAX fits in a uint64, this doesn't matter in
// practice and allows for easier fuzzing.
func (i *Import) burned(assetID ids.ID) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

var errOutputsNotSortedUnique = errors.New("outputs not sorted and unique")

func (i *Import) sanityCheck(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

var (
	errFetchingUTXOs      = errors.New("fetching UTXOs")
	errUnmarshallingUTXO  = errors.New("unmarshalling UTXO")
	errMismatchedAssetIDs = errors.New("mismatched asset IDs")
	errVerifyingTransfer  = errors.New("verifying transfer")
)

func (i *Import) verifyCredentials(sm chainsatomic.SharedMemory, creds []Credential) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(StephenButtolph): Parallelize transfer verification, which
// includes signature verification. This is non-trivial, because
// transactions frequently contain duplicate signatures, which are
// currently being cached.

var errUnexpectedInputType = errors.New("unexpected input type")

func (i *Import) numSigs() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *Import) asOp(avaxAssetID ids.ID) (op, error) {
	_ = "STUB: not implemented"
	return *new(op), nil
}

func (i *Import) atomicRequests(ids.ID) (ids.ID, *chainsatomic.Requests, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

// transferNonAVAX adds the non-AVAX balances to the statedb.
func (i *Import) transferNonAVAX(avaxAssetID ids.ID, statedb *extstate.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}
