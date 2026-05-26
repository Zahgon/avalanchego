// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build darwin

package ulimit

import (
	"github.com/ava-labs/avalanchego/utils/logging"
)

const DefaultFDLimit = 10 * 1024

// Set attempts to bump the Rlimit which has a soft (Cur) and a hard (Max) value.
// The soft limit is what is used by the kernel to report EMFILE errors. The hard
// limit is a secondary limit which the process can be bumped to without additional
// privileges. Bumping the Max limit further would require superuser privileges.
// If the value is below the recommendation warn on start.
// see: http://0pointer.net/blog/file-descriptor-limits.html
func Set(limit uint64, log logging.Logger) error { _ = "STUB: not implemented"; return nil }

// Darwin max number of FDs to allocatable for darwin systems.
// The max file limit is 10240, even though the max returned by
// Getrlimit is 1<<63-1. This is OPEN_MAX in sys/syslimits.h.
// See https://github.com/golang/go/issues/30401

// set new limit

// verify limit
