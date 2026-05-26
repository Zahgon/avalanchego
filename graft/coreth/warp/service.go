// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/common/hexutil"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/acp118"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

var (
	errNoValidators               = errors.New("cannot aggregate signatures from subnet with no validators")
	errCannotRetrieveValidatorSet = errors.New("cannot retrieve validator set")
)

// API introduces snowman specific functionality to the evm
type API struct {
	chainContext        *snow.Context
	backend             Backend
	signatureAggregator *acp118.SignatureAggregator
}

func NewAPI(chainCtx *snow.Context, backend Backend, signatureAggregator *acp118.SignatureAggregator) *API {
	_ = "STUB: not implemented"
	return nil
}

// GetMessage returns the Warp message associated with a messageID.
func (a *API) GetMessage(_ context.Context, messageID ids.ID) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetMessageSignature returns the BLS signature associated with a messageID.
func (a *API) GetMessageSignature(ctx context.Context, messageID ids.ID) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetBlockSignature returns the BLS signature associated with a blockID.
func (a *API) GetBlockSignature(ctx context.Context, blockID ids.ID) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetMessageAggregateSignature fetches the aggregate signature for the requested [messageID]
func (a *API) GetMessageAggregateSignature(ctx context.Context, messageID ids.ID, quorumNum uint64, subnetIDStr string) (signedMessageBytes hexutil.Bytes, err error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// GetBlockAggregateSignature fetches the aggregate signature for the requested [blockID]
func (a *API) GetBlockAggregateSignature(ctx context.Context, blockID ids.ID, quorumNum uint64, subnetIDStr string) (signedMessageBytes hexutil.Bytes, err error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

func (a *API) aggregateSignatures(ctx context.Context, unsignedMessage *warp.UnsignedMessage, quorumNum uint64, subnetIDStr string) (hexutil.Bytes, error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// TODO: return the signature and total weight as well to the caller for more complete details
// Need to decide on the best UI for this and write up documentation with the potential
// gotchas that could impact signed messages becoming invalid.
