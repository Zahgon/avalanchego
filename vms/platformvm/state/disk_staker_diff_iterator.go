// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"fmt"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

const (
	// startDiffKey = [subnetID] + [inverseHeight]
	startDiffKeyLength = ids.IDLen + database.Uint64Size
	// diffKey = [subnetID] + [inverseHeight] + [nodeID]
	diffKeyLength = startDiffKeyLength + ids.NodeIDLen
	// diffKeyNodeIDOffset = [subnetIDLen] + [inverseHeightLen]
	diffKeyNodeIDOffset = ids.IDLen + database.Uint64Size

	// weightValue = [isNegative] + [weight]
	weightValueLength = database.BoolSize + database.Uint64Size
)

var (
	errUnexpectedDiffKeyLength     = fmt.Errorf("expected diff key length %d", diffKeyLength)
	errUnexpectedWeightValueLength = fmt.Errorf("expected weight value length %d", weightValueLength)
)

// marshalStartDiffKeyBySubnetID is used to determine the starting key when iterating.
//
// Invariant: the result is a prefix of [marshalDiffKeyBySubnetID] when called with the
// same arguments.
func marshalStartDiffKeyBySubnetID(subnetID ids.ID, height uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// marshalStartDiffKeyByHeight is used to determine the starting key when iterating.
//
// Invariant: the result is a prefix of [marshalDiffKeyByHeight] when called with the
// same arguments.
func marshalStartDiffKeyByHeight(height uint64) []byte { _ = "STUB: not implemented"; return nil }

func marshalDiffKeyBySubnetID(subnetID ids.ID, height uint64, nodeID ids.NodeID) []byte {
	_ = "STUB: not implemented"
	return nil
}

func marshalDiffKeyByHeight(height uint64, subnetID ids.ID, nodeID ids.NodeID) []byte {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalDiffKeyBySubnetID(key []byte) (ids.ID, uint64, ids.NodeID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), 0, *new(ids.NodeID), nil
}

func unmarshalDiffKeyByHeight(key []byte) (uint64, ids.ID, ids.NodeID, error) {
	_ = "STUB: not implemented"
	return 0, *new(ids.ID), *new(ids.NodeID), nil
}

func marshalWeightDiff(diff *ValidatorWeightDiff) []byte { _ = "STUB: not implemented"; return nil }

func unmarshalWeightDiff(value []byte) (*ValidatorWeightDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: [height] is encoded as a bit flipped big endian number so that
// iterating lexicographically results in iterating in decreasing heights.
//
// Invariant: [key] has sufficient length
func packIterableHeight(key []byte, height uint64) { _ = "STUB: not implemented"; return }

// Because we bit flip the height when constructing the key, we must remember to
// bip flip again here.
//
// Invariant: [key] has sufficient length
func unpackIterableHeight(key []byte) uint64 { _ = "STUB: not implemented"; return 0 }
