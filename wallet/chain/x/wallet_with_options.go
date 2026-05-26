// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/chain/x/builder"
	"github.com/ava-labs/avalanchego/wallet/chain/x/signer"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
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
	wallet  Wallet
	options []common.Option
}

func (w *walletWithOptions) Builder() builder.Builder {
	_ = "STUB: not implemented"
	return *new(builder.Builder)
}

func (w *walletWithOptions) Signer() signer.Signer {
	_ = "STUB: not implemented"
	return *new(signer.Signer)
}

func (w *walletWithOptions) IssueBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueCreateAssetTx(
	name string,
	symbol string,
	denomination byte,
	initialState map[uint32][]verify.State,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueOperationTx(
	operations []*txs.Operation,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueOperationTxMintFT(
	outputs map[ids.ID]*secp256k1fx.TransferOutput,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueOperationTxMintNFT(
	assetID ids.ID,
	payload []byte,
	owners []*secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueOperationTxMintProperty(
	assetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueOperationTxBurnProperty(
	assetID ids.ID,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueImportTx(
	chainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueUnsignedTx(
	utx txs.UnsignedTx,
	options ...common.Option,
) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *walletWithOptions) IssueTx(
	tx *txs.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
