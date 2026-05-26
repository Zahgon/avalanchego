// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package trace

import (
	"errors"
)

const (
	Disabled ExporterType = iota
	GRPC
	HTTP
)

var (
	errUnknownExporterType = errors.New("unknown exporter type")
	errMissingQuotes       = errors.New("first and last characters should be quotes")
)

func ExporterTypeFromString(exporterTypeStr string) (ExporterType, error) {
	_ = "STUB: not implemented"
	return *new(ExporterType), nil
}

type ExporterType byte

func (t ExporterType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *ExporterType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If "null", do nothing

func (t ExporterType) String() string { _ = "STUB: not implemented"; return "" }

func (t ExporterType) toString() (string, bool) { _ = "STUB: not implemented"; return "", false }
