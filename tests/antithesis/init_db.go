// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package antithesis

import (
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

// Given a path, compose the expected path of the bootstrap node's docker compose db volume.
func getBootstrapVolumePath(targetPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Bootstraps a local process-based network, creates its subnets and chains, and copies
// the resulting db state from one of the nodes to the provided path. The path will be
// created if it does not already exist.
func initBootstrapDB(network *tmpnet.Network, destPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Since the goal is to initialize the DB, we can stop the network after it has been started successfully

// Copy the db state from the bootstrap node to the compose volume path.

// TODO(marun) Replace with os.CopyFS once we upgrade to Go 1.23
