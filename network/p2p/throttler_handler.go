// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ Handler = (*ThrottlerHandler)(nil)

func NewThrottlerHandler(handler Handler, throttler Throttler, log logging.Logger) *ThrottlerHandler {
	_ = "STUB: not implemented"
	return nil
}

type ThrottlerHandler struct {
	handler   Handler
	throttler Throttler
	log       logging.Logger
}

func (t ThrottlerHandler) AppGossip(ctx context.Context, nodeID ids.NodeID, gossipBytes []byte) {
	_ = "STUB: not implemented"
	return
}

func (t ThrottlerHandler) AppRequest(ctx context.Context, nodeID ids.NodeID, deadline time.Time, requestBytes []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}
