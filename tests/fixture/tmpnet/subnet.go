// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
)

const (
	defaultSubnetDirName = "subnets"
	jsonFileSuffix       = ".json"
)

type Chain struct {
	// Set statically
	VMID    ids.ID
	Config  string
	Genesis []byte
	// VersionArgs are the argument(s) to pass to the VM binary to receive
	// version details in json format (e.g. `--version-json`). This
	// supports checking that the rpcchainvm version of the VM binary
	// matches the version used by the configured avalanchego binary. If
	// empty, the version check will be skipped.
	VersionArgs []string

	// Set at runtime
	ChainID      ids.ID
	PreFundedKey *secp256k1.PrivateKey
}

type Subnet struct {
	// A unique string that can be used to refer to the subnet across different temporary
	// networks (since the SubnetID will be different every time the subnet is created)
	Name string

	Config ConfigMap

	// The ID of the transaction that created the subnet
	SubnetID ids.ID

	// The private key that owns the subnet
	OwningKey *secp256k1.PrivateKey

	// IDs of the nodes responsible for validating the subnet
	ValidatorIDs []ids.NodeID

	Chains []*Chain
}

// Retrieves a wallet configured for use with the subnet
func (s *Subnet) GetWallet(ctx context.Context, uri string) (*primary.Wallet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only fetch the subnet transaction if a subnet ID is present. This won't be true when
// the wallet is first used to create the subnet.

// Issues the subnet creation transaction and retains the result. The URI of a node is
// required to issue the transaction.
func (s *Subnet) Create(ctx context.Context, uri string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Subnet) CreateChains(ctx context.Context, log logging.Logger, uri string) error {
	_ = "STUB: not implemented"
	return nil
}

// Add validators to the subnet
func (s *Subnet) AddValidators(ctx context.Context, log logging.Logger, apiURI string, nodes ...*Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect the end times for current validators to reuse for subnet validators

// Write the subnet configuration to disk
func (s *Subnet) Write(subnetDir string) error { _ = "STUB: not implemented"; return nil }

// Since subnets are expected to be serialized for the first time
// without their chains having been created (i.e. chains will have
// empty IDs), use the absence of chain IDs as a prompt for a
// subnet name uniqueness check.

// Write subnet configuration for tmpnet

// HasChainConfig indicates whether at least one of the subnet's
// chains have explicit configuration. This can be used to determine
// whether validator restart is required after chain creation to
// ensure that chains are configured correctly.
func (s *Subnet) HasChainConfig() bool { _ = "STUB: not implemented"; return false }

func WaitForActiveValidators(
	ctx context.Context,
	log logging.Logger,
	pChainClient *platformvm.Client,
	subnet *Subnet,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Reads subnets from [network dir]/subnets/[subnet name].json
func readSubnets(subnetDir string) ([]*Subnet, error) { _ = "STUB: not implemented"; return nil, nil }

// Looking only for files

// Subnet files should have a .json extension
