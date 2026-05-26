// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"errors"

	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var (
	errInvalidMessageLength     = errors.New("invalid message length")
	errMaxMessageLengthExceeded = errors.New("maximum message length exceeded")
)

func writeMsgLen(msgLen uint32, maxMsgLen uint32) ([wrappers.IntLen]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readMsgLen(b []byte, maxMsgLen uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// parse the message length
