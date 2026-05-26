// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utxo

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	errSpendOverflow          = errors.New("spent amount overflows uint64")
	errInsufficientFunds      = errors.New("insufficient funds")
	errAddressesCantMintAsset = errors.New("provided addresses don't have the authority to mint the provided asset")
)

type Spender interface {
	// Spend the provided amount while deducting the provided fee.
	// Arguments:
	// - [utxos] contains assets ID and amount to be spend for each assestID
	// - [kc] are the owners of the funds
	// - [amounts] is the amount of funds that are available to be spent for each assetID
	// Returns:
	// - [amountsSpent] the amount of funds that are spent
	// - [inputs] the inputs that should be consumed to fund the outputs
	// - [signers] the proof of ownership of the funds being moved
	Spend(
		utxos []*avax.UTXO,
		kc *secp256k1fx.Keychain,
		amounts map[ids.ID]uint64,
	) (
		map[ids.ID]uint64, // amountsSpent
		[]*avax.TransferableInput, // inputs
		[][]*secp256k1.PrivateKey, // signers
		error,
	)

	SpendNFT(
		utxos []*avax.UTXO,
		kc *secp256k1fx.Keychain,
		assetID ids.ID,
		groupID uint32,
		to ids.ShortID,
	) (
		[]*txs.Operation,
		[][]*secp256k1.PrivateKey,
		error,
	)

	SpendAll(
		utxos []*avax.UTXO,
		kc *secp256k1fx.Keychain,
	) (
		map[ids.ID]uint64,
		[]*avax.TransferableInput,
		[][]*secp256k1.PrivateKey,
		error,
	)

	Mint(
		utxos []*avax.UTXO,
		kc *secp256k1fx.Keychain,
		amounts map[ids.ID]uint64,
		to ids.ShortID,
	) (
		[]*txs.Operation,
		[][]*secp256k1.PrivateKey,
		error,
	)

	MintNFT(
		utxos []*avax.UTXO,
		kc *secp256k1fx.Keychain,
		assetID ids.ID,
		payload []byte,
		to ids.ShortID,
	) (
		[]*txs.Operation,
		[][]*secp256k1.PrivateKey,
		error,
	)
}

func NewSpender(
	clk *mockable.Clock,
	codec codec.Manager,
) Spender {
	_ = "STUB: not implemented"
	return *new(Spender)
}

type spender struct {
	clock *mockable.Clock
	codec codec.Manager
}

func (s *spender) Spend(
	utxos []*avax.UTXO,
	kc *secp256k1fx.Keychain,
	amounts map[ids.ID]uint64,
) (
	map[ids.ID]uint64, // amountsSpent
	[]*avax.TransferableInput, // inputs
	[][]*secp256k1.PrivateKey, // signers
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// we already have enough inputs allocated to this asset

// this utxo can't be spent with the current keys right now

// this input doesn't have an amount, so I don't care about it here

// there was an error calculating the consumed amount, just error

// add the new input to the array

// add the required keys to the array

func (s *spender) SpendNFT(
	utxos []*avax.UTXO,
	kc *secp256k1fx.Keychain,
	assetID ids.ID,
	groupID uint32,
	to ids.ShortID,
) (
	[]*txs.Operation,
	[][]*secp256k1.PrivateKey,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we have already been able to create the operation needed

// wrong asset ID

// wrong output type

// wrong group id

// unable to spend the output

// add the new operation to the array

// add the required keys to the array

func (s *spender) SpendAll(
	utxos []*avax.UTXO,
	kc *secp256k1fx.Keychain,
) (
	map[ids.ID]uint64,
	[]*avax.TransferableInput,
	[][]*secp256k1.PrivateKey,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// this utxo can't be spent with the current keys right now

// this input doesn't have an amount, so I don't care about it here

// there was an error calculating the consumed amount, just error

// add the new input to the array

// add the required keys to the array

func (s *spender) Mint(
	utxos []*avax.UTXO,
	kc *secp256k1fx.Keychain,
	amounts map[ids.ID]uint64,
	to ids.ShortID,
) (
	[]*txs.Operation,
	[][]*secp256k1.PrivateKey,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// add the operation to the array

// add the required keys to the array

// remove the asset from the required amounts to mint

func (s *spender) MintNFT(
	utxos []*avax.UTXO,
	kc *secp256k1fx.Keychain,
	assetID ids.ID,
	payload []byte,
	to ids.ShortID,
) (
	[]*txs.Operation,
	[][]*secp256k1.PrivateKey,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we have already been able to create the operation needed

// wrong asset id

// wrong output type

// unable to spend the output

// add the operation to the array

// add the required keys to the array
