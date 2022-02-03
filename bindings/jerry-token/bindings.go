// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jerrytoken

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

// JerrytokenMetaData contains all meta data concerning the Jerrytoken contract.
var JerrytokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040805180820182526005808252642a37b5b2b760d91b60208084018281528551808701909652928552840152815191929161004f9160039161006b565b50805161006390600490602084019061006b565b50505061013f565b82805461007790610104565b90600052602060002090601f01602090048101928261009957600085556100df565b82601f106100b257805160ff19168380011785556100df565b828001600101855582156100df579182015b828111156100df5782518255916020019190600101906100c4565b506100eb9291506100ef565b5090565b5b808211156100eb57600081556001016100f0565b600181811c9082168061011857607f821691505b6020821081141561013957634e487b7160e01b600052602260045260246000fd5b50919050565b6109e98061014e6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806370a082311161007157806370a082311461014157806395d89b411461016a578063a0712d6814610172578063a457c2d714610187578063a9059cbb1461019a578063dd62ed3e146101ad57600080fd5b806306fdde03146100b9578063095ea7b3146100d757806318160ddd146100fa57806323b872dd1461010c578063313ce5671461011f578063395093511461012e575b600080fd5b6100c16101e6565b6040516100ce919061080d565b60405180910390f35b6100ea6100e536600461087e565b610278565b60405190151581526020016100ce565b6002545b6040519081526020016100ce565b6100ea61011a3660046108a8565b61028e565b604051601281526020016100ce565b6100ea61013c36600461087e565b61033d565b6100fe61014f3660046108e4565b6001600160a01b031660009081526020819052604090205490565b6100c1610379565b610185610180366004610906565b610388565b005b6100ea61019536600461087e565b610395565b6100ea6101a836600461087e565b61042e565b6100fe6101bb36600461091f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f590610952565b80601f016020809104026020016040519081016040528092919081815260200182805461022190610952565b801561026e5780601f106102435761010080835404028352916020019161026e565b820191906000526020600020905b81548152906001019060200180831161025157829003601f168201915b5050505050905090565b600061028533848461043b565b50600192915050565b600061029b84848461055f565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103255760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b610332853385840361043b565b506001949350505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161028591859061037490869061098d565b61043b565b6060600480546101f590610952565b610392338261072e565b50565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156104175760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161031c565b610424338585840361043b565b5060019392505050565b600061028533848461055f565b6001600160a01b03831661049d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161031c565b6001600160a01b0382166104fe5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161031c565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166105c35760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161031c565b6001600160a01b0382166106255760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161031c565b6001600160a01b0383166000908152602081905260409020548181101561069d5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161031c565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906106d490849061098d565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161072091815260200190565b60405180910390a350505050565b6001600160a01b0382166107845760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161031c565b8060026000828254610796919061098d565b90915550506001600160a01b038216600090815260208190526040812080548392906107c390849061098d565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600060208083528351808285015260005b8181101561083a5785810183015185820160400152820161081e565b8181111561084c576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461087957600080fd5b919050565b6000806040838503121561089157600080fd5b61089a83610862565b946020939093013593505050565b6000806000606084860312156108bd57600080fd5b6108c684610862565b92506108d460208501610862565b9150604084013590509250925092565b6000602082840312156108f657600080fd5b6108ff82610862565b9392505050565b60006020828403121561091857600080fd5b5035919050565b6000806040838503121561093257600080fd5b61093b83610862565b915061094960208401610862565b90509250929050565b600181811c9082168061096657607f821691505b6020821081141561098757634e487b7160e01b600052602260045260246000fd5b50919050565b600082198211156109ae57634e487b7160e01b600052601160045260246000fd5b50019056fea264697066735822122049c21225c0e9015712abae3476a5174cba5aa634f6b54128c01d97c0c853f2c064736f6c634300080a0033",
}

// JerrytokenABI is the input ABI used to generate the binding from.
// Deprecated: Use JerrytokenMetaData.ABI instead.
var JerrytokenABI = JerrytokenMetaData.ABI

// JerrytokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JerrytokenMetaData.Bin instead.
var JerrytokenBin = JerrytokenMetaData.Bin

// DeployJerrytoken deploys a new Ethereum contract, binding an instance of Jerrytoken to it.
func DeployJerrytoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jerrytoken, error) {
	parsed, err := JerrytokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JerrytokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jerrytoken{JerrytokenCaller: JerrytokenCaller{contract: contract}, JerrytokenTransactor: JerrytokenTransactor{contract: contract}, JerrytokenFilterer: JerrytokenFilterer{contract: contract}}, nil
}

// Jerrytoken is an auto generated Go binding around an Ethereum contract.
type Jerrytoken struct {
	JerrytokenCaller     // Read-only binding to the contract
	JerrytokenTransactor // Write-only binding to the contract
	JerrytokenFilterer   // Log filterer for contract events
}

// JerrytokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type JerrytokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JerrytokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JerrytokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JerrytokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JerrytokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JerrytokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JerrytokenSession struct {
	Contract     *Jerrytoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JerrytokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JerrytokenCallerSession struct {
	Contract *JerrytokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JerrytokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JerrytokenTransactorSession struct {
	Contract     *JerrytokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JerrytokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type JerrytokenRaw struct {
	Contract *Jerrytoken // Generic contract binding to access the raw methods on
}

// JerrytokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JerrytokenCallerRaw struct {
	Contract *JerrytokenCaller // Generic read-only contract binding to access the raw methods on
}

// JerrytokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JerrytokenTransactorRaw struct {
	Contract *JerrytokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJerrytoken creates a new instance of Jerrytoken, bound to a specific deployed contract.
func NewJerrytoken(address common.Address, backend bind.ContractBackend) (*Jerrytoken, error) {
	contract, err := bindJerrytoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jerrytoken{JerrytokenCaller: JerrytokenCaller{contract: contract}, JerrytokenTransactor: JerrytokenTransactor{contract: contract}, JerrytokenFilterer: JerrytokenFilterer{contract: contract}}, nil
}

// NewJerrytokenCaller creates a new read-only instance of Jerrytoken, bound to a specific deployed contract.
func NewJerrytokenCaller(address common.Address, caller bind.ContractCaller) (*JerrytokenCaller, error) {
	contract, err := bindJerrytoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JerrytokenCaller{contract: contract}, nil
}

// NewJerrytokenTransactor creates a new write-only instance of Jerrytoken, bound to a specific deployed contract.
func NewJerrytokenTransactor(address common.Address, transactor bind.ContractTransactor) (*JerrytokenTransactor, error) {
	contract, err := bindJerrytoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JerrytokenTransactor{contract: contract}, nil
}

// NewJerrytokenFilterer creates a new log filterer instance of Jerrytoken, bound to a specific deployed contract.
func NewJerrytokenFilterer(address common.Address, filterer bind.ContractFilterer) (*JerrytokenFilterer, error) {
	contract, err := bindJerrytoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JerrytokenFilterer{contract: contract}, nil
}

// bindJerrytoken binds a generic wrapper to an already deployed contract.
func bindJerrytoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JerrytokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jerrytoken *JerrytokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jerrytoken.Contract.JerrytokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jerrytoken *JerrytokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jerrytoken.Contract.JerrytokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jerrytoken *JerrytokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jerrytoken.Contract.JerrytokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jerrytoken *JerrytokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jerrytoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jerrytoken *JerrytokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jerrytoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jerrytoken *JerrytokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jerrytoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Jerrytoken *JerrytokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Jerrytoken *JerrytokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Jerrytoken.Contract.Allowance(&_Jerrytoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Jerrytoken *JerrytokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Jerrytoken.Contract.Allowance(&_Jerrytoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Jerrytoken *JerrytokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Jerrytoken *JerrytokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Jerrytoken.Contract.BalanceOf(&_Jerrytoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Jerrytoken *JerrytokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Jerrytoken.Contract.BalanceOf(&_Jerrytoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Jerrytoken *JerrytokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Jerrytoken *JerrytokenSession) Decimals() (uint8, error) {
	return _Jerrytoken.Contract.Decimals(&_Jerrytoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Jerrytoken *JerrytokenCallerSession) Decimals() (uint8, error) {
	return _Jerrytoken.Contract.Decimals(&_Jerrytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Jerrytoken *JerrytokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Jerrytoken *JerrytokenSession) Name() (string, error) {
	return _Jerrytoken.Contract.Name(&_Jerrytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Jerrytoken *JerrytokenCallerSession) Name() (string, error) {
	return _Jerrytoken.Contract.Name(&_Jerrytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Jerrytoken *JerrytokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Jerrytoken *JerrytokenSession) Symbol() (string, error) {
	return _Jerrytoken.Contract.Symbol(&_Jerrytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Jerrytoken *JerrytokenCallerSession) Symbol() (string, error) {
	return _Jerrytoken.Contract.Symbol(&_Jerrytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Jerrytoken *JerrytokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jerrytoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Jerrytoken *JerrytokenSession) TotalSupply() (*big.Int, error) {
	return _Jerrytoken.Contract.TotalSupply(&_Jerrytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Jerrytoken *JerrytokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Jerrytoken.Contract.TotalSupply(&_Jerrytoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Approve(&_Jerrytoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Approve(&_Jerrytoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Jerrytoken *JerrytokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Jerrytoken *JerrytokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.DecreaseAllowance(&_Jerrytoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Jerrytoken *JerrytokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.DecreaseAllowance(&_Jerrytoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Jerrytoken *JerrytokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Jerrytoken *JerrytokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.IncreaseAllowance(&_Jerrytoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Jerrytoken *JerrytokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.IncreaseAllowance(&_Jerrytoken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Jerrytoken *JerrytokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Jerrytoken *JerrytokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Mint(&_Jerrytoken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Jerrytoken *JerrytokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Mint(&_Jerrytoken.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Transfer(&_Jerrytoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.Transfer(&_Jerrytoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.TransferFrom(&_Jerrytoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Jerrytoken *JerrytokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Jerrytoken.Contract.TransferFrom(&_Jerrytoken.TransactOpts, sender, recipient, amount)
}

// JerrytokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Jerrytoken contract.
type JerrytokenApprovalIterator struct {
	Event *JerrytokenApproval // Event containing the contract specifics and raw log

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
func (it *JerrytokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JerrytokenApproval)
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
		it.Event = new(JerrytokenApproval)
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
func (it *JerrytokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JerrytokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JerrytokenApproval represents a Approval event raised by the Jerrytoken contract.
type JerrytokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JerrytokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Jerrytoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JerrytokenApprovalIterator{contract: _Jerrytoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JerrytokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Jerrytoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JerrytokenApproval)
				if err := _Jerrytoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) ParseApproval(log types.Log) (*JerrytokenApproval, error) {
	event := new(JerrytokenApproval)
	if err := _Jerrytoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JerrytokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Jerrytoken contract.
type JerrytokenTransferIterator struct {
	Event *JerrytokenTransfer // Event containing the contract specifics and raw log

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
func (it *JerrytokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JerrytokenTransfer)
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
		it.Event = new(JerrytokenTransfer)
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
func (it *JerrytokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JerrytokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JerrytokenTransfer represents a Transfer event raised by the Jerrytoken contract.
type JerrytokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JerrytokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Jerrytoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JerrytokenTransferIterator{contract: _Jerrytoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JerrytokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Jerrytoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JerrytokenTransfer)
				if err := _Jerrytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Jerrytoken *JerrytokenFilterer) ParseTransfer(log types.Log) (*JerrytokenTransfer, error) {
	event := new(JerrytokenTransfer)
	if err := _Jerrytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
