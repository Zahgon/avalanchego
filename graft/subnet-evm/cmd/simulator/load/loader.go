// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/cmd/simulator/config"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/cmd/simulator/metrics"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/cmd/simulator/txs"
)

const (
	MetricsEndpoint = "/metrics" // Endpoint for the Prometheus Metrics Server
)

// Loader executes a series of worker/tx sequence pairs.
// Each worker/txSequence pair issues [batchSize] transactions, confirms all
// of them as accepted, and then moves to the next batch until the txSequence
// is exhausted.
type Loader[T txs.THash] struct {
	clients     []txs.Worker[T]
	txSequences []txs.TxSequence[T]
	batchSize   uint64
	metrics     *metrics.Metrics
}

func New[T txs.THash](
	clients []txs.Worker[T],
	txSequences []txs.TxSequence[T],
	batchSize uint64,
	metrics *metrics.Metrics,
) *Loader[T] {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader[T]) Execute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ConfirmReachedTip finds the max height any client has reached and then ensures every client
// reaches at least that height.
//
// This allows the network to continue to roll forward and creates a synchronization point to ensure
// that every client in the loader has reached at least the max height observed of any client at
// the time this function was called.
func (l *Loader[T]) ConfirmReachedTip(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecuteLoader creates txSequences from [config] and has txAgents execute the specified simulation.
func ExecuteLoader(ctx context.Context, config config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Create buffered sigChan to receive SIGINT notifications

// Create context with cancel

// Blocks until we receive a SIGINT notification or if parent context is done

// Cancel the child context and end all processes

// Construct the arguments for the load simulator

// Ensure there are at least [config.Workers] keys and save any newly generated ones.

// Each address needs: params.GWei * MaxFeeCap * ethparams.TxGas * TxsPerWorker total wei
// to fund gas for all of their transactions.

// Print regardless of execution error
