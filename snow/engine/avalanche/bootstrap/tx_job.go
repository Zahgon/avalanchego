// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/bootstrap/queue"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var errMissingTxDependenciesOnAccept = errors.New("attempting to accept a transaction with missing dependencies")

type txParser struct {
	log         logging.Logger
	numAccepted prometheus.Counter
	vm          vertex.LinearizableVM
}

func (p *txParser) Parse(ctx context.Context, txBytes []byte) (queue.Job, error) {
	_ = "STUB: not implemented"
	return *new(queue.Job), nil
}

type txJob struct {
	log         logging.Logger
	numAccepted prometheus.Counter
	tx          snowstorm.Tx
}

func (t *txJob) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (t *txJob) MissingDependencies(context.Context) (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns true if this tx job has at least 1 missing dependency
func (t *txJob) HasMissingDependencies(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *txJob) Execute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *txJob) Bytes() []byte { _ = "STUB: not implemented"; return nil }
