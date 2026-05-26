// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"sync"
)

type SubnetSuite struct {
	blockchainIDs map[string]string
	lock          sync.RWMutex
}

func (s *SubnetSuite) GetBlockchainID(alias string) string { _ = "STUB: not implemented"; return "" }

func (s *SubnetSuite) SetBlockchainIDs(blockchainIDs map[string]string) {
	_ = "STUB: not implemented"
	return
}

// CreateSubnetsSuite creates subnets for given [genesisFiles], and registers a before suite that starts an AvalancheGo process to use for the e2e tests.
// genesisFiles is a map of test aliases to genesis file paths.
func CreateSubnetsSuite(genesisFiles map[string]string) *SubnetSuite {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of the AvalancheGo external bash script, it is null for most
// processes except the first process that starts AvalancheGo

// This is used to pass the blockchain IDs from the SynchronizedBeforeSuite() to the tests

// Our test suite runs in separate processes, ginkgo has
// SynchronizedBeforeSuite() which runs once, and its return value is passed
// over to each worker.
//
// Here an AvalancheGo node instance is started, and subnets are created for
// each test case. Each test case has its own subnet, therefore all tests
// can run in parallel without any issue.
//

// Assumes that startCmd will launch a node with HTTP Port at [utils.DefaultLocalNodeURI]

// SynchronizedAfterSuite() takes two functions, the first runs after each test suite is done and the second
// function is executed once when all the tests are done. This function is used
// to gracefully shutdown the AvalancheGo node.

// CreateNewSubnet creates a new subnet and Subnet-EVM blockchain with the given genesis file.
// returns the ID of the new created blockchain.
func CreateNewSubnet(ctx context.Context, genesisFilePath string) string {
	_ = "STUB: not implemented"
	return ""
}

// MakeWallet fetches the available UTXOs owned by [kc] on the network
// that [LocalAPIURI] is hosting.

// Confirm the new blockchain is ready by waiting for the readiness endpoint

// Return the blockchainID of the newly created blockchain

// GetDefaultChainURI returns the default chain URI for a given blockchainID
func GetDefaultChainURI(blockchainID string) string { _ = "STUB: not implemented"; return "" }

// GetFilesAndAliases returns a map of aliases to file paths in given [dir].
func GetFilesAndAliases(dir string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
