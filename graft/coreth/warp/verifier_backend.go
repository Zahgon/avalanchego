// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"

	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

const (
	ParseErrCode = iota + 1
	VerifyErrCode
)

// Verify verifies the signature of the message
// It also implements the acp118.Verifier interface
func (b *backend) Verify(ctx context.Context, unsignedMessage *avalancheWarp.UnsignedMessage, _ []byte) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}

// Known on-chain messages should be signed

// verifyBlockMessage returns nil if blockHashPayload contains the ID
// of an accepted block indicating it should be signed by the VM.
func (b *backend) verifyBlockMessage(ctx context.Context, blockHashPayload *payload.Hash) *common.AppError {
	_ = "STUB: not implemented"
	return nil
}
