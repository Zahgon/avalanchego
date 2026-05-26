// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

var secpCache = secp256k1.NewRecoverCache(2048)

type Tx struct {
	Unsigned  `serialize:"true" json:"unsigned"`
	Signature [secp256k1.SignatureLen]byte `serialize:"true" json:"signature"`
}

func Parse(bytes []byte) (*Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func Sign(utx Unsigned, key *secp256k1.PrivateKey) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) ID() (ids.ID, error) { _ = "STUB: not implemented"; return *new(ids.ID), nil }

func (tx *Tx) SenderID() (ids.ShortID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ShortID), nil
}
