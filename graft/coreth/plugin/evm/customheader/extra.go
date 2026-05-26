// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customheader

import (
	"errors"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/vms/components/gas"
)

var (
	errInvalidExtraPrefix = errors.New("invalid header.Extra prefix")
	errIncorrectFeeState  = errors.New("incorrect fee state")
	errInvalidExtraLength = errors.New("invalid header.Extra length")
)

// ExtraPrefix returns what the prefix of the header's Extra field should be
// based on the desired target excess.
//
// If the `desiredTargetExcess` is nil, the parent's target excess is used.
func ExtraPrefix(
	config *extras.ChainConfig,
	parent *types.Header,
	header *types.Header,
	desiredTargetExcess *gas.Gas,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prior to AP3 there was no expected extra prefix.

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

// By passing in the claimed target excess, we ensure that the expected
// target excess is equal to the claimed target excess if it is possible
// to have correctly set it to that value. Otherwise, the resulting
// value will be as close to the claimed value as possible, but would
// not be equal.

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
func PredicateBytesFromExtra(rules extras.AvalancheRules, extra []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Prior to Durango, the VM enforces the extra data is smaller than or equal
// to `offset`.
// After Durango, the VM pre-verifies the extra data past `offset` is valid.

// SetPredicateBytesInExtra sets the predicate result bytes in the header's extra
// data. If the extra data is not long enough (i.e., an incomplete header.Extra
// as built in the miner), it is padded with zeros.
func SetPredicateBytesInExtra(rules extras.AvalancheRules, extra []byte, predicateBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// pad extra with zeros

// truncate extra to the offset
