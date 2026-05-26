// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package protoutils

import (
	"github.com/ava-labs/avalanchego/utils/maybe"

	pb "github.com/ava-labs/avalanchego/proto/pb/sync"
)

func MaybeToProto(m maybe.Maybe[[]byte]) *pb.MaybeBytes { _ = "STUB: not implemented"; return nil }

func ProtoToMaybe(mb *pb.MaybeBytes) maybe.Maybe[[]byte] { _ = "STUB: not implemented"; return nil }
