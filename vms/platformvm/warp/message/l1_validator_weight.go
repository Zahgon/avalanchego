// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
)

var ErrNonceReservedForRemoval = errors.New("maxUint64 nonce is reserved for removal")

// L1ValidatorWeight is both received and sent by the P-chain.
//
// If the P-chain is receiving this message, it is treated as a command to
// update the weight of the validator.
//
// If the P-chain is sending this message, it is reporting the current nonce and
// weight of the validator.
type L1ValidatorWeight struct {
	payload

	ValidationID ids.ID `serialize:"true" json:"validationID"`
	Nonce        uint64 `serialize:"true" json:"nonce"`
	Weight       uint64 `serialize:"true" json:"weight"`
}

func (s *L1ValidatorWeight) Verify() error { _ = "STUB: not implemented"; return nil }

// NewL1ValidatorWeight creates a new initialized L1ValidatorWeight.
func NewL1ValidatorWeight(
	validationID ids.ID,
	nonce uint64,
	weight uint64,
) (*L1ValidatorWeight, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseL1ValidatorWeight parses bytes into an initialized L1ValidatorWeight.
func ParseL1ValidatorWeight(b []byte) (*L1ValidatorWeight, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
