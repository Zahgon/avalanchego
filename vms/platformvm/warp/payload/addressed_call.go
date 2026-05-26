// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package payload

var _ Payload = (*AddressedCall)(nil)

// AddressedCall defines the format for delivering a call across VMs including a
// source address and a payload.
//
// Note: If a destination address is expected, it should be encoded in the
// payload.
type AddressedCall struct {
	SourceAddress []byte `serialize:"true"`
	Payload       []byte `serialize:"true"`

	bytes []byte
}

// NewAddressedCall creates a new *AddressedCall and initializes it.
func NewAddressedCall(sourceAddress []byte, payload []byte) (*AddressedCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseAddressedCall converts a slice of bytes into an initialized
// AddressedCall.
func ParseAddressedCall(b []byte) (*AddressedCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bytes returns the binary representation of this payload. It assumes that the
// payload is initialized from either NewAddressedCall or Parse.
func (a *AddressedCall) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (a *AddressedCall) initialize(bytes []byte) { _ = "STUB: not implemented"; return }
