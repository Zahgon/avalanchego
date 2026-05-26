// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package address

import (
	"errors"
)

const addressSep = "-"

var (
	ErrNoSeparator = errors.New("no separator found in address")
	errBits5To8    = errors.New("unable to convert address from 5-bit to 8-bit formatting")
	errBits8To5    = errors.New("unable to convert address from 8-bit to 5-bit formatting")
)

// Parse takes in an address string and splits returns the corresponding parts.
// This returns the chain ID alias, bech32 HRP, address bytes, and an error if
// it occurs.
func Parse(addrStr string) (string, string, []byte, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

// Format takes in a chain prefix, HRP, and byte slice to produce a string for
// an address.
func Format(chainIDAlias string, hrp string, addr []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseBech32 takes a bech32 address as input and returns the HRP and data
// section of a bech32 address
func ParseBech32(addrStr string) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// FormatBech32 takes an address's bytes as input and returns a bech32 address
func FormatBech32(hrp string, payload []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
