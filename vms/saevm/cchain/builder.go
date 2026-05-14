// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cchain

import (
	"context"
	"fmt"
	"iter"
	"math/big"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/trie"
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
	"github.com/ava-labs/avalanchego/graft/evm/constants"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/acp226"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

var _ hook.BlockBuilder[*hookTx] = (*builder)(nil)

type builder struct {
	ctx          *snow.Context
	now          func() time.Time
	potentialTxs func() iter.Seq[*hookTx]
}

func (b *builder) BuildHeader(parent *types.Header) (*types.Header, error) {
	// TODO(StephenButtolph): Encode the ACP-176 target excess in the header.
	// TODO(StephenButtolph): Encode the ACP-183 min price excess in the header.
	// TODO(StephenButtolph): Enforce the minimum block time here.
	return customtypes.WithHeaderExtra(
		&types.Header{
			ParentHash:       parent.Hash(),
			Coinbase:         constants.BlackholeAddr,
			Difficulty:       big.NewInt(1),
			Number:           new(big.Int).Add(parent.Number, common.Big1),
			Time:             uint64(b.now().Unix()), //#nosec G115 -- Known non-negative
			BlobGasUsed:      new(uint64),
			ExcessBlobGas:    new(uint64),
			ParentBeaconRoot: new(common.Hash),
		},
		&customtypes.HeaderExtra{
			ExtDataGasUsed: big.NewInt(0),
			BlockGasCost:   big.NewInt(0),
			// TODO(StephenButtolph): Encode the millisecond timestamp.
			TimeMilliseconds: new(uint64),
			// TODO(StephenButtolph): Encode the min-delay excess.
			MinDelayExcess: new(acp226.DelayExcess),
		},
	), nil
}

func (b *builder) PotentialEndOfBlockOps(
	ctx context.Context,
	header *types.Header,
	settledHash common.Hash,
	source saetypes.BlockSource,
) iter.Seq[*hookTx] {
	seq := b.potentialTxs()
	return func(yield func(*hookTx) bool) {
		// Transactions are verified against the last executed state. We must
		// guarantee that they don't conflict with any transactions in blocks
		// between the block we are building and the last executed block.
		inputs, err := ancestorInputIDs(header, settledHash, source)
		if err != nil {
			b.ctx.Log.Error("failed to get ancestor input IDs",
				zap.Error(err),
			)
			return
		}

		for t := range seq {
			if inputs.Overlaps(t.inputs) {
				b.ctx.Log.Debug("tx consumes previously consumed inputs",
					zap.Stringer("txID", t.id),
				)
				continue
			}
			if err := t.tx.SanityCheck(b.ctx); err != nil {
				b.ctx.Log.Debug("tx failed sanity check",
					zap.Stringer("txID", t.id),
					zap.Error(err),
				)
				continue
			}
			if err := t.tx.VerifyCredentials(b.ctx.SharedMemory); err != nil {
				b.ctx.Log.Debug("tx failed credential verification",
					zap.Stringer("txID", t.id),
					zap.Error(err),
				)
				continue
			}

			if !yield(t) {
				return
			}
			inputs.Union(t.inputs)
		}
	}
}

// ancestorInputIDs returns the set of input IDs of all custom transactions in
// the block range (h, settled), both exclusive.
func ancestorInputIDs(h *types.Header, settled common.Hash, source saetypes.BlockSource) (set.Set[ids.ID], error) {
	var s set.Set[ids.ID]
	for h.ParentHash != settled {
		parentNumber := h.Number.Uint64() - 1
		p, ok := source(h.ParentHash, parentNumber)
		if !ok {
			return nil, fmt.Errorf("missing block: %s (%d)", h.ParentHash, parentNumber)
		}

		txs, err := tx.ParseSlice(customtypes.BlockExtData(p))
		if err != nil {
			return nil, fmt.Errorf("parsing txs in %s (%d): %w", h.ParentHash, parentNumber, err)
		}
		for _, tx := range txs {
			s.Union(tx.InputIDs())
		}
		h = p.Header()
	}
	return s, nil
}

func (*builder) BuildBlock(
	header *types.Header,
	blockCtx *block.Context,
	ethTxs []*types.Transaction,
	receipts []*types.Receipt,
	avaxTxs []*hookTx,
	settledHeight uint64,
) (*types.Block, error) {
	txs := make([]*tx.Tx, len(avaxTxs))
	for i, avaxTx := range avaxTxs {
		txs[i] = avaxTx.tx
	}
	extData, err := tx.MarshalSlice(txs)
	if err != nil {
		return nil, fmt.Errorf("marshalling txs: %w", err)
	}

	// TODO(StephenButtolph): Encode warp predicate results in the header.
	_ = blockCtx
	// TODO(StephenButtolph): Encode settledHeight in the block.
	_ = settledHeight
	// TODO(StephenButtolph): Verify the extDataHash matches the hash of extData
	// during parsing.
	return customtypes.NewBlockWithExtData(
		header,
		ethTxs,
		nil, // uncles
		receipts,
		trie.NewStackTrie(nil),
		extData,
		true, // update [customtypes.HeaderExtra.ExtDataHash]
	), nil
}
