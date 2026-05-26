// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils"
)

var _ Atomic = (*atomic)(nil)

type Atomic interface {
	common.AppHandler

	Set(common.AppHandler)
}

type atomic struct {
	handler utils.Atomic[common.AppHandler]
}

func NewAtomic(h common.AppHandler) Atomic { _ = "STUB: not implemented"; return *new(Atomic) }

func (a *atomic) AppRequest(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	deadline time.Time,
	msg []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *atomic) AppRequestFailed(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	appErr *common.AppError,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *atomic) AppResponse(
	ctx context.Context,
	nodeID ids.NodeID,
	requestID uint32,
	msg []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *atomic) AppGossip(
	ctx context.Context,
	nodeID ids.NodeID,
	msg []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *atomic) Set(h common.AppHandler) { _ = "STUB: not implemented"; return }
