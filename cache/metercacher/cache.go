// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metercacher

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
)

var _ cache.Cacher[struct{}, struct{}] = (*Cache[struct{}, struct{}])(nil)

type Cache[K comparable, V any] struct {
	cache.Cacher[K, V]

	metrics *metrics
}

func New[K comparable, V any](
	namespace string,
	registerer prometheus.Registerer,
	cache cache.Cacher[K, V],
) (*Cache[K, V], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Cache[K, _]) Evict(key K) { _ = "STUB: not implemented"; return }

func (c *Cache[_, _]) Flush() { _ = "STUB: not implemented"; return }
