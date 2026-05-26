// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import "github.com/ava-labs/avalanchego/ids"

// TestManager implements Manager but does nothing. Always returns nil error.
// To be used only in tests
var TestManager Manager = testManager{}

type testManager struct{}

func (testManager) QueueChainCreation(ChainParameters) { _ = "STUB: not implemented"; return }

func (testManager) AddRegistrant(Registrant) { _ = "STUB: not implemented"; return }

func (testManager) Aliases(ids.ID) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (testManager) PrimaryAlias(ids.ID) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (testManager) PrimaryAliasOrDefault(ids.ID) string { _ = "STUB: not implemented"; return "" }

func (testManager) Alias(ids.ID, string) error { _ = "STUB: not implemented"; return nil }

func (testManager) RemoveAliases(ids.ID) { _ = "STUB: not implemented"; return }

func (testManager) Shutdown() { _ = "STUB: not implemented"; return }

func (testManager) StartChainCreator(ChainParameters) error { _ = "STUB: not implemented"; return nil }

func (testManager) IsBootstrapped(ids.ID) bool { _ = "STUB: not implemented"; return false }

func (testManager) Lookup(s string) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (testManager) LookupVM(s string) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
