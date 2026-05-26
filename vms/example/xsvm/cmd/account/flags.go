// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package account

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/ids"
)

const (
	URIKey     = "uri"
	ChainIDKey = "chain-id"
	AddressKey = "address"
	AssetIDKey = "asset-id"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	URI     string
	ChainID string
	Address ids.ShortID
	AssetID ids.ID
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
