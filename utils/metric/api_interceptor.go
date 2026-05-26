// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metric

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/prometheus/client_golang/prometheus"
)

type APIInterceptor interface {
	InterceptRequest(i *rpc.RequestInfo) *http.Request
	AfterRequest(i *rpc.RequestInfo)
}

type contextKey int

const requestTimestampKey contextKey = iota

type apiInterceptor struct {
	requestDurationCount *prometheus.CounterVec
	requestDurationSum   *prometheus.GaugeVec
	requestErrors        *prometheus.CounterVec
}

func NewAPIInterceptor(registerer prometheus.Registerer) (APIInterceptor, error) {
	_ = "STUB: not implemented"
	return *new(APIInterceptor), nil
}

func (*apiInterceptor) InterceptRequest(i *rpc.RequestInfo) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func (apr *apiInterceptor) AfterRequest(i *rpc.RequestInfo) { _ = "STUB: not implemented"; return }
