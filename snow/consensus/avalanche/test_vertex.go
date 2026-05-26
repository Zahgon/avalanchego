// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"context"

	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
)

var _ Vertex = (*TestVertex)(nil)

// TestVertex is a useful test vertex
type TestVertex struct {
	choices.TestDecidable

	ParentsV    []Vertex
	ParentsErrV error
	HeightV     uint64
	HeightErrV  error
	TxsV        []snowstorm.Tx
	TxsErrV     error
	BytesV      []byte
}

func (v *TestVertex) Parents() ([]Vertex, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *TestVertex) Height() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *TestVertex) Txs(context.Context) ([]snowstorm.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *TestVertex) Bytes() []byte { _ = "STUB: not implemented"; return nil }
