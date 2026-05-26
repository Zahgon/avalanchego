// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

const nullStr = "null"

// JSONByteSlice represents [[]byte] that is json marshalled to hex
type JSONByteSlice []byte

func (b JSONByteSlice) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *JSONByteSlice) UnmarshalJSON(jsonBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}
