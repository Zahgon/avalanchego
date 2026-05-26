// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !prod && !nocmpopts

package gastime

import (
	"github.com/google/go-cmp/cmp"
)

// CmpOpt returns a configuration for [cmp.Diff] to compare [Time] instances in
// tests.
func CmpOpt() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }
