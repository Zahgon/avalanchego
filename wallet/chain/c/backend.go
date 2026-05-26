// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	ethcommon "github.com/ava-labs/libevm/common"
)

var (
	_ Backend = (*backend)(nil)

	errUnknownTxType = errors.New("unknown tx type")
)

// Backend defines the full interface required to support a C-chain wallet.
type Backend interface {
	common.ChainUTXOs
	BuilderBackend
	SignerBackend

	AcceptAtomicTx(ctx context.Context, tx *atomic.Tx) error
}

type backend struct {
	common.ChainUTXOs

	accountsLock sync.RWMutex
	accounts     map[ethcommon.Address]*Account
}

type Account struct {
	Balance *big.Int
	Nonce   uint64
}

func NewBackend(
	utxos common.ChainUTXOs,
	accounts map[ethcommon.Address]*Account,
) Backend {
	_ = "STUB: not implemented"
	return *new(Backend)
}

func (b *backend) AcceptAtomicTx(ctx context.Context, tx *atomic.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backend) Balance(_ context.Context, addr ethcommon.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) Nonce(_ context.Context, addr ethcommon.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
