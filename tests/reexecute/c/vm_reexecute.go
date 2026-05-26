// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/reexecute"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer"
)

var (
	blockDirArg        string
	currentStateDirArg string
	startBlockArg      uint64
	endBlockArg        uint64
	chanSizeArg        int
	executionTimeout   time.Duration
	labelsArg          string

	pprofDirArg                string
	metricsServerEnabledArg    bool
	metricsServerPortArg       uint64
	metricsCollectorEnabledArg bool

	networkUUID string = uuid.NewString()
	labels             = map[string]string{
		"job":               "c-chain-reexecution",
		"is_ephemeral_node": "false",
		"chain":             "C",
		"network_uuid":      networkUUID,
	}

	configKey         = "config"
	defaultConfigKey  = "default"
	predefinedConfigs = map[string]string{
		defaultConfigKey: `{}`,
		"archive": `{
			"pruning-enabled": false
		}`,
		"pathdb": `{
			"state-scheme": "path",
			"state-sync-enabled": false
		}`,
		"firewood": `{
			"state-scheme": "firewood",
			"snapshot-cache": 0,
			"pruning-enabled": true,
			"state-sync-enabled": false,
			"commit-interval": 4096,
			"state-history": 8192
		}`,
		"firewood-archive": `{
			"state-scheme": "firewood",
			"snapshot-cache": 0,
			"pruning-enabled": false,
			"state-sync-enabled": false
		}`,
	}

	configNameArg  string
	runnerTypeArg  string
	configBytesArg []byte

	benchmarkOutputFileArg string
)

func init() {
	evm.RegisterAllLibEVMExtras()

	flag.StringVar(&blockDirArg, "block-dir", blockDirArg, "Block DB directory to read from during re-execution.")
	flag.StringVar(&currentStateDirArg, "current-state-dir", currentStateDirArg, "Current state directory including VM DB and Chain Data Directory for re-execution.")
	flag.Uint64Var(&startBlockArg, "start-block", 101, "Start block to begin execution (exclusive).")
	flag.Uint64Var(&endBlockArg, "end-block", 200, "End block to end execution (inclusive).")
	flag.IntVar(&chanSizeArg, "chan-size", 100, "Size of the channel to use for block processing.")
	flag.DurationVar(&executionTimeout, "execution-timeout", 0, "Benchmark execution timeout. After this timeout has elapsed, terminate the benchmark without error. If 0, no timeout is applied.")

	flag.StringVar(&pprofDirArg, "pprof-dir", "", "Directory to write cpu, mem, and lock profiles. Empty to disable.")
	flag.BoolVar(&metricsServerEnabledArg, "metrics-server-enabled", false, "Whether to enable the metrics server.")
	flag.Uint64Var(&metricsServerPortArg, "metrics-server-port", 0, "The port the metrics server will listen to.")
	flag.BoolVar(&metricsCollectorEnabledArg, "metrics-collector-enabled", false, "Whether to enable the metrics collector (if true, then metrics-server-enabled must be true as well).")
	flag.StringVar(&labelsArg, "labels", "", "Comma separated KV list of metric labels to attach to all exported metrics. Ex. \"owner=tim,runner=snoopy\"")

	predefinedConfigKeys := slices.Collect(maps.Keys(predefinedConfigs))
	predefinedConfigOptionsStr := fmt.Sprintf("[%s]", strings.Join(predefinedConfigKeys, ", "))
	flag.StringVar(&configNameArg, configKey, defaultConfigKey, fmt.Sprintf("Specifies the predefined config to use for the VM. Options include %s.", predefinedConfigOptionsStr))
	flag.StringVar(&runnerTypeArg, "runner", "dev", "Type/label of the runner executing this test. Added as a metric label and to the benchmark name for grouping results.")

	flag.StringVar(&benchmarkOutputFileArg, "benchmark-output-file", benchmarkOutputFileArg, "Filepath where benchmark results will be written to.")

	flag.Parse()

	if metricsCollectorEnabledArg {
		metricsServerEnabledArg = true
	}

	customLabels, err := parseCustomLabels(labelsArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse labels: %v\n", err)
		os.Exit(1)
	}
	maps.Copy(labels, customLabels)

	// Set the config from the predefined configs and add to custom labels for the job.
	predefinedConfigStr, ok := predefinedConfigs[configNameArg]
	if !ok {
		fmt.Fprintf(os.Stderr, "invalid config name %q. Valid options include %s.\n", configNameArg, predefinedConfigOptionsStr)
		os.Exit(1)
	}
	labels[configKey] = configNameArg
	configBytesArg = []byte(predefinedConfigStr)

	// Set the runner label on the metrics.
	labels["runner"] = runnerTypeArg
}

func main() {
	tc := tests.NewTestContext(tests.NewDefaultLogger("c-chain-reexecution"))
	tc.SetDefaultContextParent(context.Background())
	defer tc.RecoverAndExit()

	benchmarkName := fmt.Sprintf(
		"BenchmarkReexecuteRange/[%d,%d]-Config-%s-Runner-%s",
		startBlockArg,
		endBlockArg,
		configNameArg,
		runnerTypeArg,
	)

	benchmarkReexecuteRange(
		tc,
		benchmarkName,
		blockDirArg,
		currentStateDirArg,
		configBytesArg,
		startBlockArg,
		endBlockArg,
		chanSizeArg,
		metricsServerEnabledArg,
		metricsServerPortArg,
		metricsCollectorEnabledArg,
		benchmarkOutputFileArg,
	)
}

func benchmarkReexecuteRange(
	tc tests.TestContext,
	benchmarkName string,
	blockDir string,
	currentStateDir string,
	configBytes []byte,
	startBlock uint64,
	endBlock uint64,
	chanSize int,
	metricsServerEnabled bool,
	metricsPort uint64,
	metricsCollectorEnabled bool,
	benchmarkOutputFile string,
) {
	_ = "STUB: not implemented"
	return
}

// Create the prefix gatherer passed to the VM and register it with the top-level,
// labeled gatherer.

// consensusRegistry includes the chain="C" label and the prefix "avalanche_snowman".
// The consensus registry is passed to the executor to mimic a subset of consensus metrics.

// Report the desired top-level metrics

type vmExecutorConfig struct {
	Log logging.Logger
	// Registry is the registry to register the metrics with.
	Registry prometheus.Registerer
	// ExecutionTimeout is the maximum timeout to continue executing blocks.
	// If 0, no timeout is applied. If non-zero, the executor will exit early
	// WITHOUT error after hitting the timeout.
	// This is useful to provide consistent duration benchmarks.
	ExecutionTimeout time.Duration

	// [StartBlock, EndBlock] defines the range (inclusive) of blocks to execute.
	StartBlock, EndBlock uint64
}

type vmExecutor struct {
	config     vmExecutorConfig
	vm         block.ChainVM
	metrics    *consensusMetrics
	etaTracker *timer.EtaTracker
}

func newVMExecutor(vm block.ChainVM, config vmExecutorConfig) (*vmExecutor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ETA tracker uses a 10-sample moving window to smooth rate estimates,
// and a 1.2 slowdown factor to slightly pad ETA early in the run,
// tapering to 1.0 as progress approaches 100%.

func (e *vmExecutor) execute(ctx context.Context, blockBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *vmExecutor) executeSequence(ctx context.Context, blkChan <-chan reexecute.BlockResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize ETA tracking with a baseline sample at 0 progress

type consensusMetrics struct {
	lastAcceptedHeight prometheus.Gauge
}

// newConsensusMetrics creates a subset of the metrics from snowman consensus
// [engine](../../snow/engine/snowman/metrics.go).
//
// The registry passed in is expected to be registered with the prefix
// "avalanche_snowman" and the chain label (ex. chain="C") that would be handled
// by the[chain manager](../../../chains/manager.go).
func newConsensusMetrics(registry prometheus.Registerer) (*consensusMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// startServer starts a Prometheus server for the provided gatherer and returns
// the server address.
func startServer(
	tc tests.TestContext,
	gatherer prometheus.Gatherer,
	port uint64,
) string {
	_ = "STUB: not implemented"
	return ""
}

// startCollector starts a Prometheus collector configured to scrape the server
// listening on serverAddr. startCollector also attaches the provided labels +
// Github labels if available to the collected metrics.
func startCollector(tc tests.TestContext, name string, labels map[string]string, serverAddr string) {
	_ = "STUB: not implemented"
	return
}

// Ensure a final metrics scrape.
// This default delay is set above the default scrape interval used by StartPrometheus.

// TODO(marun) Re-enable this check as part of a PR updating the metrics links
// r.NoError(tmpnet.CheckMetricsExist(tc.DefaultContext(), logger, networkUUID))

/* withGitHubLabels */

// benchmarkResult represents a single benchmark measurement.
type benchmarkResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

// benchmarkTool collects and manages benchmark results for a named benchmark.
// It allows adding multiple results and saving them to a file in JSON format.
type benchmarkTool struct {
	name    string
	results []benchmarkResult
}

// newBenchmarkTool creates a new benchmarkTool instance with the given name.
// The name is used as the base name for all results collected by this tool.
// When results are added, the unit is appended to this base name.
func newBenchmarkTool(name string) *benchmarkTool { _ = "STUB: not implemented"; return nil }

// addResult adds a new benchmark result with the given value and unit.
// The result name is constructed by appending the unit to the benchmark name.
// Calling `addResult` is analogous to calling `b.ReportMetric()`.
func (b *benchmarkTool) addResult(value float64, unit string) { _ = "STUB: not implemented"; return }

// saveToFile writes all collected benchmark results to a JSON file at the
// specified path. The output is formatted with indentation for readability.
// Returns an error if marshaling or file writing fails.
func (b *benchmarkTool) saveToFile(path string) error { _ = "STUB: not implemented"; return nil }

// logResults logs all collected benchmark results using the provided logger.
func (b *benchmarkTool) logResults(log logging.Logger) { _ = "STUB: not implemented"; return }

// parseCustomLabels parses a comma-separated list of key-value pairs into a map
// of custom labels.
func parseCustomLabels(labelsStr string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
