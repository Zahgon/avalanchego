// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

var ErrNilCredential = errors.New("nil credential")

type Credential struct {
	Sigs [][secp256k1.SignatureLen]byte `serialize:"true" json:"signatures"`
}

// MarshalJSON marshals [cr] to JSON
// The string representation of each signature is created using the hex formatter
func (cr *Credential) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cr *Credential) Verify() error { _ = "STUB: not implemented"; return nil }

func (cr *Credential) Self() *Credential { _ = "STUB: not implemented"; return nil }
