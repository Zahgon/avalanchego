// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ txs.Visitor = (*visitor)(nil)

	ErrUnsupportedTxType     = errors.New("unsupported tx type")
	ErrUnknownInputType      = errors.New("unknown input type")
	ErrUnknownOutputType     = errors.New("unknown output type")
	ErrInvalidUTXOSigIndex   = errors.New("invalid UTXO signature index")
	ErrUnknownAuthType       = errors.New("unknown auth type")
	ErrUnknownOwnerType      = errors.New("unknown owner type")
	ErrUnknownCredentialType = errors.New("unknown credential type")

	emptySig [secp256k1.SignatureLen]byte
)

// visitor handles signing transactions for the signer
type visitor struct {
	kc      keychain.Keychain
	backend Backend
	ctx     context.Context
	tx      *txs.Tx
}

func (*visitor) AdvanceTimeTx(*txs.AdvanceTimeTx) error { _ = "STUB: not implemented"; return nil }

func (*visitor) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) AddValidatorTx(tx *txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) AddSubnetValidatorTx(tx *txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) AddDelegatorTx(tx *txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) CreateChainTx(tx *txs.CreateChainTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) CreateSubnetTx(tx *txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) RemoveSubnetValidatorTx(tx *txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) TransformSubnetTx(tx *txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) AddPermissionlessValidatorTx(tx *txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) AddPermissionlessDelegatorTx(tx *txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) TransferSubnetOwnershipTx(tx *txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (s *visitor) ConvertSubnetToL1Tx(tx *txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) RegisterL1ValidatorTx(tx *txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) SetL1ValidatorWeightTx(tx *txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) IncreaseL1ValidatorBalanceTx(tx *txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) DisableL1ValidatorTx(tx *txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *visitor) getSigners(sourceChainID ids.ID, ins []*avax.TransferableInput) ([][]keychain.Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we don't have access to the UTXO, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func (s *visitor) getAuthSigners(ownerID ids.ID, auth verify.Verifiable) ([]keychain.Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

func sign(tx *txs.Tx, txSigners [][]keychain.Signer) error { _ = "STUB: not implemented"; return nil }

// If we don't have access to the key, then we can't sign this
// transaction. However, we can attempt to partially sign it.

// If this signature has already been populated, we can just
// copy the needed signature for the future.

// If this key has already produced a signature, we can just
// copy the previous signature.
