// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"github.com/ava-labs/avalanchego/ids"
)

// ValidatorUptime is signed when the ValidationID is known and the validator
// has been up for TotalUptime seconds.
type ValidatorUptime struct {
	ValidationID ids.ID `serialize:"true"`
	TotalUptime  uint64 `serialize:"true"` // in seconds

	bytes []byte
}

// NewValidatorUptime creates a new *ValidatorUptime and initializes it.
func NewValidatorUptime(validationID ids.ID, totalUptime uint64) (*ValidatorUptime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseValidatorUptime converts a slice of bytes into an initialized ValidatorUptime.
func ParseValidatorUptime(b []byte) (*ValidatorUptime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bytes returns the binary representation of this payload. It assumes that the
// payload is initialized from either NewValidatorUptime or Parse.
func (b *ValidatorUptime) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *ValidatorUptime) initialize(bytes []byte) { _ = "STUB: not implemented"; return }
