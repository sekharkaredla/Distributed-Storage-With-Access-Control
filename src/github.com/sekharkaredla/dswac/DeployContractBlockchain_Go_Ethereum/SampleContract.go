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

// SampleContractABI is the input ABI used to generate the binding from.
const SampleContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialMessage\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SampleContractBin is the compiled bytecode used for deploying new contracts.
const SampleContractBin = `0x608060405234801561001057600080fd5b506040516104383803806104388339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815164010000000081118282018710171561007857600080fd5b505080519093506100929250600091506020840190610099565b5050610134565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100da57805160ff1916838001178555610107565b82800160010185558215610107579182015b828111156101075782518255916020019190600101906100ec565b50610113929150610117565b5090565b61013191905b80821115610113576000815560010161011d565b90565b6102f5806101436000396000f3fe608060405260043610610045577c01000000000000000000000000000000000000000000000000000000006000350463368b8772811461004a578063e21f37ce146100ff575b600080fd5b34801561005657600080fd5b506100fd6004803603602081101561006d57600080fd5b81019060208101813564010000000081111561008857600080fd5b82018360208201111561009a57600080fd5b803590602001918460018302840111640100000000831117156100bc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610189945050505050565b005b34801561010b57600080fd5b506101146101a0565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014e578181015183820152602001610136565b50505050905090810190601f16801561017b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b805161019c90600090602084019061022e565b5050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102265780601f106101fb57610100808354040283529160200191610226565b820191906000526020600020905b81548152906001019060200180831161020957829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061026f57805160ff191683800117855561029c565b8280016001018555821561029c579182015b8281111561029c578251825591602001919060010190610281565b506102a89291506102ac565b5090565b6102c691905b808211156102a857600081556001016102b2565b9056fea165627a7a72305820575c38ba117a14a199fcfc8e79013befb2e4550d191242f64bb69ca42874cbdd0029`

// DeploySampleContract deploys a new Ethereum contract, binding an instance of SampleContract to it.
func DeploySampleContract(auth *bind.TransactOpts, backend bind.ContractBackend, initialMessage string) (common.Address, *types.Transaction, *SampleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleContractBin), backend, initialMessage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleContract{SampleContractCaller: SampleContractCaller{contract: contract}, SampleContractTransactor: SampleContractTransactor{contract: contract}, SampleContractFilterer: SampleContractFilterer{contract: contract}}, nil
}

// SampleContract is an auto generated Go binding around an Ethereum contract.
type SampleContract struct {
	SampleContractCaller     // Read-only binding to the contract
	SampleContractTransactor // Write-only binding to the contract
	SampleContractFilterer   // Log filterer for contract events
}

// SampleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleContractSession struct {
	Contract     *SampleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleContractCallerSession struct {
	Contract *SampleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SampleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleContractTransactorSession struct {
	Contract     *SampleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SampleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleContractRaw struct {
	Contract *SampleContract // Generic contract binding to access the raw methods on
}

// SampleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleContractCallerRaw struct {
	Contract *SampleContractCaller // Generic read-only contract binding to access the raw methods on
}

// SampleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleContractTransactorRaw struct {
	Contract *SampleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleContract creates a new instance of SampleContract, bound to a specific deployed contract.
func NewSampleContract(address common.Address, backend bind.ContractBackend) (*SampleContract, error) {
	contract, err := bindSampleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleContract{SampleContractCaller: SampleContractCaller{contract: contract}, SampleContractTransactor: SampleContractTransactor{contract: contract}, SampleContractFilterer: SampleContractFilterer{contract: contract}}, nil
}

// NewSampleContractCaller creates a new read-only instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractCaller(address common.Address, caller bind.ContractCaller) (*SampleContractCaller, error) {
	contract, err := bindSampleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleContractCaller{contract: contract}, nil
}

// NewSampleContractTransactor creates a new write-only instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleContractTransactor, error) {
	contract, err := bindSampleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleContractTransactor{contract: contract}, nil
}

// NewSampleContractFilterer creates a new log filterer instance of SampleContract, bound to a specific deployed contract.
func NewSampleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleContractFilterer, error) {
	contract, err := bindSampleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleContractFilterer{contract: contract}, nil
}

// bindSampleContract binds a generic wrapper to an already deployed contract.
func bindSampleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleContract *SampleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleContract.Contract.SampleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleContract *SampleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleContract.Contract.SampleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleContract *SampleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleContract.Contract.SampleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleContract *SampleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleContract *SampleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleContract *SampleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleContract.Contract.contract.Transact(opts, method, params...)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SampleContract *SampleContractCaller) Message(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleContract.contract.Call(opts, out, "message")
	return *ret0, err
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SampleContract *SampleContractSession) Message() (string, error) {
	return _SampleContract.Contract.Message(&_SampleContract.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SampleContract *SampleContractCallerSession) Message() (string, error) {
	return _SampleContract.Contract.Message(&_SampleContract.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_SampleContract *SampleContractTransactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _SampleContract.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_SampleContract *SampleContractSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _SampleContract.Contract.SetMessage(&_SampleContract.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(newMessage string) returns()
func (_SampleContract *SampleContractTransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _SampleContract.Contract.SetMessage(&_SampleContract.TransactOpts, newMessage)
}
