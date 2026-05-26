// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/ava-labs/avalanchego/ids"
)

// L1ValidatorRegistration reports if a validator is registered on the P-chain.
type L1ValidatorRegistration struct {
	payload

	ValidationID ids.ID `serialize:"true" json:"validationID"`
	// Registered being true means that validationID is currently a validator on
	// the P-chain.
	//
	// Registered being false means that validationID is not and can never
	// become a validator on the P-chain. It is possible that validationID was
	// previously a validator on the P-chain.
	Registered bool `serialize:"true" json:"registered"`
}

// NewL1ValidatorRegistration creates a new initialized L1ValidatorRegistration.
func NewL1ValidatorRegistration(
	validationID ids.ID,
	registered bool,
) (*L1ValidatorRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseL1ValidatorRegistration parses bytes into an initialized
// L1ValidatorRegistration.
func ParseL1ValidatorRegistration(b []byte) (*L1ValidatorRegistration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
