// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build linux || netbsd || openbsd

package ulimit

import (
	"github.com/ava-labs/avalanchego/utils/logging"
)

const DefaultFDLimit = 32 * 1024

// Set attempts to bump the Rlimit which has a soft (Cur) and a hard (Max) value.
// The soft limit is what is used by the kernel to report EMFILE errors. The hard
// limit is a secondary limit which the process can be bumped to without additional
// privileges. Bumping the Max limit further would require superuser privileges.
// If the current Max is below our recommendation we will warn on start.
// see: http://0pointer.net/blog/file-descriptor-limits.html
func Set(limit uint64, log logging.Logger) error { _ = "STUB: not implemented"; return nil }

// set new limit

// verify limit
