// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"github.com/ava-labs/libevm/accounts"
)

// Config represents a small collection of configuration values to fine tune the
// P2P network layer of a protocol stack. These values can be further extended by
// all registered services.
type Config struct {
	// KeyStoreDir is the file system folder that contains private keys. The directory can
	// be specified as a relative path, in which case it is resolved relative to the
	// current directory.
	//
	// If KeyStoreDir is empty, the default location is the "keystore" subdirectory of
	// DataDir. If DataDir is unspecified and KeyStoreDir is empty, an ephemeral directory
	// is created by New and destroyed when the node is stopped.
	KeyStoreDir string `toml:",omitempty"`

	// ExternalSigner specifies an external URI for a clef-type signer.
	ExternalSigner string `toml:",omitempty"`

	// UseLightweightKDF lowers the memory and CPU requirements of the key store
	// scrypt KDF at the expense of security.
	UseLightweightKDF bool `toml:",omitempty"`

	// InsecureUnlockAllowed allows user to unlock accounts in unsafe http environment.
	InsecureUnlockAllowed bool `toml:",omitempty"`

	// BatchRequestLimit is the maximum number of requests in a batch.
	BatchRequestLimit int `toml:",omitempty"`

	// BatchResponseMaxSize is the maximum number of bytes returned from a batched rpc call.
	BatchResponseMaxSize int `toml:",omitempty"`

	SubnetEVMVersion string
}

// ExtRPCEnabled returns the indicator whether node enables the external
// RPC(http, ws or graphql).
func (c *Config) ExtRPCEnabled() bool {
	_ = "STUB: not implemented"
	// In avalanche, we always disable the external RPC.
	return false
}

// KeyDirConfig determines the settings for keydirectory
func (c *Config) KeyDirConfig() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetKeyStoreDir retrieves the key directory and will create
// and ephemeral one if necessary.
func (c *Config) GetKeyStoreDir() (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// There is no datadir.

func makeAccountManager(conf *Config) (*accounts.Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Assemble the account manager and supported backends

// For now, we're using EITHER external signer OR local signers.
// If/when we implement some form of lockfile for USB and keystore wallets,
// we can have both, but it's very confusing for the user to see the same
// accounts in both externally and locally, plus very racey.
