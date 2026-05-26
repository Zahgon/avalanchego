// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package execute

import (
	"github.com/ava-labs/avalanchego/vms/example/xsvm/block"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/tx"
)

var _ tx.Visitor = (*TxExpectsContext)(nil)

func ExpectsContext(blk *block.Stateless) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type TxExpectsContext struct {
	Result bool
}

func (*TxExpectsContext) Transfer(*tx.Transfer) error { _ = "STUB: not implemented"; return nil }

func (*TxExpectsContext) Export(*tx.Export) error { _ = "STUB: not implemented"; return nil }

func (t *TxExpectsContext) Import(*tx.Import) error { _ = "STUB: not implemented"; return nil }
