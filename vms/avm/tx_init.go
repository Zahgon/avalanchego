// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"reflect"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var _ txs.Visitor = (*txInit)(nil)

// txInit initializes FxID where required
type txInit struct {
	tx            *txs.Tx
	ctx           *snow.Context
	typeToFxIndex map[reflect.Type]int
	fxs           []*fxs.ParsedFx
}

func (t *txInit) getFx(val interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// getParsedFx returns the parsedFx object for a given TransferableInput
// or TransferableOutput object
func (t *txInit) getParsedFx(val interface{}) (*fxs.ParsedFx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *txInit) init() error {
	t.tx.Unsigned.InitCtx(t.ctx)

	for _, cred := range t.tx.Creds {
		fx, err := t.getParsedFx(cred.Credential)
		if err != nil {
			return err
		}
		cred.FxID = fx.ID
	}
	return nil
}

func (t *txInit) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (t *txInit) CreateAssetTx(tx *txs.CreateAssetTx) error { _ = "STUB: not implemented"; return nil }

func (t *txInit) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (t *txInit) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (t *txInit) OperationTx(tx *txs.OperationTx) error { _ = "STUB: not implemented"; return nil }
