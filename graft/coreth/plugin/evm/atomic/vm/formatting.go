// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

// ParseServiceAddress get address ID from address string, being it either localized (using address manager,
// doing also components validations), or not localized.
// If both attempts fail, reports error from localized address parsing
func ParseServiceAddress(ctx *snow.Context, addrStr string) (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}

// ParseLocalAddress takes in an address for this chain and produces the ID
func ParseLocalAddress(ctx *snow.Context, addrStr string) (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}

// FormatLocalAddress takes in a raw address and produces the formatted address
func FormatLocalAddress(ctx *snow.Context, addr ids.ShortID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseAddress takes in an address and produces the ID of the chain it's for
// the ID of the address
func ParseAddress(ctx *snow.Context, addrStr string) (ids.ID, ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), *new(ids.ShortID), nil
}
