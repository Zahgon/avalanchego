// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package password

// Hash of a password
type Hash struct {
	Password [32]byte `serialize:"true"` // The salted, hashed password
	Salt     [16]byte `serialize:"true"` // The salt
}

// Set updates the password hash to be of the provided password
func (h *Hash) Set(password string) error { _ = "STUB: not implemented"; return nil }

// pw is the salted, hashed password

// Check returns true iff the provided password was the same as the last
// password set.
func (h *Hash) Check(password string) bool { _ = "STUB: not implemented"; return false }
