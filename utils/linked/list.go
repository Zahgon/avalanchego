// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linked

// ListElement is an element of a linked list.
type ListElement[T any] struct {
	next, prev *ListElement[T]
	list       *List[T]
	Value      T
}

// Next returns the next element or nil.
func (e *ListElement[T]) Next() *ListElement[T] { _ = "STUB: not implemented"; return nil }

// Prev returns the previous element or nil.
func (e *ListElement[T]) Prev() *ListElement[T] { _ = "STUB: not implemented"; return nil }

// List implements a doubly linked list with a sentinel node.
//
// See: https://en.wikipedia.org/wiki/Doubly_linked_list
//
// This datastructure is designed to be an almost complete drop-in replacement
// for the standard library's "container/list".
//
// The primary design change is to remove all memory allocations from the list
// definition. This allows these lists to be used in performance critical paths.
// Additionally the zero value is not useful. Lists must be created with the
// NewList method.
type List[T any] struct {
	// sentinel is only used as a placeholder to avoid complex nil checks.
	// sentinel.Value is never used.
	sentinel ListElement[T]
	length   int
}

// NewList creates a new doubly linked list.
func NewList[T any]() *List[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in l.
func (l *List[_]) Len() int {
	_ = "STUB: not implemented"

	// Front returns the element at the front of l.
	// If l is empty, nil is returned.
	return 0
}

func (l *List[T]) Front() *ListElement[T] { _ = "STUB: not implemented"; return nil }

// Back returns the element at the back of l.
// If l is empty, nil is returned.
func (l *List[T]) Back() *ListElement[T] { _ = "STUB: not implemented"; return nil }

// Remove removes e from l if e is in l.
func (l *List[T]) Remove(e *ListElement[T]) { _ = "STUB: not implemented"; return }

// PushFront inserts e at the front of l.
// If e is already in a list, l is not modified.
func (l *List[T]) PushFront(e *ListElement[T]) { _ = "STUB: not implemented"; return }

// PushBack inserts e at the back of l.
// If e is already in a list, l is not modified.
func (l *List[T]) PushBack(e *ListElement[T]) { _ = "STUB: not implemented"; return }

// InsertBefore inserts e immediately before location.
// If e is already in a list, l is not modified.
// If location is not in l, l is not modified.
func (l *List[T]) InsertBefore(e *ListElement[T], location *ListElement[T]) {
	_ = "STUB: not implemented"
	return
}

// InsertAfter inserts e immediately after location.
// If e is already in a list, l is not modified.
// If location is not in l, l is not modified.
func (l *List[T]) InsertAfter(e *ListElement[T], location *ListElement[T]) {
	_ = "STUB: not implemented"
	return
}

// MoveToFront moves e to the front of l.
// If e is not in l, l is not modified.
func (l *List[T]) MoveToFront(e *ListElement[T]) {
	_ = "STUB: not implemented"
	// If e is already at the front of l, there is nothing to do.
	return
}

// MoveToBack moves e to the back of l.
// If e is not in l, l is not modified.
func (l *List[T]) MoveToBack(e *ListElement[T]) { _ = "STUB: not implemented"; return }

// MoveBefore moves e immediately before location.
// If the elements are equal or not in l, the list is not modified.
func (l *List[T]) MoveBefore(e, location *ListElement[T]) {
	_ = "STUB: not implemented"
	// Don't introduce a cycle by moving an element before itself.
	return
}

// MoveAfter moves e immediately after location.
// If the elements are equal or not in l, the list is not modified.
func (l *List[T]) MoveAfter(e, location *ListElement[T]) { _ = "STUB: not implemented"; return }

func (l *List[T]) insertAfter(e, location *ListElement[T]) {
	_ = "STUB: not implemented"

	// Don't insert an element that is already in a list
	return
}

func (l *List[T]) moveAfter(e, location *ListElement[T]) { _ = "STUB: not implemented"; return }

// Don't modify an element that is in a different list.
// Don't introduce a cycle by moving an element after itself.

// PushFront inserts v into a new element at the front of l.
func PushFront[T any](l *List[T], v T) { _ = "STUB: not implemented"; return }

// PushBack inserts v into a new element at the back of l.
func PushBack[T any](l *List[T], v T) { _ = "STUB: not implemented"; return }

// InsertBefore inserts v into a new element immediately before location.
// If location is not in l, l is not modified.
func InsertBefore[T any](l *List[T], v T, location *ListElement[T]) {
	_ = "STUB: not implemented"
	return
}

// InsertAfter inserts v into a new element immediately after location.
// If location is not in l, l is not modified.
func InsertAfter[T any](l *List[T], v T, location *ListElement[T]) {
	_ = "STUB: not implemented"
	return
}
