// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// IPriceFeedPrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedPrice struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}

// IPriceFeedPriceResponse is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedPriceResponse struct {
	Prices   []*big.Int
	Decimals []*big.Int
}

// IPriceFeedRequest is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedRequest struct {
	Pair []string
}

// PriceFeedMetaData contains all meta data concerning the PriceFeed contract.
var PriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"pairs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"decimals\",\"type\":\"uint256[]\"}],\"name\":\"PriceFeedRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"}],\"name\":\"Feed\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"pair\",\"type\":\"string[]\"}],\"internalType\":\"structIPriceFeed.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"requestPriceFeed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"decimals\",\"type\":\"uint256[]\"}],\"internalType\":\"structIPriceFeed.PriceResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pair\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"name\":\"updatePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// PriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedMetaData.ABI instead.
var PriceFeedABI = PriceFeedMetaData.ABI

// PriceFeed is an auto generated Go binding around an Ethereum contract.
type PriceFeed struct {
	PriceFeedCaller     // Read-only binding to the contract
	PriceFeedTransactor // Write-only binding to the contract
	PriceFeedFilterer   // Log filterer for contract events
}

// PriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedSession struct {
	Contract     *PriceFeed        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedCallerSession struct {
	Contract *PriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedTransactorSession struct {
	Contract     *PriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedRaw struct {
	Contract *PriceFeed // Generic contract binding to access the raw methods on
}

// PriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedCallerRaw struct {
	Contract *PriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedTransactorRaw struct {
	Contract *PriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeed creates a new instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeed(address common.Address, backend bind.ContractBackend) (*PriceFeed, error) {
	contract, err := bindPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeed{PriceFeedCaller: PriceFeedCaller{contract: contract}, PriceFeedTransactor: PriceFeedTransactor{contract: contract}, PriceFeedFilterer: PriceFeedFilterer{contract: contract}}, nil
}

// NewPriceFeedCaller creates a new read-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedCaller, error) {
	contract, err := bindPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCaller{contract: contract}, nil
}

// NewPriceFeedTransactor creates a new write-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedTransactor, error) {
	contract, err := bindPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTransactor{contract: contract}, nil
}

// NewPriceFeedFilterer creates a new log filterer instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedFilterer, error) {
	contract, err := bindPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedFilterer{contract: contract}, nil
}

// bindPriceFeed binds a generic wrapper to an already deployed contract.
func bindPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.PriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedCaller) Feed(opts *bind.CallOpts, pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "Feed", pair)

	outstruct := new(struct {
		Pair     string
		Price    *big.Int
		Decimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pair = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedSession) Feed(pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	return _PriceFeed.Contract.Feed(&_PriceFeed.CallOpts, pair)
}

// Feed is a free data retrieval call binding the contract method 0x6d2fbecb.
//
// Solidity: function Feed(string pair) view returns(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedCallerSession) Feed(pair string) (struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
}, error) {
	return _PriceFeed.Contract.Feed(&_PriceFeed.CallOpts, pair)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedSession) Owner() (common.Address, error) {
	return _PriceFeed.Contract.Owner(&_PriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedCallerSession) Owner() (common.Address, error) {
	return _PriceFeed.Contract.Owner(&_PriceFeed.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeed.Contract.ProxiableUUID(&_PriceFeed.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeed.Contract.ProxiableUUID(&_PriceFeed.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeed *PriceFeedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeed *PriceFeedSession) Initialize() (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeed *PriceFeedTransactorSession) Initialize() (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceOwnership(&_PriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceOwnership(&_PriceFeed.TransactOpts)
}

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeed *PriceFeedTransactor) RequestPriceFeed(opts *bind.TransactOpts, request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "requestPriceFeed", request)
}

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeed *PriceFeedSession) RequestPriceFeed(request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeed.Contract.RequestPriceFeed(&_PriceFeed.TransactOpts, request)
}

// RequestPriceFeed is a paid mutator transaction binding the contract method 0x49c14ed9.
//
// Solidity: function requestPriceFeed((string[]) request) payable returns((uint256[],uint256[]) response)
func (_PriceFeed *PriceFeedTransactorSession) RequestPriceFeed(request IPriceFeedRequest) (*types.Transaction, error) {
	return _PriceFeed.Contract.RequestPriceFeed(&_PriceFeed.TransactOpts, request)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.TransferOwnership(&_PriceFeed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.TransferOwnership(&_PriceFeed.TransactOpts, newOwner)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeed *PriceFeedTransactor) UpdatePriceFeed(opts *bind.TransactOpts, price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "updatePriceFeed", price)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeed *PriceFeedSession) UpdatePriceFeed(price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdatePriceFeed(&_PriceFeed.TransactOpts, price)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x68650ca2.
//
// Solidity: function updatePriceFeed((string,uint256,uint256) price) returns()
func (_PriceFeed *PriceFeedTransactorSession) UpdatePriceFeed(price IPriceFeedPrice) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdatePriceFeed(&_PriceFeed.TransactOpts, price)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeed *PriceFeedTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeed *PriceFeedSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeTo(&_PriceFeed.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeed *PriceFeedTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeTo(&_PriceFeed.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeToAndCall(&_PriceFeed.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeToAndCall(&_PriceFeed.TransactOpts, newImplementation, data)
}

// PriceFeedAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PriceFeed contract.
type PriceFeedAdminChangedIterator struct {
	Event *PriceFeedAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedAdminChanged represents a AdminChanged event raised by the PriceFeed contract.
type PriceFeedAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeed *PriceFeedFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PriceFeedAdminChangedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedAdminChangedIterator{contract: _PriceFeed.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeed *PriceFeedFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedAdminChanged)
				if err := _PriceFeed.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeed *PriceFeedFilterer) ParseAdminChanged(log types.Log) (*PriceFeedAdminChanged, error) {
	event := new(PriceFeedAdminChanged)
	if err := _PriceFeed.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PriceFeed contract.
type PriceFeedBeaconUpgradedIterator struct {
	Event *PriceFeedBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedBeaconUpgraded represents a BeaconUpgraded event raised by the PriceFeed contract.
type PriceFeedBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeed *PriceFeedFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PriceFeedBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedBeaconUpgradedIterator{contract: _PriceFeed.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeed *PriceFeedFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedBeaconUpgraded)
				if err := _PriceFeed.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeed *PriceFeedFilterer) ParseBeaconUpgraded(log types.Log) (*PriceFeedBeaconUpgraded, error) {
	event := new(PriceFeedBeaconUpgraded)
	if err := _PriceFeed.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PriceFeed contract.
type PriceFeedInitializedIterator struct {
	Event *PriceFeedInitialized // Event containing the contract specifics and raw log

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
func (it *PriceFeedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedInitialized represents a Initialized event raised by the PriceFeed contract.
type PriceFeedInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeed *PriceFeedFilterer) FilterInitialized(opts *bind.FilterOpts) (*PriceFeedInitializedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PriceFeedInitializedIterator{contract: _PriceFeed.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeed *PriceFeedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PriceFeedInitialized) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedInitialized)
				if err := _PriceFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeed *PriceFeedFilterer) ParseInitialized(log types.Log) (*PriceFeedInitialized, error) {
	event := new(PriceFeedInitialized)
	if err := _PriceFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeed contract.
type PriceFeedOwnershipTransferredIterator struct {
	Event *PriceFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeed contract.
type PriceFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeed *PriceFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedOwnershipTransferredIterator{contract: _PriceFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeed *PriceFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedOwnershipTransferred)
				if err := _PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeed *PriceFeedFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeedOwnershipTransferred, error) {
	event := new(PriceFeedOwnershipTransferred)
	if err := _PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedPriceFeedPublishedIterator is returned from FilterPriceFeedPublished and is used to iterate over the raw logs and unpacked data for PriceFeedPublished events raised by the PriceFeed contract.
type PriceFeedPriceFeedPublishedIterator struct {
	Event *PriceFeedPriceFeedPublished // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceFeedPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceFeedPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedPriceFeedPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedPriceFeedPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceFeedPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceFeedPublished represents a PriceFeedPublished event raised by the PriceFeed contract.
type PriceFeedPriceFeedPublished struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedPublished is a free log retrieval operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) FilterPriceFeedPublished(opts *bind.FilterOpts) (*PriceFeedPriceFeedPublishedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceFeedPublished")
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceFeedPublishedIterator{contract: _PriceFeed.contract, event: "PriceFeedPublished", logs: logs, sub: sub}, nil
}

// WatchPriceFeedPublished is a free log subscription operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) WatchPriceFeedPublished(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceFeedPublished) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceFeedPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceFeedPublished)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceFeedPublished is a log parse operation binding the contract event 0xf39d648fb8a9ac771561ab993f279c7eb13f1beb860d61c9cb3e787564d06803.
//
// Solidity: event PriceFeedPublished(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) ParsePriceFeedPublished(log types.Log) (*PriceFeedPriceFeedPublished, error) {
	event := new(PriceFeedPriceFeedPublished)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedPriceFeedRequestedIterator is returned from FilterPriceFeedRequested and is used to iterate over the raw logs and unpacked data for PriceFeedRequested events raised by the PriceFeed contract.
type PriceFeedPriceFeedRequestedIterator struct {
	Event *PriceFeedPriceFeedRequested // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceFeedRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceFeedRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedPriceFeedRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedPriceFeedRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceFeedRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceFeedRequested represents a PriceFeedRequested event raised by the PriceFeed contract.
type PriceFeedPriceFeedRequested struct {
	Requester common.Address
	Pairs     []string
	Prices    []*big.Int
	Decimals  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedRequested is a free log retrieval operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeed *PriceFeedFilterer) FilterPriceFeedRequested(opts *bind.FilterOpts, requester []common.Address) (*PriceFeedPriceFeedRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceFeedRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceFeedRequestedIterator{contract: _PriceFeed.contract, event: "PriceFeedRequested", logs: logs, sub: sub}, nil
}

// WatchPriceFeedRequested is a free log subscription operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeed *PriceFeedFilterer) WatchPriceFeedRequested(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceFeedRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceFeedRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceFeedRequested)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceFeedRequested is a log parse operation binding the contract event 0x21fb481b338df488246d581cca0377c59b67cc32b0ba179ac6d2c2975432aed1.
//
// Solidity: event PriceFeedRequested(address indexed requester, string[] pairs, uint256[] prices, uint256[] decimals)
func (_PriceFeed *PriceFeedFilterer) ParsePriceFeedRequested(log types.Log) (*PriceFeedPriceFeedRequested, error) {
	event := new(PriceFeedPriceFeedRequested)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedPriceFeedUpdatedIterator is returned from FilterPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedUpdated events raised by the PriceFeed contract.
type PriceFeedPriceFeedUpdatedIterator struct {
	Event *PriceFeedPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceFeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedPriceFeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceFeedUpdated represents a PriceFeedUpdated event raised by the PriceFeed contract.
type PriceFeedPriceFeedUpdated struct {
	Pair     string
	Price    *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdated is a free log retrieval operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) FilterPriceFeedUpdated(opts *bind.FilterOpts) (*PriceFeedPriceFeedUpdatedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceFeedUpdatedIterator{contract: _PriceFeed.contract, event: "PriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdated is a free log subscription operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) WatchPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceFeedUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceFeedUpdated)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceFeedUpdated is a log parse operation binding the contract event 0x3e51a618566783aa8f6e83296eb48686664aeb521acd6a7507a35fb5fd9ac740.
//
// Solidity: event PriceFeedUpdated(string pair, uint256 price, uint256 decimals)
func (_PriceFeed *PriceFeedFilterer) ParsePriceFeedUpdated(log types.Log) (*PriceFeedPriceFeedUpdated, error) {
	event := new(PriceFeedPriceFeedUpdated)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PriceFeed contract.
type PriceFeedUpgradedIterator struct {
	Event *PriceFeedUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedUpgraded represents a Upgraded event raised by the PriceFeed contract.
type PriceFeedUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PriceFeedUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedUpgradedIterator{contract: _PriceFeed.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedUpgraded)
				if err := _PriceFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) ParseUpgraded(log types.Log) (*PriceFeedUpgraded, error) {
	event := new(PriceFeedUpgraded)
	if err := _PriceFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
