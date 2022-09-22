// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package farcaster_registry

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
)

// FarcasterRegistryMetaData contains all meta data concerning the FarcasterRegistry contract.
var FarcasterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ChangeTrustedForwarder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"DeregisterName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"ModifyName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"RegisterName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"TransferName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToUsername\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"getDirectoryUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"modify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"idx\",\"type\":\"uint32\"}],\"name\":\"usernameAtIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usernameToUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usernamesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarcasterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FarcasterRegistryMetaData.ABI instead.
var FarcasterRegistryABI = FarcasterRegistryMetaData.ABI

// FarcasterRegistry is an auto generated Go binding around an Ethereum contract.
type FarcasterRegistry struct {
	FarcasterRegistryCaller     // Read-only binding to the contract
	FarcasterRegistryTransactor // Write-only binding to the contract
	FarcasterRegistryFilterer   // Log filterer for contract events
}

// FarcasterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarcasterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarcasterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarcasterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarcasterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarcasterRegistrySession struct {
	Contract     *FarcasterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FarcasterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarcasterRegistryCallerSession struct {
	Contract *FarcasterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FarcasterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarcasterRegistryTransactorSession struct {
	Contract     *FarcasterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FarcasterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarcasterRegistryRaw struct {
	Contract *FarcasterRegistry // Generic contract binding to access the raw methods on
}

// FarcasterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarcasterRegistryCallerRaw struct {
	Contract *FarcasterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FarcasterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarcasterRegistryTransactorRaw struct {
	Contract *FarcasterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarcasterRegistry creates a new instance of FarcasterRegistry, bound to a specific deployed contract.
func NewFarcasterRegistry(address common.Address, backend bind.ContractBackend) (*FarcasterRegistry, error) {
	contract, err := bindFarcasterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistry{FarcasterRegistryCaller: FarcasterRegistryCaller{contract: contract}, FarcasterRegistryTransactor: FarcasterRegistryTransactor{contract: contract}, FarcasterRegistryFilterer: FarcasterRegistryFilterer{contract: contract}}, nil
}

// NewFarcasterRegistryCaller creates a new read-only instance of FarcasterRegistry, bound to a specific deployed contract.
func NewFarcasterRegistryCaller(address common.Address, caller bind.ContractCaller) (*FarcasterRegistryCaller, error) {
	contract, err := bindFarcasterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryCaller{contract: contract}, nil
}

// NewFarcasterRegistryTransactor creates a new write-only instance of FarcasterRegistry, bound to a specific deployed contract.
func NewFarcasterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FarcasterRegistryTransactor, error) {
	contract, err := bindFarcasterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryTransactor{contract: contract}, nil
}

// NewFarcasterRegistryFilterer creates a new log filterer instance of FarcasterRegistry, bound to a specific deployed contract.
func NewFarcasterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FarcasterRegistryFilterer, error) {
	contract, err := bindFarcasterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryFilterer{contract: contract}, nil
}

// bindFarcasterRegistry binds a generic wrapper to an already deployed contract.
func bindFarcasterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FarcasterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterRegistry *FarcasterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterRegistry.Contract.FarcasterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterRegistry *FarcasterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.FarcasterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterRegistry *FarcasterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.FarcasterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarcasterRegistry *FarcasterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarcasterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarcasterRegistry *FarcasterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarcasterRegistry *FarcasterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FarcasterRegistry.Contract.DEFAULTADMINROLE(&_FarcasterRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FarcasterRegistry.Contract.DEFAULTADMINROLE(&_FarcasterRegistry.CallOpts)
}

// AddressToUsername is a free data retrieval call binding the contract method 0xe07a0baa.
//
// Solidity: function addressToUsername(address ) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCaller) AddressToUsername(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "addressToUsername", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToUsername is a free data retrieval call binding the contract method 0xe07a0baa.
//
// Solidity: function addressToUsername(address ) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistrySession) AddressToUsername(arg0 common.Address) ([32]byte, error) {
	return _FarcasterRegistry.Contract.AddressToUsername(&_FarcasterRegistry.CallOpts, arg0)
}

// AddressToUsername is a free data retrieval call binding the contract method 0xe07a0baa.
//
// Solidity: function addressToUsername(address ) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) AddressToUsername(arg0 common.Address) ([32]byte, error) {
	return _FarcasterRegistry.Contract.AddressToUsername(&_FarcasterRegistry.CallOpts, arg0)
}

// GetDirectoryUrl is a free data retrieval call binding the contract method 0x342740c9.
//
// Solidity: function getDirectoryUrl(bytes32 username) view returns(string)
func (_FarcasterRegistry *FarcasterRegistryCaller) GetDirectoryUrl(opts *bind.CallOpts, username [32]byte) (string, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "getDirectoryUrl", username)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDirectoryUrl is a free data retrieval call binding the contract method 0x342740c9.
//
// Solidity: function getDirectoryUrl(bytes32 username) view returns(string)
func (_FarcasterRegistry *FarcasterRegistrySession) GetDirectoryUrl(username [32]byte) (string, error) {
	return _FarcasterRegistry.Contract.GetDirectoryUrl(&_FarcasterRegistry.CallOpts, username)
}

// GetDirectoryUrl is a free data retrieval call binding the contract method 0x342740c9.
//
// Solidity: function getDirectoryUrl(bytes32 username) view returns(string)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) GetDirectoryUrl(username [32]byte) (string, error) {
	return _FarcasterRegistry.Contract.GetDirectoryUrl(&_FarcasterRegistry.CallOpts, username)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FarcasterRegistry.Contract.GetRoleAdmin(&_FarcasterRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FarcasterRegistry.Contract.GetRoleAdmin(&_FarcasterRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FarcasterRegistry.Contract.HasRole(&_FarcasterRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FarcasterRegistry.Contract.HasRole(&_FarcasterRegistry.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _FarcasterRegistry.Contract.IsTrustedForwarder(&_FarcasterRegistry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _FarcasterRegistry.Contract.IsTrustedForwarder(&_FarcasterRegistry.CallOpts, forwarder)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterRegistry *FarcasterRegistrySession) Paused() (bool, error) {
	return _FarcasterRegistry.Contract.Paused(&_FarcasterRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) Paused() (bool, error) {
	return _FarcasterRegistry.Contract.Paused(&_FarcasterRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FarcasterRegistry.Contract.SupportsInterface(&_FarcasterRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FarcasterRegistry.Contract.SupportsInterface(&_FarcasterRegistry.CallOpts, interfaceId)
}

// UsernameAtIndex is a free data retrieval call binding the contract method 0x0e25fc04.
//
// Solidity: function usernameAtIndex(uint32 idx) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCaller) UsernameAtIndex(opts *bind.CallOpts, idx uint32) ([32]byte, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "usernameAtIndex", idx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UsernameAtIndex is a free data retrieval call binding the contract method 0x0e25fc04.
//
// Solidity: function usernameAtIndex(uint32 idx) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistrySession) UsernameAtIndex(idx uint32) ([32]byte, error) {
	return _FarcasterRegistry.Contract.UsernameAtIndex(&_FarcasterRegistry.CallOpts, idx)
}

// UsernameAtIndex is a free data retrieval call binding the contract method 0x0e25fc04.
//
// Solidity: function usernameAtIndex(uint32 idx) view returns(bytes32)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) UsernameAtIndex(idx uint32) ([32]byte, error) {
	return _FarcasterRegistry.Contract.UsernameAtIndex(&_FarcasterRegistry.CallOpts, idx)
}

// UsernameToUrl is a free data retrieval call binding the contract method 0x33ee3c64.
//
// Solidity: function usernameToUrl(bytes32 ) view returns(string url, bool initialized)
func (_FarcasterRegistry *FarcasterRegistryCaller) UsernameToUrl(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Url         string
	Initialized bool
}, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "usernameToUrl", arg0)

	outstruct := new(struct {
		Url         string
		Initialized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Initialized = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// UsernameToUrl is a free data retrieval call binding the contract method 0x33ee3c64.
//
// Solidity: function usernameToUrl(bytes32 ) view returns(string url, bool initialized)
func (_FarcasterRegistry *FarcasterRegistrySession) UsernameToUrl(arg0 [32]byte) (struct {
	Url         string
	Initialized bool
}, error) {
	return _FarcasterRegistry.Contract.UsernameToUrl(&_FarcasterRegistry.CallOpts, arg0)
}

// UsernameToUrl is a free data retrieval call binding the contract method 0x33ee3c64.
//
// Solidity: function usernameToUrl(bytes32 ) view returns(string url, bool initialized)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) UsernameToUrl(arg0 [32]byte) (struct {
	Url         string
	Initialized bool
}, error) {
	return _FarcasterRegistry.Contract.UsernameToUrl(&_FarcasterRegistry.CallOpts, arg0)
}

// UsernamesLength is a free data retrieval call binding the contract method 0xc6b31772.
//
// Solidity: function usernamesLength() view returns(uint256)
func (_FarcasterRegistry *FarcasterRegistryCaller) UsernamesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarcasterRegistry.contract.Call(opts, &out, "usernamesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsernamesLength is a free data retrieval call binding the contract method 0xc6b31772.
//
// Solidity: function usernamesLength() view returns(uint256)
func (_FarcasterRegistry *FarcasterRegistrySession) UsernamesLength() (*big.Int, error) {
	return _FarcasterRegistry.Contract.UsernamesLength(&_FarcasterRegistry.CallOpts)
}

// UsernamesLength is a free data retrieval call binding the contract method 0xc6b31772.
//
// Solidity: function usernamesLength() view returns(uint256)
func (_FarcasterRegistry *FarcasterRegistryCallerSession) UsernamesLength() (*big.Int, error) {
	return _FarcasterRegistry.Contract.UsernamesLength(&_FarcasterRegistry.CallOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0x5cfcdbc5.
//
// Solidity: function deregister(address owner, bytes32 username) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Deregister(opts *bind.TransactOpts, owner common.Address, username [32]byte) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "deregister", owner, username)
}

// Deregister is a paid mutator transaction binding the contract method 0x5cfcdbc5.
//
// Solidity: function deregister(address owner, bytes32 username) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Deregister(owner common.Address, username [32]byte) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Deregister(&_FarcasterRegistry.TransactOpts, owner, username)
}

// Deregister is a paid mutator transaction binding the contract method 0x5cfcdbc5.
//
// Solidity: function deregister(address owner, bytes32 username) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Deregister(owner common.Address, username [32]byte) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Deregister(&_FarcasterRegistry.TransactOpts, owner, username)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.GrantRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.GrantRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Initialize(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "initialize", _forwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Initialize(_forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Initialize(&_FarcasterRegistry.TransactOpts, _forwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Initialize(_forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Initialize(&_FarcasterRegistry.TransactOpts, _forwarder)
}

// Modify is a paid mutator transaction binding the contract method 0x7f2eec02.
//
// Solidity: function modify(string url) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Modify(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "modify", url)
}

// Modify is a paid mutator transaction binding the contract method 0x7f2eec02.
//
// Solidity: function modify(string url) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Modify(url string) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Modify(&_FarcasterRegistry.TransactOpts, url)
}

// Modify is a paid mutator transaction binding the contract method 0x7f2eec02.
//
// Solidity: function modify(string url) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Modify(url string) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Modify(&_FarcasterRegistry.TransactOpts, url)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Pause() (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Pause(&_FarcasterRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Pause(&_FarcasterRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xcf2d31fb.
//
// Solidity: function register(bytes32 username, string url) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Register(opts *bind.TransactOpts, username [32]byte, url string) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "register", username, url)
}

// Register is a paid mutator transaction binding the contract method 0xcf2d31fb.
//
// Solidity: function register(bytes32 username, string url) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Register(username [32]byte, url string) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Register(&_FarcasterRegistry.TransactOpts, username, url)
}

// Register is a paid mutator transaction binding the contract method 0xcf2d31fb.
//
// Solidity: function register(bytes32 username, string url) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Register(username [32]byte, url string) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Register(&_FarcasterRegistry.TransactOpts, username, url)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.RenounceRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.RenounceRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.RevokeRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.RevokeRole(&_FarcasterRegistry.TransactOpts, role, account)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) SetTrustedForwarder(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "setTrustedForwarder", _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.SetTrustedForwarder(&_FarcasterRegistry.TransactOpts, _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.SetTrustedForwarder(&_FarcasterRegistry.TransactOpts, _forwarder)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Transfer(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "transfer", to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Transfer(to common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Transfer(&_FarcasterRegistry.TransactOpts, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address to) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Transfer(to common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Transfer(&_FarcasterRegistry.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.TransferOwnership(&_FarcasterRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.TransferOwnership(&_FarcasterRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterRegistry *FarcasterRegistrySession) Unpause() (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Unpause(&_FarcasterRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.Unpause(&_FarcasterRegistry.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FarcasterRegistry *FarcasterRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.UpgradeTo(&_FarcasterRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.UpgradeTo(&_FarcasterRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FarcasterRegistry *FarcasterRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FarcasterRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FarcasterRegistry *FarcasterRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.UpgradeToAndCall(&_FarcasterRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FarcasterRegistry *FarcasterRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FarcasterRegistry.Contract.UpgradeToAndCall(&_FarcasterRegistry.TransactOpts, newImplementation, data)
}

// FarcasterRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FarcasterRegistry contract.
type FarcasterRegistryAdminChangedIterator struct {
	Event *FarcasterRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryAdminChanged)
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
		it.Event = new(FarcasterRegistryAdminChanged)
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
func (it *FarcasterRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryAdminChanged represents a AdminChanged event raised by the FarcasterRegistry contract.
type FarcasterRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FarcasterRegistryAdminChangedIterator, error) {

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryAdminChangedIterator{contract: _FarcasterRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryAdminChanged)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseAdminChanged(log types.Log) (*FarcasterRegistryAdminChanged, error) {
	event := new(FarcasterRegistryAdminChanged)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the FarcasterRegistry contract.
type FarcasterRegistryBeaconUpgradedIterator struct {
	Event *FarcasterRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryBeaconUpgraded)
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
		it.Event = new(FarcasterRegistryBeaconUpgraded)
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
func (it *FarcasterRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the FarcasterRegistry contract.
type FarcasterRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FarcasterRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryBeaconUpgradedIterator{contract: _FarcasterRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryBeaconUpgraded)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*FarcasterRegistryBeaconUpgraded, error) {
	event := new(FarcasterRegistryBeaconUpgraded)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryChangeTrustedForwarderIterator is returned from FilterChangeTrustedForwarder and is used to iterate over the raw logs and unpacked data for ChangeTrustedForwarder events raised by the FarcasterRegistry contract.
type FarcasterRegistryChangeTrustedForwarderIterator struct {
	Event *FarcasterRegistryChangeTrustedForwarder // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryChangeTrustedForwarderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryChangeTrustedForwarder)
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
		it.Event = new(FarcasterRegistryChangeTrustedForwarder)
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
func (it *FarcasterRegistryChangeTrustedForwarderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryChangeTrustedForwarderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryChangeTrustedForwarder represents a ChangeTrustedForwarder event raised by the FarcasterRegistry contract.
type FarcasterRegistryChangeTrustedForwarder struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeTrustedForwarder is a free log retrieval operation binding the contract event 0x78e4799a3a5df23ad7ce725d1c9379640feb1547d54e887bc70beb9c6d974068.
//
// Solidity: event ChangeTrustedForwarder(address indexed to)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterChangeTrustedForwarder(opts *bind.FilterOpts, to []common.Address) (*FarcasterRegistryChangeTrustedForwarderIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "ChangeTrustedForwarder", toRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryChangeTrustedForwarderIterator{contract: _FarcasterRegistry.contract, event: "ChangeTrustedForwarder", logs: logs, sub: sub}, nil
}

// WatchChangeTrustedForwarder is a free log subscription operation binding the contract event 0x78e4799a3a5df23ad7ce725d1c9379640feb1547d54e887bc70beb9c6d974068.
//
// Solidity: event ChangeTrustedForwarder(address indexed to)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchChangeTrustedForwarder(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryChangeTrustedForwarder, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "ChangeTrustedForwarder", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryChangeTrustedForwarder)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "ChangeTrustedForwarder", log); err != nil {
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

// ParseChangeTrustedForwarder is a log parse operation binding the contract event 0x78e4799a3a5df23ad7ce725d1c9379640feb1547d54e887bc70beb9c6d974068.
//
// Solidity: event ChangeTrustedForwarder(address indexed to)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseChangeTrustedForwarder(log types.Log) (*FarcasterRegistryChangeTrustedForwarder, error) {
	event := new(FarcasterRegistryChangeTrustedForwarder)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "ChangeTrustedForwarder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryDeregisterNameIterator is returned from FilterDeregisterName and is used to iterate over the raw logs and unpacked data for DeregisterName events raised by the FarcasterRegistry contract.
type FarcasterRegistryDeregisterNameIterator struct {
	Event *FarcasterRegistryDeregisterName // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryDeregisterNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryDeregisterName)
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
		it.Event = new(FarcasterRegistryDeregisterName)
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
func (it *FarcasterRegistryDeregisterNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryDeregisterNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryDeregisterName represents a DeregisterName event raised by the FarcasterRegistry contract.
type FarcasterRegistryDeregisterName struct {
	Owner    common.Address
	Username [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeregisterName is a free log retrieval operation binding the contract event 0x189646b792a1bc6317676c67112fd10dd6372b26fee015505adf4852c0869ea2.
//
// Solidity: event DeregisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterDeregisterName(opts *bind.FilterOpts, owner []common.Address, username [][32]byte) (*FarcasterRegistryDeregisterNameIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "DeregisterName", ownerRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryDeregisterNameIterator{contract: _FarcasterRegistry.contract, event: "DeregisterName", logs: logs, sub: sub}, nil
}

// WatchDeregisterName is a free log subscription operation binding the contract event 0x189646b792a1bc6317676c67112fd10dd6372b26fee015505adf4852c0869ea2.
//
// Solidity: event DeregisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchDeregisterName(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryDeregisterName, owner []common.Address, username [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "DeregisterName", ownerRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryDeregisterName)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "DeregisterName", log); err != nil {
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

// ParseDeregisterName is a log parse operation binding the contract event 0x189646b792a1bc6317676c67112fd10dd6372b26fee015505adf4852c0869ea2.
//
// Solidity: event DeregisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseDeregisterName(log types.Log) (*FarcasterRegistryDeregisterName, error) {
	event := new(FarcasterRegistryDeregisterName)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "DeregisterName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryModifyNameIterator is returned from FilterModifyName and is used to iterate over the raw logs and unpacked data for ModifyName events raised by the FarcasterRegistry contract.
type FarcasterRegistryModifyNameIterator struct {
	Event *FarcasterRegistryModifyName // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryModifyNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryModifyName)
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
		it.Event = new(FarcasterRegistryModifyName)
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
func (it *FarcasterRegistryModifyNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryModifyNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryModifyName represents a ModifyName event raised by the FarcasterRegistry contract.
type FarcasterRegistryModifyName struct {
	Username [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifyName is a free log retrieval operation binding the contract event 0x540d4f72bf4e259b973b675cb40ca044ae64eb1ac9369550c39aaed4b8d70400.
//
// Solidity: event ModifyName(bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterModifyName(opts *bind.FilterOpts, username [][32]byte) (*FarcasterRegistryModifyNameIterator, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "ModifyName", usernameRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryModifyNameIterator{contract: _FarcasterRegistry.contract, event: "ModifyName", logs: logs, sub: sub}, nil
}

// WatchModifyName is a free log subscription operation binding the contract event 0x540d4f72bf4e259b973b675cb40ca044ae64eb1ac9369550c39aaed4b8d70400.
//
// Solidity: event ModifyName(bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchModifyName(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryModifyName, username [][32]byte) (event.Subscription, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "ModifyName", usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryModifyName)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "ModifyName", log); err != nil {
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

// ParseModifyName is a log parse operation binding the contract event 0x540d4f72bf4e259b973b675cb40ca044ae64eb1ac9369550c39aaed4b8d70400.
//
// Solidity: event ModifyName(bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseModifyName(log types.Log) (*FarcasterRegistryModifyName, error) {
	event := new(FarcasterRegistryModifyName)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "ModifyName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the FarcasterRegistry contract.
type FarcasterRegistryPausedIterator struct {
	Event *FarcasterRegistryPaused // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryPaused)
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
		it.Event = new(FarcasterRegistryPaused)
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
func (it *FarcasterRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryPaused represents a Paused event raised by the FarcasterRegistry contract.
type FarcasterRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*FarcasterRegistryPausedIterator, error) {

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryPausedIterator{contract: _FarcasterRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryPaused)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParsePaused(log types.Log) (*FarcasterRegistryPaused, error) {
	event := new(FarcasterRegistryPaused)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryRegisterNameIterator is returned from FilterRegisterName and is used to iterate over the raw logs and unpacked data for RegisterName events raised by the FarcasterRegistry contract.
type FarcasterRegistryRegisterNameIterator struct {
	Event *FarcasterRegistryRegisterName // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryRegisterNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryRegisterName)
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
		it.Event = new(FarcasterRegistryRegisterName)
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
func (it *FarcasterRegistryRegisterNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryRegisterNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryRegisterName represents a RegisterName event raised by the FarcasterRegistry contract.
type FarcasterRegistryRegisterName struct {
	Owner    common.Address
	Username [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegisterName is a free log retrieval operation binding the contract event 0x00af56d6ef7b1de1b37e4636c3e255d78dc3e9359036fb04ca51e32d946a6834.
//
// Solidity: event RegisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterRegisterName(opts *bind.FilterOpts, owner []common.Address, username [][32]byte) (*FarcasterRegistryRegisterNameIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "RegisterName", ownerRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryRegisterNameIterator{contract: _FarcasterRegistry.contract, event: "RegisterName", logs: logs, sub: sub}, nil
}

// WatchRegisterName is a free log subscription operation binding the contract event 0x00af56d6ef7b1de1b37e4636c3e255d78dc3e9359036fb04ca51e32d946a6834.
//
// Solidity: event RegisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchRegisterName(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryRegisterName, owner []common.Address, username [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "RegisterName", ownerRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryRegisterName)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "RegisterName", log); err != nil {
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

// ParseRegisterName is a log parse operation binding the contract event 0x00af56d6ef7b1de1b37e4636c3e255d78dc3e9359036fb04ca51e32d946a6834.
//
// Solidity: event RegisterName(address indexed owner, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseRegisterName(log types.Log) (*FarcasterRegistryRegisterName, error) {
	event := new(FarcasterRegistryRegisterName)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "RegisterName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleAdminChangedIterator struct {
	Event *FarcasterRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryRoleAdminChanged)
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
		it.Event = new(FarcasterRegistryRoleAdminChanged)
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
func (it *FarcasterRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FarcasterRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryRoleAdminChangedIterator{contract: _FarcasterRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryRoleAdminChanged)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*FarcasterRegistryRoleAdminChanged, error) {
	event := new(FarcasterRegistryRoleAdminChanged)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleGrantedIterator struct {
	Event *FarcasterRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryRoleGranted)
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
		it.Event = new(FarcasterRegistryRoleGranted)
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
func (it *FarcasterRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryRoleGranted represents a RoleGranted event raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FarcasterRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryRoleGrantedIterator{contract: _FarcasterRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryRoleGranted)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseRoleGranted(log types.Log) (*FarcasterRegistryRoleGranted, error) {
	event := new(FarcasterRegistryRoleGranted)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleRevokedIterator struct {
	Event *FarcasterRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryRoleRevoked)
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
		it.Event = new(FarcasterRegistryRoleRevoked)
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
func (it *FarcasterRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryRoleRevoked represents a RoleRevoked event raised by the FarcasterRegistry contract.
type FarcasterRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FarcasterRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryRoleRevokedIterator{contract: _FarcasterRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryRoleRevoked)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseRoleRevoked(log types.Log) (*FarcasterRegistryRoleRevoked, error) {
	event := new(FarcasterRegistryRoleRevoked)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryTransferNameIterator is returned from FilterTransferName and is used to iterate over the raw logs and unpacked data for TransferName events raised by the FarcasterRegistry contract.
type FarcasterRegistryTransferNameIterator struct {
	Event *FarcasterRegistryTransferName // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryTransferNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryTransferName)
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
		it.Event = new(FarcasterRegistryTransferName)
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
func (it *FarcasterRegistryTransferNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryTransferNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryTransferName represents a TransferName event raised by the FarcasterRegistry contract.
type FarcasterRegistryTransferName struct {
	From     common.Address
	To       common.Address
	Username [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferName is a free log retrieval operation binding the contract event 0xaf3f7950be1072c7c21e4ff9b6e3a459f995d68a1c325de2eb91450a2699cbd4.
//
// Solidity: event TransferName(address indexed from, address indexed to, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterTransferName(opts *bind.FilterOpts, from []common.Address, to []common.Address, username [][32]byte) (*FarcasterRegistryTransferNameIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "TransferName", fromRule, toRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryTransferNameIterator{contract: _FarcasterRegistry.contract, event: "TransferName", logs: logs, sub: sub}, nil
}

// WatchTransferName is a free log subscription operation binding the contract event 0xaf3f7950be1072c7c21e4ff9b6e3a459f995d68a1c325de2eb91450a2699cbd4.
//
// Solidity: event TransferName(address indexed from, address indexed to, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchTransferName(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryTransferName, from []common.Address, to []common.Address, username [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "TransferName", fromRule, toRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryTransferName)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "TransferName", log); err != nil {
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

// ParseTransferName is a log parse operation binding the contract event 0xaf3f7950be1072c7c21e4ff9b6e3a459f995d68a1c325de2eb91450a2699cbd4.
//
// Solidity: event TransferName(address indexed from, address indexed to, bytes32 indexed username)
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseTransferName(log types.Log) (*FarcasterRegistryTransferName, error) {
	event := new(FarcasterRegistryTransferName)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "TransferName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the FarcasterRegistry contract.
type FarcasterRegistryUnpausedIterator struct {
	Event *FarcasterRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryUnpaused)
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
		it.Event = new(FarcasterRegistryUnpaused)
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
func (it *FarcasterRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryUnpaused represents a Unpaused event raised by the FarcasterRegistry contract.
type FarcasterRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FarcasterRegistryUnpausedIterator, error) {

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryUnpausedIterator{contract: _FarcasterRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryUnpaused)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseUnpaused(log types.Log) (*FarcasterRegistryUnpaused, error) {
	event := new(FarcasterRegistryUnpaused)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarcasterRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FarcasterRegistry contract.
type FarcasterRegistryUpgradedIterator struct {
	Event *FarcasterRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *FarcasterRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarcasterRegistryUpgraded)
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
		it.Event = new(FarcasterRegistryUpgraded)
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
func (it *FarcasterRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarcasterRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarcasterRegistryUpgraded represents a Upgraded event raised by the FarcasterRegistry contract.
type FarcasterRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FarcasterRegistry *FarcasterRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FarcasterRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FarcasterRegistryUpgradedIterator{contract: _FarcasterRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FarcasterRegistry *FarcasterRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FarcasterRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FarcasterRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarcasterRegistryUpgraded)
				if err := _FarcasterRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FarcasterRegistry *FarcasterRegistryFilterer) ParseUpgraded(log types.Log) (*FarcasterRegistryUpgraded, error) {
	event := new(FarcasterRegistryUpgraded)
	if err := _FarcasterRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
