// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapmonitor

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	initTimeout = 2 * time.Minute

	BootstrapStartingMessage = "Starting bootstrap test"
	BootstrapResumingMessage = "Resuming bootstrap test"
)

func NodeDataDir(path string) string { _ = "STUB: not implemented"; return "" }

func InitBootstrapTest(log logging.Logger, namespace string, podName string, nodeContainerName string, dataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the image uses the master tag, determine the master image id and set the container image to that

// A bootstrap is being resumed if a version file exists and the image name it contains matches the container
// image. If a bootstrap is being started, the version file should be created and the data path cleared.
