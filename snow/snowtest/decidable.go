// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowtest

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var (
	_ snow.Decidable = (*Decidable)(nil)

	ErrInvalidStateTransition = errors.New("invalid state transition")
)

type Decidable struct {
	IDV     ids.ID
	AcceptV error
	RejectV error
	Status  Status
}

func (d *Decidable) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (d *Decidable) Accept(context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *Decidable) Reject(context.Context) error { _ = "STUB: not implemented"; return nil }
