// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_ UnsignedTx             = (*CreateAssetTx)(nil)
	_ secp256k1fx.UnsignedTx = (*CreateAssetTx)(nil)
)

// CreateAssetTx is a transaction that creates a new asset.
type CreateAssetTx struct {
	BaseTx       `serialize:"true"`
	Name         string          `serialize:"true" json:"name"`
	Symbol       string          `serialize:"true" json:"symbol"`
	Denomination byte            `serialize:"true" json:"denomination"`
	States       []*InitialState `serialize:"true" json:"initialStates"`
}

func (t *CreateAssetTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// InitialStates track which virtual machines, and the initial state of these
// machines, this asset uses. The returned array should not be modified.
func (t *CreateAssetTx) InitialStates() []*InitialState { _ = "STUB: not implemented"; return nil }

func (t *CreateAssetTx) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }
