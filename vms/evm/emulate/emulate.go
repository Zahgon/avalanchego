// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package emulate provides temporary emulation of coreth (C-Chain) and
// subnet-evm (EVM L1) behaviours. All functions are safe for concurrent use
// with each other, but all hold the same mutex so their execution SHOULD be
// short-lived.
package emulate

// CChain executes `fn` as if running in a `coreth` node.
func CChain(fn func() error) error { _ = "STUB: not implemented"; return nil }

// SubnetEVM executes `fn` as if running in a `subnet-evm` node.
func SubnetEVM(fn func() error) error { _ = "STUB: not implemented"; return nil }

// CChainVal executes `fn` as if running in a `coreth` node.
func CChainVal[T any](fn func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *

	// SubnetEVMVal executes `fn` as if running in a `subnet-evm` node.
	new(T), nil
}

func SubnetEVMVal[T any](fn func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func val[T any](
	wrap func(func() error) error,
	fn func() (T, error),
) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
