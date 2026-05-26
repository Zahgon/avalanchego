// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ethapi

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"

	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/evm/rpc"
)

type DetailedExecutionResult struct {
	UsedGas    uint64        `json:"gas"`        // Total used gas but include the refunded gas
	ErrCode    int           `json:"errCode"`    // EVM error code
	Err        string        `json:"err"`        // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData hexutil.Bytes `json:"returnData"` // Data from evm(function result or data supplied with revert opcode)
}

// GetChainConfig returns the chain config.
func (api *BlockChainAPI) GetChainConfig(context.Context) *params.ChainConfig {
	_ = "STUB: not implemented"
	return nil

	// CallDetailed performs the same call as Call, but returns the full context
}

func (s *BlockChainAPI) CallDetailed(ctx context.Context, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *StateOverride) (*DetailedExecutionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the result contains a revert reason, try to unpack and return it.

// Note: this API is moved directly from ./eth/api.go to ensure that it is available under an API that is enabled by
// default without duplicating the code and serving the same API in the original location as well without creating a
// cyclic import.
//
// BadBlockArgs represents the entries in the list returned when bad blocks are queried.
type BadBlockArgs struct {
	Hash   common.Hash            `json:"hash"`
	Block  map[string]interface{} `json:"block"`
	RLP    string                 `json:"rlp"`
	Reason *core.BadBlockReason   `json:"reason"`
}

// GetBadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
// and returns them as a JSON list of block hashes.
func (s *BlockChainAPI) GetBadBlocks(context.Context) ([]*BadBlockArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hacky, but hey, it works

// stateQueryBlockNumberAllowed returns a nil error if:
//   - the node is configured to accept any state query (the query window is zero)
//   - the block given has its number within the query window before the last accepted block.
//     This query window is set to [core.TipBufferSize] when running in a non-archive mode.
//
// Otherwise, it returns a non-nil error containing block number information.
func (s *BlockChainAPI) stateQueryBlockNumberAllowed(blockNumOrHash rpc.BlockNumberOrHash) (err error) {
	_ = "STUB: not implemented"
	return nil
}
