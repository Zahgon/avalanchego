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

// AllowListTestMetaData contains all meta data concerning the AllowListTest contract.
var AllowListTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"precompileAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"deployContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610af9380380610af9833981810160405281019061003191906100d4565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f5ffd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6109ed8061010c5f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c80638c6bfb3b116100645780638c6bfb3b1461012e5780639015d3711461014a578063d0ebdbe71461017a578063eb54dae114610196578063f3ae2415146101c65761009c565b80630aaf7043146100a057806324d7806c146100bc5780636cd5c39b146100ec578063704b6c02146100f657806374a8f10314610112575b5f5ffd5b6100ba60048036038101906100b5919061082d565b6101f6565b005b6100d660048036038101906100d1919061082d565b61027f565b6040516100e39190610872565b60405180910390f35b6100f4610322565b005b610110600480360381019061010b919061082d565b61034b565b005b61012c6004803603810190610127919061082d565b6103d4565b005b6101486004803603810190610143919061082d565b6104cb565b005b610164600480360381019061015f919061082d565b610554565b6040516101719190610872565b60405180910390f35b610194600480360381019061018f919061082d565b6105f7565b005b6101b060048036038101906101ab919061082d565b610680565b6040516101bd91906108a3565b60405180910390f35b6101e060048036038101906101db919061082d565b610720565b6040516101ed9190610872565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630aaf7043826040518263ffffffff1660e01b815260040161024f91906108cb565b5f604051808303815f87803b158015610266575f5ffd5b505af1158015610278573d5f5f3e3d5ffd5b5050505050565b5f60025f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eb54dae1846040518263ffffffff1660e01b81526004016102db91906108cb565b602060405180830381865afa1580156102f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031a919061090e565b149050919050565b60405161032e906107c3565b604051809103905ff080158015610347573d5f5f3e3d5ffd5b5050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663704b6c02826040518263ffffffff1660e01b81526004016103a491906108cb565b5f604051808303815f87803b1580156103bb575f5ffd5b505af11580156103cd573d5f5f3e3d5ffd5b5050505050565b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610442576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043990610993565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c6bfb3b826040518263ffffffff1660e01b815260040161049b91906108cb565b5f604051808303815f87803b1580156104b2575f5ffd5b505af11580156104c4573d5f5f3e3d5ffd5b5050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c6bfb3b826040518263ffffffff1660e01b815260040161052491906108cb565b5f604051808303815f87803b15801561053b575f5ffd5b505af115801561054d573d5f5f3e3d5ffd5b5050505050565b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eb54dae1846040518263ffffffff1660e01b81526004016105af91906108cb565b602060405180830381865afa1580156105ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ee919061090e565b14159050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0ebdbe7826040518263ffffffff1660e01b815260040161065091906108cb565b5f604051808303815f87803b158015610667575f5ffd5b505af1158015610679573d5f5f3e3d5ffd5b5050505050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eb54dae1836040518263ffffffff1660e01b81526004016106da91906108cb565b602060405180830381865afa1580156106f5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610719919061090e565b9050919050565b5f60035f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eb54dae1846040518263ffffffff1660e01b815260040161077c91906108cb565b602060405180830381865afa158015610797573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107bb919061090e565b149050919050565b602f806109b283390190565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107fc826107d3565b9050919050565b61080c816107f2565b8114610816575f5ffd5b50565b5f8135905061082781610803565b92915050565b5f60208284031215610842576108416107cf565b5b5f61084f84828501610819565b91505092915050565b5f8115159050919050565b61086c81610858565b82525050565b5f6020820190506108855f830184610863565b92915050565b5f819050919050565b61089d8161088b565b82525050565b5f6020820190506108b65f830184610894565b92915050565b6108c5816107f2565b82525050565b5f6020820190506108de5f8301846108bc565b92915050565b6108ed8161088b565b81146108f7575f5ffd5b50565b5f81519050610908816108e4565b92915050565b5f60208284031215610923576109226107cf565b5b5f610930848285016108fa565b91505092915050565b5f82825260208201905092915050565b7f63616e6e6f74207265766f6b65206f776e20726f6c65000000000000000000005f82015250565b5f61097d601683610939565b915061098882610949565b602082019050919050565b5f6020820190508181035f8301526109aa81610971565b905091905056fe6080604052348015600e575f5ffd5b50601580601a5f395ff3fe60806040525f5ffdfea164736f6c634300081c000aa164736f6c634300081c000a",
}

// AllowListTestABI is the input ABI used to generate the binding from.
// Deprecated: Use AllowListTestMetaData.ABI instead.
var AllowListTestABI = AllowListTestMetaData.ABI

// AllowListTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllowListTestMetaData.Bin instead.
var AllowListTestBin = AllowListTestMetaData.Bin

// DeployAllowListTest deploys a new Ethereum contract, binding an instance of AllowListTest to it.
func DeployAllowListTest(auth *bind.TransactOpts, backend bind.ContractBackend, precompileAddr common.Address) (common.Address, *types.Transaction, *AllowListTest, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil, nil, nil
}

// AllowListTest is an auto generated Go binding around an Ethereum contract.
type AllowListTest struct {
	AllowListTestCaller     // Read-only binding to the contract
	AllowListTestTransactor // Write-only binding to the contract
	AllowListTestFilterer   // Log filterer for contract events
}

// AllowListTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowListTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowListTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowListTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowListTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowListTestSession struct {
	Contract     *AllowListTest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllowListTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowListTestCallerSession struct {
	Contract *AllowListTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AllowListTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowListTestTransactorSession struct {
	Contract     *AllowListTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AllowListTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowListTestRaw struct {
	Contract *AllowListTest // Generic contract binding to access the raw methods on
}

// AllowListTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowListTestCallerRaw struct {
	Contract *AllowListTestCaller // Generic read-only contract binding to access the raw methods on
}

// AllowListTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowListTestTransactorRaw struct {
	Contract *AllowListTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowListTest creates a new instance of AllowListTest, bound to a specific deployed contract.
func NewAllowListTest(address common.Address, backend bind.ContractBackend) (*AllowListTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAllowListTestCaller creates a new read-only instance of AllowListTest, bound to a specific deployed contract.
func NewAllowListTestCaller(address common.Address, caller bind.ContractCaller) (*AllowListTestCaller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAllowListTestTransactor creates a new write-only instance of AllowListTest, bound to a specific deployed contract.
func NewAllowListTestTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowListTestTransactor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAllowListTestFilterer creates a new log filterer instance of AllowListTest, bound to a specific deployed contract.
func NewAllowListTestFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowListTestFilterer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bindAllowListTest binds a generic wrapper to an already deployed contract.
func bindAllowListTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowListTest *AllowListTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowListTest *AllowListTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowListTest *AllowListTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowListTest *AllowListTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowListTest *AllowListTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowListTest *AllowListTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_AllowListTest *AllowListTestSession) IsAdmin(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCallerSession) IsAdmin(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCaller) IsEnabled(opts *bind.CallOpts, addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address addr) view returns(bool)
func (_AllowListTest *AllowListTestSession) IsEnabled(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCallerSession) IsEnabled(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCaller) IsManager(opts *bind.CallOpts, addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address addr) view returns(bool)
func (_AllowListTest *AllowListTestSession) IsManager(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address addr) view returns(bool)
func (_AllowListTest *AllowListTestCallerSession) IsManager(addr common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListTest *AllowListTestCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListTest *AllowListTestSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256)
func (_AllowListTest *AllowListTestCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns()
func (_AllowListTest *AllowListTestTransactor) DeployContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns()
func (_AllowListTest *AllowListTestSession) DeployContract() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeployContract is a paid mutator transaction binding the contract method 0x6cd5c39b.
//
// Solidity: function deployContract() returns()
func (_AllowListTest *AllowListTestTransactorSession) DeployContract() (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address addr) returns()
func (_AllowListTest *AllowListTestTransactor) Revoke(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address addr) returns()
func (_AllowListTest *AllowListTestSession) Revoke(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address addr) returns()
func (_AllowListTest *AllowListTestTransactorSession) Revoke(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListTest *AllowListTestTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListTest *AllowListTestSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_AllowListTest *AllowListTestTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListTest *AllowListTestTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListTest *AllowListTestSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_AllowListTest *AllowListTestTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_AllowListTest *AllowListTestTransactor) SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_AllowListTest *AllowListTestSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_AllowListTest *AllowListTestTransactorSession) SetManager(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListTest *AllowListTestTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListTest *AllowListTestSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_AllowListTest *AllowListTestTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
