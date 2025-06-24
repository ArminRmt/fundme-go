// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fundme

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

// FundMeMetaData contains all meta data concerning the FundMe contract.
var FundMeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FundMe__NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__WithdrawFailed\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"MINIMUM_USD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"getAddressToAmountFunded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getFunder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610e13380380610e1383398181016040528101906100329190610110565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061013d565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100dd826100b2565b9050919050565b6100ed816100d2565b81146100f857600080fd5b50565b60008151905061010a816100e4565b92915050565b600060208284031215610126576101256100ad565b5b6000610134848285016100fb565b91505092915050565b608051610cad61016660003960008181610325015281816104b201526105840152610cad6000f3fe6080604052600436106100745760003560e01c8063893d20e81161004e578063893d20e81461010c5780639e87a5cd14610137578063b60d428814610162578063d7b4750c1461016c57610083565b80630343fb251461008d5780633ccfd60b146100ca5780636b69a592146100e157610083565b36610083576100816101a9565b005b61008b6101a9565b005b34801561009957600080fd5b506100b460048036038101906100af91906107f0565b6102da565b6040516100c19190610836565b60405180910390f35b3480156100d657600080fd5b506100df610323565b005b3480156100ed57600080fd5b506100f6610574565b6040516101039190610836565b60405180910390f35b34801561011857600080fd5b50610121610580565b60405161012e9190610860565b60405180910390f35b34801561014357600080fd5b5061014c6105a8565b60405161015991906108da565b60405180910390f35b61016a6101a9565b005b34801561017857600080fd5b50610193600480360381019061018e9190610921565b6105d2565b6040516101a09190610860565b60405180910390f35b674563918244f400006101e7600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff163461061990919063ffffffff16565b101561021f576040517fb514e4ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461026e919061097d565b925050819055506000339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a8576040517f579610db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b60008054905081101561044d5760008082815481106103cd576103cc6109b1565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505080806001019150506103ab565b50600067ffffffffffffffff811115610469576104686109e0565b5b6040519080825280602002602001820160405280156104975781602001602082028036833780820191505090505b50600090805190602001906104ad9291906106e6565b5060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16476040516104f490610a40565b60006040518083038185875af1925050503d8060008114610531576040519150601f19603f3d011682016040523d82523d6000602084013e610536565b606091505b5050905080610571576040517ff02e6d3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b674563918244f4000081565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008082815481106105e7576105e66109b1565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60008061062583610655565b90506000670de0b6b3a7640000858361063e9190610a55565b6106489190610ac6565b9050809250505092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa1580156106a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c79190610b84565b5050509150506402540be400816106de9190610bff565b915050919050565b82805482825590600052602060002090810192821561075f579160200282015b8281111561075e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610706565b5b50905061076c9190610770565b5090565b5b80821115610789576000816000905550600101610771565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107bd82610792565b9050919050565b6107cd816107b2565b81146107d857600080fd5b50565b6000813590506107ea816107c4565b92915050565b6000602082840312156108065761080561078d565b5b6000610814848285016107db565b91505092915050565b6000819050919050565b6108308161081d565b82525050565b600060208201905061084b6000830184610827565b92915050565b61085a816107b2565b82525050565b60006020820190506108756000830184610851565b92915050565b6000819050919050565b60006108a061089b61089684610792565b61087b565b610792565b9050919050565b60006108b282610885565b9050919050565b60006108c4826108a7565b9050919050565b6108d4816108b9565b82525050565b60006020820190506108ef60008301846108cb565b92915050565b6108fe8161081d565b811461090957600080fd5b50565b60008135905061091b816108f5565b92915050565b6000602082840312156109375761093661078d565b5b60006109458482850161090c565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006109888261081d565b91506109938361081d565b92508282019050808211156109ab576109aa61094e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081905092915050565b50565b6000610a2a600083610a0f565b9150610a3582610a1a565b600082019050919050565b6000610a4b82610a1d565b9150819050919050565b6000610a608261081d565b9150610a6b8361081d565b9250828202610a798161081d565b91508282048414831517610a9057610a8f61094e565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610ad18261081d565b9150610adc8361081d565b925082610aec57610aeb610a97565b5b828204905092915050565b600069ffffffffffffffffffff82169050919050565b610b1681610af7565b8114610b2157600080fd5b50565b600081519050610b3381610b0d565b92915050565b6000819050919050565b610b4c81610b39565b8114610b5757600080fd5b50565b600081519050610b6981610b43565b92915050565b600081519050610b7e816108f5565b92915050565b600080600080600060a08688031215610ba057610b9f61078d565b5b6000610bae88828901610b24565b9550506020610bbf88828901610b5a565b9450506040610bd088828901610b6f565b9350506060610be188828901610b6f565b9250506080610bf288828901610b24565b9150509295509295909350565b6000610c0a82610b39565b9150610c1583610b39565b9250828202610c2381610b39565b91507f80000000000000000000000000000000000000000000000000000000000000008414600084121615610c5b57610c5a61094e565b5b8282058414831517610c7057610c6f61094e565b5b509291505056fea26469706673582212205da27a6608a2407a43ac0fe4dca3d86b7a28ecfdcb189c25cfcf01ef2877da4564736f6c634300081e0033",
}

// FundMeABI is the input ABI used to generate the binding from.
// Deprecated: Use FundMeMetaData.ABI instead.
var FundMeABI = FundMeMetaData.ABI

// FundMeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FundMeMetaData.Bin instead.
var FundMeBin = FundMeMetaData.Bin

// DeployFundMe deploys a new Ethereum contract, binding an instance of FundMe to it.
func DeployFundMe(auth *bind.TransactOpts, backend bind.ContractBackend, _priceFeed common.Address) (common.Address, *types.Transaction, *FundMe, error) {
	parsed, err := FundMeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FundMeBin), backend, _priceFeed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FundMe{FundMeCaller: FundMeCaller{contract: contract}, FundMeTransactor: FundMeTransactor{contract: contract}, FundMeFilterer: FundMeFilterer{contract: contract}}, nil
}

// FundMe is an auto generated Go binding around an Ethereum contract.
type FundMe struct {
	FundMeCaller     // Read-only binding to the contract
	FundMeTransactor // Write-only binding to the contract
	FundMeFilterer   // Log filterer for contract events
}

// FundMeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FundMeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundMeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FundMeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundMeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FundMeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundMeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FundMeSession struct {
	Contract     *FundMe           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundMeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FundMeCallerSession struct {
	Contract *FundMeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FundMeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FundMeTransactorSession struct {
	Contract     *FundMeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundMeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FundMeRaw struct {
	Contract *FundMe // Generic contract binding to access the raw methods on
}

// FundMeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FundMeCallerRaw struct {
	Contract *FundMeCaller // Generic read-only contract binding to access the raw methods on
}

// FundMeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FundMeTransactorRaw struct {
	Contract *FundMeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFundMe creates a new instance of FundMe, bound to a specific deployed contract.
func NewFundMe(address common.Address, backend bind.ContractBackend) (*FundMe, error) {
	contract, err := bindFundMe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FundMe{FundMeCaller: FundMeCaller{contract: contract}, FundMeTransactor: FundMeTransactor{contract: contract}, FundMeFilterer: FundMeFilterer{contract: contract}}, nil
}

// NewFundMeCaller creates a new read-only instance of FundMe, bound to a specific deployed contract.
func NewFundMeCaller(address common.Address, caller bind.ContractCaller) (*FundMeCaller, error) {
	contract, err := bindFundMe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FundMeCaller{contract: contract}, nil
}

// NewFundMeTransactor creates a new write-only instance of FundMe, bound to a specific deployed contract.
func NewFundMeTransactor(address common.Address, transactor bind.ContractTransactor) (*FundMeTransactor, error) {
	contract, err := bindFundMe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FundMeTransactor{contract: contract}, nil
}

// NewFundMeFilterer creates a new log filterer instance of FundMe, bound to a specific deployed contract.
func NewFundMeFilterer(address common.Address, filterer bind.ContractFilterer) (*FundMeFilterer, error) {
	contract, err := bindFundMe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FundMeFilterer{contract: contract}, nil
}

// bindFundMe binds a generic wrapper to an already deployed contract.
func bindFundMe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FundMeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FundMe *FundMeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FundMe.Contract.FundMeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FundMe *FundMeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.Contract.FundMeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FundMe *FundMeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FundMe.Contract.FundMeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FundMe *FundMeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FundMe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FundMe *FundMeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FundMe *FundMeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FundMe.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMUSD is a free data retrieval call binding the contract method 0x6b69a592.
//
// Solidity: function MINIMUM_USD() view returns(uint256)
func (_FundMe *FundMeCaller) MINIMUMUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "MINIMUM_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMUSD is a free data retrieval call binding the contract method 0x6b69a592.
//
// Solidity: function MINIMUM_USD() view returns(uint256)
func (_FundMe *FundMeSession) MINIMUMUSD() (*big.Int, error) {
	return _FundMe.Contract.MINIMUMUSD(&_FundMe.CallOpts)
}

// MINIMUMUSD is a free data retrieval call binding the contract method 0x6b69a592.
//
// Solidity: function MINIMUM_USD() view returns(uint256)
func (_FundMe *FundMeCallerSession) MINIMUMUSD() (*big.Int, error) {
	return _FundMe.Contract.MINIMUMUSD(&_FundMe.CallOpts)
}

// GetAddressToAmountFunded is a free data retrieval call binding the contract method 0x0343fb25.
//
// Solidity: function getAddressToAmountFunded(address funder) view returns(uint256)
func (_FundMe *FundMeCaller) GetAddressToAmountFunded(opts *bind.CallOpts, funder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getAddressToAmountFunded", funder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressToAmountFunded is a free data retrieval call binding the contract method 0x0343fb25.
//
// Solidity: function getAddressToAmountFunded(address funder) view returns(uint256)
func (_FundMe *FundMeSession) GetAddressToAmountFunded(funder common.Address) (*big.Int, error) {
	return _FundMe.Contract.GetAddressToAmountFunded(&_FundMe.CallOpts, funder)
}

// GetAddressToAmountFunded is a free data retrieval call binding the contract method 0x0343fb25.
//
// Solidity: function getAddressToAmountFunded(address funder) view returns(uint256)
func (_FundMe *FundMeCallerSession) GetAddressToAmountFunded(funder common.Address) (*big.Int, error) {
	return _FundMe.Contract.GetAddressToAmountFunded(&_FundMe.CallOpts, funder)
}

// GetFunder is a free data retrieval call binding the contract method 0xd7b4750c.
//
// Solidity: function getFunder(uint256 index) view returns(address)
func (_FundMe *FundMeCaller) GetFunder(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getFunder", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFunder is a free data retrieval call binding the contract method 0xd7b4750c.
//
// Solidity: function getFunder(uint256 index) view returns(address)
func (_FundMe *FundMeSession) GetFunder(index *big.Int) (common.Address, error) {
	return _FundMe.Contract.GetFunder(&_FundMe.CallOpts, index)
}

// GetFunder is a free data retrieval call binding the contract method 0xd7b4750c.
//
// Solidity: function getFunder(uint256 index) view returns(address)
func (_FundMe *FundMeCallerSession) GetFunder(index *big.Int) (common.Address, error) {
	return _FundMe.Contract.GetFunder(&_FundMe.CallOpts, index)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_FundMe *FundMeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_FundMe *FundMeSession) GetOwner() (common.Address, error) {
	return _FundMe.Contract.GetOwner(&_FundMe.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_FundMe *FundMeCallerSession) GetOwner() (common.Address, error) {
	return _FundMe.Contract.GetOwner(&_FundMe.CallOpts)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_FundMe *FundMeCaller) GetPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_FundMe *FundMeSession) GetPriceFeed() (common.Address, error) {
	return _FundMe.Contract.GetPriceFeed(&_FundMe.CallOpts)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_FundMe *FundMeCallerSession) GetPriceFeed() (common.Address, error) {
	return _FundMe.Contract.GetPriceFeed(&_FundMe.CallOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_FundMe *FundMeTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_FundMe *FundMeSession) Fund() (*types.Transaction, error) {
	return _FundMe.Contract.Fund(&_FundMe.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_FundMe *FundMeTransactorSession) Fund() (*types.Transaction, error) {
	return _FundMe.Contract.Fund(&_FundMe.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FundMe *FundMeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FundMe *FundMeSession) Withdraw() (*types.Transaction, error) {
	return _FundMe.Contract.Withdraw(&_FundMe.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FundMe *FundMeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FundMe.Contract.Withdraw(&_FundMe.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FundMe *FundMeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FundMe.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FundMe *FundMeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FundMe.Contract.Fallback(&_FundMe.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FundMe *FundMeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FundMe.Contract.Fallback(&_FundMe.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FundMe *FundMeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FundMe *FundMeSession) Receive() (*types.Transaction, error) {
	return _FundMe.Contract.Receive(&_FundMe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FundMe *FundMeTransactorSession) Receive() (*types.Transaction, error) {
	return _FundMe.Contract.Receive(&_FundMe.TransactOpts)
}
