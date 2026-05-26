// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

const (
	lastAcceptedByte byte = iota
)

var (
	lastAcceptedKey = []byte{lastAcceptedByte}

	_ ChainState = (*chainState)(nil)
)

type ChainState interface {
	SetLastAccepted(blkID ids.ID) error
	DeleteLastAccepted() error
	GetLastAccepted() (ids.ID, error)
}

type chainState struct {
	lastAccepted ids.ID
	db           database.Database
}

func NewChainState(db database.Database) ChainState {
	_ = "STUB: not implemented"
	return *new(ChainState)
}

func (s *chainState) SetLastAccepted(blkID ids.ID) error { _ = "STUB: not implemented"; return nil }

func (s *chainState) DeleteLastAccepted() error { _ = "STUB: not implemented"; return nil }

func (s *chainState) GetLastAccepted() (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
