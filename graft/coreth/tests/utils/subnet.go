// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
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

// GetDefaultChainURI returns the default chain URI for a given blockchainID
func GetDefaultChainURI(blockchainID string) string { _ = "STUB: not implemented"; return "" }

// GetFilesAndAliases returns a map of aliases to file paths in given [dir].
func GetFilesAndAliases(dir string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
