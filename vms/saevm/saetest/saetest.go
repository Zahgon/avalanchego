// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package saetest provides testing helpers for [Streaming Asynchronous
// Execution] (SAE).
//
// [Streaming Asynchronous Execution]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
package saetest

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/state/snapshot"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/params"

	"github.com/ava-labs/avalanchego/utils/lock"
	"github.com/ava-labs/avalanchego/vms/saevm/saedb"
)

var _ saedb.StateDBOpener = (*stateDBOpener)(nil)

type stateDBOpener struct {
	cache state.Database
	snaps *snapshot.Tree
}

// NewStateDBOpener provides an abstraction to create a `state.StateDB`.
// `snaps` MAY be nil.
func NewStateDBOpener(cache state.Database, snaps *snapshot.Tree) saedb.StateDBOpener {
	_ = "STUB: not implemented"
	return *new(saedb.StateDBOpener)
}

func (o *stateDBOpener) StateDB(root common.Hash) (*state.StateDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TrieHasher returns an arbitrary trie hasher.
func TrieHasher() types.TrieHasher { _ = "STUB: not implemented"; return *new(types.TrieHasher) }

// ChainConfig returns [params.MergedTestChainConfig] as it includes all EIPs
// available for testing, including post-merge upgrades. This SHOULD be used for
// all testing.
func ChainConfig() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

// Rules returns the rules associated with [ChainConfig], at height and time
// zero, and post-merge.
func Rules() params.Rules { _ = "STUB: not implemented"; return *new(params.Rules) }

// An EventCollector collects all events received from an [event.Subscription].
// All methods are safe for concurrent use.
type EventCollector[T any] struct {
	ch   chan T
	done chan struct{}
	sub  event.Subscription

	all  []T
	cond *lock.Cond
}

// NewEventCollector returns a new [EventCollector], subscribing via the
// provided function. [EventCollector.Unsubscribe] must be called to release
// resources.
func NewEventCollector[T any](subscribe func(chan<- T) event.Subscription) *EventCollector[T] {
	_ = "STUB: not implemented"
	return nil
}

func (c *EventCollector[T]) collect() { _ = "STUB: not implemented"; return }

// All returns all events received thus far.
func (c *EventCollector[T]) All() []T { _ = "STUB: not implemented"; return nil }

// Unsubscribe unsubscribes from the subscription and returns the error,
// possibly nil, received on [event.Subscription.Err].
func (c *EventCollector[T]) Unsubscribe() error { _ = "STUB: not implemented"; return nil }

// WaitForAtLeast blocks until at least `n` events have been received.
func (c *EventCollector[T]) WaitForAtLeast(ctx context.Context, n int) error {
	_ = "STUB: not implemented"
	return nil
}
