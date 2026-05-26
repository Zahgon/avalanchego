// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"
)

type getCountFunc func() (int, error)

// waitForCount waits until the provided function returns greater than zero.
func waitForCount(ctx context.Context, log logging.Logger, name string, getCount getCountFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckLogsExist checks if logs exist for the given network. If no network UUID is
// provided, an attempt will be made to derive selectors from env vars (GH_*) identifying
// a github actions run.
func CheckLogsExist(ctx context.Context, log logging.Logger, networkUUID string) error {
	_ = "STUB: not implemented"
	return nil
}

func queryLoki(
	ctx context.Context,
	config collectorConfig,
	query string,
) (int, error) {
	_ = "STUB: not implemented"
	// Compose the URL
	return 0, nil
}

// Create request

// Execute request
//nolint:bodyclose // body is closed via rpc.CleanlyCloseBody

// Read and parse response

// Parse JSON response

// Extract count value

// Convert value to a string

// Convert string to float64 first to handle scientific notation

// Round to nearest integer

// CheckMetricsExist checks if metrics exist for the given network. Github labels are also
// used as filters if provided as env vars (GH_*).
func CheckMetricsExist(ctx context.Context, log logging.Logger, networkUUID string) error {
	_ = "STUB: not implemented"
	return nil
}

func queryPrometheus(
	ctx context.Context,
	log logging.Logger,
	config collectorConfig,
	query string,
) (int, error) {
	_ = "STUB: not implemented"
	// Create client with basic auth
	return 0, nil
}

// Query Prometheus

type basicAuthRoundTripper struct {
	username, password string
	rt                 http.RoundTripper
}

func (b *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSelectors returns the comma-separated list of selectors.
func getSelectors(networkUUID string) (string, error) {
	_ = "STUB: not implemented"
	// If network UUID is provided, use it as the only selector
	return "", nil
}

// Fall back to using Github labels as selectors
