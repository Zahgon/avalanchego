// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcsigner

import (
	"context"

	"google.golang.org/grpc"

	"github.com/ava-labs/avalanchego/utils/crypto/bls"

	pb "github.com/ava-labs/avalanchego/proto/pb/signer"
)

var _ bls.Signer = (*Client)(nil)

type Client struct {
	client pb.SignerClient
	pk     *bls.PublicKey
	// grpc.ClientConn handles transient connection errors.
	connection *grpc.ClientConn
}

func NewClient(ctx context.Context, url string) (*Client, error) {
	_ = "STUB: not implemented"
	// TODO: figure out the best parameters here given the target block-time
	return nil, nil
}

// same as grpc default

// the rpc-signer client should call a proxy server (on the same machine) that forwards
// the request to the actual signer instead of relying on tls-credentials

func (c *Client) PublicKey() *bls.PublicKey { _ = "STUB: not implemented"; return nil }

func (c *Client) Sign(message []byte) (*bls.Signature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignProofOfPossession produces a ProofOfPossession signature.
// See BLS spec for more details.
func (c *Client) SignProofOfPossession(message []byte) (*bls.Signature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Shutdown() error { _ = "STUB: not implemented"; return nil }
