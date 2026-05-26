// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapmonitor

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/version"
)

// Path to write the details to on the data volume
func getTestDetailsPath(dataDir string) string { _ = "STUB: not implemented"; return "" }

// Used to serialize test details to the data volume used for a given test to
// support resuming a previously started test and tracking test duration.
type bootstrapTestDetails struct {
	Image     string    `json:"image"`
	StartTime time.Time `json:"startTime"`
}

// setImageDetails updates the pod's owning statefulset with the image of the specified container and associated version details
func setImageDetails(ctx context.Context, log logging.Logger, clientset *kubernetes.Clientset, namespace string, podName string, imageDetails *ImageDetails) error {
	_ = "STUB: not implemented"
	// Determine the name of the statefulset to update
	return nil
}

// Marshal the versions to JSON

// Create the JSON patch

// Convert patch data to JSON

// Apply the patch

// getBaseImageName removes the tag from the image name
func getBaseImageName(log logging.Logger, imageName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Image name contains a digest, remove it

// No tag or registry

// Ambiguous image name - could contain a tag or a registry

// Image name contains a registry and a tag - remove the tag

type ImageDetails struct {
	Image    string
	Versions *version.Versions
}

// getMasterImageDetails retrieves the image details for the avalanchego image with tag `master`.
func getMasterImageDetails(
	ctx context.Context,
	log logging.Logger,
	clientset *kubernetes.Clientset,
	namespace string,
	imageName string,
	containerName string,
) (*ImageDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start a new pod with the `master`-tagged avalanchego image to discover its image ID

// Ensure the latest image is always pulled for a tag other than `latest`

// Get the image id for the avalanchego image

// Get the logs for the pod

// Attempt to unmarshal the logs to a Versions instance

// Only delete the pod if successful to aid in debugging

func getClientset(log logging.Logger) (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
