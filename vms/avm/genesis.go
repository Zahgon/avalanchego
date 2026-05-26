// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	_ utils.Sortable[*GenesisAsset] = (*GenesisAsset)(nil)

	_ avax.TransferableIn  = (*secp256k1fx.TransferInput)(nil)
	_ verify.State         = (*secp256k1fx.MintOutput)(nil)
	_ avax.TransferableOut = (*secp256k1fx.TransferOutput)(nil)
	_ fxs.FxOperation      = (*secp256k1fx.MintOperation)(nil)
	_ verify.Verifiable    = (*secp256k1fx.Credential)(nil)

	_ verify.State      = (*nftfx.MintOutput)(nil)
	_ verify.State      = (*nftfx.TransferOutput)(nil)
	_ fxs.FxOperation   = (*nftfx.MintOperation)(nil)
	_ fxs.FxOperation   = (*nftfx.TransferOperation)(nil)
	_ verify.Verifiable = (*nftfx.Credential)(nil)

	_ verify.State      = (*propertyfx.MintOutput)(nil)
	_ verify.State      = (*propertyfx.OwnedOutput)(nil)
	_ fxs.FxOperation   = (*propertyfx.MintOperation)(nil)
	_ fxs.FxOperation   = (*propertyfx.BurnOperation)(nil)
	_ verify.Verifiable = (*propertyfx.Credential)(nil)
)

type Genesis struct {
	Txs []*GenesisAsset `serialize:"true"`
}

type GenesisAsset struct {
	Alias             string `serialize:"true"`
	txs.CreateAssetTx `serialize:"true"`
}

func (g *GenesisAsset) Compare(other *GenesisAsset) int { _ = "STUB: not implemented"; return 0 }

// AssetInitialState describes the initial state of an asset
type AssetInitialState struct {
	FixedCap    []Holder
	VariableCap []Owners
}

// AssetDefinition describes a genesis asset and its initial state
type AssetDefinition struct {
	Name         string
	Symbol       string
	Denomination uint8
	InitialState AssetInitialState
	Memo         []byte
}

// Holder describes how much asset is owned by an address
type Holder struct {
	Amount  uint64
	Address string
}

// Owners describes who can perform an action
type Owners struct {
	Threshold uint32
	Minters   []string
}

// NewGenesis creates a new Genesis from genesis data
func NewGenesis(
	networkID uint32,
	genesisData map[string]AssetDefinition,
) (*Genesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Should lookup secp256k1fx FxID

// Bytes serializes the Genesis to bytes using the AVM genesis codec
func (g *Genesis) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func newGenesisCodec() (codec.Manager, error) {
	_ = "STUB: not implemented"
	return *new(codec.Manager), nil
}
