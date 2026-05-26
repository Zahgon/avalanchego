// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subprocess

import (
	"context"
	"os/exec"
	"sync"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"
)

func NewStopper(logger logging.Logger, cmd *exec.Cmd) runtime.Stopper {
	_ = "STUB: not implemented"
	return *new(runtime.Stopper)
}

type stopper struct {
	once   sync.Once
	cmd    *exec.Cmd
	logger logging.Logger
}

func (s *stopper) Stop(ctx context.Context) { _ = "STUB: not implemented"; return }
