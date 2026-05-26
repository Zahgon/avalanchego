// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapmonitor

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	contextDuration = 30 * time.Second

	ImageUnchanged = "Image unchanged"
)

var nodeURL = fmt.Sprintf("http://localhost:%d", config.DefaultHTTPPort)

func WaitForCompletion(
	log logging.Logger,
	namespace string,
	podName string,
	nodeContainerName string,
	dataDir string,
	healthCheckInterval time.Duration,
	imageCheckInterval time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Avoid checking node health before it reports initial ready

// Define common fields for logging

// Check whether the node is reporting healthy which indicates that bootstrap is complete

// Statefulset will restart the pod with the new image

// Avoid exiting immediately to avoid container restart before the pod is recreated with the new image

// Determines the current disk usage for the specified directory
func getDiskUsage(log logging.Logger, dir string) string { _ = "STUB: not implemented"; return "" }

// Create a buffer to capture stderr in case an unexpected error occurs

// Exit code 1 usually indicates that files cannot be accessed. Since avalanchego will
// regularly delete files in the db dir, this can be safely ignored and the regular disk
// usage message can be printed.
