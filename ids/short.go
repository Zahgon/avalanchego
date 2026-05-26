// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"github.com/ava-labs/avalanchego/utils"
)

const ShortIDLen = 20

// ShortEmpty is a useful all zero value
var (
	ShortEmpty = ShortID{}

	_ utils.Sortable[ShortID] = ShortID{}
)

// ShortID wraps a 20 byte hash as an identifier
type ShortID [ShortIDLen]byte

// ToShortID attempt to convert a byte slice into an id
func ToShortID(bytes []byte) (ShortID, error) { _ = "STUB: not implemented"; return *new(ShortID), nil }

// ShortFromString is the inverse of ShortID.String()
func ShortFromString(idStr string) (ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ShortID), nil
}

// ShortFromStringOrPanic is the same as ShortFromString, but will panic on
// error.
func ShortFromStringOrPanic(idStr string) ShortID { _ = "STUB: not implemented"; return *new(ShortID) }

// ShortFromPrefixedString returns a ShortID assuming the cb58 format is
// prefixed
func ShortFromPrefixedString(idStr, prefix string) (ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ShortID), nil
}

func (id ShortID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ShortID) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If "null", do nothing

// Parse CB58 formatted string to bytes

func (id *ShortID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// Bytes returns the 20 byte hash as a slice. It is assumed this slice is not
// modified.
func (id ShortID) Bytes() []byte {
	_ = "STUB: not implemented"

	// Hex returns a hex encoded string of this id.
	return nil
}

func (id ShortID) Hex() string { _ = "STUB: not implemented"; return "" }

func (id ShortID) String() string {
	_ = "STUB: not implemented"
	// We assume that the maximum size of a byte slice that
	// can be stringified is at least the length of an ID
	return ""
}

// PrefixedString returns the String representation with a prefix added
func (id ShortID) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (id ShortID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id ShortID) Compare(other ShortID) int { _ = "STUB: not implemented"; return 0 }

// ShortIDsToStrings converts an array of shortIDs to an array of their string
// representations
func ShortIDsToStrings(ids []ShortID) []string { _ = "STUB: not implemented"; return nil }
