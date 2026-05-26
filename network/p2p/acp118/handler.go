// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package acp118

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

const HandlerID = p2p.SignatureRequestHandlerID

var _ p2p.Handler = (*Handler)(nil)

// Verifier verifies that a warp message should be signed
type Verifier interface {
	// Verify verifies that the provided message is valid for this node to sign
	// based on the provided justification.
	//
	// Implementations of Verify are not expected to verify the NetworkID or
	// SourceChainID fields of the message. Verification of these fields should
	// be performed externally to Verify.
	Verify(
		ctx context.Context,
		message *warp.UnsignedMessage,
		justification []byte,
	) *common.AppError
}

// NewHandler returns an instance of Handler
func NewHandler(verifier Verifier, signer warp.Signer) *Handler {
	_ = "STUB: not implemented"
	return nil
}

// NewCachedHandler returns an instance of Handler that caches successful
// signatures.
func NewCachedHandler(
	cacher cache.Cacher[ids.ID, []byte],
	verifier Verifier,
	signer warp.Signer,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

// Handler signs warp messages
type Handler struct {
	p2p.NoOpHandler

	signatureCache cache.Cacher[ids.ID, []byte]
	verifier       Verifier
	signer         warp.Signer
}

func (h *Handler) AppRequest(
	ctx context.Context,
	_ ids.NodeID,
	_ time.Time,
	requestBytes []byte,
) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify that the payload is valid to sign.

// The signer internally verifies that the NetworkID and SourceChainID are
// populated with the expected values.

func signatureToResponse(signature []byte) ([]byte, *common.AppError) {
	_ = "STUB: not implemented"
	return nil, nil
}
