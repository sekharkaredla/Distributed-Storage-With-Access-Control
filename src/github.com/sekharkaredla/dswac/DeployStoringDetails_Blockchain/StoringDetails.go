// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoringDetailsABI is the input ABI used to generate the binding from.
const StoringDetailsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getJWT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newJWT\",\"type\":\"bytes\"}],\"name\":\"setJWT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"JWT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialJWT\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// StoringDetailsBin is the compiled bytecode used for deploying new contracts.
const StoringDetailsBin = `0x608060405234801561001057600080fd5b506040516104fe3803806104fe8339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815164010000000081118282018710171561007857600080fd5b505060018054600160a060020a0319163317905580519093506100a492506000915060208401906100ab565b5050610146565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ec57805160ff1916838001178555610119565b82800160010185558215610119579182015b828111156101195782518255916020019190600101906100fe565b50610125929150610129565b5090565b61014391905b80821115610125576000815560010161012f565b90565b6103a9806101556000396000f3fe608060405260043610610050577c010000000000000000000000000000000000000000000000000000000060003504632cd38f8c81146100555780633b04e3a3146100df57806374c81a9c14610194575b600080fd5b34801561006157600080fd5b5061006a6101a9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100a457818101518382015260200161008c565b50505050905090810190601f1680156100d15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100eb57600080fd5b506101926004803603602081101561010257600080fd5b81019060208101813564010000000081111561011d57600080fd5b82018360208201111561012f57600080fd5b8035906020019184600183028401116401000000008311171561015157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610240945050505050565b005b3480156101a057600080fd5b5061006a610257565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102355780601f1061020a57610100808354040283529160200191610235565b820191906000526020600020905b81548152906001019060200180831161021857829003601f168201915b505050505090505b90565b80516102539060009060208401906102e5565b5050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102dd5780601f106102b2576101008083540402835291602001916102dd565b820191906000526020600020905b8154815290600101906020018083116102c057829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061032657805160ff1916838001178555610353565b82800160010185558215610353579182015b82811115610353578251825591602001919060010190610338565b5061035f929150610363565b5090565b61023d91905b8082111561035f576000815560010161036956fea165627a7a72305820f5fa4aa254d20e5542aa8f942cf83125f051b87e34d6f881e8ea50b98ed5a71e0029`

// DeployStoringDetails deploys a new Ethereum contract, binding an instance of StoringDetails to it.
func DeployStoringDetails(auth *bind.TransactOpts, backend bind.ContractBackend, initialJWT []byte) (common.Address, *types.Transaction, *StoringDetails, error) {
	parsed, err := abi.JSON(strings.NewReader(StoringDetailsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoringDetailsBin), backend, initialJWT)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StoringDetails{StoringDetailsCaller: StoringDetailsCaller{contract: contract}, StoringDetailsTransactor: StoringDetailsTransactor{contract: contract}, StoringDetailsFilterer: StoringDetailsFilterer{contract: contract}}, nil
}

// StoringDetails is an auto generated Go binding around an Ethereum contract.
type StoringDetails struct {
	StoringDetailsCaller     // Read-only binding to the contract
	StoringDetailsTransactor // Write-only binding to the contract
	StoringDetailsFilterer   // Log filterer for contract events
}

// StoringDetailsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoringDetailsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoringDetailsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoringDetailsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoringDetailsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoringDetailsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoringDetailsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoringDetailsSession struct {
	Contract     *StoringDetails   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoringDetailsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoringDetailsCallerSession struct {
	Contract *StoringDetailsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StoringDetailsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoringDetailsTransactorSession struct {
	Contract     *StoringDetailsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StoringDetailsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoringDetailsRaw struct {
	Contract *StoringDetails // Generic contract binding to access the raw methods on
}

// StoringDetailsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoringDetailsCallerRaw struct {
	Contract *StoringDetailsCaller // Generic read-only contract binding to access the raw methods on
}

// StoringDetailsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoringDetailsTransactorRaw struct {
	Contract *StoringDetailsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoringDetails creates a new instance of StoringDetails, bound to a specific deployed contract.
func NewStoringDetails(address common.Address, backend bind.ContractBackend) (*StoringDetails, error) {
	contract, err := bindStoringDetails(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoringDetails{StoringDetailsCaller: StoringDetailsCaller{contract: contract}, StoringDetailsTransactor: StoringDetailsTransactor{contract: contract}, StoringDetailsFilterer: StoringDetailsFilterer{contract: contract}}, nil
}

// NewStoringDetailsCaller creates a new read-only instance of StoringDetails, bound to a specific deployed contract.
func NewStoringDetailsCaller(address common.Address, caller bind.ContractCaller) (*StoringDetailsCaller, error) {
	contract, err := bindStoringDetails(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoringDetailsCaller{contract: contract}, nil
}

// NewStoringDetailsTransactor creates a new write-only instance of StoringDetails, bound to a specific deployed contract.
func NewStoringDetailsTransactor(address common.Address, transactor bind.ContractTransactor) (*StoringDetailsTransactor, error) {
	contract, err := bindStoringDetails(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoringDetailsTransactor{contract: contract}, nil
}

// NewStoringDetailsFilterer creates a new log filterer instance of StoringDetails, bound to a specific deployed contract.
func NewStoringDetailsFilterer(address common.Address, filterer bind.ContractFilterer) (*StoringDetailsFilterer, error) {
	contract, err := bindStoringDetails(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoringDetailsFilterer{contract: contract}, nil
}

// bindStoringDetails binds a generic wrapper to an already deployed contract.
func bindStoringDetails(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoringDetailsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoringDetails *StoringDetailsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StoringDetails.Contract.StoringDetailsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoringDetails *StoringDetailsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoringDetails.Contract.StoringDetailsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoringDetails *StoringDetailsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoringDetails.Contract.StoringDetailsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoringDetails *StoringDetailsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StoringDetails.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoringDetails *StoringDetailsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoringDetails.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoringDetails *StoringDetailsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoringDetails.Contract.contract.Transact(opts, method, params...)
}

// JWT is a free data retrieval call binding the contract method 0x74c81a9c.
//
// Solidity: function JWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsCaller) JWT(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _StoringDetails.contract.Call(opts, out, "JWT")
	return *ret0, err
}

// JWT is a free data retrieval call binding the contract method 0x74c81a9c.
//
// Solidity: function JWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsSession) JWT() ([]byte, error) {
	return _StoringDetails.Contract.JWT(&_StoringDetails.CallOpts)
}

// JWT is a free data retrieval call binding the contract method 0x74c81a9c.
//
// Solidity: function JWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsCallerSession) JWT() ([]byte, error) {
	return _StoringDetails.Contract.JWT(&_StoringDetails.CallOpts)
}

// GetJWT is a free data retrieval call binding the contract method 0x2cd38f8c.
//
// Solidity: function getJWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsCaller) GetJWT(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _StoringDetails.contract.Call(opts, out, "getJWT")
	return *ret0, err
}

// GetJWT is a free data retrieval call binding the contract method 0x2cd38f8c.
//
// Solidity: function getJWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsSession) GetJWT() ([]byte, error) {
	return _StoringDetails.Contract.GetJWT(&_StoringDetails.CallOpts)
}

// GetJWT is a free data retrieval call binding the contract method 0x2cd38f8c.
//
// Solidity: function getJWT() constant returns(bytes)
func (_StoringDetails *StoringDetailsCallerSession) GetJWT() ([]byte, error) {
	return _StoringDetails.Contract.GetJWT(&_StoringDetails.CallOpts)
}

// SetJWT is a paid mutator transaction binding the contract method 0x3b04e3a3.
//
// Solidity: function setJWT(newJWT bytes) returns()
func (_StoringDetails *StoringDetailsTransactor) SetJWT(opts *bind.TransactOpts, newJWT []byte) (*types.Transaction, error) {
	return _StoringDetails.contract.Transact(opts, "setJWT", newJWT)
}

// SetJWT is a paid mutator transaction binding the contract method 0x3b04e3a3.
//
// Solidity: function setJWT(newJWT bytes) returns()
func (_StoringDetails *StoringDetailsSession) SetJWT(newJWT []byte) (*types.Transaction, error) {
	return _StoringDetails.Contract.SetJWT(&_StoringDetails.TransactOpts, newJWT)
}

// SetJWT is a paid mutator transaction binding the contract method 0x3b04e3a3.
//
// Solidity: function setJWT(newJWT bytes) returns()
func (_StoringDetails *StoringDetailsTransactorSession) SetJWT(newJWT []byte) (*types.Transaction, error) {
	return _StoringDetails.Contract.SetJWT(&_StoringDetails.TransactOpts, newJWT)
}
