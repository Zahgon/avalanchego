// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

// fx exposes [secp256k1fx.Fx.VerifyTransfer] for verification of UTXO transfers
// from shared memory into the C-Chain.
var fx secp256k1fx.Fx

func init() {
	if err := fx.Initialize(fxVM{}); err != nil {
		panic(err)
	}
	// Mark the fx as bootstrapped so that it actually verifies signatures.
	if err := fx.Bootstrapped(); err != nil {
		panic(err)
	}
}

type fxVM struct {
	clock mockable.Clock
}

func (fxVM) CodecRegistry() codec.Registry { _ = "STUB: not implemented"; return *new(codec.Registry) }
func (f fxVM) Clock() *mockable.Clock      { _ = "STUB: not implemented"; return nil }
func (fxVM) Logger() logging.Logger        { _ = "STUB: not implemented"; return *new(logging.Logger) }

var _ secp256k1fx.UnsignedTx = (*fxTx)(nil)

type fxTx []byte

func toFxTx(u Unsigned) (fxTx, error) { _ = "STUB: not implemented"; return *new(fxTx), nil }

func (f fxTx) Bytes() []byte { _ = "STUB: not implemented"; return nil }
