// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mockable

import "time"

// MaxTime was taken from https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go/32620397#32620397
var MaxTime = time.Unix(1<<63-62135596801, 0) // 0 is used because we drop the nano-seconds

// Clock acts as a thin wrapper around global time that allows for easy testing
type Clock struct {
	faked bool
	time  time.Time
}

// Set the time on the clock
func (c *Clock) Set(time time.Time) { _ = "STUB: not implemented"; return }

// Sync this clock with global time
func (c *Clock) Sync() {
	_ = "STUB: not implemented"

	// Time returns the time on this clock
	return
}

func (c *Clock) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Time returns the unix time on this clock
func (c *Clock) UnixTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Unix returns the unix timestamp on this clock.
func (c *Clock) Unix() uint64 { _ = "STUB: not implemented"; return 0 }
