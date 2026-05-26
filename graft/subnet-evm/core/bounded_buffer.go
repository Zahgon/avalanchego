// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

// BoundedBuffer keeps [size] entries of type [K] in a buffer and calls
// [callback] on any item that is overwritten. This is typically used for
// dereferencing old roots during block processing.
//
// BoundedBuffer is not thread-safe and requires the caller synchronize usage.
type BoundedBuffer[K any] struct {
	lastPos  int
	size     int
	callback func(K) error
	buffer   []K

	cycled bool
}

// NewBoundedBuffer creates a new [BoundedBuffer].
func NewBoundedBuffer[K any](size int, callback func(K) error) *BoundedBuffer[K] {
	_ = "STUB: not implemented"
	return nil
}

// Insert adds a new value to the buffer. If the buffer is full, the
// oldest value will be overwritten and [callback] will be invoked.
func (b *BoundedBuffer[K]) Insert(h K) error { _ = "STUB: not implemented"; return nil }

// the first item added to the buffer will be at position 0

// Set [cycled] since we are back to the 0th element

// We ensure we have cycled through the buffer once before invoking the
// [callback] to ensure we don't call it with unset values.

// Last retrieves the last item added to the buffer.
//
// If no items have been added to the buffer, Last returns the default value of
// [K] and [false].
func (b *BoundedBuffer[K]) Last() (K, bool) { _ = "STUB: not implemented"; return *new(K), false }
