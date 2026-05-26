// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

import (
	"math/big"
	"time"

	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/chain/p/builder"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

const (
	DefaultTimeout = tests.DefaultTimeout

	DefaultPollingInterval = tmpnet.DefaultPollingInterval

	// Setting this env will disable post-test bootstrap
	// checks. Useful for speeding up iteration during test
	// development.
	SkipBootstrapChecksEnvName = "E2E_SKIP_BOOTSTRAP_CHECKS"

	DefaultValidatorStartTimeDiff = tmpnet.DefaultValidatorStartTimeDiff

	DefaultGasLimit = uint64(21000) // Standard gas limit

	// Directory used to store private networks (specific to a single test)
	// under the shared network dir.
	PrivateNetworksDirName = "private_networks"
)

// NewPrivateKey returns a new private key.
func NewPrivateKey(tc tests.TestContext) *secp256k1.PrivateKey {
	_ = "STUB: not implemented"
	return nil
}

// Create a new wallet for the provided keychain against the specified node URI.
func NewWallet(tc tests.TestContext, keychain *secp256k1fx.Keychain, nodeURI tmpnet.NodeURI) *primary.Wallet {
	_ = "STUB: not implemented"
	return nil
}

// Reducing the default from 100ms speeds up detection of tx acceptance

// OutputWalletBalances outputs the X-Chain and P-Chain balances of the provided wallet.
func OutputWalletBalances(tc tests.TestContext, wallet *primary.Wallet) {
	_ = "STUB: not implemented"
	return
}

// GetWalletBalances retrieves the X-Chain and P-Chain balances of the provided wallet.
func GetWalletBalances(tc tests.TestContext, wallet *primary.Wallet) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Create a new eth client targeting the specified node URI.
func NewEthClient(tc tests.TestContext, nodeURI tmpnet.NodeURI) *ethclient.Client {
	_ = "STUB: not implemented"
	return nil
}

// Adds an ephemeral node intended to be used by a single test.
func AddEphemeralNode(tc tests.TestContext, network *tmpnet.Network, node *tmpnet.Node) *tmpnet.Node {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the given node to report healthy.
func WaitForHealthy(t require.TestingT, node *tmpnet.Node) {
	_ = "STUB: not implemented"
	// Need to use explicit context (vs DefaultContext()) to support use with DeferCleanup
	return
}

// Sends an eth transaction and waits for the transaction receipt from the
// execution of the transaction.
func SendEthTransaction(tc tests.TestContext, ethClient *ethclient.Client, signedTx *types.Transaction) *types.Receipt {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the receipt

// Transaction is still pending

// Determines the suggested gas price for the configured client that will
// maximize the chances of transaction acceptance.
func SuggestGasPrice(tc tests.TestContext, ethClient *ethclient.Client) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

// Double the suggested gas price to maximize the chances of
// acceptance. Maybe this can be revisited pending resolution of
// https://github.com/ava-labs/avalanchego/graft/coreth/issues/314.

// Helper simplifying use via an option of a gas price appropriate for testing.
func WithSuggestedGasPrice(tc tests.TestContext, ethClient *ethclient.Client) common.Option {
	_ = "STUB: not implemented"
	return *new(common.Option)
}

// Verify that a new node can bootstrap into the network. If the check wasn't skipped,
// the node will be returned to the caller.
func CheckBootstrapIsPossible(tc tests.TestContext, network *tmpnet.Network) *tmpnet.Node {
	_ = "STUB: not implemented"
	return nil
}

// Ensure all subnets are bootstrapped

// StartNode will initiate node stop if an error is encountered during start,
// so no further cleanup effort is required if an error is seen here.

// Register a cleanup to ensure the node is stopped at the end of the test

// Check that the node becomes healthy within timeout

// Ensure that the primary validators are still healthy

// Start a temporary network with the provided avalanchego binary.
func StartNetwork(
	tc tests.TestContext,
	network *tmpnet.Network,
	rootNetworkDir string,
	shutdownDelay time.Duration,
	networkCmd NetworkCmd,
) {
	_ = "STUB: not implemented"
	return
}

// Symlink the path of the created network to the default owner path (e.g. latest_avalanchego-e2e)
// to enable easy discovery for reuse.

// NewPChainFeeCalculatorFromContext returns either a static or dynamic fee
// calculator depending on the provided context.
func NewPChainFeeCalculatorFromContext(context *builder.Context) fee.Calculator {
	_ = "STUB: not implemented"
	return *new(fee.Calculator)
}

// GetRepoRootPath strips the provided suffix from the current working
// directory. If the test binary is executed from the root of the repo, the
// result will be the repo root.
func GetRepoRootPath(suffix string) (string, error) {
	_ = "STUB: not implemented"
	// - When executed via a test binary, the working directory will be wherever
	// the binary is executed from, but scripts should require execution from
	// the repo root.
	//
	// - When executed via ginkgo (nicer for development + supports
	// parallel execution) the working directory will always be the
	// target path (e.g. [repo root]./tests/bootstrap/e2e) and getting the repo
	// root will require stripping the target path suffix.
	return "", nil
}
