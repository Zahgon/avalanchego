// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

// TODO: Consider renaming Message to, say, VMMessage

// Message is an enum of the message types that vms can send to consensus
type Message uint32

const (
	// PendingTxs notifies a consensus engine that its VM has pending
	// transactions.
	//
	// The consensus engine must eventually call BuildBlock at least once after
	// receiving this message. If the consensus engine receives multiple
	// PendingTxs messages between calls to BuildBlock, the engine may only call
	// BuildBlock once.
	PendingTxs Message = iota + 1

	// StateSyncDone notifies the state syncer engine that the VM has finishing
	// syncing the requested state summary.
	StateSyncDone
)

func (msg Message) String() string { _ = "STUB: not implemented"; return "" }
