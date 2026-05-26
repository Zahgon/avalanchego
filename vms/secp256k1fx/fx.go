// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const defaultCacheSize = 256

var (
	ErrWrongVMType                    = errors.New("wrong vm type")
	ErrWrongTxType                    = errors.New("wrong tx type")
	ErrWrongOpType                    = errors.New("wrong operation type")
	ErrWrongUTXOType                  = errors.New("wrong utxo type")
	ErrWrongInputType                 = errors.New("wrong input type")
	ErrWrongCredentialType            = errors.New("wrong credential type")
	ErrWrongOwnerType                 = errors.New("wrong owner type")
	ErrMismatchedAmounts              = errors.New("utxo amount and input amount are not equal")
	ErrWrongNumberOfUTXOs             = errors.New("wrong number of utxos for the operation")
	ErrWrongMintCreated               = errors.New("wrong mint output created from the operation")
	ErrTimelocked                     = errors.New("output is time locked")
	ErrTooManySigners                 = errors.New("input has more signers than expected")
	ErrTooFewSigners                  = errors.New("input has less signers than expected")
	ErrInputOutputIndexOutOfBounds    = errors.New("input referenced a nonexistent address in the output")
	ErrInputCredentialSignersMismatch = errors.New("input expected a different number of signers than provided in the credential")
	ErrWrongSig                       = errors.New("wrong signature")
)

// Fx describes the secp256k1 feature extension
type Fx struct {
	VM           VM
	bootstrapped bool
	recoverCache *secp256k1.RecoverCache
}

func (fx *Fx) Initialize(vmIntf interface{}) error { _ = "STUB: not implemented"; return nil }

func (fx *Fx) InitializeVM(vmIntf interface{}) error { _ = "STUB: not implemented"; return nil }

func (*Fx) Bootstrapping() error { _ = "STUB: not implemented"; return nil }

func (fx *Fx) Bootstrapped() error { _ = "STUB: not implemented"; return nil }

// VerifyPermission returns nil iff [credIntf] proves that [controlGroup] assents to [txIntf]
func (fx *Fx) VerifyPermission(txIntf, inIntf, credIntf, ownerIntf interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fx *Fx) VerifyOperation(txIntf, opIntf, credIntf interface{}, utxosIntf []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (fx *Fx) verifyOperation(tx UnsignedTx, op *MintOperation, cred *Credential, utxo *MintOutput) error {
	_ = "STUB: not implemented"
	return nil
}

func (fx *Fx) VerifyTransfer(txIntf, inIntf, credIntf, utxoIntf interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifySpend ensures that the utxo can be sent to any address
func (fx *Fx) VerifySpend(utx UnsignedTx, in *TransferInput, cred *Credential, utxo *TransferOutput) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyCredentials ensures that the output can be spent by the input with the
// credential. A nil return values means the output can be spent.
func (fx *Fx) VerifyCredentials(utx UnsignedTx, in *Input, cred *Credential, out *OutputOwners) error {
	_ = "STUB: not implemented"
	return nil
}

// disable signature verification during bootstrapping

// Make sure the input references an address that exists

// Make sure each signature in the signature list is from an owner of
// the output being consumed

// CreateOutput creates a new output with the provided control group worth
// the specified amount
func (*Fx) CreateOutput(amount uint64, ownerIntf interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
