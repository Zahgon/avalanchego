// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

const (
	// maxNumParents is the max number of parents a vertex may have
	maxNumParents = 128

	// maxTxsPerVtx is the max number of transactions a vertex may have
	maxTxsPerVtx = 128
)

var (
	errBadVersion       = errors.New("invalid version")
	errBadEpoch         = errors.New("invalid epoch")
	errTooManyParentIDs = fmt.Errorf("vertex contains more than %d parentIDs", maxNumParents)
	errNoOperations     = errors.New("vertex contains no operations")
	errTooManyTxs       = fmt.Errorf("vertex contains more than %d transactions", maxTxsPerVtx)
	errInvalidParents   = errors.New("vertex contains non-sorted or duplicated parentIDs")
	errInvalidTxs       = errors.New("vertex contains non-sorted or duplicated transactions")

	_ StatelessVertex = statelessVertex{}
)

type StatelessVertex interface {
	verify.Verifiable
	ID() ids.ID
	Bytes() []byte
	Version() uint16
	ChainID() ids.ID
	StopVertex() bool
	Height() uint64
	Epoch() uint32
	ParentIDs() []ids.ID
	Txs() [][]byte
}

type statelessVertex struct {
	// This wrapper exists so that the function calls aren't ambiguous
	innerStatelessVertex

	// cache the ID of this vertex
	id ids.ID

	// cache the binary format of this vertex
	bytes []byte
}

func (v statelessVertex) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (v statelessVertex) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (v statelessVertex) Version() uint16 { _ = "STUB: not implemented"; return 0 }

func (v statelessVertex) ChainID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (v statelessVertex) StopVertex() bool { _ = "STUB: not implemented"; return false }

func (v statelessVertex) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (v statelessVertex) Epoch() uint32 { _ = "STUB: not implemented"; return 0 }

func (v statelessVertex) ParentIDs() []ids.ID { _ = "STUB: not implemented"; return nil }

func (v statelessVertex) Txs() [][]byte { _ = "STUB: not implemented"; return nil }

type innerStatelessVertex struct {
	Version   uint16   `json:"version"`
	ChainID   ids.ID   `json:"chainID"   serializeV0:"true" serializeV1:"true"`
	Height    uint64   `json:"height"    serializeV0:"true" serializeV1:"true"`
	Epoch     uint32   `json:"epoch"     serializeV0:"true"`
	ParentIDs []ids.ID `json:"parentIDs" serializeV0:"true" serializeV1:"true"`
	Txs       [][]byte `json:"txs"       serializeV0:"true"`
}

func (v innerStatelessVertex) Verify() error { _ = "STUB: not implemented"; return nil }

func (v innerStatelessVertex) verify() error { _ = "STUB: not implemented"; return nil }

func (v innerStatelessVertex) verifyStopVertex() error { _ = "STUB: not implemented"; return nil }
