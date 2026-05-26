// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm/state"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var _ txs.Visitor = (*Executor)(nil)

type Executor struct {
	Codec          codec.Manager
	State          state.Chain // state will be modified
	Tx             *txs.Tx
	Inputs         set.Set[ids.ID]             // imported inputs
	AtomicRequests map[ids.ID]*atomic.Requests // may be nil
}

func (e *Executor) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (e *Executor) CreateAssetTx(tx *txs.CreateAssetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) OperationTx(tx *txs.OperationTx) error { _ = "STUB: not implemented"; return nil }

func (e *Executor) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (e *Executor) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }
