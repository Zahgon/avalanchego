// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package enginetest

import (
	"testing"
	"time"
)

// Timer is a test timer
type Timer struct {
	T *testing.T

	CantRegisterTimout bool

	RegisterTimeoutF func(time.Duration)
}

func (t *Timer) RegisterTimeout(delay time.Duration) { _ = "STUB: not implemented"; return }
