// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package verify

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var (
	ErrSameChainID         = errors.New("same chainID")
	ErrMismatchedSubnetIDs = errors.New("mismatched subnetIDs")
)

// SameSubnet verifies that the provided [ctx] was provided to a chain in the
// same subnet as [peerChainID], but not the same chain. If this verification
// fails, a non-nil error will be returned.
func SameSubnet(ctx context.Context, chainCtx *snow.Context, peerChainID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}
