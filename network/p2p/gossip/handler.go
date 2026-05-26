// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ p2p.Handler = (*Handler[Gossipable])(nil)

// HandlerSet exposes the ability to add new values to the set in response to
// pushed information and for responding to pull requests.
//
// TODO: Consider naming this interface based on what it provides rather than
// how its used.
type HandlerSet[T Gossipable] interface {
	// Add adds a value to the set. Returns an error if v was not added.
	Add(v T) error
	// Iterate iterates over elements until f returns false.
	Iterate(f func(v T) bool)
}

func NewHandler[T Gossipable](
	log logging.Logger,
	marshaller Marshaller[T],
	set HandlerSet[T],
	metrics Metrics,
	targetResponseSize int,
) *Handler[T] {
	_ = "STUB: not implemented"
	return nil
}

type Handler[T Gossipable] struct {
	p2p.Handler
	marshaller         Marshaller[T]
	log                logging.Logger
	set                HandlerSet[T]
	metrics            Metrics
	targetResponseSize int
}

func (h Handler[T]) AppRequest(_ context.Context, _ ids.NodeID, _ time.Time, requestBytes []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter out what the requesting peer already knows about

// check that this doesn't exceed our maximum configured target response
// size

func (h Handler[_]) AppGossip(_ context.Context, nodeID ids.NodeID, gossipBytes []byte) {
	_ = "STUB: not implemented"
	return
}
