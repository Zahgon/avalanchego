// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapper

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/set"
)

var errUnexpectedSamplerFailure = errors.New("unexpected sampler failure")

// Sample keys from [elements] uniformly by weight without replacement. The
// returned set will have size less than or equal to [maxSize]. This function
// will error if the sum of all weights overflows.
func Sample[T comparable](elements map[T]uint64, maxSize int) (set.Set[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
