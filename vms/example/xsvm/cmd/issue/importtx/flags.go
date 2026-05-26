// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package importtx

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const (
	URIKey                = "uri"
	SourceURIsKey         = "source-uris"
	SourceChainIDKey      = "source-chain-id"
	DestinationChainIDKey = "destination-chain-id"
	TxIDKey               = "tx-id"
	MaxFeeKey             = "max-fee"
	PrivateKeyKey         = "private-key"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	URI                string
	SourceURIs         []string
	SourceChainID      string
	DestinationChainID string
	TxID               ids.ID
	MaxFee             uint64
	PrivateKey         *secp256k1.PrivateKey
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
