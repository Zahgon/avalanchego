// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/go-cmd/cmd"
)

// RunCommand starts the command [bin] with the given [args] and returns the command to the caller
// TODO cmd package mentions we can do this more efficiently with cmd.NewCmdOptions rather than looping
// and calling Status().
func RunCommand(bin string, args ...string) (*cmd.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// to stream outputs

func RegisterPingTest() { _ = "STUB: not implemented"; return }

// RegisterNodeRun registers a before suite that starts an AvalancheGo process to use for the e2e tests
// and an after suite that stops the AvalancheGo process
func RegisterNodeRun() { _ = "STUB: not implemented"; return }

// BeforeSuite starts an AvalancheGo process to use for the e2e tests

// Assumes that startCmd will launch a node with HTTP Port at [utils.DefaultLocalNodeURI]

// TODO add a new node to bootstrap off of the existing node and ensure it can bootstrap all subnets
// created during the test
