// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utilstest

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ava-labs/libevm/common"
)

// Key contains an ecdsa private key field as well as an address field
// obtained from converting the ecdsa public key.
type Key struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

// NewKey generates a new key pair and returns a pointer to a [Key].
func NewKey(t *testing.T) *Key { _ = "STUB: not implemented"; return nil }
