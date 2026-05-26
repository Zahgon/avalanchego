// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package summary

import "github.com/ava-labs/avalanchego/ids"

var _ StateSummary = (*stateSummary)(nil)

type StateSummary interface {
	ID() ids.ID
	ForkHeight() uint64
	BlockBytes() []byte
	InnerSummaryBytes() []byte
	Bytes() []byte
}

type stateSummary struct {
	Height uint64 `serialize:"true"`
	// TODO: Rather than storing the full block here - we should only store
	//       proposervm information. We would then modify the StateSummary
	//       interface to expose the required information to generate the full
	//       block.
	Block        []byte `serialize:"true"`
	InnerSummary []byte `serialize:"true"`

	id    ids.ID
	bytes []byte
}

func (s *stateSummary) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (s *stateSummary) ForkHeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *stateSummary) BlockBytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *stateSummary) InnerSummaryBytes() []byte { _ = "STUB: not implemented"; return nil }

func (s *stateSummary) Bytes() []byte { _ = "STUB: not implemented"; return nil }
