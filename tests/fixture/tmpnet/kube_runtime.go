// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"errors"
	"net/netip"
	"time"

	"k8s.io/client-go/kubernetes"

	"github.com/ava-labs/avalanchego/utils/logging"

	corev1 "k8s.io/api/core/v1"
	restclient "k8s.io/client-go/rest"
)

// TODO(marun) need an easy way to cleanup stale nodes (either client side cli or a reaper)

const (
	containerName   = "avago"
	volumeName      = "data"
	volumeMountPath = "/data"

	statusCheckInterval = 500 * time.Millisecond

	// 2GB is the minimum size of a PersistentVolumeClaim used for a node's data directory:
	// - A value greater than 1GB must be used
	//   - A node will report unhealthy if it detects less than 1GiB available
	// - EBS volume sizes are in GB
	//   - The minimum number greater than 1GB is 2GB
	MinimumVolumeSizeGB = 2

	// All statefulsets configured for exclusive scheduling will use
	// anti-affinity with the following labeling to ensure their pods
	// are never scheduled to the same nodes.
	antiAffinityLabelKey   = "tmpnet-scheduling"
	antiAffinityLabelValue = "exclusive"

	// Name of config map containing tmpnet defaults
	defaultsConfigMapName = "tmpnet-defaults"
	ingressHostKey        = "ingressHost"
)

var (
	errMissingSchedulingLabels = errors.New("--kube-scheduling-label-key and --kube-scheduling-label-value are required when exclusive scheduling is enabled")
	errMissingIngressHost      = errors.New("IngressHost is a required value. Ensure the " + defaultsConfigMapName + " ConfigMap contains an entry for " + ingressHostKey)
)

type KubeRuntimeConfig struct {
	// Path to the kubeconfig file identifying the target cluster
	ConfigPath string `json:"configPath,omitempty"`
	// The context of the kubeconfig file to use
	ConfigContext string `json:"configContext,omitempty"`
	// Namespace in the target cluster in which resources will be
	// created. For simplicity all nodes are assumed to be deployed to
	// the same namespace to ensure network connectivity.
	Namespace string `json:"namespace,omitempty"`
	// The docker image to run for the node
	Image string `json:"image,omitempty"`
	// Size in gigabytes of the PersistentVolumeClaim  to allocate for the node
	VolumeSizeGB uint `json:"volumeSizeGB,omitempty"`
	// Whether to schedule each AvalancheGo node to a dedicated Kubernetes node
	UseExclusiveScheduling bool `json:"useExclusiveScheduling,omitempty"`
	// Label key to use for exclusive scheduling for node selection and toleration
	SchedulingLabelKey string `json:"schedulingLabelKey,omitempty"`
	// Label value to use for exclusive scheduling for node selection and toleration
	SchedulingLabelValue string `json:"schedulingLabelValue,omitempty"`
	// Host for ingress rules (e.g., "localhost:30791" for kind, "tmpnet.example.com" for EKS)
	IngressHost string `json:"ingressHost,omitempty"`
	// TLS secret name for ingress (empty for HTTP, populated for HTTPS)
	IngressSecret string `json:"ingressSecret,omitempty"`
}

// ensureDefaults sets cluster-specific defaults for fields not already set by flags.
func (c *KubeRuntimeConfig) ensureDefaults(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	// Only read defaults if necessary
	return nil
}

type KubeRuntime struct {
	node *Node

	kubeConfig *restclient.Config
}

// readState reads the URI and staking address for the node if the node is running.
func (p *KubeRuntime) readState(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Validate that it will be possible to construct accessible URIs when running external to the kube cluster

// GetAccessibleURI retrieves a URI for the node accessible from where
// this process is running. If the process is running inside a kube
// cluster, the node and the process will be assumed to be running in the
// same kube cluster and the node's URI be used. If the process is
// running outside of a kube cluster, a URI accessible from outside of
// the cluster will be used.
func (p *KubeRuntime) GetAccessibleURI() string { _ = "STUB: not implemented"; return "" }

// Assume tls is configured for an ingress secret

// GetAccessibleStakingAddress retrieves a StakingAddress for the node intended to be
// accessible from this process until the provided cancel function is called.
func (p *KubeRuntime) GetAccessibleStakingAddress(ctx context.Context) (netip.AddrPort, func(), error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil, nil
}

// Assume that an empty staking address indicates a need to retrieve pod state

// Use direct pod staking address if running inside the cluster

// Start the node as a Kubernetes StatefulSet.
func (p *KubeRuntime) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stateful exists - make sure it is scaled up and running

// StatefulSet does not exist - create it

// generateName

// imagePullPolicy - use default behavior

// If running outside the cluster, ensure the node's API port is accessible via ingress

// The 's-' prefix ensures DNS compatibility

// Stop the Pod by setting the replicas to zero on the StatefulSet.
func (p *KubeRuntime) InitiateStop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Waits for the node process to stop.
// TODO(marun) Consider using a watch instead
func (p *KubeRuntime) WaitForStopped(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// immediate

// Restarts the node. Does not wait for readiness or health.
func (p *KubeRuntime) Restart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(marun) Maybe optionally avoid restart if the patches will be no-op?

// Collect patches to apply to the StatefulSet

// Force a restart by scaling up and down

// immediate

// IsHealthy checks if the node is running and healthy.
func (p *KubeRuntime) IsHealthy(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ensureBootstrapIP waits for this pod to be ready if there are no other pods already
// running to ensure the availability of a bootstrap node.
func (p *KubeRuntime) ensureBootstrapIP(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Waits for the node's Pod to be ready, indicating that its API and
// staking endpoints are capable of serving traffic.
func (p *KubeRuntime) waitForPodReadiness(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Assume default ports. No reason to vary when Pods don't share port space.

// getStatefulSetName determines the name of the node's StatefulSet from the network UUID and node ID.
func (p *KubeRuntime) getStatefulSetName() string { _ = "STUB: not implemented"; return "" }

// The Pod name is the StatefulSet name with a suffix of "-0" to indicate the first Pod in the StatefulSet
func (p *KubeRuntime) getPodName() string { _ = "STUB: not implemented"; return "" }

func (p *KubeRuntime) runtimeConfig() *KubeRuntimeConfig { _ = "STUB: not implemented"; return nil }

// getKubeconfig retrieves the kubeconfig for the target cluster. It
// will be cached after the first call to avoid unnecessary logging
// when running in-cluster.
func (p *KubeRuntime) getKubeconfig() (*restclient.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(marun) Add a function that returns the kubeconfig and the clientset
func (p *KubeRuntime) getClientset() (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *KubeRuntime) forwardPort(ctx context.Context, port int) (uint16, chan struct{}, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Wait for the Pod to become ready (otherwise it won't be accepting network connections)

// Ignore stdout output

func (p *KubeRuntime) setNotRunning() { _ = "STUB: not implemented"; return }

// getFlags determines the set of avalanchego flags to configure the node with.
func (p *KubeRuntime) getFlags() (FlagsMap, error) {
	_ = "STUB: not implemented"
	return *new(FlagsMap), nil
}

// The data dir path is fixed for the Pod

// The node must bind to the Pod IP to enable the kubelet to access the http port for the readiness check

// Ensure compatibility with a non-localhost ingress host

// configureExclusiveScheduling ensures that the provided template schedules only to nodes with the provided
// labeling, tolerates a taint that matches the labeling, and uses anti-affinity to ensure only a single
// avalanchego pod is scheduled to a given target node.
func configureExclusiveScheduling(template *corev1.PodTemplateSpec, labelKey string, labelValue string) {
	_ = "STUB: not implemented"
	return
}

// Configure node selection

// Configure toleration. Nodes are assumed to have a taint with the same
// key+value as the label used to select it.

// Configure anti-affinity to ensure only one pod per node

// createNodeService creates a Kubernetes Service for the node to enable ingress routing
func (p *KubeRuntime) createNodeService(ctx context.Context, serviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// createNodeIngress creates a Kubernetes Ingress for the node to enable external access
func (p *KubeRuntime) createNodeIngress(ctx context.Context, serviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Assume nginx ingress controller
// Path pattern: /networks/<network-uuid>/<node-id>(/|$)(.*)
// Using (/|$)(.*) to properly handle trailing slashes

// Build the ingress rules

// Add host if not localhost

// Add TLS configuration if IngressSecret is set

// waitForIngressReadiness waits for the ingress to be ready and able to route traffic
// This prevents 503 errors when health checks are performed immediately after node start
func (p *KubeRuntime) waitForIngressReadiness(ctx context.Context, serviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the ingress to exist, be processed by the controller, and service endpoints to be available

// immediate

// Check if ingress exists and is processed by the controller

// Check if ingress controller has processed the ingress
// The ingress controller should populate the Status.LoadBalancer.Ingress field
// when it has successfully processed and exposed the ingress

// Validate that at least one ingress has an IP or hostname

// Check if service endpoints are available

// Check if endpoints have at least one ready address

// IsRunningInCluster detects if this code is running inside a Kubernetes cluster
// by checking for the presence of the service account token that's automatically
// mounted in every pod.
func IsRunningInCluster() bool { _ = "STUB: not implemented"; return false }
