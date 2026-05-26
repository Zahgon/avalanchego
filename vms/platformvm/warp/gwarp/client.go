// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gwarp

import (
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"

	pb "github.com/ava-labs/avalanchego/proto/pb/warp"
)

var _ warp.Signer = (*Client)(nil)

type Client struct {
	client pb.SignerClient
}

func NewClient(client pb.SignerClient) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Sign(unsignedMsg *warp.UnsignedMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
