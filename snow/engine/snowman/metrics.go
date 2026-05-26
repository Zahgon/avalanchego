// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/metric"
)

const (
	pullGossipSource = "pull_gossip"
	pushGossipSource = "push_gossip"
	builtSource      = "built"
	unknownSource    = "unknown"
)

type metrics struct {
	bootstrapFinished                     prometheus.Gauge
	numRequests                           prometheus.Gauge
	numBlocked                            prometheus.Gauge
	numBlockers                           prometheus.Gauge
	numNonVerifieds                       prometheus.Gauge
	numBuilt                              prometheus.Counter
	numBuildsFailed                       prometheus.Counter
	numUselessPutBytes                    prometheus.Counter
	numUselessPushQueryBytes              prometheus.Counter
	numMissingAcceptedBlocks              prometheus.Counter
	numProcessingAncestorFetchesFailed    prometheus.Counter
	numProcessingAncestorFetchesDropped   prometheus.Counter
	numProcessingAncestorFetchesSucceeded prometheus.Counter
	numProcessingAncestorFetchesUnneeded  prometheus.Counter
	selectedVoteIndex                     metric.Averager
	issuerStake                           metric.Averager
	issued                                *prometheus.CounterVec
	blockTimeSkew                         prometheus.Gauge
}

func newMetrics(reg prometheus.Registerer) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register the labels
