// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mempool

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/txs/mempool"
)

func New(
	namespace string,
	registerer prometheus.Registerer,
) (mempool.Mempool[*txs.Tx], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
