// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cachetest

import (
	"testing"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
)

const IntSize = ids.IDLen + 8

func IntSizeFunc(ids.ID, int64) int {
	_ = "STUB: not implemented"

	// Tests is a list of all Cacher tests
	return 0
}

var Tests = []struct {
	Size int
	Func func(t *testing.T, c cache.Cacher[ids.ID, int64])
}{
	{Size: 1, Func: Basic},
	{Size: 2, Func: Eviction},
}

func Basic(t *testing.T, cache cache.Cacher[ids.ID, int64]) { _ = "STUB: not implemented"; return }

func Eviction(t *testing.T, cache cache.Cacher[ids.ID, int64]) { _ = "STUB: not implemented"; return }
