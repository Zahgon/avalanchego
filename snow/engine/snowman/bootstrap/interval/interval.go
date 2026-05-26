// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package interval

type Interval struct {
	LowerBound uint64
	UpperBound uint64
}

func (i *Interval) Less(other *Interval) bool { _ = "STUB: not implemented"; return false }

func (i *Interval) Contains(height uint64) bool { _ = "STUB: not implemented"; return false }

// AdjacentToLowerBound returns true if height is 1 less than lowerBound.
func (i *Interval) AdjacentToLowerBound(height uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// AdjacentToUpperBound returns true if height is 1 greater than upperBound.
func (i *Interval) AdjacentToUpperBound(height uint64) bool {
	_ = "STUB: not implemented"
	return false
}
