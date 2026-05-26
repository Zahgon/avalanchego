// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package status

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/tx"
)

type TxIssuance struct {
	Tx        *tx.Tx
	TxID      ids.ID
	Nonce     uint64
	StartTime time.Time
}

func (s *TxIssuance) String() string { _ = "STUB: not implemented"; return "" }
