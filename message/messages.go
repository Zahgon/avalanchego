// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/utils/compression"
)

const (
	typeLabel      = "type"
	opLabel        = "op"
	directionLabel = "direction"

	compressionLabel   = "compression"
	decompressionLabel = "decompression"
)

var (
	_ fmt.Stringer = (*InboundMessage)(nil)

	metricLabels = []string{typeLabel, opLabel, directionLabel}

	errUnknownCompressionType = errors.New("message is compressed with an unknown compression type")
)

type InboundMessage struct {
	NodeID                ids.NodeID
	Op                    Op
	Message               fmt.Stringer
	Expiration            time.Time
	BytesSavedCompression int
	onFinishedHandling    func()
}

func (m *InboundMessage) OnFinishedHandling() { _ = "STUB: not implemented"; return }

func (m *InboundMessage) String() string { _ = "STUB: not implemented"; return "" }

// OutboundMessage represents a set of fields for an outbound message that can
// be sent over the wire
type OutboundMessage struct {
	// BypassThrottling is true if we should send this message, regardless of
	// any outbound message throttling
	BypassThrottling bool
	Op               Op
	Bytes            []byte
	// BytesSavedCompression stores the amount of bytes that this message saved
	// due to being compressed
	BytesSavedCompression int
}

// TODO: add other compression algorithms with extended interface
type msgBuilder struct {
	zstdCompressor compression.Compressor
	count          *prometheus.CounterVec // type + op + direction
	duration       *prometheus.GaugeVec   // type + op + direction

	maxMessageTimeout time.Duration
}

func newMsgBuilder(
	metrics prometheus.Registerer,
	maxMessageTimeout time.Duration,
) (*msgBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mb *msgBuilder) marshal(
	uncompressedMsg *p2p.Message,
	compressionType compression.Type,
) ([]byte, int, Op, error) {
	_ = "STUB: not implemented"
	return nil, 0, *new(Op), nil
}

// If compression is enabled, we marshal twice:
// 1. the original message
// 2. the message with compressed bytes
//
// This recursive packing allows us to avoid an extra compression on/off
// field in the message.

func (mb *msgBuilder) unmarshal(b []byte) (*p2p.Message, int, Op, error) {
	_ = "STUB: not implemented"
	return nil, 0, *new(Op), nil
}

// Figure out what compression type, if any, was used to compress the message.

// The message wasn't compressed

// Record decompression time metric

func (mb *msgBuilder) createOutbound(m *p2p.Message, compressionType compression.Type, bypassThrottling bool) (*OutboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mb *msgBuilder) parseInbound(
	bytes []byte,
	nodeID ids.NodeID,
	onFinishedHandling func(),
) (*InboundMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
