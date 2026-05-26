// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/tests"
)

type metricKind uint

const (
	counter metricKind = iota + 1
	gauge
)

var (
	gasMetric = topLevelMetric{
		name:  "gas",
		query: "avalanche_evm_eth_chain_block_gas_used_processed",
		kind:  counter,
	}
	meterVMMetrics = []topLevelMetric{
		{
			name:  "block_parse",
			query: "avalanche_meterchainvm_parse_block_sum",
			kind:  gauge,
		},
		{
			name:  "block_verify",
			query: "avalanche_meterchainvm_verify_sum",
			kind:  gauge,
		},
		{
			name:  "block_accept",
			query: "avalanche_meterchainvm_accept_sum",
			kind:  gauge,
		},
	}
)

type topLevelMetric struct {
	name  string
	query string
	kind  metricKind
}

func getMetricValue(registry prometheus.Gatherer, metric topLevelMetric) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getTopLevelMetrics(tc tests.TestContext, tool *benchmarkTool, registry prometheus.Gatherer, elapsed time.Duration) {
	_ = "STUB: not implemented"
	return
}

// MeterVM counters are in terms of nanoseconds
