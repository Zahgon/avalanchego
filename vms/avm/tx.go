// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var (
	_ snowstorm.Tx = (*Tx)(nil)

	errTxNotProcessing  = errors.New("transaction is not processing")
	errUnexpectedReject = errors.New("attempting to reject transaction")
)

type Tx struct {
	vm *VM
	tx *txs.Tx
}

func (tx *Tx) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (tx *Tx) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Tx) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) Status() choices.Status { _ = "STUB: not implemented"; return *new(choices.Status) }

func (tx *Tx) MissingDependencies() (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tx was already accepted

func (tx *Tx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (tx *Tx) Verify(context.Context) error { _ = "STUB: not implemented"; return nil }
