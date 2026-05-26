// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	collectorTickerInterval = 1 * time.Second

	// TODO(marun) Maybe use dynamic HTTP ports to avoid the possibility of them being already bound?

	// Prometheus configuration
	prometheusCmd            = "prometheus"
	prometheusScrapeInterval = 10 * time.Second
	prometheusListenAddress  = "127.0.0.1:9090"
	prometheusReadinessURL   = "http://" + prometheusListenAddress + "/-/ready"

	// Promtail configuration
	promtailCmd          = "promtail"
	promtailHTTPPort     = "3101"
	promtailReadinessURL = "http://127.0.0.1:" + promtailHTTPPort + "/ready"

	// Use a delay slightly longer than the scrape interval to ensure a final scrape before shutdown
	NetworkShutdownDelay = prometheusScrapeInterval + 2*time.Second
)

// StartPrometheus ensures prometheus is running to collect metrics from local nodes.
func StartPrometheus(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// StartPromtail ensures promtail is running to collect logs from local nodes.
func StartPromtail(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPromtailReadiness waits until prometheus is ready. It can only succeed after
// one or more nodes have written their service discovery configuration.
func WaitForPromtailReadiness(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// StopMetricsCollector ensures prometheus is not running.
func StopMetricsCollector(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// StopLogsCollector ensures promtail is not running.
func StopLogsCollector(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// stopCollector stops the collector process if it is running.
func stopCollector(ctx context.Context, log logging.Logger, cmdName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Determine if the process is running

// Process is no longer running

// Attempt to clear the PID file. Not critical that it is removed, just good housekeeping.

// startPrometheus ensures an agent-mode prometheus process is running to collect metrics from local nodes.
func startPrometheus(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// startPromtail ensures a promtail process is running to collect logs from local nodes.
func startPromtail(ctx context.Context, log logging.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func getWorkingDir(cmdName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetPrometheusServiceDiscoveryDir returns the path for prometheus file-based
// service discovery configuration.
func GetPrometheusServiceDiscoveryDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getServiceDiscoveryDir(cmdName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SDConfig represents a Prometheus service discovery config entry.
//
// file_sd_config docs: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#file_sd_config
type SDConfig struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

// WritePrometheusSDConfig writes the SDConfig with the provided name
// to the location expected by the prometheus instance start by tmpnet.
//
// If withGitHubLabels is true, checks env vars for GitHub-specific labels
// and adds them as labels if present before writing the SDConfig.
//
// Returns the path to the written configuration file.
func WritePrometheusSDConfig(name string, sdConfig SDConfig, withGitHubLabels bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func applyGitHubLabels(sdConfig SDConfig) SDConfig {
	_ = "STUB: not implemented"
	return *new(SDConfig)
}

func getLogFilename(cmdName string) string { _ = "STUB: not implemented"; return "" }

func getLogPath(cmdName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getPIDPath(workingDir string) string { _ = "STUB: not implemented"; return "" }

// startCollector starts a collector process if it is not already running.
func startCollector(
	ctx context.Context,
	log logging.Logger,
	cmdName string,
	args string,
	config string,
) error {
	_ = "STUB: not implemented"
	// Determine paths
	return nil
}

// Ensure required paths exist

// Check if the process is already running

// Clear any stale pid file

// Check if the specified command is available in the path

// Write the collector config file

// Start the process

// processFromPIDFile attempts to retrieve a running process from the specified PID file.
func processFromPIDFile(cmdName string, pidPath string) (*os.Process, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getPID attempts to read the PID of the collector from a PID file.
func getPID(cmdName string, pidPath string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// clearStalePIDFile remove an existing pid file to avoid conflicting with a new process.
func clearStalePIDFile(log logging.Logger, cmdName string, pidPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// collectorConfig represents the configuration for a collector of metrics or logs
type collectorConfig struct {
	// Credentials for basic auth
	username string
	password string
	// URL to push to or to query
	url string
}

// getCollectorConfigForQuery retrieves the url, username and password to query with for the given command.
func getCollectorConfigForQuery(cmdName string) (collectorConfig, error) {
	_ = "STUB: not implemented"
	return *new(collectorConfig), nil
}

// getCollectorConfigForPush retrieves the url, username and password to push with for the given command.
func getCollectorConfigForPush(cmdName string) (collectorConfig, error) {
	_ = "STUB: not implemented"
	return *new(collectorConfig), nil
}

// getCollectorConfig retrieves the url, username and password for the
// command. The urlSuffix will determine whether the returned URL is
// used for pushing data or verifying collection.
func getCollectorConfig(cmdName string, urlSuffix string) (collectorConfig, error) {
	_ = "STUB: not implemented"
	return *new(collectorConfig), nil
}

// Start a collector process. Use bash to execute the command in the background and enable
// stderr and stdout redirection to a log file.
//
// Ideally this would be possible without bash, but it does not seem possible to
// have this process open a log file, set cmd.Stdout cmd.Stderr to that file, and
// then have the child process be able to write to that file once the parent
// process exits. Attempting to do so resulted in an empty log file.
func startCollectorProcess(
	ctx context.Context,
	log logging.Logger,
	cmdName string,
	args string,
	workingDir string,
	pidPath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure the child process will outlive its parent

// Wait for PID file

// Wait for non-empty log file. An empty log file should only occur if the command
// invocation is not correctly redirecting stderr and stdout to the expected file.

// checkReadiness retrieves the provided URL and indicates whether it returned 200
func checkReadiness(ctx context.Context, url string) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

//nolint:bodyclose // body is closed via rpc.CleanlyCloseBody

// waitForReadiness waits until the given readiness URL returns 200
func waitForReadiness(ctx context.Context, log logging.Logger, cmdName string, readinessURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func pollUntilContextCancel(ctx context.Context, condition wait.ConditionWithContextFunc) error {
	_ = "STUB: not implemented"
	return nil
}

/* immediate */
