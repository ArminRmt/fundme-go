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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FundMe__FundingClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__NoFundsToRefund\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__RefundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__RefundPeriodOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundMe__WithdrawFailed\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FUNDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_USD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"getAddressToAmountFunded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getFunder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundingDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"getRefundAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161154438038061154483398181016040528101906100329190610118565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260a0818152505050610145565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e5826100ba565b9050919050565b6100f5816100da565b811461010057600080fd5b50565b600081519050610112816100ec565b92915050565b60006020828403121561012e5761012d6100b5565b5b600061013c84828501610103565b91505092915050565b60805160a0516113a961019b600039600081816102bb0152818161056a0152818161079f01528181610825015281816108590152610c530152600081816104e1015281816106d50152610bbc01526113a96000f3fe6080604052600436106100ab5760003560e01c8063893d20e811610064578063893d20e8146101c25780639e87a5cd146101ed578063b60d428814610218578063d7b4750c14610222578063e54a51a61461025f578063f90c3f271461028a576100ba565b80630343fb25146100c45780633ccfd60b146101015780634772be5a1461011857806356b6491a14610155578063590e1ae3146101805780636b69a59214610197576100ba565b366100ba576100b86102b5565b005b6100c26102b5565b005b3480156100d057600080fd5b506100eb60048036038101906100e69190610e53565b610496565b6040516100f89190610e99565b60405180910390f35b34801561010d57600080fd5b506101166104df565b005b34801561012457600080fd5b5061013f600480360381019061013a9190610e53565b610797565b60405161014c9190610ecf565b60405180910390f35b34801561016157600080fd5b5061016a61081d565b6040516101779190610e99565b60405180910390f35b34801561018c57600080fd5b50610195610853565b005b3480156101a357600080fd5b506101ac610bac565b6040516101b99190610e99565b60405180910390f35b3480156101ce57600080fd5b506101d7610bb8565b6040516101e49190610ef9565b60405180910390f35b3480156101f957600080fd5b50610202610be0565b60405161020f9190610f73565b60405180910390f35b6102206102b5565b005b34801561022e57600080fd5b5061024960048036038101906102449190610fba565b610c0a565b6040516102569190610ef9565b60405180910390f35b34801561026b57600080fd5b50610274610c51565b6040516102819190610e99565b60405180910390f35b34801561029657600080fd5b5061029f610c75565b6040516102ac9190610e99565b60405180910390f35b62093a807f00000000000000000000000000000000000000000000000000000000000000006102e49190611016565b42111561031d576040517f3b5de1d900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b674563918244f4000061035b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1634610c7c90919063ffffffff16565b1015610393576040517fb514e4ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361043e576000339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461048d9190611016565b92505081905550565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610564576040517f579610db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a807f00000000000000000000000000000000000000000000000000000000000000006105939190611016565b42116105cb576040517f71a9a10b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b6000805490508110156106705760008082815481106105f0576105ef61104a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505080806001019150506105ce565b50600067ffffffffffffffff81111561068c5761068b611079565b5b6040519080825280602002602001820160405280156106ba5781602001602082028036833780820191505090505b50600090805190602001906106d0929190610d49565b5060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1647604051610717906110d9565b60006040518083038185875af1925050503d8060008114610754576040519150601f19603f3d011682016040523d82523d6000602084013e610759565b606091505b5050905080610794576040517ff02e6d3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b600062093a807f00000000000000000000000000000000000000000000000000000000000000006107c89190611016565b421115801561081657506000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054115b9050919050565b600062093a807f000000000000000000000000000000000000000000000000000000000000000061084e9190611016565b905090565b62093a807f00000000000000000000000000000000000000000000000000000000000000006108829190611016565b4211156108bb576040517f3b5de1d900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008103610939576040517fd477458000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060005b600080549050811015610b04573373ffffffffffffffffffffffffffffffffffffffff16600082815481106109b9576109b861104a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610af75760006001600080549050610a1391906110ee565b81548110610a2457610a2361104a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008281548110610a6357610a6261104a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000805480610abd57610abc611122565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055610b04565b8080600101915050610981565b5060003373ffffffffffffffffffffffffffffffffffffffff1682604051610b2b906110d9565b60006040518083038185875af1925050503d8060008114610b68576040519150601f19603f3d011682016040523d82523d6000602084013e610b6d565b606091505b5050905080610ba8576040517f5ad8b8e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b674563918244f4000081565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000808281548110610c1f57610c1e61104a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b62093a8081565b600080610c8883610cb8565b90506000670de0b6b3a76400008583610ca19190611151565b610cab91906111c2565b9050809250505092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610d06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2a9190611280565b5050509150506402540be40081610d4191906112fb565b915050919050565b828054828255906000526020600020908101928215610dc2579160200282015b82811115610dc15782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610d69565b5b509050610dcf9190610dd3565b5090565b5b80821115610dec576000816000905550600101610dd4565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e2082610df5565b9050919050565b610e3081610e15565b8114610e3b57600080fd5b50565b600081359050610e4d81610e27565b92915050565b600060208284031215610e6957610e68610df0565b5b6000610e7784828501610e3e565b91505092915050565b6000819050919050565b610e9381610e80565b82525050565b6000602082019050610eae6000830184610e8a565b92915050565b60008115159050919050565b610ec981610eb4565b82525050565b6000602082019050610ee46000830184610ec0565b92915050565b610ef381610e15565b82525050565b6000602082019050610f0e6000830184610eea565b92915050565b6000819050919050565b6000610f39610f34610f2f84610df5565b610f14565b610df5565b9050919050565b6000610f4b82610f1e565b9050919050565b6000610f5d82610f40565b9050919050565b610f6d81610f52565b82525050565b6000602082019050610f886000830184610f64565b92915050565b610f9781610e80565b8114610fa257600080fd5b50565b600081359050610fb481610f8e565b92915050565b600060208284031215610fd057610fcf610df0565b5b6000610fde84828501610fa5565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061102182610e80565b915061102c83610e80565b925082820190508082111561104457611043610fe7565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081905092915050565b50565b60006110c36000836110a8565b91506110ce826110b3565b600082019050919050565b60006110e4826110b6565b9150819050919050565b60006110f982610e80565b915061110483610e80565b925082820390508181111561111c5761111b610fe7565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600061115c82610e80565b915061116783610e80565b925082820261117581610e80565b9150828204841483151761118c5761118b610fe7565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006111cd82610e80565b91506111d883610e80565b9250826111e8576111e7611193565b5b828204905092915050565b600069ffffffffffffffffffff82169050919050565b611212816111f3565b811461121d57600080fd5b50565b60008151905061122f81611209565b92915050565b6000819050919050565b61124881611235565b811461125357600080fd5b50565b6000815190506112658161123f565b92915050565b60008151905061127a81610f8e565b92915050565b600080600080600060a0868803121561129c5761129b610df0565b5b60006112aa88828901611220565b95505060206112bb88828901611256565b94505060406112cc8882890161126b565b93505060606112dd8882890161126b565b92505060806112ee88828901611220565b9150509295509295909350565b600061130682611235565b915061131183611235565b925082820261131f81611235565b91507f8000000000000000000000000000000000000000000000000000000000000000841460008412161561135757611356610fe7565b5b828205841483151761136c5761136b610fe7565b5b509291505056fea264697066735822122047b7508b901e5dc84feac0b7fb044efb26005de2af643eab89d84dfa77858b2164736f6c634300081e0033",
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

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_FundMe *FundMeCaller) FUNDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "FUNDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_FundMe *FundMeSession) FUNDINGPERIOD() (*big.Int, error) {
	return _FundMe.Contract.FUNDINGPERIOD(&_FundMe.CallOpts)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_FundMe *FundMeCallerSession) FUNDINGPERIOD() (*big.Int, error) {
	return _FundMe.Contract.FUNDINGPERIOD(&_FundMe.CallOpts)
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

// GetFundingDeadline is a free data retrieval call binding the contract method 0x56b6491a.
//
// Solidity: function getFundingDeadline() view returns(uint256)
func (_FundMe *FundMeCaller) GetFundingDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getFundingDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingDeadline is a free data retrieval call binding the contract method 0x56b6491a.
//
// Solidity: function getFundingDeadline() view returns(uint256)
func (_FundMe *FundMeSession) GetFundingDeadline() (*big.Int, error) {
	return _FundMe.Contract.GetFundingDeadline(&_FundMe.CallOpts)
}

// GetFundingDeadline is a free data retrieval call binding the contract method 0x56b6491a.
//
// Solidity: function getFundingDeadline() view returns(uint256)
func (_FundMe *FundMeCallerSession) GetFundingDeadline() (*big.Int, error) {
	return _FundMe.Contract.GetFundingDeadline(&_FundMe.CallOpts)
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

// GetRefundAvailable is a free data retrieval call binding the contract method 0x4772be5a.
//
// Solidity: function getRefundAvailable(address funder) view returns(bool)
func (_FundMe *FundMeCaller) GetRefundAvailable(opts *bind.CallOpts, funder common.Address) (bool, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "getRefundAvailable", funder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRefundAvailable is a free data retrieval call binding the contract method 0x4772be5a.
//
// Solidity: function getRefundAvailable(address funder) view returns(bool)
func (_FundMe *FundMeSession) GetRefundAvailable(funder common.Address) (bool, error) {
	return _FundMe.Contract.GetRefundAvailable(&_FundMe.CallOpts, funder)
}

// GetRefundAvailable is a free data retrieval call binding the contract method 0x4772be5a.
//
// Solidity: function getRefundAvailable(address funder) view returns(bool)
func (_FundMe *FundMeCallerSession) GetRefundAvailable(funder common.Address) (bool, error) {
	return _FundMe.Contract.GetRefundAvailable(&_FundMe.CallOpts, funder)
}

// SStartTime is a free data retrieval call binding the contract method 0xe54a51a6.
//
// Solidity: function s_startTime() view returns(uint256)
func (_FundMe *FundMeCaller) SStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FundMe.contract.Call(opts, &out, "s_startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SStartTime is a free data retrieval call binding the contract method 0xe54a51a6.
//
// Solidity: function s_startTime() view returns(uint256)
func (_FundMe *FundMeSession) SStartTime() (*big.Int, error) {
	return _FundMe.Contract.SStartTime(&_FundMe.CallOpts)
}

// SStartTime is a free data retrieval call binding the contract method 0xe54a51a6.
//
// Solidity: function s_startTime() view returns(uint256)
func (_FundMe *FundMeCallerSession) SStartTime() (*big.Int, error) {
	return _FundMe.Contract.SStartTime(&_FundMe.CallOpts)
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

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_FundMe *FundMeTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FundMe.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_FundMe *FundMeSession) Refund() (*types.Transaction, error) {
	return _FundMe.Contract.Refund(&_FundMe.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_FundMe *FundMeTransactorSession) Refund() (*types.Transaction, error) {
	return _FundMe.Contract.Refund(&_FundMe.TransactOpts)
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
