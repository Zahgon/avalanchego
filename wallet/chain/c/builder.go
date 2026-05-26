// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"

	ethcommon "github.com/ava-labs/libevm/common"
)

const avaxConversionRateInt = 1_000_000_000

var (
	_ Builder = (*builder)(nil)

	errInsufficientFunds = errors.New("insufficient funds")

	// avaxConversionRate is the conversion rate between the smallest
	// denomination on the X-Chain and P-chain, 1 nAVAX, and the smallest
	// denomination on the C-Chain 1 wei. Where 1 nAVAX = 1 gWei.
	//
	// This is only required for AVAX because the denomination of 1 AVAX is 9
	// decimal places on the X and P chains, but is 18 decimal places within the
	// EVM.
	avaxConversionRate = big.NewInt(avaxConversionRateInt)
)

// Builder provides a convenient interface for building unsigned C-chain
// transactions.
type Builder interface {
	// Context returns the configuration of the chain that this builder uses to
	// create transactions.
	Context() *Context

	// GetBalance calculates the amount of AVAX that this builder has control
	// over.
	GetBalance(
		options ...common.Option,
	) (*big.Int, error)

	// GetImportableBalance calculates the amount of AVAX that this builder
	// could import from the provided chain.
	//
	// - [chainID] specifies the chain the funds are from.
	GetImportableBalance(
		chainID ids.ID,
		options ...common.Option,
	) (uint64, error)

	// NewImportTx creates an import transaction that attempts to consume all
	// the available UTXOs and import the funds to [to].
	//
	// - [chainID] specifies the chain to be importing funds from.
	// - [to] specifies where to send the imported funds to.
	// - [baseFee] specifies the fee price willing to be paid by this tx.
	NewImportTx(
		chainID ids.ID,
		to ethcommon.Address,
		baseFee *big.Int,
		options ...common.Option,
	) (*atomic.UnsignedImportTx, error)

	// NewExportTx creates an export transaction that attempts to send all the
	// provided [outputs] to the requested [chainID].
	//
	// - [chainID] specifies the chain to be exporting the funds to.
	// - [outputs] specifies the outputs to send to the [chainID].
	// - [baseFee] specifies the fee price willing to be paid by this tx.
	NewExportTx(
		chainID ids.ID,
		outputs []*secp256k1fx.TransferOutput,
		baseFee *big.Int,
		options ...common.Option,
	) (*atomic.UnsignedExportTx, error)
}

// BuilderBackend specifies the required information needed to build unsigned
// C-chain transactions.
type BuilderBackend interface {
	UTXOs(ctx context.Context, sourceChainID ids.ID) ([]*avax.UTXO, error)
	Balance(ctx context.Context, addr ethcommon.Address) (*big.Int, error)
	Nonce(ctx context.Context, addr ethcommon.Address) (uint64, error)
}

type builder struct {
	avaxAddrs set.Set[ids.ShortID]
	ethAddrs  set.Set[ethcommon.Address]
	context   *Context
	backend   BuilderBackend
}

// NewBuilder returns a new transaction builder.
//
//   - [avaxAddrs] is the set of addresses in the AVAX format that the builder
//     assumes can be used when signing the transactions in the future.
//   - [ethAddrs] is the set of addresses in the Eth format that the builder
//     assumes can be used when signing the transactions in the future.
//   - [backend] provides the required access to the chain's context and state
//     to build out the transactions.
func NewBuilder(
	avaxAddrs set.Set[ids.ShortID],
	ethAddrs set.Set[ethcommon.Address],
	context *Context,
	backend BuilderBackend,
) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) Context() *Context { _ = "STUB: not implemented"; return nil }

func (b *builder) GetBalance(
	options ...common.Option,
) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *builder) GetImportableBalance(
	chainID ids.ID,
	options ...common.Option,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *builder) NewImportTx(
	chainID ids.ID,
	to ethcommon.Address,
	baseFee *big.Int,
	options ...common.Option,
) (*atomic.UnsignedImportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We must initialize the bytes of the tx to calculate the initial cost

/*=IsApricotPhase5*/

func (b *builder) NewExportTx(
	chainID ids.ID,
	outputs []*secp256k1fx.TransferOutput,
	baseFee *big.Int,
	options ...common.Option,
) (*atomic.UnsignedExportTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We must initialize the bytes of the tx to calculate the initial cost

/*=IsApricotPhase5*/

// Since the asset is AVAX, we divide by the avaxConversionRate to
// convert back to the correct denomination of AVAX that can be
// exported.

// If the balance for [addr] is insufficient to cover the additional
// cost of adding an input to the transaction, skip adding the input
// altogether.

// Update the cost for the next iteration

func getSpendableAmount(
	utxo *avax.UTXO,
	addrs set.Set[ids.ShortID],
	minIssuanceTime uint64,
	avaxAssetID ids.ID,
) (uint64, []uint32, bool) {
	_ = "STUB: not implemented"
	return 0, nil, false

	// Only AVAX can be imported
}

// Can't import an unknown transfer output type
