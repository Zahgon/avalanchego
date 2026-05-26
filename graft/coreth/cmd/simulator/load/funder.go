// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/cmd/simulator/key"
	"github.com/ava-labs/avalanchego/graft/coreth/cmd/simulator/metrics"
	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
)

// DistributeFunds ensures that each address in keys has at least [minFundsPerAddr] by sending funds
// from the key with the highest starting balance.
// This function returns a set of at least [numKeys] keys, each having a minimum balance [minFundsPerAddr].
func DistributeFunds(ctx context.Context, client *ethclient.Client, keys []*key.Key, numKeys int, minFundsPerAddr *big.Int, m *metrics.Metrics) ([]*key.Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: clean up fund distribution.

// If there are not enough funded keys, cut [needFundsAddrs] to the number of keys that
// must be funded to reach [numKeys] required.

// Generate a sequence of transactions to distribute the required funds.
