// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocktest

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	_ block.StateSyncableVM = (*StateSyncableVM)(nil)

	errStateSyncEnabled           = errors.New("unexpectedly called StateSyncEnabled")
	errStateSyncGetOngoingSummary = errors.New("unexpectedly called StateSyncGetOngoingSummary")
	errGetLastStateSummary        = errors.New("unexpectedly called GetLastStateSummary")
	errParseStateSummary          = errors.New("unexpectedly called ParseStateSummary")
	errGetStateSummary            = errors.New("unexpectedly called GetStateSummary")
)

type StateSyncableVM struct {
	T *testing.T

	CantStateSyncEnabled,
	CantStateSyncGetOngoingSummary,
	CantGetLastStateSummary,
	CantParseStateSummary,
	CantGetStateSummary bool

	StateSyncEnabledF           func(context.Context) (bool, error)
	GetOngoingSyncStateSummaryF func(context.Context) (block.StateSummary, error)
	GetLastStateSummaryF        func(context.Context) (block.StateSummary, error)
	ParseStateSummaryF          func(ctx context.Context, summaryBytes []byte) (block.StateSummary, error)
	GetStateSummaryF            func(ctx context.Context, summaryHeight uint64) (block.StateSummary, error)
}

func (vm *StateSyncableVM) StateSyncEnabled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (vm *StateSyncableVM) GetOngoingSyncStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *StateSyncableVM) GetLastStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *StateSyncableVM) ParseStateSummary(ctx context.Context, summaryBytes []byte) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (vm *StateSyncableVM) GetStateSummary(ctx context.Context, summaryHeight uint64) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}
