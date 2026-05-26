// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"github.com/ava-labs/avalanchego/ids"
)

type Metadata struct {
	id            ids.ID // The ID of this data
	unsignedBytes []byte // Unsigned byte representation of this data
	bytes         []byte // Byte representation of this data
}

// Initialize set the bytes and ID
func (md *Metadata) Initialize(unsignedBytes, bytes []byte) { _ = "STUB: not implemented"; return }

// ID returns the unique ID of this data
func (md *Metadata) ID() ids.ID {
	_ = "STUB: not implemented"

	// UnsignedBytes returns the unsigned binary representation of this data
	return *new(ids.ID)
}

func (md *Metadata) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Bytes returns the binary representation of this data
func (md *Metadata) SignedBytes() []byte { _ = "STUB: not implemented"; return nil }
