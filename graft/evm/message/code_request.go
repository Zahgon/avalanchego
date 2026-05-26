// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"context"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/ids"
)

var _ Request = CodeRequest{}

// CodeRequest is a request to retrieve a contract code with specified Hash
type CodeRequest struct {
	// Hashes is a list of contract code hashes
	Hashes []common.Hash `serialize:"true"`
}

func (c CodeRequest) String() string { _ = "STUB: not implemented"; return "" }

func (c CodeRequest) Handle(ctx context.Context, nodeID ids.NodeID, requestID uint32, handler RequestHandler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCodeRequest(hashes []common.Hash) CodeRequest {
	_ = "STUB: not implemented"
	return *new(CodeRequest)
}

// CodeResponse is a response to a CodeRequest
// crypto.Keccak256Hash of each element in Data is expected to equal
// the corresponding element in CodeRequest.Hashes
// handler: handlers.CodeRequestHandler
type CodeResponse struct {
	Data [][]byte `serialize:"true"`
}
