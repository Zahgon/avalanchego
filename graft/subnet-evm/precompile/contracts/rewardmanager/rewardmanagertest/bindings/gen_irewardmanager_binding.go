// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/accounts/abi/bind"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRewardManagerMetaData contains all meta data concerning the IRewardManager contract.
var IRewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"FeeRecipientsAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRewardAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRewardAddress\",\"type\":\"address\"}],\"name\":\"RewardAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RewardsDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRole\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowFeeRecipients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"areFeeRecipientsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardManagerMetaData.ABI instead.
var IRewardManagerABI = IRewardManagerMetaData.ABI

// IRewardManager is an auto generated Go binding around an Ethereum contract.
type IRewardManager struct {
	IRewardManagerCaller     // Read-only binding to the contract
	IRewardManagerTransactor // Write-only binding to the contract
	IRewardManagerFilterer   // Log filterer for contract events
}

// IRewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardManagerSession struct {
	Contract     *IRewardManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardManagerCallerSession struct {
	Contract *IRewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IRewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardManagerTransactorSession struct {
	Contract     *IRewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IRewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardManagerRaw struct {
	Contract *IRewardManager // Generic contract binding to access the raw methods on
}

// IRewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardManagerCallerRaw struct {
	Contract *IRewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardManagerTransactorRaw struct {
	Contract *IRewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardManager creates a new instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManager(address common.Address, backend bind.ContractBackend) (*IRewardManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIRewardManagerCaller creates a new read-only instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerCaller(address common.Address, caller bind.ContractCaller) (*IRewardManagerCaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIRewardManagerTransactor creates a new write-only instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardManagerTransactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIRewardManagerFilterer creates a new log filterer instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardManagerFilterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindIRewardManager binds a generic wrapper to an already deployed contract.
func bindIRewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardManager *IRewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardManager *IRewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardManager *IRewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardManager *IRewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardManager *IRewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardManager *IRewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AreFeeRecipientsAllowed is a free data retrieval call binding the contract method 0xf6542b2e.
//
// Solidity: function areFeeRecipientsAllowed() view returns(bool isAllowed)
func (_IRewardManager *IRewardManagerCaller) AreFeeRecipientsAllowed(opts *bind.CallOpts) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AreFeeRecipientsAllowed is a free data retrieval call binding the contract method 0xf6542b2e.
//
// Solidity: function areFeeRecipientsAllowed() view returns(bool isAllowed)
func (_IRewardManager *IRewardManagerSession) AreFeeRecipientsAllowed() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AreFeeRecipientsAllowed is a free data retrieval call binding the contract method 0xf6542b2e.
//
// Solidity: function areFeeRecipientsAllowed() view returns(bool isAllowed)
func (_IRewardManager *IRewardManagerCallerSession) AreFeeRecipientsAllowed() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CurrentRewardAddress is a free data retrieval call binding the contract method 0xe915608b.
//
// Solidity: function currentRewardAddress() view returns(address rewardAddress)
func (_IRewardManager *IRewardManagerCaller) CurrentRewardAddress(opts *bind.CallOpts) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// CurrentRewardAddress is a free data retrieval call binding the contract method 0xe915608b.
//
// Solidity: function currentRewardAddress() view returns(address rewardAddress)
func (_IRewardManager *IRewardManagerSession) CurrentRewardAddress() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// CurrentRewardAddress is a free data retrieval call binding the contract method 0xe915608b.
//
// Solidity: function currentRewardAddress() view returns(address rewardAddress)
func (_IRewardManager *IRewardManagerCallerSession) CurrentRewardAddress() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IRewardManager *IRewardManagerCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IRewardManager *IRewardManagerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IRewardManager *IRewardManagerCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllowFeeRecipients is a paid mutator transaction binding the contract method 0x0329099f.
//
// Solidity: function allowFeeRecipients() returns()
func (_IRewardManager *IRewardManagerTransactor) AllowFeeRecipients(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllowFeeRecipients is a paid mutator transaction binding the contract method 0x0329099f.
//
// Solidity: function allowFeeRecipients() returns()
func (_IRewardManager *IRewardManagerSession) AllowFeeRecipients() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllowFeeRecipients is a paid mutator transaction binding the contract method 0x0329099f.
//
// Solidity: function allowFeeRecipients() returns()
func (_IRewardManager *IRewardManagerTransactorSession) AllowFeeRecipients() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DisableRewards is a paid mutator transaction binding the contract method 0xbc178628.
//
// Solidity: function disableRewards() returns()
func (_IRewardManager *IRewardManagerTransactor) DisableRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DisableRewards is a paid mutator transaction binding the contract method 0xbc178628.
//
// Solidity: function disableRewards() returns()
func (_IRewardManager *IRewardManagerSession) DisableRewards() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DisableRewards is a paid mutator transaction binding the contract method 0xbc178628.
//
// Solidity: function disableRewards() returns()
func (_IRewardManager *IRewardManagerTransactorSession) DisableRewards() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IRewardManager *IRewardManagerTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IRewardManager *IRewardManagerSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IRewardManager *IRewardManagerTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IRewardManager *IRewardManagerSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IRewardManager *IRewardManagerTransactor) SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IRewardManager *IRewardManagerSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IRewardManager *IRewardManagerTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IRewardManager *IRewardManagerSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address addr) returns()
func (_IRewardManager *IRewardManagerTransactor) SetRewardAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address addr) returns()
func (_IRewardManager *IRewardManagerSession) SetRewardAddress(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetRewardAddress is a paid mutator transaction binding the contract method 0x5e00e679.
//
// Solidity: function setRewardAddress(address addr) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetRewardAddress(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IRewardManagerFeeRecipientsAllowedIterator is returned from FilterFeeRecipientsAllowed and is used to iterate over the raw logs and unpacked data for FeeRecipientsAllowed events raised by the IRewardManager contract.
type IRewardManagerFeeRecipientsAllowedIterator struct {
	Event *IRewardManagerFeeRecipientsAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRewardManagerFeeRecipientsAllowedIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRewardManagerFeeRecipientsAllowedIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IRewardManagerFeeRecipientsAllowedIterator) Close() error {
	_ = "STUB: not implemented"
	return nil
}

// IRewardManagerFeeRecipientsAllowed represents a FeeRecipientsAllowed event raised by the IRewardManager contract.
type IRewardManagerFeeRecipientsAllowed struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientsAllowed is a free log retrieval operation binding the contract event 0xabb1949bd129fef9b84601a48aee89d600d90074ca10216a02ce43996be55991.
//
// Solidity: event FeeRecipientsAllowed(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) FilterFeeRecipientsAllowed(opts *bind.FilterOpts, sender []common.Address) (*IRewardManagerFeeRecipientsAllowedIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchFeeRecipientsAllowed is a free log subscription operation binding the contract event 0xabb1949bd129fef9b84601a48aee89d600d90074ca10216a02ce43996be55991.
//
// Solidity: event FeeRecipientsAllowed(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) WatchFeeRecipientsAllowed(opts *bind.WatchOpts, sink chan<- *IRewardManagerFeeRecipientsAllowed, sender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseFeeRecipientsAllowed is a log parse operation binding the contract event 0xabb1949bd129fef9b84601a48aee89d600d90074ca10216a02ce43996be55991.
//
// Solidity: event FeeRecipientsAllowed(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) ParseFeeRecipientsAllowed(log types.Log) (*IRewardManagerFeeRecipientsAllowed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IRewardManagerRewardAddressChangedIterator is returned from FilterRewardAddressChanged and is used to iterate over the raw logs and unpacked data for RewardAddressChanged events raised by the IRewardManager contract.
type IRewardManagerRewardAddressChangedIterator struct {
	Event *IRewardManagerRewardAddressChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRewardManagerRewardAddressChangedIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRewardManagerRewardAddressChangedIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IRewardManagerRewardAddressChangedIterator) Close() error {
	_ = "STUB: not implemented"
	return nil
}

// IRewardManagerRewardAddressChanged represents a RewardAddressChanged event raised by the IRewardManager contract.
type IRewardManagerRewardAddressChanged struct {
	Sender           common.Address
	OldRewardAddress common.Address
	NewRewardAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardAddressChanged is a free log retrieval operation binding the contract event 0xc2a9e07cba6f4920acaa5933bd0406949d5dbef7ee698e786ea23e8708f32a6c.
//
// Solidity: event RewardAddressChanged(address indexed sender, address indexed oldRewardAddress, address indexed newRewardAddress)
func (_IRewardManager *IRewardManagerFilterer) FilterRewardAddressChanged(opts *bind.FilterOpts, sender []common.Address, oldRewardAddress []common.Address, newRewardAddress []common.Address) (*IRewardManagerRewardAddressChangedIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchRewardAddressChanged is a free log subscription operation binding the contract event 0xc2a9e07cba6f4920acaa5933bd0406949d5dbef7ee698e786ea23e8708f32a6c.
//
// Solidity: event RewardAddressChanged(address indexed sender, address indexed oldRewardAddress, address indexed newRewardAddress)
func (_IRewardManager *IRewardManagerFilterer) WatchRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *IRewardManagerRewardAddressChanged, sender []common.Address, oldRewardAddress []common.Address, newRewardAddress []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseRewardAddressChanged is a log parse operation binding the contract event 0xc2a9e07cba6f4920acaa5933bd0406949d5dbef7ee698e786ea23e8708f32a6c.
//
// Solidity: event RewardAddressChanged(address indexed sender, address indexed oldRewardAddress, address indexed newRewardAddress)
func (_IRewardManager *IRewardManagerFilterer) ParseRewardAddressChanged(log types.Log) (*IRewardManagerRewardAddressChanged, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IRewardManagerRewardsDisabledIterator is returned from FilterRewardsDisabled and is used to iterate over the raw logs and unpacked data for RewardsDisabled events raised by the IRewardManager contract.
type IRewardManagerRewardsDisabledIterator struct {
	Event *IRewardManagerRewardsDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRewardManagerRewardsDisabledIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRewardManagerRewardsDisabledIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IRewardManagerRewardsDisabledIterator) Close() error {
	_ = "STUB: not implemented"
	return nil
}

// IRewardManagerRewardsDisabled represents a RewardsDisabled event raised by the IRewardManager contract.
type IRewardManagerRewardsDisabled struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsDisabled is a free log retrieval operation binding the contract event 0xeb121f0335efe8f4b8ebef7793c18c171834696989656a8c345acc558359fabf.
//
// Solidity: event RewardsDisabled(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) FilterRewardsDisabled(opts *bind.FilterOpts, sender []common.Address) (*IRewardManagerRewardsDisabledIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchRewardsDisabled is a free log subscription operation binding the contract event 0xeb121f0335efe8f4b8ebef7793c18c171834696989656a8c345acc558359fabf.
//
// Solidity: event RewardsDisabled(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) WatchRewardsDisabled(opts *bind.WatchOpts, sink chan<- *IRewardManagerRewardsDisabled, sender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseRewardsDisabled is a log parse operation binding the contract event 0xeb121f0335efe8f4b8ebef7793c18c171834696989656a8c345acc558359fabf.
//
// Solidity: event RewardsDisabled(address indexed sender)
func (_IRewardManager *IRewardManagerFilterer) ParseRewardsDisabled(log types.Log) (*IRewardManagerRewardsDisabled, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IRewardManagerRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the IRewardManager contract.
type IRewardManagerRoleSetIterator struct {
	Event *IRewardManagerRoleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRewardManagerRoleSetIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRewardManagerRoleSetIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IRewardManagerRoleSetIterator) Close() error { _ = "STUB: not implemented"; return nil }

// IRewardManagerRoleSet represents a RoleSet event raised by the IRewardManager contract.
type IRewardManagerRoleSet struct {
	Role    *big.Int
	Account common.Address
	Sender  common.Address
	OldRole *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IRewardManager *IRewardManagerFilterer) FilterRoleSet(opts *bind.FilterOpts, role []*big.Int, account []common.Address, sender []common.Address) (*IRewardManagerRoleSetIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IRewardManager *IRewardManagerFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *IRewardManagerRoleSet, role []*big.Int, account []common.Address, sender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseRoleSet is a log parse operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IRewardManager *IRewardManagerFilterer) ParseRoleSet(log types.Log) (*IRewardManagerRoleSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
