// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utxotest

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

func NewDeterministicChainUTXOs(t *testing.T, utxoSets map[ids.ID][]*avax.UTXO) *DeterministicChainUTXOs {
	_ = "STUB: not implemented"
	return nil
}

type DeterministicChainUTXOs struct {
	common.ChainUTXOs
}

func (c *DeterministicChainUTXOs) UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
