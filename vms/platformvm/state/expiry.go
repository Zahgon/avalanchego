// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"fmt"

	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/iterator"
)

// expiryEntry = [timestamp] + [validationID]
const expiryEntryLength = database.Uint64Size + ids.IDLen

var (
	errUnexpectedExpiryEntryLength = fmt.Errorf("expected expiry entry length %d", expiryEntryLength)

	_ btree.LessFunc[ExpiryEntry] = ExpiryEntry.Less
	_ utils.Sortable[ExpiryEntry] = ExpiryEntry{}
)

type Expiry interface {
	// GetExpiryIterator returns an iterator of all the expiry entries in order
	// of lowest to highest timestamp.
	GetExpiryIterator() (iterator.Iterator[ExpiryEntry], error)

	// HasExpiry returns true if the database has the specified entry.
	HasExpiry(ExpiryEntry) (bool, error)

	// PutExpiry adds the entry to the database. If the entry already exists, it
	// is a noop.
	PutExpiry(ExpiryEntry)

	// DeleteExpiry removes the entry from the database. If the entry doesn't
	// exist, it is a noop.
	DeleteExpiry(ExpiryEntry)
}

type ExpiryEntry struct {
	Timestamp    uint64
	ValidationID ids.ID
}

func (e *ExpiryEntry) Marshal() []byte { _ = "STUB: not implemented"; return nil }

func (e *ExpiryEntry) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (e ExpiryEntry) Less(o ExpiryEntry) bool { _ = "STUB: not implemented"; return false }

// Invariant: Compare produces the same ordering as the marshalled bytes.
func (e ExpiryEntry) Compare(o ExpiryEntry) int { _ = "STUB: not implemented"; return 0 }

type expiryDiff struct {
	modified map[ExpiryEntry]bool // bool represents isAdded
	added    *btree.BTreeG[ExpiryEntry]
}

func newExpiryDiff() *expiryDiff { _ = "STUB: not implemented"; return nil }

func (e *expiryDiff) PutExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

func (e *expiryDiff) DeleteExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

func (e *expiryDiff) getExpiryIterator(parentIterator iterator.Iterator[ExpiryEntry]) iterator.Iterator[ExpiryEntry] {
	_ = "STUB: not implemented"
	return nil
}
