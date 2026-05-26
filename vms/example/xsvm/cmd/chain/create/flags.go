// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package create

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const (
	URIKey        = "uri"
	SubnetIDKey   = "subnet-id"
	AddressKey    = "address"
	BalanceKey    = "balance"
	NameKey       = "name"
	PrivateKeyKey = "private-key"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	URI        string
	SubnetID   ids.ID
	Address    ids.ShortID
	Balance    uint64
	Name       string
	PrivateKey *secp256k1.PrivateKey
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
