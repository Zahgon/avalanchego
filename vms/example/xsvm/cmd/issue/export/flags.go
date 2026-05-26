// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package export

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const (
	URIKey                = "uri"
	SourceChainIDKey      = "source-chain-id"
	DestinationChainIDKey = "destination-chain-id"
	MaxFeeKey             = "max-fee"
	IsReturnKey           = "is-return"
	AmountKey             = "amount"
	ToKey                 = "to"
	PrivateKeyKey         = "private-key"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	URI                string
	SourceChainID      ids.ID
	DestinationChainID ids.ID
	MaxFee             uint64
	IsReturn           bool
	Amount             uint64
	To                 ids.ShortID
	PrivateKey         *secp256k1.PrivateKey
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
