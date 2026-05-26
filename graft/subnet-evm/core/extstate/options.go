// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extstate

import (
	"github.com/ava-labs/libevm/core/state"

	"github.com/ava-labs/avalanchego/graft/evm/utils"
)

type workerPool struct {
	*utils.BoundedWorkers
}

func (wp *workerPool) Done() {
	_ = "STUB: not implemented"
	// Done is guaranteed to only be called after all work is already complete,
	// so we call Wait for goroutines to finish before returning.
	return
}

func WithConcurrentWorkers(prefetchers int) state.PrefetcherOption {
	_ = "STUB: not implemented"
	return *new(state.PrefetcherOption)
}
