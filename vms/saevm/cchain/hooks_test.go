// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/trie"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/utils/set"
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

func TestAncestorInputIDs(t *testing.T) {
	// Three sequential blocks chained by hash, each contributing distinct
	// InputIDs via a single Export tx. The wallet auto-increments the input
	// nonce on each call, so the txs are guaranteed to have disjoint inputs.
	w := newWallet(txtest.NewKey(t), snowtest.Context(t, snowtest.CChainID), nil)
	export := func() *tx.Tx {
		signed, _ := w.newExportTx(
			t,
			snowtest.XChainID,
			[]*secp256k1fx.TransferOutput{
				txtest.NewTransferOutput(50, w.sk.Address()),
			},
			0,
		)
		return signed
	}

	var (
		settled   = common.Hash{0xff}
		txA       = export()
		txB       = export()
		txC       = export()
		blockA    = blockWithTxs(t, 1, settled, []*tx.Tx{txA})
		blockB    = blockWithTxs(t, 2, blockA.Hash(), []*tx.Tx{txB})
		blockC    = blockWithTxs(t, 3, blockB.Hash(), []*tx.Tx{txC})
		nextBlock = blockWithTxs(t, 4, blockC.Hash(), nil) // header only; not in source
	)

	tests := []struct {
		name    string
		header  *types.Header
		settled common.Hash
		want    set.Set[ids.ID]
		wantErr error
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
			want:    union(txA.InputIDs(), txB.InputIDs(), txC.InputIDs()),
		},
		{
			name:    "missing block",
			header:  blockB.Header(),
			settled: common.Hash{0xaa}, // never matches
			wantErr: errMissingBlock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := func(hash common.Hash, number uint64) (*types.Block, bool) {
				for _, b := range []*types.Block{blockA, blockB, blockC} {
					if b.Hash() == hash && b.NumberU64() == number {
						return b, true
					}
				}
				return nil, false
			}

			got, err := ancestorInputIDs(tt.header, tt.settled, source)
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}

func union(sets ...set.Set[ids.ID]) set.Set[ids.ID] {
	var out set.Set[ids.ID]
	for _, s := range sets {
		out.Union(s)
	}
	return out
}
