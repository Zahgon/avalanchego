// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import "github.com/ava-labs/avalanchego/ids"

// readMainnetBonusBlocks returns maps of bonus block numbers to block IDs.
// Note bonus blocks are indexed in the atomic trie.
func readMainnetBonusBlocks() (map[uint64]ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
