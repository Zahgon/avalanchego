// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"

	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/trace"
)

var _ snowstorm.Tx = (*tracedTx)(nil)

type tracedTx struct {
	snowstorm.Tx

	tracer trace.Tracer
}

func (t *tracedTx) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *tracedTx) Accept(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *tracedTx) Reject(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
