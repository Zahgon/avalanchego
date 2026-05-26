// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"errors"

	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	errWrongTxType         = errors.New("wrong tx type")
	errWrongUTXOType       = errors.New("wrong utxo type")
	errWrongOperationType  = errors.New("wrong operation type")
	errWrongCredentialType = errors.New("wrong credential type")
	errWrongNumberOfUTXOs  = errors.New("wrong number of UTXOs for the operation")
	errWrongMintOutput     = errors.New("wrong mint output provided")
	errCantTransfer        = errors.New("cant transfer with this fx")
)

type Fx struct{ secp256k1fx.Fx }

func (fx *Fx) Initialize(vmIntf interface{}) error { _ = "STUB: not implemented"; return nil }

func (fx *Fx) VerifyOperation(txIntf, opIntf, credIntf interface{}, utxosIntf []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fx *Fx) VerifyMintOperation(tx secp256k1fx.UnsignedTx, op *MintOperation, cred *Credential, utxoIntf interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fx *Fx) VerifyTransferOperation(tx secp256k1fx.UnsignedTx, op *BurnOperation, cred *Credential, utxoIntf interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Fx) VerifyTransfer(_, _, _, _ interface{}) error { _ = "STUB: not implemented"; return nil }
