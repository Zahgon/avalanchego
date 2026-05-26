// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package antithesis

import (
	"time"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

const (
	URIsKey     = "uris"
	ChainIDsKey = "chain-ids"
	DurationKey = "duration"

	EnvPrefix = "avawl"
)

type Config struct {
	URIs     []string
	ChainIDs []string
	Duration time.Duration
}

type SubnetsForNodesFunc func(nodes ...*tmpnet.Node) []*tmpnet.Subnet

func NewConfig(tc tests.TestContext, defaultNetwork *tmpnet.Network) *Config {
	_ = "STUB: not implemented"
	return nil
}

func NewConfigWithSubnets(tc tests.TestContext, defaultNetwork *tmpnet.Network, getSubnets SubnetsForNodesFunc) *Config {
	_ = "STUB: not implemented"
	// tmpnet configuration
	return nil
}

// Accept a list of chain IDs, assume they each belong to a separate subnet
// TODO(marun) Revisit how chain IDs are provided when 1:n subnet:chain configuration is required.

// Env vars take priority over flags

//nolint:errcheck // CSV.Set doesn't actually return an error

//nolint:errcheck // CSV.Set doesn't actually return an error

// Use the network configuration provided

// Create a new network

// configForNewNetwork creates a new network and returns the resulting config.
func configForNewNetwork(
	tc tests.TestContext,
	defaultNetwork *tmpnet.Network,
	getSubnets SubnetsForNodesFunc,
	flagVars *e2e.FlagVars,
	duration time.Duration,
) *Config {
	_ = "STUB: not implemented"
	return nil
}

// CSV is a custom type that implements the flag.Value interface
type CSV []string

// String returns the string representation of the CSV type
func (c *CSV) String() string { _ = "STUB: not implemented"; return "" }

// Set splits the input string by commas and sets the CSV type
func (c *CSV) Set(value string) error { _ = "STUB: not implemented"; return nil }
