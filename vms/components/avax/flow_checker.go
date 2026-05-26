// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type FlowChecker struct {
	consumed, produced map[ids.ID]uint64
	errs               wrappers.Errs
}

func NewFlowChecker() *FlowChecker { _ = "STUB: not implemented"; return nil }

func (fc *FlowChecker) Consume(assetID ids.ID, amount uint64) { _ = "STUB: not implemented"; return }

func (fc *FlowChecker) Produce(assetID ids.ID, amount uint64) { _ = "STUB: not implemented"; return }

func (fc *FlowChecker) add(value map[ids.ID]uint64, assetID ids.ID, amount uint64) {
	_ = "STUB: not implemented"
	return
}

func (fc *FlowChecker) Verify() error { _ = "STUB: not implemented"; return nil }
