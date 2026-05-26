// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"io"
	"time"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	restclient "k8s.io/client-go/rest"
)

// NewNodeStatefulSet returns a statefulset for an avalanchego node.
func NewNodeStatefulSet(
	name string,
	generateName bool,
	imageName string,
	imagePullPolicy corev1.PullPolicy,
	containerName string,
	volumeName string,
	volumeSize string,
	volumeMountPath string,
	flags FlagsMap,
	labels map[string]string,
) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

// The following sets of annotations can coexist safely since each
// collection agent only reads the annotations it recognizes.

// Configure collection by prometheus and promtail

// Configure collection by the grafana cloud agent

// These labels may contain values invalid for use in labels. Set them as annotations instead.

// stringMapToEnvVarSlice converts a string map to a kube EnvVar slice.
func flagsToEnvVarSlice(flags FlagsMap) []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

func envVarsToJSONValue(envVars []corev1.EnvVar) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func sortEnvVars(envVars []corev1.EnvVar) { _ = "STUB: not implemented"; return }

// WaitForNodeHealthy waits for the node running in the specified pod to report healthy.
func WaitForNodeHealthy(
	ctx context.Context,
	log logging.Logger,
	kubeconfig *restclient.Config,
	namespace string,
	podName string,
	healthCheckInterval time.Duration,
	out io.Writer,
	outErr io.Writer,
) (ids.NodeID, error) {
	_ = "STUB: not implemented"
	// A forwarded connection enables connectivity without exposing the node external to the kube cluster
	return *new(ids.NodeID), nil
}

// TODO(marun) A node started with tmpnet should know the node ID before start

// Error is potentially recoverable - log and continue

// WaitForPodCondition watches the specified pod until the status includes the specified condition.
func WaitForPodCondition(ctx context.Context, clientset *kubernetes.Clientset, namespace string, podName string, conditionType corev1.PodConditionType) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPodStatus watches the specified pod until the status is deemed acceptable by the provided test function.
func WaitForPodStatus(
	ctx context.Context,
	clientset *kubernetes.Clientset,
	namespace string,
	name string,
	acceptable func(*corev1.PodStatus) bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// enableLocalForwardForPod enables traffic forwarding from a local port to the specified pod with client-go. The returned
// stop channel should be closed to stop the port forwarding.
func enableLocalForwardForPod(
	kubeconfig *restclient.Config,
	namespace string,
	name string,
	port int,
	out, errOut io.Writer,
) (uint16, chan struct{}, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// TODO(marun) Need better error handling here

// Wait for port forwarding to be ready

// Retrieve the dynamically allocated local port

// GetClientConfig replicates the behavior of clientcmd.BuildConfigFromFlags with zap logging and
// support for an optional config context. If path is not provided, use of in-cluster config will
// be attempted.
func GetClientConfig(log logging.Logger, path string, context string) (*restclient.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClientset returns a kubernetes clientset for the provided kubeconfig path and context.
func GetClientset(log logging.Logger, path string, context string) (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// applyManifest creates or updates the resources defined by the provided manifest using server-side apply.
// If namespace is empty, the namespace from the manifest will be used for namespaced resources.
func applyManifest(
	ctx context.Context,
	log logging.Logger,
	dynamicClient dynamic.Interface,
	manifest []byte,
	namespace string,
) error {
	_ = "STUB: not implemented"
	// Split the manifest into individual resources
	return nil
}

// Determine namespace for the resource

// Use namespace from the manifest if not provided

// Convert object to JSON for server-side apply

// Use server-side apply to create or update the resource

// TODO(marun) Check that the resources are running and healthy
