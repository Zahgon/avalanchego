// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subprocess

import (
	"context"
	"io"
	"net"
	"os/exec"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"
)

type Config struct {
	// Stderr of the VM process written to this writer.
	Stderr io.Writer
	// Stdout of the VM process written to this writer.
	Stdout io.Writer
	// Duration engine server will wait for handshake success.
	HandshakeTimeout time.Duration
	Log              logging.Logger
}

type Status struct {
	// Id of the process.
	Pid int
	// Address of the VM gRPC service.
	Addr string
}

// Bootstrap starts a VM as a subprocess after initialization completes and
// pipes the IO to the appropriate writers.
//
// The subprocess is expected to be stopped by the caller if a non-nil error is
// returned. If piping the IO fails then the subprocess will be stopped.
//
// TODO: create the listener inside this method once we refactor the tests
func Bootstrap(
	ctx context.Context,
	listener net.Listener,
	cmd *exec.Cmd,
	config *Config,
) (*Status, runtime.Stopper, error) {
	_ = "STUB: not implemented"
	return nil, *new(runtime.Stopper), nil
}

// pass golang debug env to subprocess

// start subproccess

// start stdout collector

// start stderr collector

// wait for handshake success
