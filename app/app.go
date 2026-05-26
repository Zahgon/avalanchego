// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package app

import (
	"sync"

	"github.com/ava-labs/avalanchego/node"
	"github.com/ava-labs/avalanchego/utils/logging"

	nodeconfig "github.com/ava-labs/avalanchego/config/node"
)

const Header = `     _____               .__                       .__
    /  _  \___  _______  |  | _____    ____   ____ |  |__   ____    ,_ o
   /  /_\  \  \/ /\__  \ |  | \__  \  /    \_/ ___\|  |  \_/ __ \   / //\,
  /    |    \   /  / __ \|  |__/ __ \|   |  \  \___|   Y  \  ___/    \>> |
  \____|__  /\_/  (____  /____(____  /___|  /\___  >___|  /\___  >    \\
          \/           \/          \/     \/     \/     \/     \/`

var _ App = (*app)(nil)

type App interface {
	// Start kicks off the application and returns immediately.
	// Start should only be called once.
	Start()

	// Stop notifies the application to exit and returns immediately.
	// Stop should only be called after [Start].
	// It is safe to call Stop multiple times.
	Stop()

	// ExitCode should only be called after [Start] returns. It
	// should block until the application finishes
	ExitCode() int
}

func New(config nodeconfig.Config) (App, error) {
	_ = "STUB: not implemented"
	// Set the data directory permissions to be read write.
	return *new(App), nil
}

// update fd limit

func Run(app App) int {
	_ = "STUB: not implemented"
	// start running the application
	return 0
}

// register terminationSignals to kill the application

// start up a new go routine to handle attempts to kill the application

// start a goroutine to listen on SIGABRT signals,
// to print the stack trace to standard error.

// wait for the app to exit and get the exit code response

// shut down the termination signal go routine

// shut down the stack trace go routine

// return the exit code that the application reported

// app is a wrapper around a node that runs in this process
type app struct {
	node       *node.Node
	log        logging.Logger
	logFactory logging.Factory
	exitWG     sync.WaitGroup
}

// Start the business logic of the node (as opposed to config reading, etc).
// Does not block until the node is done.
func (a *app) Start() {
	_ = "STUB: not implemented"
	// [p.ExitCode] will block until [p.exitWG.Done] is called
	return
}

// If [p.node.Dispatch()] panics, then we should log the panic and
// then re-raise the panic. This is why the above defer is broken
// into two parts.

// Stop attempts to shutdown the currently running node. This function will
// block until Shutdown returns.
func (a *app) Stop() {
	_ = "STUB: not implemented"

	// ExitCode returns the exit code that the node is reporting. This function
	// blocks until the node has been shut down.
	return
}

func (a *app) ExitCode() int { _ = "STUB: not implemented"; return 0 }
