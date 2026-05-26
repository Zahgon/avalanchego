// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/graft/evm/message"
)

var (
	_ message.ResponseHandler = (*waitingResponseHandler)(nil)

	errRequestFailed = errors.New("request failed")
)

// waitingResponseHandler implements the ResponseHandler interface
// Internally used to wait for response after making a request synchronously
// responseChan may contain response bytes if the original request has not failed
// responseChan is closed in either fail or success scenario
type waitingResponseHandler struct {
	responseChan chan []byte // blocking channel with response bytes
	failed       bool        // whether the original request is failed
}

// newWaitingResponseHandler returns new instance of the waitingResponseHandler
func newWaitingResponseHandler() *waitingResponseHandler { _ = "STUB: not implemented"; return nil }

// Make buffer length 1 so that OnResponse can complete
// even if no goroutine is waiting on the channel (i.e.
// the context of a request is cancelled.)

// OnResponse passes the response bytes to the responseChan and closes the channel
func (w *waitingResponseHandler) OnResponse(response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// OnFailure sets the failed flag to true and closes the channel
func (w *waitingResponseHandler) OnFailure() error { _ = "STUB: not implemented"; return nil }

func (w *waitingResponseHandler) WaitForResult(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
