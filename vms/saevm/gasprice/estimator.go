// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package gasprice provides gas price statistics and suggestions for timely
// transaction inclusion.
package gasprice

import (
	"context"
	"errors"
	"io"
	"math/big"
	"sync"
	"time"

	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
)

// Backend that the [Estimator] depends on for chain data.
type Backend interface {
	ResolveBlockNumber(bn rpc.BlockNumber) (uint64, error)
	BlockByNumber(bn rpc.BlockNumber) (*types.Block, error)
	SubscribeAcceptedBlocks(ch chan<- *blocks.Block) event.Subscription
	LastAcceptedBlock() *blocks.Block
}

// Config allows parameterizing an [Estimator].
type Config struct {
	// Now returns the current time.
	Now func() time.Time

	// MinSuggestedTip is the minimum suggested tip and the default tip if no
	// better estimate can be made.
	MinSuggestedTip *big.Int
	// SuggestedTipPercentile, in the range (0, 100], specifies what percentile of
	// tips is used when suggesting based on recent transactions.
	SuggestedTipPercentile uint64
	MaxSuggestedTip        *big.Int

	// SuggestedTipMaxBlocks specifies the maximum number of recent blocks to fetch
	// for [Estimator.SuggestTipCap].
	SuggestedTipMaxBlocks uint64
	// SuggestedTipMaxDuration specifies how long a block is considered recent
	// for [Estimator.SuggestGasTipCap].
	SuggestedTipMaxDuration time.Duration

	// HistoryMaxBlocksFromHead specifies the furthest lastBlock behind the last
	// accepted block that can be requested by [Estimator.FeeHistory].
	HistoryMaxBlocksFromHead uint64
	// HistoryMaxBlocks specifies the maximum number of blocks that can be
	// fetched in a single call to [Estimator.FeeHistory].
	HistoryMaxBlocks uint64
}

// DefaultConfig returns a [Config] with all fields set to their default values.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// SuggestedTipPercentile is chosen to be a value that is lower than the median of the tips in the recent blocks.
// This is to prevent suggesting a tip that could cause a self-induced fee spiral.

// Chosen to be larger than the fee lookback window that MetaMask uses (20k blocks).

var (
	errNilNow             = errors.New("config Now must be non-nil")
	errNilMinSuggestedTip = errors.New("config MinSuggestedTip must be non-nil")
	errNilMaxSuggestedTip = errors.New("config MaxSuggestedTip must be non-nil")
	errBadTipPercentile   = errors.New("config SuggestedTipPercentile must be in (0, 100]")
	errMinTipExceedsMax   = errors.New("config MinSuggestedTip must be <= MaxSuggestedTip")
)

// validate returns an error if the config is invalid.
func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }

type last struct {
	lock   sync.RWMutex
	number uint64
	price  *big.Int
}

// Estimator provides gas-price suggestions and fee-history data for SAE by
// analyzing recently accepted blocks.
type Estimator struct {
	backend Backend
	c       Config

	last last

	acceptedBlocks event.Subscription
	blockCache     *blockCache
}

// NewEstimator creates an Estimator for gas tips and fee history.
func NewEstimator(backend Backend, log logging.Logger, c Config) (*Estimator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New blocks are cached in the background upon acceptance to avoid slow
// responses after long periods of no requests to the estimator. This
// allows us to avoid parallelizing reads inside individual API calls.

// Additional slots in the cache allows processing queries for previous
// blocks while new blocks are added concurrently.

//#nosec G115 -- Overflow would require misconfiguration

// `Unsubscribe` can fire twice on Close(), but it's safe to call multiple times.

// SuggestGasTipCap recommends a priority-fee (tip) for new transactions based on
// tips from recently accepted transactions.
func (e *Estimator) SuggestGasTipCap(ctx context.Context) (tip *big.Int, _ error) {
	_ = "STUB: not implemented"

	// Tip is modified by callers of this function, so we must ensure that
	// it is copied.
	return nil, nil
}

// A different goroutine might have beaten us when upgrading to a write lock.

//#nosec G115 -- Known non-negative

// getBlock does not return an error if the context is cancelled.
// We don't want to early return from `SuggestGasTipCap` if the context is cancelled.
// Instead we continue to fetch the blocks and cache them.

//#nosec G115 -- Known to be between (0, 100]

var (
	errHistoryDepthExhausted = errors.New("requested block is too far behind accepted head")
	errMissingBlock          = errors.New("missing block")
)

// FeeHistory returns data relevant for fee estimation based on the specified
// range of blocks.
//
// The range can be specified either with absolute block numbers or ending with
// the latest or pending block.
//
// This function returns:
//
//   - The first block of the actually processed range.
//   - The tips paid for each percentile of the cumulative gas limits of the
//     transactions in each block.
//   - The baseFee of each block and the next block's upperbound baseFee.
//   - The portion that each block was filled.
func (e *Estimator) FeeHistory(
	ctx context.Context,
	blocks uint64,
	lastBlock rpc.BlockNumber,
	rewardPercentiles []float64,
) (
	lowestHeight *big.Int,
	rewards [][]*big.Int,
	baseFees []*big.Int,
	portionFull []float64,
	_ error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

// requested value
// DoS protection
// Underflow protection for "first" calculation

var _ io.Closer = (*Estimator)(nil)

// Close releases allocated resources.
func (e *Estimator) Close() error { _ = "STUB: not implemented"; return nil }

const maxPercentiles = 100

var errBadPercentile = errors.New("percentile out of range or misordered")

func validatePercentiles(percentiles []float64) error { _ = "STUB: not implemented"; return nil }
