// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snow

import (
	"errors"

	"github.com/ava-labs/avalanchego/proto/pb/p2p"
)

const (
	Initializing State = iota
	StateSyncing
	Bootstrapping
	NormalOp
)

var ErrUnknownState = errors.New("unknown state")

type State uint8

func (st State) String() string { _ = "STUB: not implemented"; return "" }

type EngineState struct {
	Type  p2p.EngineType
	State State
}
