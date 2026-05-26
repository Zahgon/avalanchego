// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
)

var errTxExecutionFailed = errors.New("transaction accepted but failed to execute entirely")

type Wallet struct {
	privKey *ecdsa.PrivateKey
	nonce   uint64
	chainID *big.Int
	signer  types.Signer
	client  *ethclient.Client
	metrics metrics
}

func newWallet(
	privKey *ecdsa.PrivateKey,
	nonce uint64,
	chainID *big.Int,
	client *ethclient.Client,
	metrics metrics,
) *Wallet {
	_ = "STUB: not implemented"
	return nil
}

func (w *Wallet) SendTx(
	ctx context.Context,
	tx *types.Transaction,
) error {
	_ = "STUB: not implemented"
	// start listening for blocks
	return nil
}

// wait for err chan to close before safely closing headers

func (w Wallet) awaitTx(
	ctx context.Context,
	headers chan *types.Header,
	errs <-chan error,
	txHash common.Hash,
) error {
	_ = "STUB: not implemented"
	return nil
}
