// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package key

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/libevm/common"
)

type Key struct {
	PrivKey *ecdsa.PrivateKey
	Address common.Address
}

func CreateKey(pk *ecdsa.PrivateKey) *Key { _ = "STUB: not implemented"; return nil }

// Load attempts to open a [Key] stored at [file].
func Load(file string) (*Key, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadAll loads all keys in [dir].
func LoadAll(ctx context.Context, dir string) ([]*Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save persists a [Key] to [dir] (where the filename is the hex-encoded
// address).
func (k *Key) Save(dir string) error { _ = "STUB: not implemented"; return nil }

// Generate creates a new [Key] and returns it.
func Generate() (*Key, error) { _ = "STUB: not implemented"; return nil, nil }
