// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import (
	"github.com/ava-labs/avalanchego/database"
)

func (v *view) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (v *view) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (v *view) NewIteratorWithPrefix(prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (v *view) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// Binary search for [startKey] index.

// Since [sortedKeyChanges] is sorted, if the prefix isnt found anymore after we
// added at least one [KeyChange], we can stop.

// viewIterator walks over both the in memory database and the underlying database
// at the same time.
type viewIterator struct {
	view       *view
	parentIter database.Iterator

	key, value []byte
	err        error

	sortedChanges []KeyChange

	initialized, parentIterExhausted bool
}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted. We must pay careful attention to set the proper values
// based on if the in memory changes or the underlying db should be read next
func (it *viewIterator) Next() bool { _ = "STUB: not implemented"; return false }

// there are no more changes or underlying key/values

// there are no more underlying key/values, so use the local changes

// move to next change

// If current change is not a deletion, return it.
// Otherwise go to next loop iteration.

// The current change has a smaller key than the parent key.
// Move to the next change.

// If current change is not a deletion, return it.
// Otherwise, go to next loop iteration.

// The parent key is smaller, so return it and iterate the parent iterator

// the keys are the same, so use the local change and
// iterate both the sorted changes and the parent iterator

func (it *viewIterator) Error() error { _ = "STUB: not implemented"; return nil }

func (it *viewIterator) Key() []byte { _ = "STUB: not implemented"; return nil }

func (it *viewIterator) Value() []byte { _ = "STUB: not implemented"; return nil }

func (it *viewIterator) Release() { _ = "STUB: not implemented"; return }
