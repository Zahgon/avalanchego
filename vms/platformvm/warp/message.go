// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

// Message defines the standard format for a Warp message.
type Message struct {
	UnsignedMessage `serialize:"true"`
	Signature       Signature `serialize:"true"`

	bytes []byte
}

// NewMessage creates a new *Message and initializes it.
func NewMessage(
	unsignedMsg *UnsignedMessage,
	signature Signature,
) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseMessage converts a slice of bytes into an initialized *Message.
func ParseMessage(b []byte) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

// Initialize recalculates the result of Bytes(). It does not call Initialize()
// on the UnsignedMessage.
func (m *Message) Initialize() error { _ = "STUB: not implemented"; return nil }

// Bytes returns the binary representation of this message. It assumes that the
// message is initialized from either New, Parse, or an explicit call to
// Initialize.
func (m *Message) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }
