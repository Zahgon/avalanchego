// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var (
	errNoChangeAddress   = errors.New("no possible change address")
	errInsufficientFunds = errors.New("insufficient funds")

	fxIndexToID = map[uint32]ids.ID{
		SECP256K1FxIndex: secp256k1fx.ID,
		NFTFxIndex:       nftfx.ID,
		PropertyFxIndex:  propertyfx.ID,
	}

	_ Builder = (*builder)(nil)
)

// Builder provides a convenient interface for building unsigned X-chain
// transactions.
type Builder interface {
	// Context returns the configuration of the chain that this builder uses to
	// create transactions.
	Context() *Context

	// GetFTBalance calculates the amount of each fungible asset that this
	// builder has control over.
	GetFTBalance(
		options ...common.Option,
	) (map[ids.ID]uint64, error)

	// GetImportableBalance calculates the amount of each fungible asset that
	// this builder could import from the provided chain.
	//
	// - [chainID] specifies the chain the funds are from.
	GetImportableBalance(
		chainID ids.ID,
		options ...common.Option,
	) (map[ids.ID]uint64, error)

	// NewBaseTx creates a new simple value transfer.
	//
	// - [outputs] specifies all the recipients and amounts that should be sent
	//   from this transaction.
	NewBaseTx(
		outputs []*avax.TransferableOutput,
		options ...common.Option,
	) (*txs.BaseTx, error)

	// NewCreateAssetTx creates a new asset.
	//
	// - [name] specifies a human readable name for this asset.
	// - [symbol] specifies a human readable abbreviation for this asset.
	// - [denomination] specifies how many times the asset can be split. For
	//   example, a denomination of [4] would mean that the smallest unit of the
	//   asset would be 0.001 units.
	// - [initialState] specifies the supported feature extensions for this
	//   asset as well as the initial outputs for the asset.
	NewCreateAssetTx(
		name string,
		symbol string,
		denomination byte,
		initialState map[uint32][]verify.State,
		options ...common.Option,
	) (*txs.CreateAssetTx, error)

	// NewOperationTx performs state changes on the UTXO set. These state
	// changes may be more complex than simple value transfers.
	//
	// - [operations] specifies the state changes to perform.
	NewOperationTx(
		operations []*txs.Operation,
		options ...common.Option,
	) (*txs.OperationTx, error)

	// NewOperationTxMintFT performs a set of state changes that mint new tokens
	// for the requested assets.
	//
	// - [outputs] maps the assetID to the output that should be created for the
	//   asset.
	NewOperationTxMintFT(
		outputs map[ids.ID]*secp256k1fx.TransferOutput,
		options ...common.Option,
	) (*txs.OperationTx, error)

	// NewOperationTxMintNFT performs a state change that mints new NFTs for the
	// requested asset.
	//
	// - [assetID] specifies the asset to mint the NFTs under.
	// - [payload] specifies the payload to provide each new NFT.
	// - [owners] specifies the new owners of each NFT.
	NewOperationTxMintNFT(
		assetID ids.ID,
		payload []byte,
		owners []*secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.OperationTx, error)

	// NewOperationTxMintProperty performs a state change that mints a new
	// property for the requested asset.
	//
	// - [assetID] specifies the asset to mint the property under.
	// - [owner] specifies the new owner of the property.
	NewOperationTxMintProperty(
		assetID ids.ID,
		owner *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.OperationTx, error)

	// NewOperationTxBurnProperty performs state changes that burns all the
	// properties of the requested asset.
	//
	// - [assetID] specifies the asset to burn the property of.
	NewOperationTxBurnProperty(
		assetID ids.ID,
		options ...common.Option,
	) (*txs.OperationTx, error)

	// NewImportTx creates an import transaction that attempts to consume all
	// the available UTXOs and import the funds to [to].
	//
	// - [chainID] specifies the chain to be importing funds from.
	// - [to] specifies where to send the imported funds to.
	NewImportTx(
		chainID ids.ID,
		to *secp256k1fx.OutputOwners,
		options ...common.Option,
	) (*txs.ImportTx, error)

	// NewExportTx creates an export transaction that attempts to send all the
	// provided [outputs] to the requested [chainID].
	//
	// - [chainID] specifies the chain to be exporting the funds to.
	// - [outputs] specifies the outputs to send to the [chainID].
	NewExportTx(
		chainID ids.ID,
		outputs []*avax.TransferableOutput,
		options ...common.Option,
	) (*txs.ExportTx, error)
}

type Backend interface {
	UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error)
}

type builder struct {
	addrs   set.Set[ids.ShortID]
	context *Context
	backend Backend
}

// New returns a new transaction builder.
//
//   - [addrs] is the set of addresses that the builder assumes can be used when
//     signing the transactions in the future.
//   - [context] provides the chain's configuration.
//   - [backend] provides the chain's state.
func New(
	addrs set.Set[ids.ShortID],
	context *Context,
	backend Backend,
) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) Context() *Context { _ = "STUB: not implemented"; return nil }

func (b *builder) GetFTBalance(
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) GetImportableBalance(
	chainID ids.ID,
	options ...common.Option,
) (map[ids.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewBaseTx(
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.BaseTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort the outputs

func (b *builder) NewCreateAssetTx(
	name string,
	symbol string,
	denomination byte,
	initialState map[uint32][]verify.State,
	options ...common.Option,
) (*txs.CreateAssetTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort the outputs

// sort the initial states

func (b *builder) NewOperationTx(
	operations []*txs.Operation,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewOperationTxMintFT(
	outputs map[ids.ID]*secp256k1fx.TransferOutput,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewOperationTxMintNFT(
	assetID ids.ID,
	payload []byte,
	owners []*secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewOperationTxMintProperty(
	assetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewOperationTxBurnProperty(
	assetID ids.ID,
	options ...common.Option,
) (*txs.OperationTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) NewImportTx(
	chainID ids.ID,
	to *secp256k1fx.OutputOwners,
	options ...common.Option,
) (*txs.ImportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate over the unlocked UTXOs

// Can't import an unknown transfer output type

// We couldn't spend this UTXO, so we skip to the next one

// sort imported inputs

// imported amount goes toward paying tx fee

func (b *builder) NewExportTx(
	chainID ids.ID,
	outputs []*avax.TransferableOutput,
	options ...common.Option,
) (*txs.ExportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) getBalance(
	chainID ids.ID,
	options *common.Options,
) (
	balance map[ids.ID]uint64,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate over the UTXOs

// We only support [secp256k1fx.TransferOutput]s.

// We couldn't spend this UTXO, so we skip to the next one

func (b *builder) spend(
	amountsToBurn map[ids.ID]uint64,
	options *common.Options,
) (
	inputs []*avax.TransferableInput,
	outputs []*avax.TransferableOutput,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Iterate over the UTXOs

// If we have consumed enough of the asset, then we have no need burn
// more.

// We only support burning [secp256k1fx.TransferOutput]s.

// We couldn't spend this UTXO, so we skip to the next one

// Burn any value that should be burned

// Amount we still need to burn
// Amount available to burn

// This input had extra value, so some of it must be returned

// sort inputs
// sort the change outputs

func (b *builder) mintFTs(
	outputs map[ids.ID]*secp256k1fx.TransferOutput,
	options *common.Options,
) (
	operations []*txs.Operation,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add the operation to the array

// remove the asset from the required outputs to mint

// TODO: make this able to generate multiple NFT groups
func (b *builder) mintNFTs(
	assetID ids.ID,
	payload []byte,
	owners []*secp256k1fx.OutputOwners,
	options *common.Options,
) (
	operations []*txs.Operation,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// wrong output type

// add the operation to the array

func (b *builder) mintProperty(
	assetID ids.ID,
	owner *secp256k1fx.OutputOwners,
	options *common.Options,
) (
	operations []*txs.Operation,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// wrong output type

// add the operation to the array

func (b *builder) burnProperty(
	assetID ids.ID,
	options *common.Options,
) (
	operations []*txs.Operation,
	err error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// wrong output type

// add the operation to the array

func (b *builder) initCtx(tx txs.UnsignedTx) error { _ = "STUB: not implemented"; return nil }
