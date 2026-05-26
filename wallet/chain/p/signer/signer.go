// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"

	stdcontext "context"
)

var _ Signer = (*txSigner)(nil)

type Signer interface {
	// Sign adds as many missing signatures as possible to the provided
	// transaction.
	//
	// If there are already some signatures on the transaction, those signatures
	// will not be removed.
	//
	// If the signer doesn't have the ability to provide a required signature,
	// the signature slot will be skipped without reporting an error.
	Sign(ctx stdcontext.Context, tx *txs.Tx) error
}

type Backend interface {
	GetUTXO(ctx stdcontext.Context, chainID, utxoID ids.ID) (*avax.UTXO, error)
	GetOwner(ctx stdcontext.Context, ownerID ids.ID) (fx.Owner, error)
}

type txSigner struct {
	kc      keychain.Keychain
	backend Backend
}

func New(kc keychain.Keychain, backend Backend) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

func (s *txSigner) Sign(ctx stdcontext.Context, tx *txs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func SignUnsigned(
	ctx stdcontext.Context,
	signer Signer,
	utx txs.UnsignedTx,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
