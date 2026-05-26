// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"context"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/ids"
)

// SubnetEVMLeafsRequest preserves the original subnet-evm wire format where NodeType is not serialized.
//
// Note: The fields Root, Account, Start, End, and Limit are duplicated in CorethLeafsRequest.
// This duplication is intentional because the avalanchego codec requires all serialized fields
// to be exported (start with uppercase). Using an embedded struct would require exporting it,
// which would unnecessarily expose internal types in the public API. By duplicating fields,
// we keep the implementation details private while ensuring correct serialization.
//
// NOTE: NodeType is not serialized to maintain backward compatibility with subnet-evm nodes.
//
// TODO: In a future network upgrade, there will be only one unified wire format for both
// coreth and subnet-evm. At that point, this distinction between CorethLeafsRequest and
// SubnetEVMLeafsRequest will be removed.
type SubnetEVMLeafsRequest struct {
	Root     common.Hash `serialize:"true"`
	Account  common.Hash `serialize:"true"`
	Start    []byte      `serialize:"true"`
	End      []byte      `serialize:"true"`
	Limit    uint16      `serialize:"true"`
	NodeType NodeType
}

func (s SubnetEVMLeafsRequest) Handle(ctx context.Context, nodeID ids.NodeID, requestID uint32, handler RequestHandler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s SubnetEVMLeafsRequest) String() string { _ = "STUB: not implemented"; return "" }

func (s SubnetEVMLeafsRequest) RootHash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}
func (s SubnetEVMLeafsRequest) AccountHash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}
func (s SubnetEVMLeafsRequest) StartKey() []byte { _ = "STUB: not implemented"; return nil }
func (s SubnetEVMLeafsRequest) EndKey() []byte   { _ = "STUB: not implemented"; return nil }
func (s SubnetEVMLeafsRequest) KeyLimit() uint16 { _ = "STUB: not implemented"; return 0 }
func (s SubnetEVMLeafsRequest) LeafType() NodeType {
	_ = "STUB: not implemented"
	return *new(NodeType)
}
