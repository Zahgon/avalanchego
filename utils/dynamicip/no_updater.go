// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dynamicip

import "github.com/ava-labs/avalanchego/utils/logging"

var _ Updater = noUpdater{}

func NewNoUpdater() Updater { _ = "STUB: not implemented"; return *new(Updater) }

type noUpdater struct{}

func (noUpdater) Dispatch(logging.Logger) { _ = "STUB: not implemented"; return }

func (noUpdater) Stop() { _ = "STUB: not implemented"; return }
