// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapmonitor

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/client-go/kubernetes"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/version"

	corev1 "k8s.io/api/core/v1"
)

// The sync mode of a bootstrap configuration should be as explicit as possible to ensure
// unambiguous names if/when new sync modes become supported.
//
// For example, at time of writing only the C-Chain supports state sync. The temptation
// might be to call this mode that tests C-Chain state sync `state-sync`. But if the P-
// or X-Chains start supporting state sync in the future, a name like
// `c-chain-state-sync` ensures that logs and metrics for historical results can be
// differentiated from new results that involve state sync of the other chains.
type SyncMode string

const (
	FullSync           SyncMode = "full-sync"
	CChainStateSync    SyncMode = "c-chain-state-sync"     // aka state sync
	OnlyPChainFullSync SyncMode = "p-chain-full-sync-only" // aka partial sync

	VersionsAnnotationKey = "avalanche.avax.network/avalanchego-versions"
)

var (
	chainConfigContentEnvName        = config.EnvVarName(config.EnvPrefix, config.ChainConfigContentKey)
	networkEnvName                   = config.EnvVarName(config.EnvPrefix, config.NetworkNameKey)
	partialSyncPrimaryNetworkEnvName = config.EnvVarName(config.EnvPrefix, config.PartialSyncPrimaryNetworkKey)

	// Errors for bootstrapTestConfigForPod
	errContainerNotFound          = errors.New("container not found")
	errInvalidNetworkEnvVar       = fmt.Errorf("missing or empty %s env var", networkEnvName)
	errFailedToUnmarshalAnnoation = errors.New("failed to unmarshal versions annotation")

	// Errors for stateSyncEnabledFromEnvVars
	errFailedToDecodeChainConfigContent    = errors.New("failed to decode chain config content")
	errFailedToUnmarshalChainConfigContent = errors.New("failed to unmarshal chain config content")
	errFailedToUnmarshalCChainConfig       = errors.New("failed to unmarshal C-Chain config")
	errFailedToCastToBool                  = errors.New("failed to cast to bool")
)

type BootstrapTestConfig struct {
	Network  string            `json:"network"`
	SyncMode SyncMode          `json:"syncMode"`
	Image    string            `json:"image"`
	Versions *version.Versions `json:"versions,omitempty"`
}

// GetBootstrapTestConfigFromPod extracts the bootstrap test configuration from the specified pod.
func GetBootstrapTestConfigFromPod(ctx context.Context, clientset *kubernetes.Clientset, namespace string, podName string, nodeContainerName string) (*BootstrapTestConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bootstrapTestConfigForPod collects the details for a bootstrap test configuration from the provided pod.
func bootstrapTestConfigForPod(pod *corev1.Pod, nodeContainerName string) (*BootstrapTestConfig, error) {
	_ = "STUB: not implemented"
	// Find the node container
	return nil, nil
}

// Get the network ID from the container's environment

// Determine the sync mode from the env vars

// Attempt to retrieve the image versions from a pod annotation. The annotation may not be populated in
// the case of a newly-created bootstrap test using an image tagged `master` that hasn't yet had a
// chance to discover the versions.

// syncModeFromEnvVars derives the bootstrap sync mode from the provided environment variables.
func syncModeFromEnvVars(env []corev1.EnvVar) (SyncMode, error) {
	_ = "STUB: not implemented"
	return *new(SyncMode), nil
}

// If partial sync is enabled, only the P-Chain will be synced so the state sync
// configuration of the C-Chain is irrelevant.

// Full sync is enabled

// C-Chain state sync is assumed if the other modes are not explicitly enabled

// partialSyncEnabledFromEnvVars determines whether the env vars configure partial sync
// for a node container. Partial sync is assumed to be enabled if the
// AVAGO_PARTIAL_SYNC_PRIMARY_NETWORK env var is set and evaluates to true.
func partialSyncEnabledFromEnvVars(env []corev1.EnvVar) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// stateSyncEnabledFromEnvVars determines whether the env vars configure state sync for a
// node container. State sync is assumed to be enabled if the chain config content is
// missing, does not contain C-Chain configuration, or the C-Chain configuration does not
// configure state-sync-enabled.
func stateSyncEnabledFromEnvVars(env []corev1.EnvVar) (bool, error) {
	_ = "STUB: not implemented"
	// Look for chain config content in the env vars
	return false, nil
}

// Attempt to unmarshal

// Attempt to unmarshal the C-Chain config

// Attempt to read the value from the C-Chain config
