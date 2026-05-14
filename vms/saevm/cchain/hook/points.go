// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hook

import (
	"fmt"
	"iter"
	"slices"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/libevm"
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/graft/coreth/core/extstate"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/txpool"
	"github.com/ava-labs/avalanchego/vms/saevm/gastime"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/x/blockdb"

	saestate "github.com/ava-labs/avalanchego/vms/saevm/cchain/state"
	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
	ethparams "github.com/ava-labs/libevm/params"
)

var _ hook.PointsG[*transaction] = (*Points)(nil)

type Points struct {
	builder
	state *saestate.State
}

func NewPoints(
	ctx *snow.Context,
	state *saestate.State,
	pool *txpool.Pending,
) *Points {
	poolTxs := func(yield func(*transaction) bool) {
		for rawTx := range pool.Iter() {
			t, err := newTx(rawTx, ctx.AVAXAssetID)
			if err != nil {
				ctx.Log.Warn("failed to wrap tx",
					zap.Stringer("txID", rawTx.ID()),
					zap.Error(err),
				)
				continue
			}
			if !yield(t) {
				return
			}
		}
	}
	return &Points{
		builder{
			ctx,
			time.Now,
			func() iter.Seq[*transaction] {
				return poolTxs
			},
		},
		state,
	}
}

func (p *Points) BlockRebuilderFrom(b *types.Block) (hook.BlockBuilder[*transaction], error) {
	rawTxs, err := tx.ParseSlice(customtypes.BlockExtData(b))
	if err != nil {
		return nil, fmt.Errorf("parsing txs: %w", err)
	}

	txs := make([]*transaction, len(rawTxs))
	for i, rawTx := range rawTxs {
		tx, err := newTx(rawTx, p.ctx.AVAXAssetID)
		if err != nil {
			return nil, fmt.Errorf("converting tx %s (%d): %w", rawTx.ID(), i, err)
		}
		txs[i] = tx
	}

	now := p.BlockTime(b.Header())
	potentialTxs := slices.Values(txs)
	return &builder{
		p.ctx,
		func() time.Time {
			return now
		},
		func() iter.Seq[*transaction] {
			return potentialTxs
		},
	}, nil
}

func (p *Points) ExecutionResultsDB(dataDir string) (saetypes.ExecutionResults, error) {
	db, err := blockdb.New(
		blockdb.DefaultConfig().WithDir(dataDir),
		p.ctx.Log,
	)
	if err != nil {
		return saetypes.ExecutionResults{}, fmt.Errorf("creating execution results db: %w", err)
	}
	return saetypes.ExecutionResults{HeightIndex: db}, nil
}

func (*Points) GasConfigAfter(h *types.Header) (gas.Gas, gastime.GasPriceConfig) {
	// TODO(StephenButtolph): Extract from the header.
	return 1_000_000, gastime.GasPriceConfig{
		TargetToExcessScaling: 87,
		MinPrice:              1,
	}
}

func (*Points) SettledHeight(h *types.Header) uint64 {
	// TODO(StephenButtolph): Extract from the header.
	return 0
}

func (*Points) BlockTime(h *types.Header) time.Time {
	// TODO(StephenButtolph): Include milliseconds.
	return time.Unix(int64(h.Time), 0) //#nosec G115 -- Won't overflow for a few millennia
}

func (p *Points) EndOfBlockOps(b *types.Block) ([]hook.Op, error) {
	txs, err := tx.ParseSlice(customtypes.BlockExtData(b))
	if err != nil {
		return nil, fmt.Errorf("parsing txs: %w", err)
	}

	ops := make([]hook.Op, len(txs))
	for i, tx := range txs {
		op, err := tx.AsOp(p.ctx.AVAXAssetID)
		if err != nil {
			return nil, fmt.Errorf("converting tx %s (%d): %w", tx.ID(), i, err)
		}
		ops[i] = op
	}
	return ops, nil
}

func (*Points) CanExecuteTransaction(common.Address, *common.Address, libevm.StateReader) error {
	return nil
}

func (*Points) BeforeExecutingBlock(ethparams.Rules, *state.StateDB, *types.Block) error {
	return nil
}

func (p *Points) AfterExecutingBlock(statedb *state.StateDB, b *types.Block, receipts types.Receipts) error {
	txs, err := tx.ParseSlice(customtypes.BlockExtData(b))
	if err != nil {
		return fmt.Errorf("parsing txs: %w", err)
	}

	extstatedb := extstate.New(statedb)
	for i, tx := range txs {
		if err := tx.TransferNonAVAX(p.ctx.AVAXAssetID, extstatedb); err != nil {
			return fmt.Errorf("transferring non-AVAX assets of tx %s (%d): %w", tx.ID(), i, err)
		}
	}

	if err := p.state.Apply(b.NumberU64(), txs); err != nil {
		return fmt.Errorf("applying cross-chain state: %w", err)
	}
	return nil
}
