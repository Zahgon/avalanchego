// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorstest

import (
	"testing"

	"github.com/ava-labs/avalanchego/snow/validators"
)

func NewWarp(t testing.TB, weight uint64) *validators.Warp { _ = "STUB: not implemented"; return nil }

func NewWarpSet(t testing.TB, n uint64) validators.WarpSet {
	_ = "STUB: not implemented"
	return *new(validators.WarpSet)
}

func WarpToOutput(w *validators.Warp) *validators.GetValidatorOutput {
	_ = "STUB: not implemented"
	return nil
}
