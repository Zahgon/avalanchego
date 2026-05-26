// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package acp118

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

var errFailedVerification = errors.New("failed verification")

type indexedValidator struct {
	*validators.Warp
	Index int
}

type result struct {
	NodeID    ids.NodeID
	Validator indexedValidator
	Signature *bls.Signature
	Err       error
}

// NewSignatureAggregator returns an instance of SignatureAggregator
func NewSignatureAggregator(log logging.Logger, client *p2p.Client) *SignatureAggregator {
	_ = "STUB: not implemented"
	return nil
}

// SignatureAggregator aggregates validator signatures for warp messages
type SignatureAggregator struct {
	log    logging.Logger
	client *p2p.Client
}

// AggregateSignatures blocks until quorumNum/quorumDen signatures from
// validators are requested to be aggregated into a warp message or the context
// is canceled. Returns the signed message and the amount of stake that signed
// the message. Caller is responsible for providing a well-formed canonical
// validator set corresponding to the signer bitset in the message.
func (s *SignatureAggregator) AggregateSignatures(
	ctx context.Context,
	message *warp.Message,
	justification []byte,
	validators []*validators.Warp,
	quorumNum uint64,
	quorumDen uint64,
) (
	_ *warp.Message,
	aggregatedStake *big.Int,
	totalStake *big.Int,
	_ error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// TODO expose concrete type to avoid type casting

// Only try to aggregate signatures from validators that are not already in
// the signer bit set

// Account for requested signatures + the signature that was provided

// Block until:
// 1. The context is cancelled
// 2. We get responses from all validators
// 3. The specified security threshold is reached

// Try to return whatever progress we have if the context is cancelled

// Validators may share public keys so drop any duplicate signatures

func newWarpMessage(
	message *warp.Message,
	signerBitSet set.Bits,
	signatures []*bls.Signature,
) (*warp.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type responseHandler struct {
	message             *warp.Message
	nodeIDsToValidators map[ids.NodeID]indexedValidator
	results             chan result
}

func (r *responseHandler) HandleResponse(
	_ context.Context,
	nodeID ids.NodeID,
	responseBytes []byte,
	err error,
) {
	_ = "STUB: not implemented"
	return
}
