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

// IGasPriceManagerGasPriceConfig is an auto generated low-level Go binding around an user-defined struct.
type IGasPriceManagerGasPriceConfig struct {
	ValidatorTargetGas bool
	TargetGas          uint64
	StaticPricing      bool
	MinGasPrice        uint64
	TimeToDouble       uint64
}

// IGasPriceManagerMetaData contains all meta data concerning the IGasPriceManager contract.
var IGasPriceManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"validatorTargetGas\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"targetGas\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"staticPricing\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeToDouble\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIGasPriceManager.GasPriceConfig\",\"name\":\"oldGasPriceConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"validatorTargetGas\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"targetGas\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"staticPricing\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeToDouble\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIGasPriceManager.GasPriceConfig\",\"name\":\"newGasPriceConfig\",\"type\":\"tuple\"}],\"name\":\"GasPriceConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRole\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getGasPriceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"validatorTargetGas\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"targetGas\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"staticPricing\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeToDouble\",\"type\":\"uint64\"}],\"internalType\":\"structIGasPriceManager.GasPriceConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasPriceConfigLastChangedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"validatorTargetGas\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"targetGas\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"staticPricing\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timeToDouble\",\"type\":\"uint64\"}],\"internalType\":\"structIGasPriceManager.GasPriceConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setGasPriceConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGasPriceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasPriceManagerMetaData.ABI instead.
var IGasPriceManagerABI = IGasPriceManagerMetaData.ABI

// IGasPriceManager is an auto generated Go binding around an Ethereum contract.
type IGasPriceManager struct {
	IGasPriceManagerCaller     // Read-only binding to the contract
	IGasPriceManagerTransactor // Write-only binding to the contract
	IGasPriceManagerFilterer   // Log filterer for contract events
}

// IGasPriceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasPriceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasPriceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasPriceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasPriceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasPriceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasPriceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasPriceManagerSession struct {
	Contract     *IGasPriceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasPriceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasPriceManagerCallerSession struct {
	Contract *IGasPriceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IGasPriceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasPriceManagerTransactorSession struct {
	Contract     *IGasPriceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IGasPriceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasPriceManagerRaw struct {
	Contract *IGasPriceManager // Generic contract binding to access the raw methods on
}

// IGasPriceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasPriceManagerCallerRaw struct {
	Contract *IGasPriceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IGasPriceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasPriceManagerTransactorRaw struct {
	Contract *IGasPriceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasPriceManager creates a new instance of IGasPriceManager, bound to a specific deployed contract.
func NewIGasPriceManager(address common.Address, backend bind.ContractBackend) (*IGasPriceManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIGasPriceManagerCaller creates a new read-only instance of IGasPriceManager, bound to a specific deployed contract.
func NewIGasPriceManagerCaller(address common.Address, caller bind.ContractCaller) (*IGasPriceManagerCaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIGasPriceManagerTransactor creates a new write-only instance of IGasPriceManager, bound to a specific deployed contract.
func NewIGasPriceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasPriceManagerTransactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIGasPriceManagerFilterer creates a new log filterer instance of IGasPriceManager, bound to a specific deployed contract.
func NewIGasPriceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasPriceManagerFilterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindIGasPriceManager binds a generic wrapper to an already deployed contract.
func bindIGasPriceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasPriceManager *IGasPriceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasPriceManager *IGasPriceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasPriceManager *IGasPriceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasPriceManager *IGasPriceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasPriceManager *IGasPriceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasPriceManager *IGasPriceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGasPriceConfig is a free data retrieval call binding the contract method 0x445821e3.
//
// Solidity: function getGasPriceConfig() view returns((bool,uint64,bool,uint64,uint64) config)
func (_IGasPriceManager *IGasPriceManagerCaller) GetGasPriceConfig(opts *bind.CallOpts) (IGasPriceManagerGasPriceConfig, error) {
	_ = "STUB: not implemented"
	return *new(IGasPriceManagerGasPriceConfig), nil
}

// GetGasPriceConfig is a free data retrieval call binding the contract method 0x445821e3.
//
// Solidity: function getGasPriceConfig() view returns((bool,uint64,bool,uint64,uint64) config)
func (_IGasPriceManager *IGasPriceManagerSession) GetGasPriceConfig() (IGasPriceManagerGasPriceConfig, error) {
	_ = "STUB: not implemented"
	return *new(IGasPriceManagerGasPriceConfig), nil
}

// GetGasPriceConfig is a free data retrieval call binding the contract method 0x445821e3.
//
// Solidity: function getGasPriceConfig() view returns((bool,uint64,bool,uint64,uint64) config)
func (_IGasPriceManager *IGasPriceManagerCallerSession) GetGasPriceConfig() (IGasPriceManagerGasPriceConfig, error) {
	_ = "STUB: not implemented"
	return *new(IGasPriceManagerGasPriceConfig), nil
}

// GetGasPriceConfigLastChangedAt is a free data retrieval call binding the contract method 0xeb8df396.
//
// Solidity: function getGasPriceConfigLastChangedAt() view returns(uint256 blockNumber)
func (_IGasPriceManager *IGasPriceManagerCaller) GetGasPriceConfigLastChangedAt(opts *bind.CallOpts) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGasPriceConfigLastChangedAt is a free data retrieval call binding the contract method 0xeb8df396.
//
// Solidity: function getGasPriceConfigLastChangedAt() view returns(uint256 blockNumber)
func (_IGasPriceManager *IGasPriceManagerSession) GetGasPriceConfigLastChangedAt() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGasPriceConfigLastChangedAt is a free data retrieval call binding the contract method 0xeb8df396.
//
// Solidity: function getGasPriceConfigLastChangedAt() view returns(uint256 blockNumber)
func (_IGasPriceManager *IGasPriceManagerCallerSession) GetGasPriceConfigLastChangedAt() (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IGasPriceManager *IGasPriceManagerCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IGasPriceManager *IGasPriceManagerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IGasPriceManager *IGasPriceManagerCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGasPriceConfig is a paid mutator transaction binding the contract method 0x65c97a28.
//
// Solidity: function setGasPriceConfig((bool,uint64,bool,uint64,uint64) config) returns()
func (_IGasPriceManager *IGasPriceManagerTransactor) SetGasPriceConfig(opts *bind.TransactOpts, config IGasPriceManagerGasPriceConfig) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGasPriceConfig is a paid mutator transaction binding the contract method 0x65c97a28.
//
// Solidity: function setGasPriceConfig((bool,uint64,bool,uint64,uint64) config) returns()
func (_IGasPriceManager *IGasPriceManagerSession) SetGasPriceConfig(config IGasPriceManagerGasPriceConfig) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGasPriceConfig is a paid mutator transaction binding the contract method 0x65c97a28.
//
// Solidity: function setGasPriceConfig((bool,uint64,bool,uint64,uint64) config) returns()
func (_IGasPriceManager *IGasPriceManagerTransactorSession) SetGasPriceConfig(config IGasPriceManagerGasPriceConfig) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactor) SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactorSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IGasPriceManager *IGasPriceManagerTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IGasPriceManagerGasPriceConfigUpdatedIterator is returned from FilterGasPriceConfigUpdated and is used to iterate over the raw logs and unpacked data for GasPriceConfigUpdated events raised by the IGasPriceManager contract.
type IGasPriceManagerGasPriceConfigUpdatedIterator struct {
	Event *IGasPriceManagerGasPriceConfigUpdated // Event containing the contract specifics and raw log

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
func (it *IGasPriceManagerGasPriceConfigUpdatedIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGasPriceManagerGasPriceConfigUpdatedIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IGasPriceManagerGasPriceConfigUpdatedIterator) Close() error {
	_ = "STUB: not implemented"
	return nil
}

// IGasPriceManagerGasPriceConfigUpdated represents a GasPriceConfigUpdated event raised by the IGasPriceManager contract.
type IGasPriceManagerGasPriceConfigUpdated struct {
	Sender            common.Address
	OldGasPriceConfig IGasPriceManagerGasPriceConfig
	NewGasPriceConfig IGasPriceManagerGasPriceConfig
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGasPriceConfigUpdated is a free log retrieval operation binding the contract event 0x9456356f1b11d07d47f9bd3d568279a926ea8caceec0e6078e426eaea579be66.
//
// Solidity: event GasPriceConfigUpdated(address indexed sender, (bool,uint64,bool,uint64,uint64) oldGasPriceConfig, (bool,uint64,bool,uint64,uint64) newGasPriceConfig)
func (_IGasPriceManager *IGasPriceManagerFilterer) FilterGasPriceConfigUpdated(opts *bind.FilterOpts, sender []common.Address) (*IGasPriceManagerGasPriceConfigUpdatedIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchGasPriceConfigUpdated is a free log subscription operation binding the contract event 0x9456356f1b11d07d47f9bd3d568279a926ea8caceec0e6078e426eaea579be66.
//
// Solidity: event GasPriceConfigUpdated(address indexed sender, (bool,uint64,bool,uint64,uint64) oldGasPriceConfig, (bool,uint64,bool,uint64,uint64) newGasPriceConfig)
func (_IGasPriceManager *IGasPriceManagerFilterer) WatchGasPriceConfigUpdated(opts *bind.WatchOpts, sink chan<- *IGasPriceManagerGasPriceConfigUpdated, sender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseGasPriceConfigUpdated is a log parse operation binding the contract event 0x9456356f1b11d07d47f9bd3d568279a926ea8caceec0e6078e426eaea579be66.
//
// Solidity: event GasPriceConfigUpdated(address indexed sender, (bool,uint64,bool,uint64,uint64) oldGasPriceConfig, (bool,uint64,bool,uint64,uint64) newGasPriceConfig)
func (_IGasPriceManager *IGasPriceManagerFilterer) ParseGasPriceConfigUpdated(log types.Log) (*IGasPriceManagerGasPriceConfigUpdated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IGasPriceManagerRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the IGasPriceManager contract.
type IGasPriceManagerRoleSetIterator struct {
	Event *IGasPriceManagerRoleSet // Event containing the contract specifics and raw log

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
func (it *IGasPriceManagerRoleSetIterator) Next() bool {
	_ = "STUB: not implemented"
	// If the iterator failed, stop iterating
	return false
}

// If the iterator completed, deliver directly whatever's available

// Iterator still in progress, wait for either a data or an error event

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGasPriceManagerRoleSetIterator) Error() error {
	_ = "STUB: not implemented"

	// Close terminates the iteration process, releasing any pending underlying
	// resources.
	return nil
}

func (it *IGasPriceManagerRoleSetIterator) Close() error { _ = "STUB: not implemented"; return nil }

// IGasPriceManagerRoleSet represents a RoleSet event raised by the IGasPriceManager contract.
type IGasPriceManagerRoleSet struct {
	Role    *big.Int
	Account common.Address
	Sender  common.Address
	OldRole *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IGasPriceManager *IGasPriceManagerFilterer) FilterRoleSet(opts *bind.FilterOpts, role []*big.Int, account []common.Address, sender []common.Address) (*IGasPriceManagerRoleSetIterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IGasPriceManager *IGasPriceManagerFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *IGasPriceManagerRoleSet, role []*big.Int, account []common.Address, sender []common.Address) (event.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(event.Subscription), nil
}

// New log arrived, parse the event and forward to the user

// ParseRoleSet is a log parse operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IGasPriceManager *IGasPriceManagerFilterer) ParseRoleSet(log types.Log) (*IGasPriceManagerRoleSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
