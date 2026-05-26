// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
)

const (
	maximumExtraDataSize = 64 // Maximum size extra data may be after Genesis.
)

var (
	errInvalidExtraPrefix = errors.New("invalid header.Extra prefix")
	errInvalidExtraLength = errors.New("invalid header.Extra length")
)

// ExtraPrefix takes the previous header and the timestamp of its child
// block and calculates the expected extra prefix for the child block.
func ExtraPrefix(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prior to SubnetEVM there was no expected extra prefix.

// VerifyExtraPrefix verifies that the header's Extra field is correctly
// formatted.
func VerifyExtraPrefix(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyExtra verifies that the header's Extra field is correctly formatted for
// rules.
//
// TODO: Should this be merged with VerifyExtraPrefix?
func VerifyExtra(rules extras.AvalancheRules, extra []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// PredicateBytesFromExtra returns the predicate result bytes from the header's
// extra data. If the extra data is not long enough, an empty slice is returned.
func PredicateBytesFromExtra(extra []byte) []byte { _ = "STUB: not implemented"; return nil }

// Prior to Durango, the VM enforces the extra data is smaller than or equal
// to `offset`.
// After Durango, the VM pre-verifies the extra data past `offset` is valid.

// SetPredicateBytesInExtra sets the predicate result bytes in the header's extra
// data. If the extra data is not long enough (i.e., an incomplete header.Extra
// as built in the miner), it is padded with zeros.
func SetPredicateBytesInExtra(extra []byte, predicateBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// pad extra with zeros

// truncate extra to the offset
