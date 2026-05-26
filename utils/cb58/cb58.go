// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cb58

import (
	"errors"
)

const checksumLen = 4

var (
	ErrBase58Decoding   = errors.New("base58 decoding error")
	ErrMissingChecksum  = errors.New("input string is smaller than the checksum size")
	ErrBadChecksum      = errors.New("invalid input checksum")
	errEncodingOverFlow = errors.New("encoding overflow")
)

// Encode [bytes] to a string using cb58 format.
// [bytes] may be nil, in which case it will be treated the same as an empty
// slice.
func Encode(bytes []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Decode [str] to bytes from cb58.
func Decode(str string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Verify the checksum
