// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var _ atomic.Visitor = (*semanticVerifier)(nil)

var (
	ErrAssetIDMismatch            = errors.New("asset IDs in the input don't match the utxo")
	ErrConflictingAtomicInputs    = errors.New("invalid block due to conflicting atomic inputs")
	errFailedToFetchImportUTXOs   = errors.New("failed to fetch import UTXOs")
	errFailedToUnmarshalUTXO      = errors.New("failed to unmarshal UTXO")
	errRejectedParent             = errors.New("rejected parent")
	errIncorrectNumCredentials    = errors.New("incorrect number of credentials")
	errIncorrectNumSignatures     = errors.New("incorrect number of signatures")
	errPublicKeySignatureMismatch = errors.New("signature doesn't match public key")
)

type BlockFetcher interface {
	// GetExtendedBlock returns the ExtendedBlock for the given ID or an error if the block is not found
	GetExtendedBlock(context.Context, ids.ID) (extension.ExtendedBlock, error)
	// LastAcceptedExtendedBlock returns the last accepted VM block
	LastAcceptedExtendedBlock() extension.ExtendedBlock
}

type VerifierBackend struct {
	Ctx          *snow.Context
	Fx           *secp256k1fx.Fx
	Rules        extras.Rules
	Bootstrapped bool
	BlockFetcher BlockFetcher
	SecpCache    *secp256k1.RecoverCache
}

func NewVerifierBackend(vm *VM, rules extras.Rules) *VerifierBackend {
	_ = "STUB: not implemented"
	return nil
}

// SemanticVerify checks the semantic validity of atomic transactions.
func (b *VerifierBackend) SemanticVerify(tx *atomic.Tx, parent extension.ExtendedBlock, baseFee *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// semanticVerifier is a visitor that checks the semantic validity of atomic transactions.
type semanticVerifier struct {
	backend *VerifierBackend
	tx      *atomic.Tx
	parent  extension.ExtendedBlock
	baseFee *big.Int
}

// ImportTx verifies this transaction is valid.
func (s *semanticVerifier) ImportTx(utx *atomic.UnsignedImportTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Check the transaction consumes and produces the right amounts

// Apply dynamic fees to import transactions as of Apricot Phase 3

// Apply fees to import transactions as of Apricot Phase 2

// Allow for force committing during bootstrapping

// allUTXOBytes is guaranteed to be the same length as utxoIDs

// conflicts returns an error if [inputs] conflicts with any of the atomic inputs contained in [ancestor]
// or any of its ancestor blocks going back to the last accepted block in its ancestry. If [ancestor] is
// accepted, then nil will be returned immediately.
// If the ancestry of [ancestor] cannot be fetched, then [errRejectedParent] may be returned.
func conflicts(backend *VerifierBackend, inputs set.Set[ids.ID], ancestor extension.ExtendedBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// If any of the atomic transactions in the ancestor conflict with [inputs]
// return an error.

// Move up the chain.

// If the ancestor is unknown, then the parent failed
// verification when it was called.
// If the ancestor is rejected, then this block shouldn't be
// inserted into the canonical chain because the parent is
// will be missing.
// If the ancestor is processing, then the block may have
// been verified.

// ExportTx verifies this transaction is valid.
func (s *semanticVerifier) ExportTx(utx *atomic.UnsignedExportTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Check the transaction consumes and produces the right amounts

// Apply dynamic fees to export transactions as of Apricot Phase 3

// Apply fees to export transactions before Apricot Phase 3
