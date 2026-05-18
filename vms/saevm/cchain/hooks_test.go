// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/trie"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx/txtest"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

// blockWithTxs returns a minimal [*types.Block] whose ExtData encodes txs and
// whose header is configured for ancestor traversal (parent hash + number).
func blockWithTxs(tb testing.TB, number uint64, parent common.Hash, txs []*tx.Tx) *types.Block {
	tb.Helper()

	extData, err := tx.MarshalSlice(txs)
	require.NoErrorf(tb, err, "tx.MarshalSlice(%d txs)", len(txs))

	header := customtypes.WithHeaderExtra(
		&types.Header{
			ParentHash: parent,
			Number:     new(big.Int).SetUint64(number),
		},
		&customtypes.HeaderExtra{},
	)
	return customtypes.NewBlockWithExtData(header, nil, nil, nil, trie.NewStackTrie(nil), extData, true)
}

// exportTx builds a signed [tx.Export] that consumes one input from sk's
// EthAddress at the given nonce. The exact contents are unimportant — only
// InputIDs() matters for ancestor-tracking tests.
func exportTx(tb testing.TB, nonce uint64) *tx.Tx {
	tb.Helper()

	sk := txtest.NewKey(tb)
	avaxAssetID := snowtest.AVAXAssetID
	export := &tx.Export{
		NetworkID:        constants.UnitTestID,
		BlockchainID:     snowtest.CChainID,
		DestinationChain: snowtest.XChainID,
		Ins: []tx.Input{{
			Address: sk.EthAddress(),
			Amount:  100,
			AssetID: avaxAssetID,
			Nonce:   nonce,
		}},
		ExportedOutputs: []*avax.TransferableOutput{{
			Asset: avax.Asset{ID: avaxAssetID},
			Out:   txtest.NewTransferOutput(50, sk.Address()),
		}},
	}
	sig := txtest.Sign(tb, export, sk)
	return &tx.Tx{
		Unsigned: export,
		Creds:    []tx.Credential{&secp256k1fx.Credential{Sigs: []txtest.Signature{sig}}},
	}
}

func TestAncestorInputIDs(t *testing.T) {
	// Three sequential blocks chained by hash, each contributing distinct
	// InputIDs via a single Export tx.
	var (
		settled   = common.Hash{0xff}
		txA       = exportTx(t, 0)
		txB       = exportTx(t, 1)
		txC       = exportTx(t, 2)
		blockA    = blockWithTxs(t, 1, settled, []*tx.Tx{txA})
		blockB    = blockWithTxs(t, 2, blockA.Hash(), []*tx.Tx{txB})
		blockC    = blockWithTxs(t, 3, blockB.Hash(), []*tx.Tx{txC})
		nextBlock = blockWithTxs(t, 4, blockC.Hash(), nil) // header only; not in source
	)

	source := func(hash common.Hash, number uint64) (*types.Block, bool) {
		for _, b := range []*types.Block{blockA, blockB, blockC} {
			if b.Hash() == hash && b.NumberU64() == number {
				return b, true
			}
		}
		return nil, false
	}

	tests := []struct {
		name    string
		header  *types.Header
		settled common.Hash
		want    set.Set[ids.ID]
		wantErr string
	}{
		{
			name:    "empty range",
			header:  blockA.Header(),
			settled: settled,
			want:    nil,
		},
		{
			name:    "single ancestor",
			header:  blockB.Header(),
			settled: settled,
			want:    txA.InputIDs(),
		},
		{
			name:    "multiple ancestors",
			header:  nextBlock.Header(),
			settled: settled,
			want:    setUnion(txA.InputIDs(), txB.InputIDs(), txC.InputIDs()),
		},
		{
			name:    "missing block",
			header:  blockB.Header(),
			settled: common.Hash{0xaa}, // never matches
			wantErr: "missing block",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ancestorInputIDs(tt.header, tt.settled, source)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("ancestorInputIDs() (-want +got):\n%s", diff)
			}
		})
	}
}

func setUnion(sets ...set.Set[ids.ID]) set.Set[ids.ID] {
	var out set.Set[ids.ID]
	for _, s := range sets {
		out.Union(s)
	}
	return out
}
