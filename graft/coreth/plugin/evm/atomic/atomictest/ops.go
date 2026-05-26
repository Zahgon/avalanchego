// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

func ConvertToAtomicOps(tx *atomic.Tx) (map[ids.ID]*avalancheatomic.Requests, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
