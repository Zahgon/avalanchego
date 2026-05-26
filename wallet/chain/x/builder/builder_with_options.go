// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ Builder = (*builderWithOptions)(nil)

type builderWithOptions struct {
	builder Builder
	options []common.Option
}

// NewWithOptions returns a new transaction builder that will use the given
// options by default.
//
//   - [builder] is the builder that will be called to perform the underlying
//     operations.
//   - [options] will be provided to the builder in addition to the options
//     provided in the method calls.
func NewWithOptions(builder Builder, options ...common.Option) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builderWithOptions) Context() *Context { _ = "STUB: not implemented"; return nil }

func (b *builderWithOptions) GetFTBalance(
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) GetImportableBalance(
	chainID ids.ID,
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.BaseTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewCreateAssetTx(
	name string,
	symbol string,
	denomination byte,
	initialState map[uint32][]verify.State,
	options ...common.Option,
) (*txs.CreateAssetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewOperationTx(
	operations []*txs.Operation,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewOperationTxMintFT(
	outputs map[ids.ID]*secp256k1fx.TransferOutput,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewOperationTxMintNFT(
	assetID ids.ID,
	payload []byte,
	owners []*secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewOperationTxMintProperty(
	assetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewOperationTxBurnProperty(
	assetID ids.ID,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewImportTx(
	chainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.ImportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builderWithOptions) NewExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.ExportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
