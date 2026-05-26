// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build linux

// ^ SIGTERM signal is not available on Windows
// ^ syscall.SysProcAttr only has field Pdeathsig on Linux

package subprocess

import (
	"context"
	"os/exec"

	"github.com/ava-labs/avalanchego/utils/logging"
)

func NewCmd(path string, args ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

func stop(ctx context.Context, log logging.Logger, cmd *exec.Cmd) {
	_ = "STUB: not implemented"
	return
}

// attempt graceful shutdown

// force kill
