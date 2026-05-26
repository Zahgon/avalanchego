// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"errors"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errDuplicatePut    = errors.New("duplicate put")
	errDuplicateRemove = errors.New("duplicate remove")
)

type dbElement struct {
	// Present indicates the value was removed before existing.
	// If set to false, when this element is added to the shared memory, it will
	// be immediately removed.
	// If set to true, then this element will be removed normally when remove is
	// called.
	Present bool `serialize:"true"`

	// Value is the body of this element.
	Value []byte `serialize:"true"`

	// Traits are a collection of features that can be used to lookup this
	// element.
	Traits [][]byte `serialize:"true"`
}

// state is used to maintain a mapping from keys to dbElements and an index that
// maps traits to the keys with those traits.
type state struct {
	// valueDB contains a mapping from key to the corresponding dbElement.
	valueDB database.Database

	// indexDB stores the trait -> key mappings.
	// To get this mapping, we construct a prefixdb using the trait as a prefix
	// and then construct a linkeddb on the result.
	// The linkeddb contains the keys that the trait maps to as the key and map
	// to nil values.
	indexDB database.Database
}

// Value returns the Element associated with [key].
func (s *state) Value(key []byte) (*Element, error) { _ = "STUB: not implemented"; return nil, nil }

// If [key] is indexed, but has been marked as deleted, return
// [database.ErrNotFound].

// SetValue places the element [e] into the state and maps each of the traits of
// the element to its key, so that the element can be looked up by any of its
// traits.
//
// If the value was preemptively marked as deleted, then the operations cancel
// and the removal marker is deleted.
func (s *state) SetValue(e *Element) error { _ = "STUB: not implemented"; return nil }

// The key was already registered with the state.

// This was previously optimistically deleted from the database, so
// it should be immediately removed.

// This key was written twice, which is invalid

// An unexpected error occurred, so we should propagate that error

// RemoveValue removes [key] from the state.
// This operation removes the element associated with the key from both the
// valueDB and the indexDB.
// If [key] has already been marked as removed, an error is returned as the same
// element cannot be removed twice.
// If [key] is not present in the db, then it is marked as deleted so that it
// will be removed immediately when it gets added in the future.
//
// This ensures that we can consume a UTXO before it has been added into shared
// memory in bootstrapping.
// Ex. P-Chain attempts to consume atomic UTXO from the C-Chain in block 100.
// P-Chain executes before the C-Chain, so when bootstrapping it must be able to
// verify and accept this block before the node has processed the block on the
// C-Chain where the atomic UTXO is added to shared memory.
// Additionally, when the C-Chain actually does add the atomic UTXO to shared
// memory, RemoveValue must handle the case that the atomic UTXO was marked as
// deleted before it was actually added.
// To do this, the node essentially adds a tombstone marker when the P-Chain
// consumes the non-existent UTXO, which is deleted when the C-Chain actually
// adds the atomic UTXO to shared memory.
//
// This implies that chains interacting with shared memory must be able to
// generate their chain state without actually performing the read of shared
// memory. Shared memory should only be used to verify that the transition
// being performed is valid. That ensures that such verification can be skipped
// during bootstrapping. It is up to the chain to ensure this based on the
// current engine state.
func (s *state) RemoveValue(key []byte) error { _ = "STUB: not implemented"; return nil }

// The value doesn't exist, so we should optimistically delete it

// Don't allow the removal of something that was already removed.

// Remove [key] from the indexDB for each trait that has indexed this key.

// loadValue retrieves the dbElement corresponding to [key] from the value
// database.
func (s *state) loadValue(key []byte) (*dbElement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The key was in the database

// getKeys returns up to [limit] keys starting at [startTrait]/[startKey] which
// possess at least one of the specified [traits].
// Returns the set of keys possessing traits and the ending [lastTrait] and
// [lastKey] to use as an index for pagination.
func (s *state) getKeys(traits [][]byte, startTrait, startKey []byte, limit int) ([][]byte, []byte, []byte, error) {
	_ = "STUB: not implemented"
	// Note: We use a reference to [keySet] since otherwise, depending on how
	//       this variable is declared, the map may not be initialized from the
	//       start. The first add to the underlying map of the set would then
	//       result in the map being initialized.
	return nil, nil, nil, nil
}

// Iterate over the traits in order appending all of the keys that possess
// the given [traits].

// Skip the trait, if we have already paginated past it.

// We are already past [startTrait]/[startKey], so we should now
// start indexing all of [trait].

// Add any additional keys that possess [trait] to [keys].

// Return the [keys] that we found as well as the index given by [lastTrait]
// and [lastKey].

// appendTraitKeys iterates over the indexDB of [trait] starting at [startKey]
// and adds keys that possess [trait] to [keys] until the iteration completes or
// limit hits 0. If a key possesses multiple traits, it will be de-duplicated
// with [keySet].
func (s *state) appendTraitKeys(keys *[][]byte, keySet *set.Set[ids.ID], limit *int, trait, startKey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create the prefixDB for the specified trait.
		nil
}

// Interpret [traitDB] as a linkeddb that contains a set of keys.

// Iterate over the keys that possess [trait].

// Calculate the hash of the key to check against the set and ensure
// we don't add the same element twice.
