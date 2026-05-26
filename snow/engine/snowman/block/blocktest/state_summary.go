// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocktest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ block.StateSummary = (*StateSummary)(nil)

	errAccept = errors.New("unexpectedly called Accept")
)

type StateSummary struct {
	IDV     ids.ID
	HeightV uint64
	BytesV  []byte

	T          *testing.T
	CantAccept bool
	AcceptF    func(context.Context) (block.StateSyncMode, error)
}

func (s *StateSummary) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *StateSummary) Height() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StateSummary) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *StateSummary) Accept(ctx context.Context) (block.StateSyncMode, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSyncMode), nil
}
