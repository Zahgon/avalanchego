// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_ UnsignedTx             = (*BaseTx)(nil)
	_ secp256k1fx.UnsignedTx = (*BaseTx)(nil)
)

// BaseTx is the basis of all transactions.
type BaseTx struct {
	avax.BaseTx `serialize:"true"`

	bytes []byte
}

func (t *BaseTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (t *BaseTx) SetBytes(bytes []byte) { _ = "STUB: not implemented"; return }

func (t *BaseTx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (t *BaseTx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (t *BaseTx) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }
