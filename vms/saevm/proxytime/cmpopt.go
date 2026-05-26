// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !prod && !nocmpopts

package proxytime

import (
	"github.com/google/go-cmp/cmp"
)

// CmpOpt returns a configuration for [cmp.Diff] to compare [Time] instances in
// tests. The option will only be applied to the specific [Duration] type.
func CmpOpt[D Duration]() cmp.Option { _ = "STUB: not implemented"; return *new(cmp.Option) }
