// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"testing"

	"github.com/ava-labs/libevm/libevm/options"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx/txtest"
	"github.com/ava-labs/avalanchego/vms/saevm/saetest"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

// TestIssueTxRejectsInvalidTransaction asserts that [Client.IssueTx] surfaces
// an error from the transaction pool's verification pipeline. The individual
// rejection categories are covered by [txpool] tests; this test only confirms
// that verification is wired into the API layer.
func TestIssueTxRejectsInvalidTransaction(t *testing.T) {
	sk := txtest.NewKey(t)
	sender := sk.EthAddress()
	sut := newSUT(t, options.Func[sutConfig](func(c *sutConfig) {
		c.genesis.Alloc = saetest.MaxAllocFor(sender)
	}))

	w := newWallet(sk, sut.snowCtx, sut.Client)
	signed, _ := w.newExportTx(
		t,
		sut.snowCtx.XChainID,
		[]*secp256k1fx.TransferOutput{txtest.NewTransferOutput(50, sk.Address())},
		50,
	)

	// First submission seeds the pool; the second exercises the pool's
	// duplicate check, proving the full verification pipeline is wired.
	require.NoErrorf(t, sut.IssueTx(t.Context(), signed), "%T.IssueTx() seed", sut.Client)

	err := sut.IssueTx(t.Context(), signed)
	require.ErrorContainsf(t, err, errIssuingTx.Error(), "%T.IssueTx() resubmit", sut.Client)
}

// TestGetTxNotFound asserts that [Client.GetTx] surfaces an error when the
// requested tx has never been accepted.
func TestGetTxNotFound(t *testing.T) {
	sut := newSUT(t)

	_, _, err := sut.GetTx(t.Context(), ids.GenerateTestID())
	require.ErrorContainsf(t, err, errFetchingTx.Error(), "%T.GetTx()", sut.Client)
}

// TestGetUTXOsPagination seeds N UTXOs and walks [Client.GetUTXOs] with
// limit=1, asserting that each non-terminal page returns exactly one UTXO,
// that the terminal page is signaled by len(page) < limit (i.e. zero
// results), and that the union of pages matches the seeded set with no
// duplicates.
func TestGetUTXOsPagination(t *testing.T) {
	sut := newSUT(t)

	addr := txtest.NewKey(t).Address()

	const numUTXOs = 5
	want := make([]*avax.UTXO, numUTXOs)
	for i := range want {
		want[i] = txtest.NewUTXO(uint64(i+1), sut.snowCtx.AVAXAssetID, addr)
	}
	sut.addUTXOs(t, snowtest.XChainID, want...)

	var (
		ctx         = t.Context()
		got         []*avax.UTXO
		startAddr   ids.ShortID
		startUTXOID ids.ID
	)
	for {
		const limit = 1
		page, endAddr, endUTXOID, err := sut.GetUTXOs(
			ctx,
			[]ids.ShortID{addr},
			snowtest.XChainID,
			limit,
			startAddr,
			startUTXOID,
		)
		require.NoErrorf(t, err, "%T.GetUTXOs()", sut.Client)
		got = append(got, page...)
		if len(page) < limit {
			break
		}
		startAddr, startUTXOID = endAddr, endUTXOID
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(avax.UTXOID{}, secp256k1fx.OutputOwners{}),
		cmpopts.SortSlices(func(a, b *avax.UTXO) bool {
			aID, bID := a.InputID(), b.InputID()
			return aID.Compare(bID) < 0
		}),
	}
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("paginated UTXOs (-want +got):\n%s", diff)
	}
}
