// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/metric"
)

type blockMetrics struct {
	buildBlock,
	buildBlockErr,
	parseBlock,
	parseBlockErr,
	getBlock,
	getBlockErr,
	setPreference,
	lastAccepted,
	verify,
	verifyErr,
	accept,
	reject,
	// Height metrics
	getBlockIDAtHeight,
	// Block verification with context metrics
	shouldVerifyWithContext,
	verifyWithContext,
	verifyWithContextErr,
	// Block building with context metrics
	buildBlockWithContext,
	buildBlockWithContextErr,
	// Setting preference with context metrics
	setPreferenceWithContext,
	// Batched metrics
	getAncestors,
	batchedParseBlock,
	// State sync metrics
	stateSyncEnabled,
	getOngoingSyncStateSummary,
	getLastStateSummary,
	parseStateSummary,
	parseStateSummaryErr,
	getStateSummary,
	getStateSummaryErr metric.Averager
}

func (m *blockMetrics) Initialize(
	supportsBlockBuildingWithContext bool,
	supportsSettingPreferenceWithContext bool,
	supportsBatchedFetching bool,
	supportsStateSync bool,
	reg prometheus.Registerer,
) error {
	_ = "STUB: not implemented"
	return nil
}
