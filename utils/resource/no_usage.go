// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package resource

// NoUsage implements Usage() by always returning 0.
var NoUsage User = noUsage{}

type noUsage struct{}

func (noUsage) AvailableDiskPercentage() uint64 { _ = "STUB: not implemented"; return 0 }

func (noUsage) CPUUsage() float64 { _ = "STUB: not implemented"; return 0 }

func (noUsage) DiskUsage() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func (noUsage) AvailableDiskBytes() uint64 { _ = "STUB: not implemented"; return 0 }
