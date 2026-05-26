// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build !windows && !openbsd

package storage

import (
	"errors"
)

var errZeroAvailableBytes = errors.New("available blocks is reported as 0")

func AvailableBytes(storagePath string) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
