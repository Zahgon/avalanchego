// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

var offset = uint64(0)

// GenerateTestID returns a new ID that should only be used for testing
func GenerateTestID() ID { _ = "STUB: not implemented"; return *new(ID) }

// GenerateTestShortID returns a new ID that should only be used for testing
func GenerateTestShortID() ShortID { _ = "STUB: not implemented"; return *new(ShortID) }

// GenerateTestNodeID returns a new ID that should only be used for testing
func GenerateTestNodeID() NodeID { _ = "STUB: not implemented"; return *new(NodeID) }

// BuildTestNodeID is an utility to build NodeID from bytes in UTs
// It must not be used in production code. In production code we should
// use ToNodeID, which performs proper length checking.
func BuildTestNodeID(src []byte) NodeID { _ = "STUB: not implemented"; return *new(NodeID) }
