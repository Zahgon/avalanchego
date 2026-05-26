// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package syncer

import (
	"context"

	"github.com/ava-labs/firewood-go-ethhash/ffi"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database/merkle/sync"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

var (
	_ sync.DB[*RangeProof, struct{}] = (*database)(nil)

	defaultSimultaneousWorkLimit = 8
)

// database wraps a Firewood [ffi.Database] to implement the xsync.DB interface.
type database struct {
	db *ffi.Database
}

type Config struct {
	SimultaneousWorkLimit int
	Log                   logging.Logger
	StateSyncNodes        []ids.NodeID
	Registerer            prometheus.Registerer
}

func New(config Config, db *ffi.Database, targetRoot ids.ID, proofClient *p2p.Client) (*sync.Syncer[*RangeProof, struct{}], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWithDB(config Config, db sync.DB[*RangeProof, struct{}], targetRoot ids.ID, proofClient *p2p.Client) (*sync.Syncer[*RangeProof, struct{}], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *database) GetMerkleRoot(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (db *database) GetRangeProofAtRoot(_ context.Context, rootID ids.ID, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], maxLength int) (*RangeProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*database) VerifyRangeProof(_ context.Context, proof *RangeProof, start maybe.Maybe[[]byte], end maybe.Maybe[[]byte], expectedEndRootID ids.ID, maxLength int) error {
	_ = "STUB: not implemented"
	// Extra data must be provided at commit time.
	// TODO: remove this once the FFI no longer requires it.
	return nil
}

func (db *database) CommitRangeProof(_ context.Context, start, end maybe.Maybe[[]byte], proof *RangeProof) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No error indicates the range is complete.

// TODO: This will eventually be handled by `FindNextKey`.

// TODO: Use change proofs to optimize syncing.
// Returning the sentinel error suggests to the server handler to serve a full range proof instead.
func (*database) GetChangeProof(context.Context, ids.ID, ids.ID, maybe.Maybe[[]byte], maybe.Maybe[[]byte], int) (struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: implement this method.
func (*database) VerifyChangeProof(context.Context, struct{}, maybe.Maybe[[]byte], maybe.Maybe[[]byte], ids.ID, int) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: implement this method.
func (*database) CommitChangeProof(context.Context, maybe.Maybe[[]byte], struct{}) (maybe.Maybe[[]byte], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *database) Clear() error {
	_ = "STUB: not implemented"
	// Prefix delete key of length 0.
	return nil
}
