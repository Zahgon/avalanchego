// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package formatting

import (
	"errors"
)

const (
	hexPrefix   = "0x"
	checksumLen = 4
)

var (
	errEncodingOverFlow            = errors.New("encoding overflow")
	errInvalidEncoding             = errors.New("invalid encoding")
	errUnsupportedEncodingInMethod = errors.New("unsupported encoding in method")
	errMissingChecksum             = errors.New("input string is smaller than the checksum size")
	errBadChecksum                 = errors.New("invalid input checksum")
	errMissingHexPrefix            = errors.New("missing 0x prefix to hex encoding")
)

// Encoding defines how bytes are converted to a string and vice versa
type Encoding uint8

const (
	// Hex specifies a hex plus 4 byte checksum encoding format
	Hex Encoding = iota
	// HexNC specifies a hex encoding format
	HexNC
	// HexC specifies a hex plus 4 byte checksum encoding format
	HexC
	// JSON specifies the JSON encoding format
	JSON
)

func (enc Encoding) String() string { _ = "STUB: not implemented"; return "" }

func (enc Encoding) valid() bool { _ = "STUB: not implemented"; return false }

func (enc Encoding) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (enc *Encoding) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Encode [bytes] to a string using the given encoding format [bytes] may be
// nil, in which case it will be treated the same as an empty slice.
func Encode(encoding Encoding, bytes []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// JSON Marshal does not support []byte input and we rely on the
// router's json marshalling to marshal our interface{} into JSON
// in response. Therefore it is not supported in this call.

// Decode [str] to bytes using the given encoding
// If [str] is the empty string, returns a nil byte slice and nil error
func Decode(encoding Encoding, str string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: remove the empty string check and enforce the correct format.

// JSON unmarshalling requires interface and has no return values
// contrary to this method, therefore it is not supported in this call

// Verify the checksum
