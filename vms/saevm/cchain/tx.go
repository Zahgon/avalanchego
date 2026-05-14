// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
)

var _ hook.Transaction = (*transaction)(nil)

// transaction adapts a [tx.Tx] to the [hook.Transaction] interface.
type transaction struct {
	id     ids.ID
	tx     *tx.Tx
	inputs set.Set[ids.ID]
	op     hook.Op
}

func newTx(t *tx.Tx, avaxAssetID ids.ID) (*transaction, error) {
	op, err := t.AsOp(avaxAssetID)
	if err != nil {
		return nil, err
	}
	return &transaction{
		id:     op.ID,
		tx:     t,
		inputs: t.InputIDs(),
		op:     op,
	}, nil
}

func (t *transaction) AsOp() hook.Op { return t.op }
