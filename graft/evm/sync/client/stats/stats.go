// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stats

import (
	"time"

	"github.com/ava-labs/libevm/metrics"

	"github.com/ava-labs/avalanchego/graft/evm/message"
)

var (
	_ ClientSyncerStats = (*clientSyncerStats)(nil)
	_ ClientSyncerStats = (*noopStats)(nil)
)

type ClientSyncerStats interface {
	GetMetric(message.Request) (MessageMetric, error)
}

type MessageMetric interface {
	IncRequested()
	IncSucceeded()
	IncFailed()
	IncInvalidResponse()
	IncReceived(int64)
	UpdateRequestLatency(time.Duration)
}

type messageMetric struct {
	requested       metrics.Counter // Number of times a request has been sent
	succeeded       metrics.Counter // Number of times a request has succeeded
	failed          metrics.Counter // Number of times a request failed (does not include invalid responses)
	invalidResponse metrics.Counter // Number of times a request failed due to an invalid response
	received        metrics.Counter // Number of items that have been received

	requestLatency metrics.Timer // Latency for this request
}

func NewMessageMetric(name string) MessageMetric {
	_ = "STUB: not implemented"
	return *new(MessageMetric)
}

func (m *messageMetric) IncRequested() { _ = "STUB: not implemented"; return }

func (m *messageMetric) IncSucceeded() { _ = "STUB: not implemented"; return }

func (m *messageMetric) IncFailed() { _ = "STUB: not implemented"; return }

func (m *messageMetric) IncInvalidResponse() { _ = "STUB: not implemented"; return }

func (m *messageMetric) IncReceived(size int64) { _ = "STUB: not implemented"; return }

func (m *messageMetric) UpdateRequestLatency(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

type clientSyncerStats struct {
	leafMetrics        map[message.NodeType]MessageMetric
	codeRequestMetric  MessageMetric
	blockRequestMetric MessageMetric
}

// NewClientSyncerStats returns stats for the client syncer
func NewClientSyncerStats(leafMetricNames map[message.NodeType]string) *clientSyncerStats {
	_ = "STUB: not implemented"
	return nil
}

// GetMetric returns the appropriate messaage metric for the given request
func (c *clientSyncerStats) GetMetric(msgIntf message.Request) (MessageMetric, error) {
	_ = "STUB: not implemented"
	return *new(MessageMetric), nil
}

// no-op implementation of ClientSyncerStats
type noopStats struct {
	noop noopMsgMetric
}

type noopMsgMetric struct{}

func (noopMsgMetric) IncRequested()                      { _ = "STUB: not implemented"; return }
func (noopMsgMetric) IncSucceeded()                      { _ = "STUB: not implemented"; return }
func (noopMsgMetric) IncFailed()                         { _ = "STUB: not implemented"; return }
func (noopMsgMetric) IncInvalidResponse()                { _ = "STUB: not implemented"; return }
func (noopMsgMetric) IncReceived(int64)                  { _ = "STUB: not implemented"; return }
func (noopMsgMetric) UpdateRequestLatency(time.Duration) { _ = "STUB: not implemented"; return }

func NewNoOpStats() ClientSyncerStats { _ = "STUB: not implemented"; return *new(ClientSyncerStats) }

func (n noopStats) GetMetric(_ message.Request) (MessageMetric, error) {
	_ = "STUB: not implemented"
	return *new(MessageMetric), nil
}
