// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/keychain"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	errCantSpend = errors.New("unable to spend this UTXO")

	_ keychain.Keychain = (*Keychain)(nil)
)

// Keychain is a collection of keys that can be used to spend outputs
type Keychain struct {
	avaxAddrToKeyIndex map[ids.ShortID]int
	ethAddrToKeyIndex  map[common.Address]int

	// These can be used to iterate over. However, they should not be modified
	// externally.
	Addrs    set.Set[ids.ShortID]
	EthAddrs set.Set[common.Address]
	Keys     []*secp256k1.PrivateKey
}

// NewKeychain returns a new keychain containing [keys]
func NewKeychain(keys ...*secp256k1.PrivateKey) *Keychain { _ = "STUB: not implemented"; return nil }

// Add a new key to the key chain
func (kc *Keychain) Add(key *secp256k1.PrivateKey) { _ = "STUB: not implemented"; return }

// Get a key from the keychain and return whether the key existed.
func (kc Keychain) Get(id ids.ShortID) (keychain.Signer, bool) {
	_ = "STUB: not implemented"

	// Get a key from the keychain and return whether the key existed.
	return *new(keychain.Signer), false
}

func (kc Keychain) GetEth(addr common.Address) (keychain.Signer, bool) {
	_ = "STUB: not implemented"
	return *new(keychain.Signer), false
}

// Addresses returns a list of addresses this keychain manages
func (kc Keychain) Addresses() set.Set[ids.ShortID] {
	_ = "STUB: not implemented"

	// EthAddresses returns a list of addresses this keychain manages
	return nil
}

func (kc Keychain) EthAddresses() set.Set[common.Address] {
	_ = "STUB: not implemented"

	// New returns a newly generated private key
	return nil
}

func (kc *Keychain) New() (*secp256k1.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Spend attempts to create an input
func (kc *Keychain) Spend(out verify.Verifiable, time uint64) (verify.Verifiable, []*secp256k1.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(verify.Verifiable), nil, nil
}

// Match attempts to match a list of addresses up to the provided threshold
func (kc *Keychain) Match(owners *OutputOwners, time uint64) ([]uint32, []*secp256k1.PrivateKey, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// PrefixedString returns the key chain as a string representation with [prefix]
// added before every line.
func (kc *Keychain) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

// We assume that the maximum size of a byte slice that
// can be stringified is at least the length of a SECP256K1 private key

func (kc *Keychain) String() string { _ = "STUB: not implemented"; return "" }

// to avoid internals type assertions
func (kc Keychain) get(id ids.ShortID) (*secp256k1.PrivateKey, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
