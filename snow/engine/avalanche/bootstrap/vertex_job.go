// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/avalanche"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/bootstrap/queue"
	"github.com/ava-labs/avalanchego/snow/engine/avalanche/vertex"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errMissingVtxDependenciesOnAccept = errors.New("attempting to execute blocked vertex")
	errTxNotAcceptedInVtxOnAccept     = errors.New("attempting to execute vertex with non-accepted transaction")
)

type vtxParser struct {
	log         logging.Logger
	numAccepted prometheus.Counter
	manager     vertex.Manager
}

func (p *vtxParser) Parse(ctx context.Context, vtxBytes []byte) (queue.Job, error) {
	_ = "STUB: not implemented"
	return *new(queue.Job), nil
}

type vertexJob struct {
	log         logging.Logger
	numAccepted prometheus.Counter
	vtx         avalanche.Vertex
}

func (v *vertexJob) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (v *vertexJob) MissingDependencies(context.Context) (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns true if this vertex job has at least 1 missing dependency
func (v *vertexJob) HasMissingDependencies(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *vertexJob) Execute(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (v *vertexJob) Bytes() []byte { _ = "STUB: not implemented"; return nil }
