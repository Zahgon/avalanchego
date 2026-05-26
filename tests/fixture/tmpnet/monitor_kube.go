// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	_ "embed"

	"github.com/ava-labs/avalanchego/utils/logging"
)

//go:embed yaml/promtail-daemonset.yaml
var promtailManifest []byte

//go:embed yaml/prometheus-agent.yaml
var prometheusManifest []byte

// This must match the namespace defined in the manifests
const monitoringNamespace = "ci-monitoring"

// Configuration for a kube-hosted collector
type kubeCollectorConfig struct {
	name         string
	target       string
	secretPrefix string
	manifest     []byte
}

// deployKubeCollectors deploys collectors of logs and metrics to a Kubernetes cluster.
func deployKubeCollectors(
	ctx context.Context,
	log logging.Logger,
	configPath string,
	configContext string,
	startMetricsCollector bool,
	startLogsCollector bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to do

// deployKubeCollector deploys a named collector to a Kubernetes cluster via the provided manifest bytes.
func deployKubeCollector(
	ctx context.Context,
	log logging.Logger,
	clientset *kubernetes.Clientset,
	dynamicClient dynamic.Interface,
	kubeConfig kubeCollectorConfig,
) error {
	_ = "STUB: not implemented"
	// Source the collector url and auth creds from the environment
	return nil
}

// createCollectorConfigSecret creates a secret with the provided collector config
func createCollectorConfigSecret(
	ctx context.Context,
	log logging.Logger,
	clientset *kubernetes.Clientset,
	namePrefix string,
	config collectorConfig,
) error {
	_ = "STUB: not implemented"
	return nil
}
