// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dbtest

import (
	"testing"

	"github.com/ava-labs/avalanchego/database"
)

// Tests is a list of all database tests
var Tests = []struct {
	Name string
	Test func(t *testing.T, newDB func() database.HeightIndex)
}{
	{"TestPutGet", TestPutGet},
	{"TestHas", TestHas},
	{"TestSync", TestSync},
	{"TestCloseAndPut", TestCloseAndPut},
	{"TestCloseAndGet", TestCloseAndGet},
	{"TestCloseAndHas", TestCloseAndHas},
	{"TestCloseAndSync", TestCloseAndSync},
	{"TestClose", TestClose},
}

type putArgs struct {
	height uint64
	data   []byte
}

func TestPutGet(t *testing.T, newDB func() database.HeightIndex) { _ = "STUB: not implemented"; return }

func TestHas(t *testing.T, newDB func() database.HeightIndex) { _ = "STUB: not implemented"; return }

// Perform all puts

func TestCloseAndPut(t *testing.T, newDB func() database.HeightIndex) {
	_ = "STUB: not implemented"
	return
}

// Try to put after close - should return error

func TestCloseAndGet(t *testing.T, newDB func() database.HeightIndex) {
	_ = "STUB: not implemented"
	return
}

// Try to get after close - should return error

func TestCloseAndHas(t *testing.T, newDB func() database.HeightIndex) {
	_ = "STUB: not implemented"
	return
}

// Try to has after close - should return error

func TestCloseAndSync(t *testing.T, newDB func() database.HeightIndex) {
	_ = "STUB: not implemented"
	return
}

// Try to sync after close - should return error

func TestClose(t *testing.T, newDB func() database.HeightIndex) { _ = "STUB: not implemented"; return }

// Second close should return error

func TestSync(t *testing.T, newDB func() database.HeightIndex) { _ = "STUB: not implemented"; return }
