// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package syncer

import (
	"context"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/database/merkle/sync"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

var _ sync.DB[*RangeProof, struct{}] = (*evmDB)(nil)

type CodeQueue interface {
	AddCode(context.Context, []common.Hash) error
}

type evmDB struct {
	db        *database
	codeQueue CodeQueue
}

func NewEVM(
	config Config,
	db *ffi.Database,
	codeQueue CodeQueue,
	targetRoot ids.ID,
	proofClient *p2p.Client,
) (*sync.Syncer[*RangeProof, struct{}], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *evmDB) CommitRangeProof(ctx context.Context, start, end maybe.Maybe[[]byte], proof *RangeProof) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	// First enqueue any code hashes found in the proof.
	// If an error occurs here, we don't want to commit the proof to the database, because the
	// code hashes will never be requested.
	return nil, nil
}

func (*evmDB) CommitChangeProof(context.Context, maybe.Maybe[[]byte], struct{}) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *evmDB) GetMerkleRoot(ctx context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (e *evmDB) Clear() error { _ = "STUB: not implemented"; return nil }

func (e *evmDB) GetChangeProof(ctx context.Context, startRootID ids.ID, endRootID ids.ID, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], maxLength int) (struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *evmDB) GetRangeProofAtRoot(ctx context.Context, rootID ids.ID, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], maxLength int) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *evmDB) VerifyChangeProof(ctx context.Context, proof struct{}, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], expectedEndRootID ids.ID, maxLength int) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *evmDB) VerifyRangeProof(ctx context.Context, proof *RangeProof, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], expectedEndRootID ids.ID, maxLength int) error {
	_ = "STUB: not implemented"
	return nil
}
