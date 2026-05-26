// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"github.com/ava-labs/avalanchego/ids"
)

// UnsignedMessage defines the standard format for an unsigned Warp message.
type UnsignedMessage struct {
	NetworkID     uint32 `serialize:"true"`
	SourceChainID ids.ID `serialize:"true"`
	Payload       []byte `serialize:"true"`

	bytes []byte
	id    ids.ID
}

// NewUnsignedMessage creates a new *UnsignedMessage and initializes it.
func NewUnsignedMessage(
	networkID uint32,
	sourceChainID ids.ID,
	payload []byte,
) (*UnsignedMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseUnsignedMessage converts a slice of bytes into an initialized
// *UnsignedMessage.
func ParseUnsignedMessage(b []byte) (*UnsignedMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize recalculates the result of Bytes().
func (m *UnsignedMessage) Initialize() error { _ = "STUB: not implemented"; return nil }

// Bytes returns the binary representation of this message. It assumes that the
// message is initialized from either New, Parse, or an explicit call to
// Initialize.
func (m *UnsignedMessage) Bytes() []byte {
	_ = "STUB: not implemented"

	// ID returns an identifier for this message. It assumes that the
	// message is initialized from either New, Parse, or an explicit call to
	// Initialize.
	return nil
}

func (m *UnsignedMessage) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (m *UnsignedMessage) String() string { _ = "STUB: not implemented"; return "" }
