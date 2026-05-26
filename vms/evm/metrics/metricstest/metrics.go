// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metricstest

import (
	"sync"
	"testing"
)

var metricsLock sync.Mutex

// WithMetrics enables [metrics.Enabled] for the test and prevents any other
// tests with metrics from running concurrently.
//
// [metrics.Enabled] is restored to its original value during testing cleanup.
func WithMetrics(t testing.TB) { _ = "STUB: not implemented"; return }
