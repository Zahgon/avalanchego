// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/state"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ txs.Visitor = (*SemanticVerifier)(nil)

	errAssetIDMismatch = errors.New("asset IDs in the input don't match the utxo")
	errNotAnAsset      = errors.New("not an asset")
	errIncompatibleFx  = errors.New("incompatible feature extension")
	errUnknownFx       = errors.New("unknown feature extension")
)

type SemanticVerifier struct {
	*Backend
	State state.ReadOnlyChain
	Tx    *txs.Tx
}

func (v *SemanticVerifier) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

// Note: Verification of the length of [t.tx.Creds] happens during
// syntactic verification, which happens before semantic verification.

func (v *SemanticVerifier) CreateAssetTx(tx *txs.CreateAssetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SemanticVerifier) OperationTx(tx *txs.OperationTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: Verification of the length of [t.tx.Creds] happens during
// syntactic verification, which happens before semantic verification.

func (v *SemanticVerifier) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

// Note: Verification of the length of [t.tx.Creds] happens during
// syntactic verification, which happens before semantic verification.

func (v *SemanticVerifier) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (v *SemanticVerifier) verifyTransfer(
	tx txs.UnsignedTx,
	in *avax.TransferableInput,
	cred verify.Verifiable,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SemanticVerifier) verifyTransferOfUTXO(
	tx txs.UnsignedTx,
	in *avax.TransferableInput,
	cred verify.Verifiable,
	utxo *avax.UTXO,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SemanticVerifier) verifyOperation(
	tx *txs.OperationTx,
	op *txs.Operation,
	cred verify.Verifiable,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SemanticVerifier) verifyFxUsage(
	fxID int,
	assetID ids.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *SemanticVerifier) getFx(val interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
