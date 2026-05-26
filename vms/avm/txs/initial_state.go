// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	ErrNilInitialState  = errors.New("nil initial state is not valid")
	ErrNilFxOutput      = errors.New("nil feature extension output is not valid")
	ErrOutputsNotSorted = errors.New("outputs not sorted")
	ErrUnknownFx        = errors.New("unknown feature extension")

	_ utils.Sortable[*InitialState] = (*InitialState)(nil)
)

type InitialState struct {
	FxIndex uint32         `serialize:"true"  json:"fxIndex"`
	FxID    ids.ID         `serialize:"false" json:"fxID"`
	Outs    []verify.State `serialize:"true"  json:"outputs"`
}

func (is *InitialState) InitCtx(ctx *snow.Context) { _ = "STUB: not implemented"; return }

func (is *InitialState) Verify(c codec.Manager, numFxs int) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *InitialState) Compare(other *InitialState) int { _ = "STUB: not implemented"; return 0 }

func (is *InitialState) Sort(c codec.Manager) { _ = "STUB: not implemented"; return }

type innerSortState struct {
	vers  []verify.State
	codec codec.Manager
}

func (vers *innerSortState) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (vers *innerSortState) Len() int { _ = "STUB: not implemented"; return 0 }

func (vers *innerSortState) Swap(i, j int) { _ = "STUB: not implemented"; return }

func sortState(vers []verify.State, c codec.Manager) { _ = "STUB: not implemented"; return }

func isSortedState(vers []verify.State, c codec.Manager) bool {
	_ = "STUB: not implemented"
	return false
}
