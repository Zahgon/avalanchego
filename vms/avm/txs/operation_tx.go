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
	_ UnsignedTx             = (*OperationTx)(nil)
	_ secp256k1fx.UnsignedTx = (*OperationTx)(nil)
)

// OperationTx is a transaction with no credentials.
type OperationTx struct {
	BaseTx `serialize:"true"`

	Ops []*Operation `serialize:"true" json:"operations"`
}

func (t *OperationTx) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

// Operations track which ops this transaction is performing. The returned array
// should not be modified.
func (t *OperationTx) Operations() []*Operation { _ = "STUB: not implemented"; return nil }

func (t *OperationTx) InputUTXOs() []*avax.UTXOID { _ = "STUB: not implemented"; return nil }

func (t *OperationTx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

// NumCredentials returns the number of expected credentials
func (t *OperationTx) NumCredentials() int { _ = "STUB: not implemented"; return 0 }

func (t *OperationTx) Visit(v Visitor) error { _ = "STUB: not implemented"; return nil }
