// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saetest

import (
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/vms/saevm/types"
)

// A ClonableHeightIndex extends [database.HeightIndex] with the ability to
// clone itself.
type ClonableHeightIndex interface {
	database.HeightIndex
	Clone() ClonableHeightIndex
}

// NewHeightIndexDB returns an in-memory [database.HeightIndex]; its additional
// `Clone()` method can be called before or after closing, and the clone will
// not be closed in either circumstance. Only heights for which `Sync()` has
// returned without error will be cloned.
func NewHeightIndexDB() ClonableHeightIndex {
	_ = "STUB: not implemented"
	return *new(ClonableHeightIndex)
}

// NewExecutionResultsDB wraps and returns a [NewHeightIndexDB].
func NewExecutionResultsDB() types.ExecutionResults {
	_ = "STUB: not implemented"
	return *new(types.ExecutionResults)
}

type hIndex struct {
	mu      sync.RWMutex
	pending map[uint64]bool
	data    map[uint64][]byte
	closed  bool
}

func readHIndex[T any](h *hIndex, fn func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (h *hIndex) write(fn func() error) error { _ = "STUB: not implemented"; return nil }

func (h *hIndex) Put(n uint64, b []byte) error { _ = "STUB: not implemented"; return nil }

func (h *hIndex) Get(n uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *hIndex) Clone() ClonableHeightIndex {
	_ = "STUB: not implemented"
	return *new(ClonableHeightIndex)
}

//nolint:forcetypeassert // Internal invariant

func (h *hIndex) Has(n uint64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (h *hIndex) Sync(from, to uint64) error { _ = "STUB: not implemented"; return nil }

func (h *hIndex) Close() error { _ = "STUB: not implemented"; return nil }
