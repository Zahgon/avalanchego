// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"sync"

	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

func (b *backend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// A number of subscriptions don't make sense in SAE so are no-ops. The lack of
// reorgs makes chain-side and log-removal events impossible. As "pending"
// refers to accepted but not executed blocks, pending logs are an oxymoron.

func (*backend) SubscribeChainSideEvent(chan<- core.ChainSideEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (*backend) SubscribeRemovedLogsEvent(chan<- core.RemovedLogsEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

func (*backend) SubscribePendingLogsEvent(chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

type noopSubscription struct {
	once sync.Once
	err  chan error
}

func newNoopSubscription() *noopSubscription { _ = "STUB: not implemented"; return nil }

func (s *noopSubscription) Err() <-chan error { _ = "STUB: not implemented"; return nil }

func (s *noopSubscription) Unsubscribe() { _ = "STUB: not implemented"; return }
