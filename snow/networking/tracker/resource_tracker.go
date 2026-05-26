// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracker

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/math/meter"
	"github.com/ava-labs/avalanchego/utils/resource"
)

const epsilon = 1e-9

var _ ResourceTracker = (*resourceTracker)(nil)

type Tracker interface {
	// Returns the current usage for the given node.
	Usage(nodeID ids.NodeID, now time.Time) float64
	// Returns the current usage by all nodes.
	TotalUsage() float64
	// Returns the duration between [now] and when the usage of [nodeID] reaches
	// [value], assuming that the node uses no more resources.
	// If the node's usage isn't known, or is already <= [value], returns the
	// zero duration.
	TimeUntilUsage(nodeID ids.NodeID, now time.Time, value float64) time.Duration
}

type DiskTracker interface {
	Tracker
	AvailableDiskBytes() uint64
	AvailableDiskPercentage() uint64
}

// ResourceTracker is an interface for tracking peers' usage of resources
type ResourceTracker interface {
	CPUTracker() Tracker
	DiskTracker() DiskTracker
	// Registers that the given node started processing at the given time.
	StartProcessing(ids.NodeID, time.Time)
	// Registers that the given node stopped processing at the given time.
	StopProcessing(ids.NodeID, time.Time)
}

type cpuResourceTracker struct {
	t *resourceTracker
}

func (t *cpuResourceTracker) Usage(nodeID ids.NodeID, now time.Time) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (t *cpuResourceTracker) TotalUsage() float64 { _ = "STUB: not implemented"; return 0 }

func (t *cpuResourceTracker) TimeUntilUsage(nodeID ids.NodeID, now time.Time, value float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type diskResourceTracker struct {
	t *resourceTracker
}

func (t *diskResourceTracker) Usage(nodeID ids.NodeID, now time.Time) float64 {
	_ = "STUB: not implemented"
	return 0
}

// [realWriteUsage] is only used for metrics.

func (t *diskResourceTracker) AvailableDiskBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *diskResourceTracker) AvailableDiskPercentage() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *diskResourceTracker) TotalUsage() float64 { _ = "STUB: not implemented"; return 0 }

func (t *diskResourceTracker) TimeUntilUsage(nodeID ids.NodeID, now time.Time, value float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// [realWriteUsage] is only used for metrics.

type resourceTracker struct {
	lock sync.RWMutex

	resources resource.User
	factory   meter.Factory
	// Tracks total number of current processing requests by all nodes.
	processingMeter meter.Meter
	halflife        time.Duration
	// Each element is a meter that tracks the number of current processing
	// requests by a node. [meters] is ordered by the last time that a meter was
	// utilized. This doesn't necessarily result in the meters being sorted
	// based on their usage. However, in practice the nodes that are not being
	// utilized will move towards the oldest elements where they can be deleted.
	meters  *linked.Hashmap[ids.NodeID, meter.Meter]
	metrics *trackerMetrics
}

func NewResourceTracker(
	reg prometheus.Registerer,
	resources resource.User,
	factory meter.Factory,
	halflife time.Duration,
) (ResourceTracker, error) {
	_ = "STUB: not implemented"
	return *new(ResourceTracker), nil
}

func (rt *resourceTracker) CPUTracker() Tracker { _ = "STUB: not implemented"; return *new(Tracker) }

func (rt *resourceTracker) DiskTracker() DiskTracker {
	_ = "STUB: not implemented"
	return *new(DiskTracker)
}

func (rt *resourceTracker) StartProcessing(nodeID ids.NodeID, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (rt *resourceTracker) StopProcessing(nodeID ids.NodeID, now time.Time) {
	_ = "STUB: not implemented"
	return
}

// getMeter returns the meter used to measure CPU time spent processing
// messages from [nodeID].
// assumes [rt.lock] is held.
func (rt *resourceTracker) getMeter(nodeID ids.NodeID) meter.Meter {
	_ = "STUB: not implemented"
	return *new(meter.Meter)
}

// prune attempts to remove meters that currently show a value less than
// [epsilon].
//
// Because [rt.meters] isn't guaranteed to be sorted by their values, this
// doesn't guarantee that all meters showing less than [epsilon] are removed.
func (rt *resourceTracker) prune(now time.Time) { _ = "STUB: not implemented"; return }

type trackerMetrics struct {
	processingTimeMetric    prometheus.Gauge
	cpuMetric               prometheus.Gauge
	diskReadsMetric         prometheus.Gauge
	diskWritesMetric        prometheus.Gauge
	diskSpaceAvailable      prometheus.Gauge
	diskPercentageAvailable prometheus.Gauge
}

func newCPUTrackerMetrics(reg prometheus.Registerer) (*trackerMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
