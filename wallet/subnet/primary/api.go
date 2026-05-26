// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package primary

import (
	"context"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/coreth/ethclient"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/client"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/wallet/chain/c"

	pbuilder "github.com/ava-labs/avalanchego/wallet/chain/p/builder"
	xbuilder "github.com/ava-labs/avalanchego/wallet/chain/x/builder"
	walletcommon "github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
	ethcommon "github.com/ava-labs/libevm/common"
)

const (
	MainnetAPIURI = "https://api.avax.network"
	FujiAPIURI    = "https://api.avax-test.network"
	LocalAPIURI   = "http://localhost:9650"

	fetchLimit = 1024
)

var (
	_ UTXOClient = (*platformvm.Client)(nil)
	_ UTXOClient = (*avm.Client)(nil)
	_ UTXOClient = (*client.Client)(nil)
)

type UTXOClient interface {
	GetAtomicUTXOs(
		ctx context.Context,
		addrs []ids.ShortID,
		sourceChain string,
		limit uint32,
		startAddress ids.ShortID,
		startUTXOID ids.ID,
		options ...rpc.Option,
	) ([][]byte, ids.ShortID, ids.ID, error)
}

type AVAXState struct {
	PClient *platformvm.Client
	PCTX    *pbuilder.Context
	XClient *avm.Client
	XCTX    *xbuilder.Context
	CClient *client.Client
	CCTX    *c.Context
	UTXOs   walletcommon.UTXOs
}

func FetchState(
	ctx context.Context,
	uri string,
	addrs set.Set[ids.ShortID],
) (
	*AVAXState,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FetchPState(
	ctx context.Context,
	uri string,
	addrs set.Set[ids.ShortID],
) (
	*platformvm.Client,
	*pbuilder.Context,
	walletcommon.UTXOs,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, *new(walletcommon.UTXOs), nil
}

type EthState struct {
	Client   *ethclient.Client
	Accounts map[ethcommon.Address]*c.Account
}

func FetchEthState(
	ctx context.Context,
	uri string,
	addrs set.Set[ethcommon.Address],
) (*EthState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddAllUTXOs fetches all the UTXOs referenced by [addresses] that were sent
// from [sourceChainID] to [destinationChainID] from the [client]. It then uses
// [codec] to parse the returned UTXOs and it adds them into [utxos]. If [ctx]
// expires, then the returned error will be immediately reported.
func AddAllUTXOs(
	ctx context.Context,
	utxos walletcommon.UTXOs,
	client UTXOClient,
	codec codec.Manager,
	sourceChainID ids.ID,
	destinationChainID ids.ID,
	addrs []ids.ShortID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the vars to query the next page of UTXOs.
