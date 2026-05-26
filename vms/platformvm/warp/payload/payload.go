// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package payload

import (
	"errors"
)

var ErrWrongType = errors.New("wrong payload type")

// Payload provides a common interface for all payloads implemented by this
// package.
type Payload interface {
	// Bytes returns the binary representation of this payload.
	Bytes() []byte

	// initialize the payload with the provided binary representation.
	initialize(b []byte)
}

func Parse(bytes []byte) (Payload, error) { _ = "STUB: not implemented"; return *new(Payload), nil }

func initialize(p Payload) error { _ = "STUB: not implemented"; return nil }
