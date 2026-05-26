// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build freebsd

package ulimit

import (
	"github.com/ava-labs/avalanchego/utils/logging"
)

const DefaultFDLimit = 32 * 1024

// Set attempts to bump the Rlimit which has a soft (Cur) and a hard (Max) value.
// The soft limit is what is used by the kernel to report EMFILE errors. The hard
// limit is a secondary limit which the process can be bumped to without additional
// privileges. Bumping the Max limit further would require superuser privileges.
// If the value is below the recommendation warn on start.
// see: http://0pointer.net/blog/file-descriptor-limits.html
func Set(limit uint64, log logging.Logger) error {
	_ = "STUB: not implemented"
	// Note: BSD Rlimit is type int64
	// ref: https://cs.opensource.google/go/x/sys/+/b874c991:unix/ztypes_freebsd_amd64.go
	return nil
}

// set new limit

// verify limit
