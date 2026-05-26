// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package transfer

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const (
	URIKey               = "uri"
	ChainIDKey           = "chain-id"
	MaxFeeKey            = "max-fee"
	AssetIDKey           = "asset-id"
	AmountKey            = "amount"
	ToKey                = "to"
	PrivateKeyKey        = "private-key"
	WaitForAcceptanceKey = "wait-for-acceptance"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	URI               string
	ChainID           ids.ID
	MaxFee            uint64
	AssetID           ids.ID
	Amount            uint64
	To                ids.ShortID
	PrivateKey        *secp256k1.PrivateKey
	WaitForAcceptance bool
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
