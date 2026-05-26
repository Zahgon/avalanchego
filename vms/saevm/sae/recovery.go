// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"context"
	"iter"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/params"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/saexec"
	"github.com/ava-labs/avalanchego/vms/saevm/types"
)

type recovery struct {
	db              ethdb.Database
	xdb             types.ExecutionResults
	chainConfig     *params.ChainConfig
	log             logging.Logger
	hooks           hook.Points
	config          Config
	lastSynchronous *blocks.Block
}

func (rec *recovery) newCanonicalBlock(num uint64, parent *blocks.Block) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rec *recovery) lastCommittedBlock() (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rec *recovery) canonicalAfter(parent *blocks.Block) iter.Seq2[*blocks.Block, error] {
	_ = "STUB: not implemented"
	return nil
}

func (rec *recovery) executeAllAccepted(ctx context.Context, exec *saexec.Executor) error {
	_ = "STUB: not implemented"
	return nil
}

// Consensus only requires post-execution state after and including the
// last-settled block.

// lastOf returns the lastOf element in a slice, which MUST NOT be empty.
func lastOf[E any](s []E) E {
	_ = "STUB: not implemented"

	// consensusCriticalBlocks returns a block-hash-keyed map of all blocks from the
	// last executed back to, and including, the block that it settled. Said settled
	// block is returned separately, for convenience.
	return *new(E)
}

func (rec *recovery) consensusCriticalBlocks(exec *saexec.Executor) (_ *syncMap[common.Hash, *blocks.Block], lastSettled *blocks.Block, _ error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// reverse height order

// extend appends to the chain all the blocks in settler's ancestry up to
// and including the block that it settled.

// The post-execution root is tracked by the [saexec.Executor] as
// soon as it's known. In the case of database recovery, this
// occurred in [recovery.executeAllAccepted].

// i.e. deleted due to settlement not rejection
