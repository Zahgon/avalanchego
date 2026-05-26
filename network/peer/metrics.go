// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/message"
)

const (
	ioLabel         = "io"
	opLabel         = "op"
	compressedLabel = "compressed"

	sentLabel     = "sent"
	receivedLabel = "received"
)

var (
	opLabels             = []string{opLabel}
	ioOpLabels           = []string{ioLabel, opLabel}
	ioOpCompressedLabels = []string{ioLabel, opLabel, compressedLabel}
)

type Metrics struct {
	ClockSkewCount prometheus.Counter
	ClockSkewSum   prometheus.Gauge

	RTTCount prometheus.Counter
	RTTSum   prometheus.Gauge

	NumFailedToParse prometheus.Counter
	NumSendFailed    *prometheus.CounterVec // op

	Messages   *prometheus.CounterVec // io + op + compressed
	Bytes      *prometheus.CounterVec // io + op
	BytesSaved *prometheus.GaugeVec   // io + op
}

func NewMetrics(registerer prometheus.Registerer) (*Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sent updates the metrics for having sent [msg].
func (m *Metrics) Sent(msg *message.OutboundMessage) { _ = "STUB: not implemented"; return }

// assume that if [saved] == 0, [msg] wasn't compressed

func (m *Metrics) MultipleSendsFailed(op message.Op, count int) { _ = "STUB: not implemented"; return }

// SendFailed updates the metrics for having failed to send [msg].
func (m *Metrics) SendFailed(msg *message.OutboundMessage) { _ = "STUB: not implemented"; return }

func (m *Metrics) Received(msg *message.InboundMessage, msgLen uint32) {
	_ = "STUB: not implemented"
	return
}

// assume that if [saved] == 0, [msg] wasn't compressed
