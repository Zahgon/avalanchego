// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/client"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	ethcommon "github.com/ava-labs/libevm/common"
)

var _ Wallet = (*wallet)(nil)

type Wallet interface {
	// Builder returns the builder that will be used to create the transactions.
	Builder() Builder

	// Signer returns the signer that will be used to sign the transactions.
	Signer() Signer

	// IssueImportTx creates, signs, and issues an import transaction that
	// attempts to consume all the available UTXOs and import the funds to [to].
	//
	// - [chainID] specifies the chain to be importing funds from.
	// - [to] specifies where to send the imported funds to.
	IssueImportTx(
		chainID ids.ID,
		to ethcommon.Address,
		options ...common.Option,
	) (*atomic.Tx, error)

	// IssueExportTx creates, signs, and issues an export transaction that
	// attempts to send all the provided [outputs] to the requested [chainID].
	//
	// - [chainID] specifies the chain to be exporting the funds to.
	// - [outputs] specifies the outputs to send to the [chainID].
	IssueExportTx(
		chainID ids.ID,
		outputs []*secp256k1fx.TransferOutput,
		options ...common.Option,
	) (*atomic.Tx, error)

	// IssueUnsignedAtomicTx signs and issues the unsigned tx.
	IssueUnsignedAtomicTx(
		utx atomic.UnsignedAtomicTx,
		options ...common.Option,
	) (*atomic.Tx, error)

	// IssueAtomicTx issues the signed tx.
	IssueAtomicTx(
		tx *atomic.Tx,
		options ...common.Option,
	) error
}

func NewWallet(
	builder Builder,
	signer Signer,
	avaxClient *client.Client,
	ethClient *ethclient.Client,
	backend Backend,
) Wallet {
	_ = "STUB: not implemented"
	return *new(Wallet)
}

type wallet struct {
	Backend
	builder    Builder
	signer     Signer
	avaxClient *client.Client
	ethClient  *ethclient.Client
}

func (w *wallet) Builder() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (w *wallet) Signer() Signer { _ = "STUB: not implemented"; return *new(Signer) }

func (w *wallet) IssueImportTx(
	chainID ids.ID,
	to ethcommon.Address,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *wallet) IssueExportTx(
	chainID ids.ID,
	outputs []*secp256k1fx.TransferOutput,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *wallet) IssueUnsignedAtomicTx(
	utx atomic.UnsignedAtomicTx,
	options ...common.Option,
) (*atomic.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *wallet) IssueAtomicTx(
	tx *atomic.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *wallet) baseFee(options []common.Option) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
