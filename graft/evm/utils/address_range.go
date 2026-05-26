// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/ava-labs/libevm/common"
)

// AddressRange represents a continuous range of addresses
type AddressRange struct {
	Start common.Address
	End   common.Address
}

// Contains returns true iff [addr] is contained within the (inclusive)
// range of addresses defined by [a].
func (a *AddressRange) Contains(addr common.Address) bool { _ = "STUB: not implemented"; return false }
