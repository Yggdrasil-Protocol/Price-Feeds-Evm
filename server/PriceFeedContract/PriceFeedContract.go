// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PriceFeedContract

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

// PriceFeedContractMetaData contains all meta data concerning the PriceFeedContract contract.
var PriceFeedContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"PriceNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"TooManyAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeePerAsset\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"assets\",\"type\":\"bytes32[]\"}],\"name\":\"PricesRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"TrustedSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"MAX_ASSETS_PER_REQUEST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeePerAsset\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_assets\",\"type\":\"bytes32[]\"},{\"internalType\":\"function(uint8[],uint256[])external\",\"name\":\"_callback\",\"type\":\"function\"}],\"name\":\"requestPrices\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeePerAsset\",\"type\":\"uint256\"}],\"name\":\"setFeePerAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSigner\",\"type\":\"address\"}],\"name\":\"setTrustedSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_assets\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PriceFeedContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedContractMetaData.ABI instead.
var PriceFeedContractABI = PriceFeedContractMetaData.ABI

// PriceFeedContract is an auto generated Go binding around an Ethereum contract.
type PriceFeedContract struct {
	PriceFeedContractCaller     // Read-only binding to the contract
	PriceFeedContractTransactor // Write-only binding to the contract
	PriceFeedContractFilterer   // Log filterer for contract events
}

// PriceFeedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedContractSession struct {
	Contract     *PriceFeedContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceFeedContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedContractCallerSession struct {
	Contract *PriceFeedContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceFeedContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedContractTransactorSession struct {
	Contract     *PriceFeedContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceFeedContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedContractRaw struct {
	Contract *PriceFeedContract // Generic contract binding to access the raw methods on
}

// PriceFeedContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedContractCallerRaw struct {
	Contract *PriceFeedContractCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactorRaw struct {
	Contract *PriceFeedContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedContract creates a new instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContract(address common.Address, backend bind.ContractBackend) (*PriceFeedContract, error) {
	contract, err := bindPriceFeedContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContract{PriceFeedContractCaller: PriceFeedContractCaller{contract: contract}, PriceFeedContractTransactor: PriceFeedContractTransactor{contract: contract}, PriceFeedContractFilterer: PriceFeedContractFilterer{contract: contract}}, nil
}

// NewPriceFeedContractCaller creates a new read-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedContractCaller, error) {
	contract, err := bindPriceFeedContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractCaller{contract: contract}, nil
}

// NewPriceFeedContractTransactor creates a new write-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedContractTransactor, error) {
	contract, err := bindPriceFeedContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractTransactor{contract: contract}, nil
}

// NewPriceFeedContractFilterer creates a new log filterer instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedContractFilterer, error) {
	contract, err := bindPriceFeedContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractFilterer{contract: contract}, nil
}

// bindPriceFeedContract binds a generic wrapper to an already deployed contract.
func bindPriceFeedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.PriceFeedContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transact(opts, method, params...)
}

// MAXASSETSPERREQUEST is a free data retrieval call binding the contract method 0x24b335cf.
//
// Solidity: function MAX_ASSETS_PER_REQUEST() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) MAXASSETSPERREQUEST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "MAX_ASSETS_PER_REQUEST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXASSETSPERREQUEST is a free data retrieval call binding the contract method 0x24b335cf.
//
// Solidity: function MAX_ASSETS_PER_REQUEST() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) MAXASSETSPERREQUEST() (*big.Int, error) {
	return _PriceFeedContract.Contract.MAXASSETSPERREQUEST(&_PriceFeedContract.CallOpts)
}

// MAXASSETSPERREQUEST is a free data retrieval call binding the contract method 0x24b335cf.
//
// Solidity: function MAX_ASSETS_PER_REQUEST() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) MAXASSETSPERREQUEST() (*big.Int, error) {
	return _PriceFeedContract.Contract.MAXASSETSPERREQUEST(&_PriceFeedContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x5f18aa0c.
//
// Solidity: function decimals(bytes32 ) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractCaller) Decimals(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "decimals", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x5f18aa0c.
//
// Solidity: function decimals(bytes32 ) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractSession) Decimals(arg0 [32]byte) (uint8, error) {
	return _PriceFeedContract.Contract.Decimals(&_PriceFeedContract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x5f18aa0c.
//
// Solidity: function decimals(bytes32 ) view returns(uint8)
func (_PriceFeedContract *PriceFeedContractCallerSession) Decimals(arg0 [32]byte) (uint8, error) {
	return _PriceFeedContract.Contract.Decimals(&_PriceFeedContract.CallOpts, arg0)
}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) FeePerAsset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "feePerAsset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) FeePerAsset() (*big.Int, error) {
	return _PriceFeedContract.Contract.FeePerAsset(&_PriceFeedContract.CallOpts)
}

// FeePerAsset is a free data retrieval call binding the contract method 0x6eefe8fa.
//
// Solidity: function feePerAsset() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) FeePerAsset() (*big.Int, error) {
	return _PriceFeedContract.Contract.FeePerAsset(&_PriceFeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCallerSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractSession) Paused() (bool, error) {
	return _PriceFeedContract.Contract.Paused(&_PriceFeedContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceFeedContract *PriceFeedContractCallerSession) Paused() (bool, error) {
	return _PriceFeedContract.Contract.Paused(&_PriceFeedContract.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.Prices(&_PriceFeedContract.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) Prices(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.Prices(&_PriceFeedContract.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeedContract.Contract.ProxiableUUID(&_PriceFeedContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeedContract *PriceFeedContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeedContract.Contract.ProxiableUUID(&_PriceFeedContract.CallOpts)
}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCaller) TrustedSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "trustedSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_PriceFeedContract *PriceFeedContractSession) TrustedSigner() (common.Address, error) {
	return _PriceFeedContract.Contract.TrustedSigner(&_PriceFeedContract.CallOpts)
}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCallerSession) TrustedSigner() (common.Address, error) {
	return _PriceFeedContract.Contract.TrustedSigner(&_PriceFeedContract.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Initialize(opts *bind.TransactOpts, _newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "initialize", _newFeePerAsset)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractSession) Initialize(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Initialize(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Initialize(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Initialize(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RenounceOwnership(&_PriceFeedContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RenounceOwnership(&_PriceFeedContract.TransactOpts)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) RequestPrices(opts *bind.TransactOpts, _assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "requestPrices", _assets, _callback)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractSession) RequestPrices(_assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPrices(&_PriceFeedContract.TransactOpts, _assets, _callback)
}

// RequestPrices is a paid mutator transaction binding the contract method 0xb2ad4c2e.
//
// Solidity: function requestPrices(bytes32[] _assets, function _callback) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) RequestPrices(_assets [][32]byte, _callback [24]byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.RequestPrices(&_PriceFeedContract.TransactOpts, _assets, _callback)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) SetFeePerAsset(opts *bind.TransactOpts, _newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "setFeePerAsset", _newFeePerAsset)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractSession) SetFeePerAsset(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetFeePerAsset(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// SetFeePerAsset is a paid mutator transaction binding the contract method 0x6f1518f8.
//
// Solidity: function setFeePerAsset(uint256 _newFeePerAsset) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) SetFeePerAsset(_newFeePerAsset *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetFeePerAsset(&_PriceFeedContract.TransactOpts, _newFeePerAsset)
}

// SetTrustedSigner is a paid mutator transaction binding the contract method 0x56a1c701.
//
// Solidity: function setTrustedSigner(address _newSigner) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) SetTrustedSigner(opts *bind.TransactOpts, _newSigner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "setTrustedSigner", _newSigner)
}

// SetTrustedSigner is a paid mutator transaction binding the contract method 0x56a1c701.
//
// Solidity: function setTrustedSigner(address _newSigner) returns()
func (_PriceFeedContract *PriceFeedContractSession) SetTrustedSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetTrustedSigner(&_PriceFeedContract.TransactOpts, _newSigner)
}

// SetTrustedSigner is a paid mutator transaction binding the contract method 0x56a1c701.
//
// Solidity: function setTrustedSigner(address _newSigner) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) SetTrustedSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetTrustedSigner(&_PriceFeedContract.TransactOpts, _newSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.TransferOwnership(&_PriceFeedContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.TransferOwnership(&_PriceFeedContract.TransactOpts, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x6fe0a68d.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices, bytes _signature) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpdatePrice(opts *bind.TransactOpts, _assets [][32]byte, _decimals []uint8, _prices []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "updatePrice", _assets, _decimals, _prices, _signature)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x6fe0a68d.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices, bytes _signature) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpdatePrice(_assets [][32]byte, _decimals []uint8, _prices []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePrice(&_PriceFeedContract.TransactOpts, _assets, _decimals, _prices, _signature)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x6fe0a68d.
//
// Solidity: function updatePrice(bytes32[] _assets, uint8[] _decimals, uint256[] _prices, bytes _signature) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpdatePrice(_assets [][32]byte, _decimals []uint8, _prices []*big.Int, _signature []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdatePrice(&_PriceFeedContract.TransactOpts, _assets, _decimals, _prices, _signature)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeTo(&_PriceFeedContract.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeTo(&_PriceFeedContract.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeToAndCall(&_PriceFeedContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpgradeToAndCall(&_PriceFeedContract.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractSession) Withdraw() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Withdraw(&_PriceFeedContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Withdraw(&_PriceFeedContract.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Fallback(&_PriceFeedContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Fallback(&_PriceFeedContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractSession) Receive() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Receive(&_PriceFeedContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) Receive() (*types.Transaction, error) {
	return _PriceFeedContract.Contract.Receive(&_PriceFeedContract.TransactOpts)
}

// PriceFeedContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PriceFeedContract contract.
type PriceFeedContractAdminChangedIterator struct {
	Event *PriceFeedContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractAdminChanged)
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
		it.Event = new(PriceFeedContractAdminChanged)
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
func (it *PriceFeedContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractAdminChanged represents a AdminChanged event raised by the PriceFeedContract contract.
type PriceFeedContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PriceFeedContractAdminChangedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractAdminChangedIterator{contract: _PriceFeedContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractAdminChanged)
				if err := _PriceFeedContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_PriceFeedContract *PriceFeedContractFilterer) ParseAdminChanged(log types.Log) (*PriceFeedContractAdminChanged, error) {
	event := new(PriceFeedContractAdminChanged)
	if err := _PriceFeedContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PriceFeedContract contract.
type PriceFeedContractBeaconUpgradedIterator struct {
	Event *PriceFeedContractBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractBeaconUpgraded)
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
		it.Event = new(PriceFeedContractBeaconUpgraded)
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
func (it *PriceFeedContractBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractBeaconUpgraded represents a BeaconUpgraded event raised by the PriceFeedContract contract.
type PriceFeedContractBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PriceFeedContractBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractBeaconUpgradedIterator{contract: _PriceFeedContract.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedContractBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractBeaconUpgraded)
				if err := _PriceFeedContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_PriceFeedContract *PriceFeedContractFilterer) ParseBeaconUpgraded(log types.Log) (*PriceFeedContractBeaconUpgraded, error) {
	event := new(PriceFeedContractBeaconUpgraded)
	if err := _PriceFeedContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractFeeUpdatedIterator struct {
	Event *PriceFeedContractFeeUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractFeeUpdated)
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
		it.Event = new(PriceFeedContractFeeUpdated)
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
func (it *PriceFeedContractFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractFeeUpdated represents a FeeUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractFeeUpdated struct {
	NewFeePerAsset *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*PriceFeedContractFeeUpdatedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractFeeUpdatedIterator{contract: _PriceFeedContract.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractFeeUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFeePerAsset)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseFeeUpdated(log types.Log) (*PriceFeedContractFeeUpdated, error) {
	event := new(PriceFeedContractFeeUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PriceFeedContract contract.
type PriceFeedContractInitializedIterator struct {
	Event *PriceFeedContractInitialized // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractInitialized)
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
		it.Event = new(PriceFeedContractInitialized)
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
func (it *PriceFeedContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractInitialized represents a Initialized event raised by the PriceFeedContract contract.
type PriceFeedContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*PriceFeedContractInitializedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractInitializedIterator{contract: _PriceFeedContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PriceFeedContractInitialized) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractInitialized)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PriceFeedContract *PriceFeedContractFilterer) ParseInitialized(log types.Log) (*PriceFeedContractInitialized, error) {
	event := new(PriceFeedContractInitialized)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeedContract contract.
type PriceFeedContractOwnershipTransferredIterator struct {
	Event *PriceFeedContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractOwnershipTransferred)
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
		it.Event = new(PriceFeedContractOwnershipTransferred)
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
func (it *PriceFeedContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeedContract contract.
type PriceFeedContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeedContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractOwnershipTransferredIterator{contract: _PriceFeedContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeedContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractOwnershipTransferred)
				if err := _PriceFeedContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceFeedContract *PriceFeedContractFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeedContractOwnershipTransferred, error) {
	event := new(PriceFeedContractOwnershipTransferred)
	if err := _PriceFeedContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceFeedContract contract.
type PriceFeedContractPausedIterator struct {
	Event *PriceFeedContractPaused // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPaused)
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
		it.Event = new(PriceFeedContractPaused)
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
func (it *PriceFeedContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPaused represents a Paused event raised by the PriceFeedContract contract.
type PriceFeedContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPaused(opts *bind.FilterOpts) (*PriceFeedContractPausedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPausedIterator{contract: _PriceFeedContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPaused) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPaused)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePaused(log types.Log) (*PriceFeedContractPaused, error) {
	event := new(PriceFeedContractPaused)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractPriceUpdatedIterator struct {
	Event *PriceFeedContractPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceUpdated)
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
		it.Event = new(PriceFeedContractPriceUpdated)
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
func (it *PriceFeedContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceUpdated represents a PriceUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractPriceUpdated struct {
	Asset   [32]byte
	Price   *big.Int
	Decimal uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts, asset [][32]byte) (*PriceFeedContractPriceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceUpdatedIterator{contract: _PriceFeedContract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceUpdated, asset [][32]byte) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x8c62fe32113aae6ed87fdea7c7da35e9d2b99790bf4f577b0869fe9bbe12d58f.
//
// Solidity: event PriceUpdated(bytes32 indexed asset, uint256 price, uint8 decimal)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceUpdated(log types.Log) (*PriceFeedContractPriceUpdated, error) {
	event := new(PriceFeedContractPriceUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractPricesRequestedIterator is returned from FilterPricesRequested and is used to iterate over the raw logs and unpacked data for PricesRequested events raised by the PriceFeedContract contract.
type PriceFeedContractPricesRequestedIterator struct {
	Event *PriceFeedContractPricesRequested // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractPricesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPricesRequested)
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
		it.Event = new(PriceFeedContractPricesRequested)
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
func (it *PriceFeedContractPricesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPricesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPricesRequested represents a PricesRequested event raised by the PriceFeedContract contract.
type PriceFeedContractPricesRequested struct {
	Requester common.Address
	Assets    [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPricesRequested is a free log retrieval operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPricesRequested(opts *bind.FilterOpts, requester []common.Address) (*PriceFeedContractPricesRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PricesRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPricesRequestedIterator{contract: _PriceFeedContract.contract, event: "PricesRequested", logs: logs, sub: sub}, nil
}

// WatchPricesRequested is a free log subscription operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPricesRequested(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPricesRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PricesRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPricesRequested)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PricesRequested", log); err != nil {
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

// ParsePricesRequested is a log parse operation binding the contract event 0xaa33eb4cfd36e0ee1d8d133a19a2620fb088f12c0141f8bc7e2d48c220be343f.
//
// Solidity: event PricesRequested(address indexed requester, bytes32[] assets)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePricesRequested(log types.Log) (*PriceFeedContractPricesRequested, error) {
	event := new(PriceFeedContractPricesRequested)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PricesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractTrustedSignerUpdatedIterator is returned from FilterTrustedSignerUpdated and is used to iterate over the raw logs and unpacked data for TrustedSignerUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractTrustedSignerUpdatedIterator struct {
	Event *PriceFeedContractTrustedSignerUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractTrustedSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractTrustedSignerUpdated)
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
		it.Event = new(PriceFeedContractTrustedSignerUpdated)
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
func (it *PriceFeedContractTrustedSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractTrustedSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractTrustedSignerUpdated represents a TrustedSignerUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractTrustedSignerUpdated struct {
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedSignerUpdated is a free log retrieval operation binding the contract event 0xced828023e9f94aa2a4373d4254f9bab903a9b34b822ab716d6c3688b2d0e17b.
//
// Solidity: event TrustedSignerUpdated(address newSigner)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterTrustedSignerUpdated(opts *bind.FilterOpts) (*PriceFeedContractTrustedSignerUpdatedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "TrustedSignerUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractTrustedSignerUpdatedIterator{contract: _PriceFeedContract.contract, event: "TrustedSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedSignerUpdated is a free log subscription operation binding the contract event 0xced828023e9f94aa2a4373d4254f9bab903a9b34b822ab716d6c3688b2d0e17b.
//
// Solidity: event TrustedSignerUpdated(address newSigner)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchTrustedSignerUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractTrustedSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "TrustedSignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractTrustedSignerUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "TrustedSignerUpdated", log); err != nil {
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

// ParseTrustedSignerUpdated is a log parse operation binding the contract event 0xced828023e9f94aa2a4373d4254f9bab903a9b34b822ab716d6c3688b2d0e17b.
//
// Solidity: event TrustedSignerUpdated(address newSigner)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseTrustedSignerUpdated(log types.Log) (*PriceFeedContractTrustedSignerUpdated, error) {
	event := new(PriceFeedContractTrustedSignerUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "TrustedSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceFeedContract contract.
type PriceFeedContractUnpausedIterator struct {
	Event *PriceFeedContractUnpaused // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractUnpaused)
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
		it.Event = new(PriceFeedContractUnpaused)
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
func (it *PriceFeedContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractUnpaused represents a Unpaused event raised by the PriceFeedContract contract.
type PriceFeedContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceFeedContractUnpausedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractUnpausedIterator{contract: _PriceFeedContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceFeedContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractUnpaused)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseUnpaused(log types.Log) (*PriceFeedContractUnpaused, error) {
	event := new(PriceFeedContractUnpaused)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PriceFeedContract contract.
type PriceFeedContractUpgradedIterator struct {
	Event *PriceFeedContractUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractUpgraded)
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
		it.Event = new(PriceFeedContractUpgraded)
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
func (it *PriceFeedContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractUpgraded represents a Upgraded event raised by the PriceFeedContract contract.
type PriceFeedContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PriceFeedContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractUpgradedIterator{contract: _PriceFeedContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractUpgraded)
				if err := _PriceFeedContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PriceFeedContract *PriceFeedContractFilterer) ParseUpgraded(log types.Log) (*PriceFeedContractUpgraded, error) {
	event := new(PriceFeedContractUpgraded)
	if err := _PriceFeedContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
