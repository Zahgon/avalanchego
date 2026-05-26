// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow/engine/common"
)

var (
	_ common.BootstrapableEngine = (*Bootstrapper)(nil)

	errClear = errors.New("unexpectedly called Clear")
)

type Bootstrapper struct {
	Engine

	CantClear bool

	ClearF func(ctx context.Context) error
}

func (b *Bootstrapper) Default(cant bool) { _ = "STUB: not implemented"; return }

func (b *Bootstrapper) Clear(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
