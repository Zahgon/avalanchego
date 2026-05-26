// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/spf13/pflag"

	xsgenesis "github.com/ava-labs/avalanchego/vms/example/xsvm/genesis"
)

const (
	TimeKey     = "time"
	AddressKey  = "address"
	BalanceKey  = "balance"
	EncodingKey = "encoding"

	binaryEncoding = "binary"
	hexEncoding    = "hex"
)

func AddFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Config struct {
	Genesis  *xsgenesis.Genesis
	Encoding string
}

func ParseFlags(flags *pflag.FlagSet, args []string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
