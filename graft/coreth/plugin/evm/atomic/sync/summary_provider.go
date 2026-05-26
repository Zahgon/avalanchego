// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic/state"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

var _ message.SyncSummaryProvider = (*SummaryProvider)(nil)

// SummaryProvider provides and parses state sync summaries for the atomic trie.
type SummaryProvider struct {
	trie *state.AtomicTrie
}

// Initialize initializes the summary provider with the atomic trie.
func (a *SummaryProvider) Initialize(trie *state.AtomicTrie) {
	_ = "STUB: not implemented"

	// StateSummaryAtBlock returns the block state summary at [blk] if valid.
	return
}

func (a *SummaryProvider) StateSummaryAtBlock(blk *ethtypes.Block) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// Parse parses the summary bytes into a Syncable summary.
func (*SummaryProvider) Parse(summaryBytes []byte, acceptImpl message.AcceptImplFn) (message.Syncable, error) {
	_ = "STUB: not implemented"
	return *new(message.Syncable), nil
}
