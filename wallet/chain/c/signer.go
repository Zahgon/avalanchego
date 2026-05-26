// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

const version = 0

var (
	_ Signer = (*txSigner)(nil)

	errUnknownInputType      = errors.New("unknown input type")
	errUnknownCredentialType = errors.New("unknown credential type")
	errUnknownOutputType     = errors.New("unknown output type")
	errInvalidUTXOSigIndex   = errors.New("invalid UTXO signature index")

	emptySig [secp256k1.SignatureLen]byte
)

type Signer interface {
	// SignAtomic adds as many missing signatures as possible to the provided
	// transaction.
	//
	// If there are already some signatures on the transaction, those signatures
	// will not be removed.
	//
	// If the signer doesn't have the ability to provide a required signature,
	// the signature slot will be skipped without reporting an error.
	SignAtomic(ctx context.Context, tx *atomic.Tx) error
}

type EthKeychain interface {
	// The returned Signer can provide a signature for [addr]
	GetEth(addr common.Address) (keychain.Signer, bool)
	// Returns the set of addresses for which the accessor keeps an associated
	// signer
	EthAddresses() set.Set[common.Address]
}

type SignerBackend interface {
	GetUTXO(ctx context.Context, chainID, utxoID ids.ID) (*avax.UTXO, error)
}

type txSigner struct {
	avaxKC  keychain.Keychain
	ethKC   EthKeychain
	backend SignerBackend
}

func NewSigner(avaxKC keychain.Keychain, ethKC EthKeychain, backend SignerBackend) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

func (s *txSigner) SignAtomic(ctx context.Context, tx *atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *txSigner) getImportSigners(ctx context.Context, sourceChainID ids.ID, ins []*avax.TransferableInput) ([][]keychain.Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we don't have access to the UTXO, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func (s *txSigner) getExportSigners(ins []atomic.EVMInput) [][]keychain.Signer {
	_ = "STUB: not implemented"
	return nil
}

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func SignUnsignedAtomic(ctx context.Context, signer Signer, utx atomic.UnsignedAtomicTx) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sign(tx *atomic.Tx, txSigners [][]keychain.Signer) error {
	_ = "STUB: not implemented"
	return nil
}

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If this signature has already been populated, we can just
// copy the needed signature for the future.

// If this key has already produced a signature, we can just
// copy the previous signature.
