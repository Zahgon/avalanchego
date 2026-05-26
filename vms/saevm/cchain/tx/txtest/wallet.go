// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txtest

import (
	"testing"

	// Imported for [secp256k1fx.Credential] comment resolution.
	_ "github.com/ava-labs/avalanchego/vms/secp256k1fx"

	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
)

// Signature can be used within a [secp256k1fx.Credential] to authorize a
// transaction.
type Signature = [secp256k1.SignatureLen]byte

// Sign signs u with s and returns the signature.
func Sign(tb testing.TB, u tx.Unsigned, s keychain.Signer) Signature {
	_ = "STUB: not implemented"
	return *new(Signature)
}
