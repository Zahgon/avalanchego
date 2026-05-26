// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package statstest

import (
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers/stats"
)

var _ stats.HandlerStats = (*TestHandlerStats)(nil)

// TestHandlerStats is test for capturing and asserting on handler metrics in test
type TestHandlerStats struct {
	lock sync.Mutex

	BlockRequestCount,
	MissingBlockHashCount,
	BlocksReturnedSum uint32
	BlockRequestProcessingTimeSum time.Duration

	CodeRequestCount,
	MissingCodeHashCount,
	TooManyHashesRequested,
	DuplicateHashesRequested,
	CodeBytesReturnedSum uint32
	CodeReadTimeSum time.Duration

	LeafsRequestCount,
	InvalidLeafsRequestCount,
	LeafsReturnedSum,
	MissingRootCount,
	TrieErrorCount,
	ProofErrorCount,
	SnapshotReadErrorCount,
	SnapshotReadAttemptCount,
	SnapshotReadSuccessCount,
	SnapshotSegmentValidCount,
	SnapshotSegmentInvalidCount uint32
	ProofValsReturned int64
	LeafsReadTime,
	SnapshotReadTime,
	GenerateRangeProofTime,
	LeafRequestProcessingTimeSum time.Duration
}

func (t *TestHandlerStats) Reset() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncBlockRequest() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncMissingBlockHash() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) UpdateBlocksReturned(num uint16) { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) UpdateBlockRequestProcessingTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) IncCodeRequest() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncMissingCodeHash() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncTooManyHashesRequested() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncDuplicateHashesRequested() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) UpdateCodeReadTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) UpdateCodeBytesReturned(bytes uint32) { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncLeafsRequest() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncInvalidLeafsRequest() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) UpdateLeafsReturned(numLeafs uint16) { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) UpdateLeafsRequestProcessingTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) UpdateReadLeafsTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) UpdateGenerateRangeProofTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) UpdateSnapshotReadTime(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) UpdateRangeProofValsReturned(numProofVals int64) {
	_ = "STUB: not implemented"
	return
}

func (t *TestHandlerStats) IncMissingRoot() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncTrieError() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncProofError() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncSnapshotReadError() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncSnapshotReadAttempt() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncSnapshotReadSuccess() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncSnapshotSegmentValid() { _ = "STUB: not implemented"; return }

func (t *TestHandlerStats) IncSnapshotSegmentInvalid() { _ = "STUB: not implemented"; return }
