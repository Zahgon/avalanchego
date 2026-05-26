// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package wrappers

type Errs struct{ Err error }

func (errs *Errs) Errored() bool { _ = "STUB: not implemented"; return false }

func (errs *Errs) Add(errors ...error) { _ = "STUB: not implemented"; return }
