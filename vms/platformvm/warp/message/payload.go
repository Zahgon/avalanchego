// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"
)

var ErrWrongType = errors.New("wrong payload type")

// Payload provides a common interface for all payloads implemented by this
// package.
type Payload interface {
	// Bytes returns the binary representation of this payload.
	//
	// If the payload is not initialized, this method will return nil.
	Bytes() []byte

	// initialize the payload with the provided binary representation.
	initialize(b []byte)
}

// payload is embedded by all the payloads to provide the common implementation
// of Payload.
type payload []byte

func (p payload) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (p *payload) initialize(bytes []byte) { _ = "STUB: not implemented"; return }

func Parse(bytes []byte) (Payload, error) { _ = "STUB: not implemented"; return *new(Payload), nil }

func Initialize(p Payload) error { _ = "STUB: not implemented"; return nil }
