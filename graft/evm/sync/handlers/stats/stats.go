// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stats

import (
	"time"

	"github.com/ava-labs/libevm/metrics"
)

// HandlerStats reports prometheus metrics for the state sync handlers
type HandlerStats interface {
	BlockRequestHandlerStats
	CodeRequestHandlerStats
	LeafsRequestHandlerStats
}

type BlockRequestHandlerStats interface {
	IncBlockRequest()
	IncMissingBlockHash()
	UpdateBlocksReturned(num uint16)
	UpdateBlockRequestProcessingTime(duration time.Duration)
}

type CodeRequestHandlerStats interface {
	IncCodeRequest()
	IncMissingCodeHash()
	IncTooManyHashesRequested()
	IncDuplicateHashesRequested()
	UpdateCodeReadTime(duration time.Duration)
	UpdateCodeBytesReturned(bytes uint32)
}

type LeafsRequestHandlerStats interface {
	IncLeafsRequest()
	IncInvalidLeafsRequest()
	UpdateLeafsReturned(numLeafs uint16)
	UpdateLeafsRequestProcessingTime(duration time.Duration)
	UpdateReadLeafsTime(duration time.Duration)
	UpdateSnapshotReadTime(duration time.Duration)
	UpdateGenerateRangeProofTime(duration time.Duration)
	UpdateRangeProofValsReturned(numProofVals int64)
	IncMissingRoot()
	IncTrieError()
	IncProofError()
	IncSnapshotReadError()
	IncSnapshotReadAttempt()
	IncSnapshotReadSuccess()
	IncSnapshotSegmentValid()
	IncSnapshotSegmentInvalid()
}

type handlerStats struct {
	// BlockRequestHandler metrics
	blockRequest               metrics.Counter
	missingBlockHash           metrics.Counter
	blocksReturned             metrics.Histogram
	blockRequestProcessingTime metrics.Timer

	// CodeRequestHandler stats
	codeRequest              metrics.Counter
	missingCodeHash          metrics.Counter
	tooManyHashesRequested   metrics.Counter
	duplicateHashesRequested metrics.Counter
	codeBytesReturned        metrics.Histogram
	codeReadDuration         metrics.Timer

	// LeafsRequestHandler stats
	leafsRequest               metrics.Counter
	invalidLeafsRequest        metrics.Counter
	leafsReturned              metrics.Histogram
	leafsRequestProcessingTime metrics.Timer
	leafsReadTime              metrics.Timer
	snapshotReadTime           metrics.Timer
	generateRangeProofTime     metrics.Timer
	proofValsReturned          metrics.Histogram
	missingRoot                metrics.Counter
	trieError                  metrics.Counter
	proofError                 metrics.Counter
	snapshotReadError          metrics.Counter
	snapshotReadAttempt        metrics.Counter
	snapshotReadSuccess        metrics.Counter
	snapshotSegmentValid       metrics.Counter
	snapshotSegmentInvalid     metrics.Counter
}

func (h *handlerStats) IncBlockRequest() { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncMissingBlockHash() { _ = "STUB: not implemented"; return }

func (h *handlerStats) UpdateBlocksReturned(num uint16) { _ = "STUB: not implemented"; return }

func (h *handlerStats) UpdateBlockRequestProcessingTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) IncCodeRequest() { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncMissingCodeHash() { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncTooManyHashesRequested() { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncDuplicateHashesRequested() { _ = "STUB: not implemented"; return }

func (h *handlerStats) UpdateCodeReadTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) UpdateCodeBytesReturned(bytesLen uint32) { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncLeafsRequest() { _ = "STUB: not implemented"; return }

func (h *handlerStats) IncInvalidLeafsRequest() { _ = "STUB: not implemented"; return }

func (h *handlerStats) UpdateLeafsRequestProcessingTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) UpdateLeafsReturned(numLeafs uint16) { _ = "STUB: not implemented"; return }

func (h *handlerStats) UpdateReadLeafsTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) UpdateSnapshotReadTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) UpdateGenerateRangeProofTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) UpdateRangeProofValsReturned(numProofVals int64) {
	_ = "STUB: not implemented"
	return
}

func (h *handlerStats) IncMissingRoot()            { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncTrieError()              { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncProofError()             { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncSnapshotReadError()      { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncSnapshotReadAttempt()    { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncSnapshotReadSuccess()    { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncSnapshotSegmentValid()   { _ = "STUB: not implemented"; return }
func (h *handlerStats) IncSnapshotSegmentInvalid() { _ = "STUB: not implemented"; return }

// GetOrRegisterHandlerStats returns a [HandlerStats] to track state sync handler metrics.
// If `enabled` is false, a no-op implementation is returned.
// if `enabled` is true, calling this multiple times will return the same registered metrics.
func GetOrRegisterHandlerStats(enabled bool) HandlerStats {
	_ = "STUB: not implemented"
	return *new(HandlerStats)
}

// initialize block request stats

// initialize code request stats

// initialize leafs request stats

// no op implementation
type noopHandlerStats struct{}

func NewNoopHandlerStats() HandlerStats { _ = "STUB: not implemented"; return *new(HandlerStats) }

// all operations are no-ops
func (*noopHandlerStats) IncBlockRequest()            { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncMissingBlockHash()        { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateBlocksReturned(uint16) { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateBlockRequestProcessingTime(time.Duration) {
	_ = "STUB: not implemented"
	return
}
func (*noopHandlerStats) IncCodeRequest()                  { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncMissingCodeHash()              { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncTooManyHashesRequested()       { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncDuplicateHashesRequested()     { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateCodeReadTime(time.Duration) { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateCodeBytesReturned(uint32)   { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncLeafsRequest()                 { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncInvalidLeafsRequest()          { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateLeafsRequestProcessingTime(time.Duration) {
	_ = "STUB: not implemented"
	return
}
func (*noopHandlerStats) UpdateLeafsReturned(uint16)             { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateReadLeafsTime(_ time.Duration)    { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateSnapshotReadTime(_ time.Duration) { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) UpdateGenerateRangeProofTime(_ time.Duration) {
	_ = "STUB: not implemented"
	return
}
func (*noopHandlerStats) UpdateRangeProofValsReturned(_ int64) { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncMissingRoot()                      { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncTrieError()                        { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncProofError()                       { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncSnapshotReadError()                { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncSnapshotReadAttempt()              { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncSnapshotReadSuccess()              { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncSnapshotSegmentValid()             { _ = "STUB: not implemented"; return }
func (*noopHandlerStats) IncSnapshotSegmentInvalid()           { _ = "STUB: not implemented"; return }
