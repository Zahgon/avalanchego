// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !prod && !nocmpopts

package blocks

import (
	"github.com/google/go-cmp/cmp"
)

// CmpOpt returns a configuration for [cmp.Diff] to compare [Block] instances in
// tests.
func CmpOpt() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }

func (e *executionResults) equalForTests(f *executionResults) bool {
	_ = "STUB: not implemented"
	return false
}
