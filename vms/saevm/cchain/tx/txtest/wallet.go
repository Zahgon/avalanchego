// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txtest

import (
	"testing"

	"github.com/stretchr/testify/require"

	// Imported for [secp256k1fx.Credential] comment resolution.
	_ "github.com/ava-labs/avalanchego/vms/secp256k1fx"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
)

// Signature can be used within a [secp256k1fx.Credential] to authorize a
// transaction.
type Signature = [secp256k1.SignatureLen]byte

// NewKey returns a freshly-generated [secp256k1.PrivateKey], failing tb on any
// error.
func NewKey(tb testing.TB) *secp256k1.PrivateKey {
	tb.Helper()

	sk, err := secp256k1.NewPrivateKey()
	require.NoError(tb, err, "secp256k1.NewPrivateKey()")
	return sk
}

// Sign signs u with s and returns the signature.
func Sign(tb testing.TB, u tx.Unsigned, s keychain.Signer) Signature {
	tb.Helper()

	b, err := tx.UnsignedBytes(u)
	require.NoErrorf(tb, err, "tx.UnsignedBytes(%T)", u)
	sig, err := s.Sign(b)
	require.NoErrorf(tb, err, "%T.Sign(%T)", s, u)
	require.Lenf(tb, sig, len(Signature{}), "len(%T.Sign(%T))", s, u)
	return Signature(sig)
}

// MustMarshalUTXO returns the canonical binary format of utxo, failing tb on
// any error.
func MustMarshalUTXO(tb testing.TB, utxo *avax.UTXO) []byte {
	tb.Helper()

	b, err := tx.MarshalUTXO(utxo)
	require.NoError(tb, err, "tx.MarshalUTXO()")
	return b
}

// ExportedUTXOs returns the UTXOs produced by e when wrapped in a [tx.Tx]
// with the given txID.
func ExportedUTXOs(txID ids.ID, e *tx.Export) []*avax.UTXO {
	utxos := make([]*avax.UTXO, len(e.ExportedOutputs))
	for i, out := range e.ExportedOutputs {
		utxos[i] = &avax.UTXO{
			UTXOID: avax.UTXOID{
				TxID:        txID,
				OutputIndex: uint32(i), //#nosec G115 -- Won't overflow
			},
			Asset: out.Asset,
			Out:   out.Out,
		}
	}
	return utxos
}
