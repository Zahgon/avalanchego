// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utilstest

import (
	"context"
	"sync"
	"testing"
	"time"
)

func WaitGroupWithContext(t *testing.T, ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// include context error for easier debugging

func SleepWithContext(ctx context.Context, d time.Duration) { _ = "STUB: not implemented"; return }
