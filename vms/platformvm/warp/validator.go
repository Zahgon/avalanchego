// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	ErrUnknownValidator = errors.New("unknown validator")
	ErrWeightOverflow   = errors.New("weight overflowed")
)

// ValidatorState defines the functions that must be implemented to get
// the canonical validator set for warp message validation.
type ValidatorState interface {
	GetValidatorSet(ctx context.Context, height uint64, subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error)
}

type (
	// Deprecated: use [validators.WarpSet] instead.
	CanonicalValidatorSet = validators.WarpSet
	// Deprecated: use [validators.Warp] instead.
	Validator = validators.Warp
)

// Deprecated: use [validators.FlattenValidatorSet] instead.
var FlattenValidatorSet = validators.FlattenValidatorSet

// GetCanonicalValidatorSetFromSubnetID returns the CanonicalValidatorSet of [subnetID] at
// [pChainHeight]. The returned CanonicalValidatorSet includes the validator set in a canonical ordering
// and the total weight.
//
// Deprecated: Use [validators.State.GetWarpValidatorSet] instead.
func GetCanonicalValidatorSetFromSubnetID(
	ctx context.Context,
	pChainState ValidatorState,
	pChainHeight uint64,
	subnetID ids.ID,
) (validators.WarpSet, error) {
	_ = "STUB: not implemented"
	// Get the validator set at the given height.
	return *new(validators.WarpSet), nil
}

// Convert the validator set into the canonical ordering.

// FilterValidators returns the validators in [vdrs] whose bit is set to 1 in
// [indices].
//
// Returns an error if [indices] references an unknown validator.
func FilterValidators(
	indices set.Bits,
	vdrs []*validators.Warp,
) ([]*validators.Warp, error) {
	_ = "STUB: not implemented"
	// Verify that all alleged signers exist
	return nil, nil
}

// -1 to convert from length to index

// SumWeight returns the total weight of the provided validators.
func SumWeight(vdrs []*validators.Warp) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// AggregatePublicKeys returns the public key of the provided validators.
//
// Invariant: All of the public keys in [vdrs] are valid.
func AggregatePublicKeys(vdrs []*validators.Warp) (*bls.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCanonicalValidatorSetFromChainID returns the canonical validator set given a validators.State, pChain height and a sourceChainID.
func GetCanonicalValidatorSetFromChainID(ctx context.Context,
	pChainState validators.State,
	pChainHeight uint64,
	sourceChainID ids.ID,
) (validators.WarpSet, error) {
	_ = "STUB: not implemented"
	return *new(validators.WarpSet), nil
}
