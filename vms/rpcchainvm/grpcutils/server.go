// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package grpcutils

import (
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// MinTime is the minimum amount of time a client should wait before sending
	// a keepalive ping. grpc-go default 5 mins
	defaultServerKeepAliveMinTime = 5 * time.Second
	// After a duration of this time if the server doesn't see any activity it
	// pings the client to see if the transport is still alive.
	// If set below 1s, a minimum value of 1s will be used instead.
	// grpc-go default 2h
	defaultServerKeepAliveInterval = 2 * time.Hour
	// After having pinged for keepalive check, the server waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed. grpc-go default 20s
	defaultServerKeepAliveTimeout = 20 * time.Second
)

var DefaultServerOptions = []grpc.ServerOption{
	grpc.MaxRecvMsgSize(math.MaxInt),
	grpc.MaxSendMsgSize(math.MaxInt),
	grpc.MaxConcurrentStreams(math.MaxUint32),
	grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             defaultServerKeepAliveMinTime,
		PermitWithoutStream: defaultPermitWithoutStream,
	}),
	grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    defaultServerKeepAliveInterval,
		Timeout: defaultServerKeepAliveTimeout,
	}),
}

// NewServer will return a gRPC server with server options as defined by
// DefaultServerOptions. ServerOption can also optionally be passed.
func NewServer(opts ...ServerOption) *grpc.Server { _ = "STUB: not implemented"; return nil }

type ServerOptions struct {
	opts []grpc.ServerOption
}

// append(DefaultServerOptions, ...) will always allocate a new slice and will
// not overwrite any potential data that may have previously been appended to
// DefaultServerOptions https://go.dev/ref/spec#Composite_literals
func newServerOpts(opts []ServerOption) []grpc.ServerOption { _ = "STUB: not implemented"; return nil }

func (s *ServerOptions) applyOpts(opts []ServerOption) { _ = "STUB: not implemented"; return }

// ServerOption are options which can be applied to a gRPC server in addition to
// the defaults set by DefaultServerOPtions.
type ServerOption func(*ServerOptions)

// WithUnaryInterceptor adds a single unary interceptor to the gRPC server
// options.
func WithUnaryInterceptor(unaryInterceptor grpc.UnaryServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithStreamInterceptor adds a single stream interceptor to the gRPC server
// options.
func WithStreamInterceptor(streamInterceptor grpc.StreamServerInterceptor) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// NewListener returns a TCP listener listening against the next available port
// on the system bound to localhost.
func NewListener() (net.Listener, error) { _ = "STUB: not implemented"; return *new(net.Listener), nil }

// Serve will start a gRPC server and block until it errors or is shutdown.
func Serve(listener net.Listener, grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	// TODO: While errors will be reported later, it could be useful to somehow
	//
	//	log this if it is the primary error.
	//
	// There is nothing to with the error returned by serve here. Later requests
	// will propagate their error if they occur.
	return
}

// Similarly, there is nothing to with an error when the listener is closed.
