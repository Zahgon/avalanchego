// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package galiasreader

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"

	aliasreaderpb "github.com/ava-labs/avalanchego/proto/pb/aliasreader"
)

var _ aliasreaderpb.AliasReaderServer = (*Server)(nil)

// Server enables alias lookups over RPC.
type Server struct {
	aliasreaderpb.UnsafeAliasReaderServer
	aliaser ids.AliaserReader
}

// NewServer returns an alias lookup connected to a remote alias lookup
func NewServer(aliaser ids.AliaserReader) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Lookup(
	_ context.Context,
	req *aliasreaderpb.Alias,
) (*aliasreaderpb.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) PrimaryAlias(
	_ context.Context,
	req *aliasreaderpb.ID,
) (*aliasreaderpb.Alias, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Aliases(
	_ context.Context,
	req *aliasreaderpb.ID,
) (*aliasreaderpb.AliasList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
