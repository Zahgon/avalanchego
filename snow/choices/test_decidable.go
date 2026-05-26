// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package choices

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
)

var _ Decidable = (*TestDecidable)(nil)

// TestDecidable is a test Decidable
type TestDecidable struct {
	IDV              ids.ID
	AcceptV, RejectV error
	StatusV          Status
}

func (d *TestDecidable) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (d *TestDecidable) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *TestDecidable) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *TestDecidable) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (d *TestDecidable) SetStatus(status Status) { _ = "STUB: not implemented"; return }
