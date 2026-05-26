// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"

	"k8s.io/client-go/kubernetes"

	_ "embed"

	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	// TODO(marun) This should be configurable
	DefaultTmpnetNamespace = "tmpnet"

	KindKubeconfigContext = "kind-kind"

	// TODO(marun) Check for the presence of the context rather than string matching on this error
	missingContextMsg = `context "` + KindKubeconfigContext + `" does not exist`

	// Ingress controller constants
	ingressNamespace      = "ingress-nginx"
	ingressReleaseName    = "ingress-nginx"
	ingressChartRepo      = "https://kubernetes.github.io/ingress-nginx"
	ingressChartName      = "ingress-nginx/ingress-nginx"
	ingressControllerName = "ingress-nginx-controller"
	// This must match the nodePort configured in scripts/kind-with-registry.sh
	ingressNodePort = 30791

	// Chaos Mesh constants
	chaosMeshNamespace      = "chaos-mesh"
	chaosMeshReleaseName    = "chaos-mesh"
	chaosMeshChartRepo      = "https://charts.chaos-mesh.org"
	chaosMeshChartName      = "chaos-mesh/chaos-mesh"
	chaosMeshChartVersion   = "2.7.2"
	chaosMeshControllerName = "chaos-controller-manager"
	chaosMeshDashboardName  = "chaos-dashboard"
	chaosMeshDashboardHost  = "chaos-mesh.localhost"
)

//go:embed yaml/tmpnet-rbac.yaml
var tmpnetRBACManifest []byte

// StartKindCluster starts a new kind cluster with integrated registry if one is not already running.
func StartKindCluster(
	ctx context.Context,
	log logging.Logger,
	configPath string,
	startMetricsCollector bool,
	startLogsCollector bool,
	installChaosMesh bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Deploy RBAC resources for tmpnet

// Create service account kubeconfig context to enable checking that RBAC permissions are sufficient

// isKindClusterRunning determines if a kind cluster is running
func isKindClusterRunning(log logging.Logger, configPath string, configContext string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// All other errors are assumed fatal

// Assume any errors in discovery indicate the cluster is not running
//
// TODO(marun) Maybe differentiate between configuration and endpoint errors?

// ensureNamespace ensures that the specified namespace exists in cluster targeted by the clientset.
func ensureNamespace(ctx context.Context, log logging.Logger, clientset *kubernetes.Clientset, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// deployRBAC deploys the RBAC resources for tmpnet to a Kubernetes cluster.
func deployRBAC(
	ctx context.Context,
	log logging.Logger,
	configPath string,
	configContext string,
	namespace string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Apply the RBAC manifest

// createServiceAccountKubeconfig creates a kubeconfig that uses the tmpnet service account token.
// It only creates the context if it doesn't already exist.
// This function is called from StartKindCluster after the kubeconfig and context have been verified.
func createServiceAccountKubeconfig(
	ctx context.Context,
	log logging.Logger,
	configPath string,
	configContext string,
	namespace string,
	newContextName string,
) error {
	_ = "STUB: not implemented"
	// Get the existing kubeconfig
	return nil
}

// Get the current context (already verified to exist by StartKindCluster)

// Get clientset to retrieve service account token

// Create a token for the service account (Kubernetes 1.24+)

// Token will be valid for 1 year

// Create new context with the token

// Create new context

// Save the updated kubeconfig

// deployIngressController deploys the nginx ingress controller using Helm.
func deployIngressController(ctx context.Context, log logging.Logger, configPath string, configContext string) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the helm repo for ingress-nginx

// Install nginx-ingress with values set directly via flags
// Using fixed nodePort 30791 for cross-platform compatibility

// This port value must match the port configured in scripts/kind-with-registry.sh

// isIngressControllerRunning checks if the nginx ingress controller is already running.
func isIngressControllerRunning(ctx context.Context, log logging.Logger, configPath string, configContext string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// TODO(marun) Handle the case of the deployment being in a failed state

// runHelmCommand runs a Helm command with the given arguments.
func runHelmCommand(ctx context.Context, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// createDefaultsConfigMap creates a ConfigMap containing defaults for the tmpnet namespace.
func createDefaultsConfigMap(ctx context.Context, log logging.Logger, configPath string, configContext string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if configmap already exists

// deployChaosMesh deploys Chaos Mesh using Helm.
func deployChaosMesh(ctx context.Context, log logging.Logger, configPath string, configContext string) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the helm repo for chaos-mesh

// Install Chaos Mesh with all required settings including ingress

// Wait for Chaos Mesh to be ready

// Log access information

// isChaosMeshRunning checks if Chaos Mesh is already running.
func isChaosMeshRunning(ctx context.Context, log logging.Logger, configPath string, configContext string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if controller manager deployment exists

// waitForChaosMesh waits for Chaos Mesh components to be ready.
func waitForChaosMesh(ctx context.Context, log logging.Logger, configPath string, configContext string) error {
	_ = "STUB: not implemented"
	// Wait for controller manager
	return nil
}

// Wait for dashboard

// waitForDeployment waits for a deployment to have at least one ready replica.
func waitForDeployment(ctx context.Context, log logging.Logger, configPath string, configContext string, namespace string, deploymentName string, displayName string) error {
	_ = "STUB: not implemented"
	return nil
}

/* immediate */
