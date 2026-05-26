// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils"
)

const (
	IDLen   = 32
	nullStr = "null"
)

var (
	// Empty is a useful all zero value
	Empty = ID{}

	errMissingQuotes = errors.New("first and last characters should be quotes")

	_ utils.Sortable[ID] = ID{}
)

// ID wraps a 32 byte hash used as an identifier
type ID [IDLen]byte

// ToID attempt to convert a byte slice into an id
func ToID(bytes []byte) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

// FromString is the inverse of ID.String()
func FromString(idStr string) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

// FromStringOrPanic is the same as FromString, but will panic on error
func FromStringOrPanic(idStr string) ID { _ = "STUB: not implemented"; return *new(ID) }

func (id ID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ID) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If "null", do nothing

// Parse CB58 formatted string to bytes

func (id *ID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// Prefix this id to create a more selective id. This can be used to store
// multiple values under the same key. For example:
// prefix1(id) -> confidence
// prefix2(id) -> vertex
// This will return a new id and not modify the original id.
func (id ID) Prefix(prefixes ...uint64) ID { _ = "STUB: not implemented"; return *new(ID) }

// Append this id with the provided suffixes and re-hash the result. This
// returns a new ID and does not modify the original ID.
//
// This is used to generate ACP-77 validationIDs.
//
// Ref: https://github.com/avalanche-foundation/ACPs/tree/e333b335c34c8692d84259d21bd07b2bb849dc2c/ACPs/77-reinventing-subnets#convertsubnettol1tx
func (id ID) Append(suffixes ...uint32) ID { _ = "STUB: not implemented"; return *new(ID) }

// XOR this id and the provided id and return the resulting id.
//
// Note: this id is not modified.
func (id ID) XOR(other ID) ID { _ = "STUB: not implemented"; return *new(ID) }

// Bit returns the bit value at the ith index of the byte array. Returns 0 or 1
func (id ID) Bit(i uint) int { _ = "STUB: not implemented"; return 0 }

// b = [7, 6, 5, 4, 3, 2, 1, 0]

// b = [0, ..., bitIndex + 1, bitIndex]
// 1 = [0, 0, 0, 0, 0, 0, 0, 1]

// b = [0, 0, 0, 0, 0, 0, 0, bitIndex]

// Hex returns a hex encoded string of this id.
func (id ID) Hex() string { _ = "STUB: not implemented"; return "" }

func (id ID) String() string {
	_ = "STUB: not implemented"
	// We assume that the maximum size of a byte slice that
	// can be stringified is at least the length of an ID
	return ""
}

func (id ID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id ID) Compare(other ID) int { _ = "STUB: not implemented"; return 0 }
