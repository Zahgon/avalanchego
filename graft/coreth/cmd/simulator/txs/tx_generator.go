// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
)

var _ TxSequence[*types.Transaction] = (*txSequence)(nil)

type CreateTx func(key *ecdsa.PrivateKey, nonce uint64) (*types.Transaction, error)

func GenerateTxSequence(ctx context.Context, generator CreateTx, client *ethclient.Client, key *ecdsa.PrivateKey, numTxs uint64, async bool) (TxSequence[*types.Transaction], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateTxSequences(ctx context.Context, generator CreateTx, client *ethclient.Client, keys []*ecdsa.PrivateKey, txsPerKey uint64, async bool) ([]TxSequence[*types.Transaction], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addTxs(ctx context.Context, txSequence *txSequence, generator CreateTx, client *ethclient.Client, key *ecdsa.PrivateKey, numTxs uint64) error {
	_ = "STUB: not implemented"
	return nil
}

type txSequence struct {
	txChan chan *types.Transaction
}

func ConvertTxSliceToSequence(txs []*types.Transaction) TxSequence[*types.Transaction] {
	_ = "STUB: not implemented"
	return nil
}

func (t *txSequence) Chan() <-chan *types.Transaction { _ = "STUB: not implemented"; return nil }
