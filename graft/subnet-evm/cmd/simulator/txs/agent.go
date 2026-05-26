// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"context"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/cmd/simulator/metrics"
)

type THash interface {
	Hash() common.Hash
}

// TxSequence provides an interface to return a channel of transactions.
// The sequence is responsible for closing the channel when there are no further
// transactions.
type TxSequence[T THash] interface {
	Chan() <-chan T
}

// Worker defines the interface for issuance and confirmation of transactions.
// The caller is responsible for calling Close to cleanup resources used by the
// worker at the end of the simulation.
type Worker[T THash] interface {
	IssueTx(ctx context.Context, tx T) error
	ConfirmTx(ctx context.Context, tx T) error
	LatestHeight(ctx context.Context) (uint64, error)
}

// Execute the work of the given agent.
type Agent[T THash] interface {
	Execute(ctx context.Context) error
}

// issueNAgent issues and confirms a batch of N transactions at a time.
type issueNAgent[T THash] struct {
	sequence TxSequence[T]
	worker   Worker[T]
	n        uint64
	metrics  *metrics.Metrics
}

// NewIssueNAgent creates a new issueNAgent
func NewIssueNAgent[T THash](sequence TxSequence[T], worker Worker[T], n uint64, metrics *metrics.Metrics) Agent[T] {
	_ = "STUB: not implemented"
	return nil
}

// Execute issues txs in batches of N and waits for them to confirm
func (a issueNAgent[T]) Execute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Tracks the total amount of time waiting for issuing and confirming txs

// Start time for execution

// Start issuance batch

// Get the batch's issuance time and add it to totalIssuedTime

// Wait for txs in this batch to confirm

// Get the batch's confirmation time and add it to totalConfirmedTime

// Check if this is the last batch, if so write the final log and return
