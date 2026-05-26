// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

import (
	"errors"
	"sync"
)

var (
	ErrNoIDWithAlias      = errors.New("there is no ID with alias")
	errNoAliasForID       = errors.New("there is no alias for ID")
	errAliasAlreadyMapped = errors.New("alias already mapped to an ID")
)

// AliaserReader allows one to lookup the aliases given to an ID.
type AliaserReader interface {
	// Lookup returns the ID associated with alias
	Lookup(alias string) (ID, error)

	// PrimaryAlias returns the first alias of [id]
	PrimaryAlias(id ID) (string, error)

	// Aliases returns the aliases of an ID
	Aliases(id ID) ([]string, error)
}

// AliaserWriter allows one to give an ID aliases. An ID can have arbitrarily
// many aliases; two IDs may not have the same alias.
type AliaserWriter interface {
	// Alias gives [id] the alias [alias]
	Alias(id ID, alias string) error

	// RemoveAliases of the provided ID
	RemoveAliases(id ID)
}

// Aliaser allows one to give an ID aliases and lookup the aliases given to an
// ID.
type Aliaser interface {
	AliaserReader
	AliaserWriter

	// PrimaryAliasOrDefault returns the first alias of [id], or ID string as a
	// default if no alias exists
	PrimaryAliasOrDefault(id ID) string
}

type aliaser struct {
	lock    sync.RWMutex
	dealias map[string]ID
	aliases map[ID][]string
}

func NewAliaser() Aliaser { _ = "STUB: not implemented"; return *new(Aliaser) }

func (a *aliaser) Lookup(alias string) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func (a *aliaser) PrimaryAlias(id ID) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (a *aliaser) PrimaryAliasOrDefault(id ID) string { _ = "STUB: not implemented"; return "" }

func (a *aliaser) Aliases(id ID) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *aliaser) Alias(id ID, alias string) error { _ = "STUB: not implemented"; return nil }

func (a *aliaser) RemoveAliases(id ID) { _ = "STUB: not implemented"; return }

// GetRelevantAliases returns the aliases with the redundant identity alias
// removed (each id is aliased to at least itself).
func GetRelevantAliases(aliaser Aliaser, ids []ID) (map[ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove the redundant alias where alias = id.
