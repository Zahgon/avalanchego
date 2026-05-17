// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txtest

import (
	"math"
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

const _x2cRate = 1_000_000_000

var x2cRate = uint256.NewInt(_x2cRate)

// ScaleAVAX converts an amount denominated in nAVAX into the C-Chain's aAVAX
// denomination.
func ScaleAVAX(nAVAX uint64) uint256.Int {
	var aAVAX uint256.Int
	aAVAX.SetUint64(nAVAX)
	aAVAX.Mul(&aAVAX, x2cRate)
	return aAVAX
}

// AddNAVAX returns balance + nAVAXDelta nAVAX (scaled to aAVAX). The delta
// may be negative. It panics if the result does not fit in a uint256.
func AddNAVAX(balance uint256.Int, nAVAXDelta int64) uint256.Int {
	delta := new(big.Int).Mul(big.NewInt(nAVAXDelta), big.NewInt(_x2cRate))
	sum := new(big.Int).Add(balance.ToBig(), delta)
	result, overflow := uint256.FromBig(sum)
	if overflow {
		panic("AddNAVAX: result overflows uint256")
	}
	return *result
}

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

// MustParseUTXO deserializes an [avax.UTXO] from its canonical binary format,
// failing tb on any error.
func MustParseUTXO(tb testing.TB, b []byte) *avax.UTXO {
	tb.Helper()

	utxo, err := tx.ParseUTXO(b)
	require.NoError(tb, err, "tx.ParseUTXO()")
	return utxo
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

// A Wallet builds and signs cross-chain transactions on behalf of a single
// key. It is the analog of [wallet/chain/c.Wallet] for SAE.
type Wallet struct {
	sk      *secp256k1.PrivateKey
	snowCtx *snow.Context
	memory  *atomic.Memory
	nonce   uint64
}

// NewWallet returns a [*Wallet] backed by sk for the chain described by
// snowCtx. memory is consulted when building imports to discover spendable
// UTXOs.
func NewWallet(sk *secp256k1.PrivateKey, snowCtx *snow.Context, memory *atomic.Memory) *Wallet {
	return &Wallet{
		sk:      sk,
		snowCtx: snowCtx,
		memory:  memory,
	}
}

// NewExportTx builds and signs an [tx.Export] sending outputs to
// destinationChain. The wallet contributes a single AVAX input from its eth
// address with Amount = sum(outputs.Amt) + fee, using its next nonce.
func (w *Wallet) NewExportTx(
	tb testing.TB,
	destinationChain ids.ID,
	outputs []*secp256k1fx.TransferOutput,
	fee uint64,
) (*tx.Tx, *tx.Export) {
	tb.Helper()

	avaxAssetID := w.snowCtx.AVAXAssetID
	var exportedAmount uint64
	transferable := make([]*avax.TransferableOutput, len(outputs))
	for i, out := range outputs {
		transferable[i] = &avax.TransferableOutput{
			Asset: avax.Asset{ID: avaxAssetID},
			Out:   out,
		}
		exportedAmount += out.Amt
	}

	export := &tx.Export{
		NetworkID:        w.snowCtx.NetworkID,
		BlockchainID:     w.snowCtx.ChainID,
		DestinationChain: destinationChain,
		Ins: []tx.Input{{
			Address: w.sk.EthAddress(),
			Amount:  exportedAmount + fee,
			AssetID: avaxAssetID,
			Nonce:   w.nonce,
		}},
		ExportedOutputs: transferable,
	}
	w.nonce++

	return w.sign(tb, export, 1), export
}

// NewImportTx builds and signs an [tx.Import] consuming all spendable AVAX
// UTXOs in shared memory between sourceChain and the C-Chain owned by the
// wallet, crediting the total imported (minus fee) to `to` on the C-Chain.
func (w *Wallet) NewImportTx(
	tb testing.TB,
	sourceChain ids.ID,
	to common.Address,
	fee uint64,
) (*tx.Tx, *tx.Import) {
	tb.Helper()

	cMemory := w.memory.NewSharedMemory(w.snowCtx.ChainID)
	utxoBytes, _, _, err := cMemory.Indexed(
		sourceChain,
		[][]byte{w.sk.Address().Bytes()},
		nil,
		nil,
		math.MaxInt,
	)
	require.NoErrorf(tb, err, "%T.Indexed()", cMemory)

	var (
		avaxAssetID  = w.snowCtx.AVAXAssetID
		importedAVAX uint64
		inputs       = make([]*avax.TransferableInput, 0, len(utxoBytes))
	)
	for _, b := range utxoBytes {
		utxo := MustParseUTXO(tb, b)
		if utxo.Asset.ID != avaxAssetID {
			continue
		}

		out, ok := utxo.Out.(*secp256k1fx.TransferOutput)
		require.Truef(tb, ok, "unexpected UTXO output type %T", utxo.Out)

		importedAVAX += out.Amt
		inputs = append(inputs, &avax.TransferableInput{
			UTXOID: utxo.UTXOID,
			Asset:  utxo.Asset,
			In: &secp256k1fx.TransferInput{
				Amt: out.Amt,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{0},
				},
			},
		})
	}
	require.Greaterf(tb, importedAVAX, fee, "imported AVAX insufficient to cover fee")

	imp := &tx.Import{
		NetworkID:      w.snowCtx.NetworkID,
		BlockchainID:   w.snowCtx.ChainID,
		SourceChain:    sourceChain,
		ImportedInputs: inputs,
		Outs: []tx.Output{{
			Address: to,
			Amount:  importedAVAX - fee,
			AssetID: avaxAssetID,
		}},
	}
	return w.sign(tb, imp, len(inputs)), imp
}

// sign wraps u in a [tx.Tx] with numCreds copies of a single-sig credential
// over u.
func (w *Wallet) sign(tb testing.TB, u tx.Unsigned, numCreds int) *tx.Tx {
	tb.Helper()

	sig := Sign(tb, u, w.sk)
	creds := make([]tx.Credential, numCreds)
	for i := range creds {
		creds[i] = &secp256k1fx.Credential{Sigs: []Signature{sig}}
	}
	return &tx.Tx{
		Unsigned: u,
		Creds:    creds,
	}
}
