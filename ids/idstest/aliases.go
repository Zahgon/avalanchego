// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package codectest provides a test suite for testing functionality related to
// IDs.
package idstest

import (
	"testing"

	"github.com/ava-labs/avalanchego/ids"
)

// An AliasTest couples a test in the Aliaser suite with a human-readable name.
type AliasTest struct {
	Name string
	Test func(testing.TB, ids.AliaserReader, ids.AliaserWriter)
}

// Run runs the test on the Aliaser{Reader+Writer} pair.
func (tt *AliasTest) Run(t *testing.T, r ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

// RunAllAlias runs all [AliasTests], constructing a new GeneralCodec for each.
func RunAllAlias(t *testing.T, ctor func() (ids.AliaserReader, ids.AliaserWriter)) {
	_ = "STUB: not implemented"
	return
}

var AliasTests = []AliasTest{
	{"Lookup Error}", TestAliaserLookupError},
	{"Lookup}", TestAliaserLookup},
	{"Aliases Empty}", TestAliaserAliasesEmpty},
	{"Aliases}", TestAliaserAliases},
	{"Primary Alias}", TestAliaserPrimaryAlias},
	{"Alias Clash}", TestAliaserAliasClash},
	{"Remove Alias}", TestAliaserRemoveAlias},
}

func TestAliaserLookupError(tb testing.TB, r ids.AliaserReader, _ ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errNoIDWithAlias
//nolint:forbidigo // currently returns grpc errors too

func TestAliaserLookup(tb testing.TB, r ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

func TestAliaserAliasesEmpty(tb testing.TB, r ids.AliaserReader, _ ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

func TestAliaserAliases(tb testing.TB, r ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

func TestAliaserPrimaryAlias(tb testing.TB, r ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errNoAliasForID
//nolint:forbidigo // currently returns grpc errors too

func TestAliaserAliasClash(tb testing.TB, _ ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errAliasAlreadyMapped
//nolint:forbidigo // currently returns grpc errors too

func TestAliaserRemoveAlias(tb testing.TB, r ids.AliaserReader, w ids.AliaserWriter) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errNoAliasForID
//nolint:forbidigo // currently returns grpc errors too
