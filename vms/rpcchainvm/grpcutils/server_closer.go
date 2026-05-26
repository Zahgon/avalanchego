// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package grpcutils

import (
	"sync"

	"google.golang.org/grpc"
)

type ServerCloser struct {
	lock    sync.Mutex
	closed  bool
	servers []*grpc.Server
}

func (s *ServerCloser) Add(server *grpc.Server) { _ = "STUB: not implemented"; return }

func (s *ServerCloser) Stop() { _ = "STUB: not implemented"; return }

func (s *ServerCloser) GracefulStop() { _ = "STUB: not implemented"; return }
