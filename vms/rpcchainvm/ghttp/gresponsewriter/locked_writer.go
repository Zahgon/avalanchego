// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gresponsewriter

import (
	"bufio"
	"net"
	"net/http"
	"sync"
)

var (
	_ http.ResponseWriter = (*lockedWriter)(nil)
	_ http.Flusher        = (*lockedWriter)(nil)
	_ http.Hijacker       = (*lockedWriter)(nil)
)

type lockedWriter struct {
	lock          sync.Mutex
	writer        http.ResponseWriter
	headerWritten bool
}

func NewLockedWriter(w http.ResponseWriter) http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func (lw *lockedWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (lw *lockedWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (lw *lockedWriter) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

// Skip writing the header if it has already been written once.

func (lw *lockedWriter) Flush() { _ = "STUB: not implemented"; return }

func (lw *lockedWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}
