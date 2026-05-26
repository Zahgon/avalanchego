// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	ethcommon "github.com/ava-labs/libevm/common"
)

var _ Wallet = (*walletWithOptions)(nil)

func NewWalletWithOptions(
	wallet Wallet,
	options ...common.Option,
) Wallet {
	_ = "STUB: not implemented"
	return *new(Wallet)
}

type walletWithOptions struct {
	Wallet
	options []common.Option
}

func (w *walletWithOptions) Builder() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (w *walletWithOptions) IssueImportTx(
	chainID ids.ID,
	to ethcommon.Address,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueExportTx(
	chainID ids.ID,
	outputs []*secp256k1fx.TransferOutput,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueUnsignedAtomicTx(
	utx atomic.UnsignedAtomicTx,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueAtomicTx(
	tx *atomic.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
