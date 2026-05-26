// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ txs.Visitor = (*visitor)(nil)

	ErrUnknownInputType      = errors.New("unknown input type")
	ErrUnknownOpType         = errors.New("unknown operation type")
	ErrInvalidNumUTXOsInOp   = errors.New("invalid number of UTXOs in operation")
	ErrUnknownCredentialType = errors.New("unknown credential type")
	ErrUnknownOutputType     = errors.New("unknown output type")
	ErrInvalidUTXOSigIndex   = errors.New("invalid UTXO signature index")

	emptySig [secp256k1.SignatureLen]byte
)

// visitor handles signing transactions for the signer
type visitor struct {
	kc      keychain.Keychain
	backend Backend
	ctx     context.Context
	tx      *txs.Tx
}

func (s *visitor) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) CreateAssetTx(tx *txs.CreateAssetTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) OperationTx(tx *txs.OperationTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) getSigners(ctx context.Context, sourceChainID ids.ID, ins []*avax.TransferableInput) ([]verify.Verifiable, [][]keychain.Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If we don't have access to the UTXO, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func (s *visitor) getOpsSigners(ctx context.Context, sourceChainID ids.ID, ops []*txs.Operation) ([]verify.Verifiable, [][]keychain.Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If we don't have access to the UTXO, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func sign(tx *txs.Tx, creds []verify.Verifiable, txSigners [][]keychain.Signer) error {
	_ = "STUB: not implemented"
	return nil
}

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If this signature has already been populated, we can just
// copy the needed signature for the future.

// If this key has already produced a signature, we can just
// copy the previous signature.
