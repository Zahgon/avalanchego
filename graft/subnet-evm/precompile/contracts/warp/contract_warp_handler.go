// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

var (
	_ messageHandler = addressedPayloadHandler{}
	_ messageHandler = blockHashHandler{}
)

var (
	getVerifiedWarpMessageInvalidOutput   []byte
	getVerifiedWarpBlockHashInvalidOutput []byte
)

func init() {
	res, err := PackGetVerifiedWarpMessageOutput(GetVerifiedWarpMessageOutput{Valid: false})
	if err != nil {
		panic(err)
	}
	getVerifiedWarpMessageInvalidOutput = res

	res, err = PackGetVerifiedWarpBlockHashOutput(GetVerifiedWarpBlockHashOutput{Valid: false})
	if err != nil {
		panic(err)
	}
	getVerifiedWarpBlockHashInvalidOutput = res
}

type messageHandler interface {
	packFailed() []byte
	handleMessage(msg *warp.Message) ([]byte, error)
}

func handleWarpMessage(accessibleState contract.AccessibleState, input []byte, suppliedGas uint64, handler messageHandler) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// This conversion is safe even if int is 32 bits because we checked above.

// Note: we charge for the size of the message during both predicate verification and each time the message is read during
// EVM execution because each execution incurs an additional read cost.

// Note: since the predicate is verified in advance of execution, the precompile should not
// hit an error during execution.

type addressedPayloadHandler struct{}

func (addressedPayloadHandler) packFailed() []byte { _ = "STUB: not implemented"; return nil }

func (addressedPayloadHandler) handleMessage(warpMessage *warp.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type blockHashHandler struct{}

func (blockHashHandler) packFailed() []byte { _ = "STUB: not implemented"; return nil }

func (blockHashHandler) handleMessage(warpMessage *warp.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
