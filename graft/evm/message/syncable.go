// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	errParseSyncableSummary = errors.New("cannot parse syncable summary")
	errComputeSummaryID     = errors.New("cannot compute summary ID")
)

type Syncable interface {
	block.StateSummary
	GetBlockHash() common.Hash
	GetBlockRoot() common.Hash
}

// SyncSummaryProvider provides and parses state sync summaries.
type SyncSummaryProvider interface {
	// StateSummaryAtBlock returns the state summary for the given block.
	StateSummaryAtBlock(ethBlock *types.Block) (block.StateSummary, error)
	// Parse parses a syncable summary from bytes.
	Parse(summaryBytes []byte, acceptImpl AcceptImplFn) (Syncable, error)
}

type AcceptImplFn func(Syncable) (block.StateSyncMode, error)

// ParseSyncableSummary unmarshals `summaryBytes` into `summary` and returns its ID.
func ParseSyncableSummary(c codec.Manager, summaryBytes []byte, summary any) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
