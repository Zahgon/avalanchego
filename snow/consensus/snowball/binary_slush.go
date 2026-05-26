// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

func newBinarySlush(choice int) binarySlush { _ = "STUB: not implemented"; return *new(binarySlush) }

// binarySlush is the implementation of a binary slush instance
type binarySlush struct {
	// preference is the choice that last had a successful poll. Unless there
	// hasn't been a successful poll, in which case it is the initially provided
	// choice.
	preference int
}

func (sl *binarySlush) Preference() int { _ = "STUB: not implemented"; return 0 }

func (sl *binarySlush) RecordSuccessfulPoll(choice int) { _ = "STUB: not implemented"; return }

func (sl *binarySlush) String() string { _ = "STUB: not implemented"; return "" }
