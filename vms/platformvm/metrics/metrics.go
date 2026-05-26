// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
)

const (
	ResourceLabel   = "resource"
	GasLabel        = "gas"
	ValidatorsLabel = "validators"
)

var (
	gasLabels = prometheus.Labels{
		ResourceLabel: GasLabel,
	}
	validatorsLabels = prometheus.Labels{
		ResourceLabel: ValidatorsLabel,
	}
)

var _ Metrics = (*metrics)(nil)

type Block struct {
	Block block.Block

	GasConsumed gas.Gas
	GasState    gas.State
	GasPrice    gas.Price

	ActiveL1Validators   int
	ValidatorExcess      gas.Gas
	ValidatorPrice       gas.Price
	AccruedValidatorFees uint64
}

type Metrics interface {
	metric.APIInterceptor

	// Mark that the given block was accepted.
	MarkAccepted(Block) error

	// Mark that a validator set was created.
	IncValidatorSetsCreated()
	// Mark that a validator set was cached.
	IncValidatorSetsCached()
	// Mark that we spent the given time computing validator diffs.
	AddValidatorSetsDuration(time.Duration)
	// Mark that we computed a validator diff at a height with the given
	// difference from the top.
	AddValidatorSetsHeightDiff(uint64)

	// Mark that this much stake is staked on the node.
	SetLocalStake(uint64)
	// Mark that this much stake is staked in the network.
	SetTotalStake(uint64)
	// Mark when this node will unstake from the Primary Network.
	SetTimeUntilUnstake(time.Duration)
	// Mark when this node will unstake from a subnet.
	SetTimeUntilSubnetUnstake(subnetID ids.ID, timeUntilUnstake time.Duration)
}

func New(registerer prometheus.Registerer) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

type metrics struct {
	metric.APIInterceptor

	blockMetrics *blockMetrics

	// Staking metrics
	timeUntilUnstake       prometheus.Gauge
	timeUntilSubnetUnstake *prometheus.GaugeVec
	localStake             prometheus.Gauge
	totalStake             prometheus.Gauge

	gasConsumed          prometheus.Counter
	gasCapacity          prometheus.Gauge
	activeL1Validators   prometheus.Gauge
	excess               *prometheus.GaugeVec
	price                *prometheus.GaugeVec
	accruedValidatorFees prometheus.Gauge

	// Validator set diff metrics
	validatorSetsCached     prometheus.Counter
	validatorSetsCreated    prometheus.Counter
	validatorSetsHeightDiff prometheus.Gauge
	validatorSetsDuration   prometheus.Gauge
}

func (m *metrics) MarkAccepted(b Block) error { _ = "STUB: not implemented"; return nil }

func (m *metrics) IncValidatorSetsCreated() { _ = "STUB: not implemented"; return }

func (m *metrics) IncValidatorSetsCached() { _ = "STUB: not implemented"; return }

func (m *metrics) AddValidatorSetsDuration(d time.Duration) { _ = "STUB: not implemented"; return }

func (m *metrics) AddValidatorSetsHeightDiff(d uint64) { _ = "STUB: not implemented"; return }

func (m *metrics) SetLocalStake(s uint64) { _ = "STUB: not implemented"; return }

func (m *metrics) SetTotalStake(s uint64) { _ = "STUB: not implemented"; return }

func (m *metrics) SetTimeUntilUnstake(timeUntilUnstake time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *metrics) SetTimeUntilSubnetUnstake(subnetID ids.ID, timeUntilUnstake time.Duration) {
	_ = "STUB: not implemented"
	return
}
