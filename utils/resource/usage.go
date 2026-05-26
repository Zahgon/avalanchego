// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package resource

import (
	"math"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/process"

	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	lnHalf = math.Log(.5)

	_ Manager = (*manager)(nil)
)

type CPUUser interface {
	// CPUUsage returns the number of CPU cores of usage this user has attributed
	// to it.
	//
	// For example, if this user is reporting a process's CPU utilization and
	// that process is currently using 150% CPU (i.e. one and a half cores of
	// compute) then the return value will be 1.5.
	CPUUsage() float64
}

type DiskUser interface {
	// DiskUsage returns the number of bytes per second read from/written to
	// disk recently.
	DiskUsage() (read float64, write float64)

	// returns number of bytes available in the db volume
	AvailableDiskBytes() uint64

	// returns percentage free in the db volume
	AvailableDiskPercentage() uint64
}

type User interface {
	CPUUser
	DiskUser
}

type ProcessTracker interface {
	// TrackProcess adds [pid] to the list of processes that this tracker is
	// currently managing. Duplicate requests are dropped.
	TrackProcess(pid int)

	// UntrackProcess removes [pid] from the list of processes that this tracker
	// is currently managing. Untracking a currently untracked [pid] is a noop.
	UntrackProcess(pid int)
}

type Manager interface {
	User
	ProcessTracker

	// Shutdown allocated resources and stop tracking all processes.
	Shutdown()
}

type manager struct {
	log            logging.Logger
	processMetrics *metrics

	processesLock sync.Mutex
	processes     map[int]*proc

	usageLock sync.RWMutex
	cpuUsage  float64
	// [readUsage] is the number of bytes/second read from disk recently.
	readUsage float64
	// [writeUsage] is the number of bytes/second written to disk recently.
	writeUsage float64

	availableDiskBytes uint64

	availableDiskPercent uint64

	closeOnce sync.Once
	onClose   chan struct{}
}

func NewManager(
	log logging.Logger,
	diskPath string,
	frequency,
	cpuHalflife,
	diskHalflife time.Duration,
	metricsRegisterer prometheus.Registerer,
) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

func (m *manager) CPUUsage() float64 { _ = "STUB: not implemented"; return 0 }

func (m *manager) DiskUsage() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func (m *manager) AvailableDiskBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (m *manager) AvailableDiskPercentage() uint64 { _ = "STUB: not implemented"; return 0 }

func (m *manager) TrackProcess(pid int) { _ = "STUB: not implemented"; return }

func (m *manager) UntrackProcess(pid int) { _ = "STUB: not implemented"; return }

func (m *manager) Shutdown() { _ = "STUB: not implemented"; return }

func (m *manager) update(diskPath string, frequency, cpuHalflife, diskHalflife time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Returns:
// 1. Current CPU usage by all processes.
// 2. Current bytes/sec read from disk by all processes.
// 3. Current bytes/sec written to disk by all processes.
func (m *manager) getActiveUsage(secondsSinceLastUpdate float64) (float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

type proc struct {
	p   *process.Process
	log logging.Logger

	initialized bool

	// [lastTotalCPU] is the most recent measurement of total CPU usage.
	lastTotalCPU float64

	// [numReads] is the total number of disk reads performed.
	numReads uint64
	// [lastReadBytes] is the most recent measurement of total disk bytes read.
	lastReadBytes uint64

	// [numWrites] is the total number of disk writes performed.
	numWrites uint64
	// [lastWriteBytes] is the most recent measurement of total disk bytes
	// written.
	lastWriteBytes uint64
}

func (p *proc) getActiveUsage(secondsSinceLastUpdate float64) (float64, float64, float64) {
	_ = "STUB: not implemented"
	// If there is an error tracking the CPU/disk utilization of a process,
	// assume that the utilization is 0.
	return 0, 0, 0
}

// Note: IOCounters is not implemented on macos and therefore always returns
// an error on macos.

// getSampleWeights converts the frequency of CPU sampling and the halflife of
// the CPU sample's usefulness into weights to scale the newly sampled point and
// previously samples.
func getSampleWeights(frequency, halflife time.Duration) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}
