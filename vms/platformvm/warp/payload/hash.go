// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package payload

import (
	"github.com/ava-labs/avalanchego/ids"
)

var _ Payload = (*Hash)(nil)

type Hash struct {
	Hash ids.ID `serialize:"true"`

	bytes []byte
}

// NewHash creates a new *Hash and initializes it.
func NewHash(hash ids.ID) (*Hash, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseHash converts a slice of bytes into an initialized Hash.
func ParseHash(b []byte) (*Hash, error) { _ = "STUB: not implemented"; return nil, nil }

// Bytes returns the binary representation of this payload. It assumes that the
// payload is initialized from either NewHash or Parse.
func (b *Hash) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Hash) initialize(bytes []byte) { _ = "STUB: not implemented"; return }
