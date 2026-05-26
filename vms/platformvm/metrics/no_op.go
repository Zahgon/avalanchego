// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"net/http"
	"time"

	"github.com/gorilla/rpc/v2"

	"github.com/ava-labs/avalanchego/ids"
)

var Noop Metrics = noopMetrics{}

type noopMetrics struct{}

func (noopMetrics) MarkOptionVoteWon() { _ = "STUB: not implemented"; return }

func (noopMetrics) MarkOptionVoteLost() { _ = "STUB: not implemented"; return }

func (noopMetrics) MarkAccepted(Block) error { _ = "STUB: not implemented"; return nil }

func (noopMetrics) InterceptRequest(i *rpc.RequestInfo) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func (noopMetrics) AfterRequest(*rpc.RequestInfo) { _ = "STUB: not implemented"; return }

func (noopMetrics) IncValidatorSetsCreated() { _ = "STUB: not implemented"; return }

func (noopMetrics) IncValidatorSetsCached() { _ = "STUB: not implemented"; return }

func (noopMetrics) AddValidatorSetsDuration(time.Duration) { _ = "STUB: not implemented"; return }

func (noopMetrics) AddValidatorSetsHeightDiff(uint64) { _ = "STUB: not implemented"; return }

func (noopMetrics) SetLocalStake(uint64) { _ = "STUB: not implemented"; return }

func (noopMetrics) SetTotalStake(uint64) { _ = "STUB: not implemented"; return }

func (noopMetrics) SetTimeUntilUnstake(time.Duration) { _ = "STUB: not implemented"; return }

func (noopMetrics) SetTimeUntilSubnetUnstake(ids.ID, time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (noopMetrics) SetSubnetPercentConnected(ids.ID, float64) { _ = "STUB: not implemented"; return }

func (noopMetrics) SetPercentConnected(float64) { _ = "STUB: not implemented"; return }
