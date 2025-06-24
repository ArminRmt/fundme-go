// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockaggregator

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

// MockV3AggregatorMetaData contains all meta data concerning the MockV3Aggregator contract.
var MockV3AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"_initialAnswer\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"}],\"name\":\"updateAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161076e38038061076e8339818101604052810190610032919061014d565b816000806101000a81548160ff021916908360ff16021790555061005b8161006260201b60201c565b505061020e565b806001819055504260028190555060036000815480929190610083906101c6565b919050555080600460006003548152602001908152602001600020819055504260056000600354815260200190815260200160002081905550426006600060035481526020019081526020016000208190555050565b600080fd5b600060ff82169050919050565b6100f4816100de565b81146100ff57600080fd5b50565b600081519050610111816100eb565b92915050565b6000819050919050565b61012a81610117565b811461013557600080fd5b50565b60008151905061014781610121565b92915050565b60008060408385031215610164576101636100d9565b5b600061017285828601610102565b925050602061018385828601610138565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006101d1826101bc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036102035761020261018d565b5b600182019050919050565b6105518061021d6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638205bf6a116100665780638205bf6a14610110578063a87a20ce1461012e578063b5ab58dc1461014a578063b633620c1461017a578063feaf968c146101aa57610093565b8063313ce5671461009857806350d25bcd146100b657806354fd4d50146100d4578063668a0f02146100f2575b600080fd5b6100a06101cc565b6040516100ad91906102f2565b60405180910390f35b6100be6101dd565b6040516100cb9190610326565b60405180910390f35b6100dc6101e3565b6040516100e9919061035a565b60405180910390f35b6100fa6101e8565b604051610107919061035a565b60405180910390f35b6101186101ee565b604051610125919061035a565b60405180910390f35b610148600480360381019061014391906103a6565b6101f4565b005b610164600480360381019061015f91906103ff565b61026b565b6040516101719190610326565b60405180910390f35b610194600480360381019061018f91906103ff565b610283565b6040516101a1919061035a565b60405180910390f35b6101b261029b565b6040516101c3959493929190610451565b60405180910390f35b60008054906101000a900460ff1681565b60015481565b600081565b60035481565b60025481565b806001819055504260028190555060036000815480929190610215906104d3565b919050555080600460006003548152602001908152602001600020819055504260056000600354815260200190815260200160002081905550426006600060035481526020019081526020016000208190555050565b60046020528060005260406000206000915090505481565b60056020528060005260406000206000915090505481565b600080600080600060035460015460066000600354815260200190815260200160002054600254600354945094509450945094509091929394565b600060ff82169050919050565b6102ec816102d6565b82525050565b600060208201905061030760008301846102e3565b92915050565b6000819050919050565b6103208161030d565b82525050565b600060208201905061033b6000830184610317565b92915050565b6000819050919050565b61035481610341565b82525050565b600060208201905061036f600083018461034b565b92915050565b600080fd5b6103838161030d565b811461038e57600080fd5b50565b6000813590506103a08161037a565b92915050565b6000602082840312156103bc576103bb610375565b5b60006103ca84828501610391565b91505092915050565b6103dc81610341565b81146103e757600080fd5b50565b6000813590506103f9816103d3565b92915050565b60006020828403121561041557610414610375565b5b6000610423848285016103ea565b91505092915050565b600069ffffffffffffffffffff82169050919050565b61044b8161042c565b82525050565b600060a0820190506104666000830188610442565b6104736020830187610317565b610480604083018661034b565b61048d606083018561034b565b61049a6080830184610442565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006104de82610341565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036105105761050f6104a4565b5b60018201905091905056fea2646970667358221220c4ee6705470aaf88a0a562c2000d7cf8301a7f638366122ef9a3bf8e3405584c64736f6c634300081e0033",
}

// MockV3AggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use MockV3AggregatorMetaData.ABI instead.
var MockV3AggregatorABI = MockV3AggregatorMetaData.ABI

// MockV3AggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockV3AggregatorMetaData.Bin instead.
var MockV3AggregatorBin = MockV3AggregatorMetaData.Bin

// DeployMockV3Aggregator deploys a new Ethereum contract, binding an instance of MockV3Aggregator to it.
func DeployMockV3Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _decimals uint8, _initialAnswer *big.Int) (common.Address, *types.Transaction, *MockV3Aggregator, error) {
	parsed, err := MockV3AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockV3AggregatorBin), backend, _decimals, _initialAnswer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockV3Aggregator{MockV3AggregatorCaller: MockV3AggregatorCaller{contract: contract}, MockV3AggregatorTransactor: MockV3AggregatorTransactor{contract: contract}, MockV3AggregatorFilterer: MockV3AggregatorFilterer{contract: contract}}, nil
}

// MockV3Aggregator is an auto generated Go binding around an Ethereum contract.
type MockV3Aggregator struct {
	MockV3AggregatorCaller     // Read-only binding to the contract
	MockV3AggregatorTransactor // Write-only binding to the contract
	MockV3AggregatorFilterer   // Log filterer for contract events
}

// MockV3AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockV3AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockV3AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockV3AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockV3AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockV3AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockV3AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockV3AggregatorSession struct {
	Contract     *MockV3Aggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockV3AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockV3AggregatorCallerSession struct {
	Contract *MockV3AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MockV3AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockV3AggregatorTransactorSession struct {
	Contract     *MockV3AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MockV3AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockV3AggregatorRaw struct {
	Contract *MockV3Aggregator // Generic contract binding to access the raw methods on
}

// MockV3AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockV3AggregatorCallerRaw struct {
	Contract *MockV3AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// MockV3AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockV3AggregatorTransactorRaw struct {
	Contract *MockV3AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockV3Aggregator creates a new instance of MockV3Aggregator, bound to a specific deployed contract.
func NewMockV3Aggregator(address common.Address, backend bind.ContractBackend) (*MockV3Aggregator, error) {
	contract, err := bindMockV3Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockV3Aggregator{MockV3AggregatorCaller: MockV3AggregatorCaller{contract: contract}, MockV3AggregatorTransactor: MockV3AggregatorTransactor{contract: contract}, MockV3AggregatorFilterer: MockV3AggregatorFilterer{contract: contract}}, nil
}

// NewMockV3AggregatorCaller creates a new read-only instance of MockV3Aggregator, bound to a specific deployed contract.
func NewMockV3AggregatorCaller(address common.Address, caller bind.ContractCaller) (*MockV3AggregatorCaller, error) {
	contract, err := bindMockV3Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorCaller{contract: contract}, nil
}

// NewMockV3AggregatorTransactor creates a new write-only instance of MockV3Aggregator, bound to a specific deployed contract.
func NewMockV3AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MockV3AggregatorTransactor, error) {
	contract, err := bindMockV3Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorTransactor{contract: contract}, nil
}

// NewMockV3AggregatorFilterer creates a new log filterer instance of MockV3Aggregator, bound to a specific deployed contract.
func NewMockV3AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MockV3AggregatorFilterer, error) {
	contract, err := bindMockV3Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorFilterer{contract: contract}, nil
}

// bindMockV3Aggregator binds a generic wrapper to an already deployed contract.
func bindMockV3Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockV3AggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockV3Aggregator *MockV3AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockV3Aggregator.Contract.MockV3AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockV3Aggregator *MockV3AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.MockV3AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockV3Aggregator *MockV3AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.MockV3AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockV3Aggregator *MockV3AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockV3Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockV3Aggregator *MockV3AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockV3Aggregator *MockV3AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockV3Aggregator *MockV3AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockV3Aggregator *MockV3AggregatorSession) Decimals() (uint8, error) {
	return _MockV3Aggregator.Contract.Decimals(&_MockV3Aggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) Decimals() (uint8, error) {
	return _MockV3Aggregator.Contract.Decimals(&_MockV3Aggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 ) view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorCaller) GetAnswer(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "getAnswer", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 ) view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorSession) GetAnswer(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetAnswer(&_MockV3Aggregator.CallOpts, arg0)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 ) view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) GetAnswer(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetAnswer(&_MockV3Aggregator.CallOpts, arg0)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 ) view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCaller) GetTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "getTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 ) view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorSession) GetTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetTimestamp(&_MockV3Aggregator.CallOpts, arg0)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 ) view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) GetTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetTimestamp(&_MockV3Aggregator.CallOpts, arg0)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestAnswer(&_MockV3Aggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestAnswer(&_MockV3Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorSession) LatestRound() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestRound(&_MockV3Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestRound(&_MockV3Aggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockV3Aggregator *MockV3AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockV3Aggregator *MockV3AggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.LatestRoundData(&_MockV3Aggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.LatestRoundData(&_MockV3Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestTimestamp(&_MockV3Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestTimestamp(&_MockV3Aggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorSession) Version() (*big.Int, error) {
	return _MockV3Aggregator.Contract.Version(&_MockV3Aggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MockV3Aggregator *MockV3AggregatorCallerSession) Version() (*big.Int, error) {
	return _MockV3Aggregator.Contract.Version(&_MockV3Aggregator.CallOpts)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0xa87a20ce.
//
// Solidity: function updateAnswer(int256 _answer) returns()
func (_MockV3Aggregator *MockV3AggregatorTransactor) UpdateAnswer(opts *bind.TransactOpts, _answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.contract.Transact(opts, "updateAnswer", _answer)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0xa87a20ce.
//
// Solidity: function updateAnswer(int256 _answer) returns()
func (_MockV3Aggregator *MockV3AggregatorSession) UpdateAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateAnswer(&_MockV3Aggregator.TransactOpts, _answer)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0xa87a20ce.
//
// Solidity: function updateAnswer(int256 _answer) returns()
func (_MockV3Aggregator *MockV3AggregatorTransactorSession) UpdateAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateAnswer(&_MockV3Aggregator.TransactOpts, _answer)
}
