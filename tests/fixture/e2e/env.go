// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

import (
	"time"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

// Env is used to access shared test fixture. Intended to be
// initialized from SynchronizedBeforeSuite. Not exported to limit
// access to the shared env to GetEnv which adds a test context.
var env *TestEnvironment

func InitSharedTestEnvironment(tc tests.TestContext, envBytes []byte) {
	_ = "STUB: not implemented"
	return
}

// Ginkgo parallelization is at the process level, so a given key
// can safely be used by all tests in a given process without fear
// of conflicting usage.

type TestEnvironment struct {
	// The parent directory of network directories
	RootNetworkDir string
	// The directory where the test network configuration is stored
	NetworkDir string
	// Pre-funded key for this ginkgo process
	PreFundedKey *secp256k1.PrivateKey
	// The duration to wait before shutting down private networks. A
	// non-zero value may be useful to ensure all metrics can be
	// scraped before shutdown.
	PrivateNetworkShutdownDelay time.Duration

	testContext tests.TestContext
}

// Retrieve the test environment configured with the provided test context.
func GetEnv(tc tests.TestContext) *TestEnvironment { _ = "STUB: not implemented"; return nil }

func (te *TestEnvironment) Marshal() []byte { _ = "STUB: not implemented"; return nil }

// Initialize a new test environment with a shared network (either pre-existing or newly created).
func NewTestEnvironment(tc tests.TestContext, flagVars *FlagVars, desiredNetwork *tmpnet.Network) *TestEnvironment {
	_ = "STUB: not implemented"
	return nil
}

// Consider monitoring flags for any command but stop

// Register cleanups before network start to ensure they run after the network is stopped (LIFO)

// Attempt to load the network if it may already be running

// If populated, prompts removal of the referenced symlink if --stop-network is specified

// Attempt to reuse the network at the default owner path

// Try to load the existing network

// Enable removal of the referenced symlink if --stop-network is specified

// Remove the symlink to avoid attempts to reuse the stopped network

// Start a new network

// TODO(marun) Maybe accept a factory function for the desired network
// that is only run when a new network will be started?

// Once one or more nodes are running it should be safe to wait for promtail to report readiness

// TODO(marun) Maybe make this configurable to enable the check for a test suite that writes service
// discovery configuration for its own metrics endpoint?

// Retrieve a random URI to naively attempt to spread API load across nodes.
func (te *TestEnvironment) GetRandomNodeURI() tmpnet.NodeURI {
	_ = "STUB: not implemented"
	return *new(tmpnet.NodeURI)
}

//#nosec G404

// Avoid returning URIs for nodes whose lifespan is indeterminate

// Only running nodes have URIs

// Retrieve the network to target for testing.
func (te *TestEnvironment) GetNetwork() *tmpnet.Network { _ = "STUB: not implemented"; return nil }

// Create a new keychain with the process's pre-funded key.
func (te *TestEnvironment) NewKeychain() *secp256k1fx.Keychain {
	_ = "STUB: not implemented"
	return nil
}

// Create a new private network that is not shared with other tests.
func (te *TestEnvironment) StartPrivateNetwork(network *tmpnet.Network) {
	_ = "STUB: not implemented"
	return
}

// Use the same configuration as the shared network
