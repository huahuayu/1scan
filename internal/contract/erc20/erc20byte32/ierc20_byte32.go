// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20byte32

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Ierc20byte32ABI is the input ABI used to generate the binding from.
const Ierc20byte32ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ierc20byte32 is an auto generated Go binding around an Ethereum contract.
type Ierc20byte32 struct {
	Ierc20byte32Caller     // Read-only binding to the contract
	Ierc20byte32Transactor // Write-only binding to the contract
	Ierc20byte32Filterer   // Log filterer for contract events
}

// Ierc20byte32Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc20byte32Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20byte32Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc20byte32Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20byte32Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc20byte32Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20byte32Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc20byte32Session struct {
	Contract     *Ierc20byte32     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ierc20byte32CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc20byte32CallerSession struct {
	Contract *Ierc20byte32Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ierc20byte32TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc20byte32TransactorSession struct {
	Contract     *Ierc20byte32Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ierc20byte32Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc20byte32Raw struct {
	Contract *Ierc20byte32 // Generic contract binding to access the raw methods on
}

// Ierc20byte32CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc20byte32CallerRaw struct {
	Contract *Ierc20byte32Caller // Generic read-only contract binding to access the raw methods on
}

// Ierc20byte32TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc20byte32TransactorRaw struct {
	Contract *Ierc20byte32Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc20byte32 creates a new instance of Ierc20byte32, bound to a specific deployed contract.
func NewIerc20byte32(address common.Address, backend bind.ContractBackend) (*Ierc20byte32, error) {
	contract, err := bindIerc20byte32(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32{Ierc20byte32Caller: Ierc20byte32Caller{contract: contract}, Ierc20byte32Transactor: Ierc20byte32Transactor{contract: contract}, Ierc20byte32Filterer: Ierc20byte32Filterer{contract: contract}}, nil
}

// NewIerc20byte32Caller creates a new read-only instance of Ierc20byte32, bound to a specific deployed contract.
func NewIerc20byte32Caller(address common.Address, caller bind.ContractCaller) (*Ierc20byte32Caller, error) {
	contract, err := bindIerc20byte32(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32Caller{contract: contract}, nil
}

// NewIerc20byte32Transactor creates a new write-only instance of Ierc20byte32, bound to a specific deployed contract.
func NewIerc20byte32Transactor(address common.Address, transactor bind.ContractTransactor) (*Ierc20byte32Transactor, error) {
	contract, err := bindIerc20byte32(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32Transactor{contract: contract}, nil
}

// NewIerc20byte32Filterer creates a new log filterer instance of Ierc20byte32, bound to a specific deployed contract.
func NewIerc20byte32Filterer(address common.Address, filterer bind.ContractFilterer) (*Ierc20byte32Filterer, error) {
	contract, err := bindIerc20byte32(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32Filterer{contract: contract}, nil
}

// bindIerc20byte32 binds a generic wrapper to an already deployed contract.
func bindIerc20byte32(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ierc20byte32ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20byte32 *Ierc20byte32Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20byte32.Contract.Ierc20byte32Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20byte32 *Ierc20byte32Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Ierc20byte32Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20byte32 *Ierc20byte32Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Ierc20byte32Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20byte32 *Ierc20byte32CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20byte32.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20byte32 *Ierc20byte32TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20byte32 *Ierc20byte32TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ierc20byte32.Contract.Allowance(&_Ierc20byte32.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ierc20byte32.Contract.Allowance(&_Ierc20byte32.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ierc20byte32.Contract.BalanceOf(&_Ierc20byte32.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ierc20byte32.Contract.BalanceOf(&_Ierc20byte32.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20byte32 *Ierc20byte32Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20byte32 *Ierc20byte32Session) Decimals() (uint8, error) {
	return _Ierc20byte32.Contract.Decimals(&_Ierc20byte32.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ierc20byte32 *Ierc20byte32CallerSession) Decimals() (uint8, error) {
	return _Ierc20byte32.Contract.Decimals(&_Ierc20byte32.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32Caller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32Session) Name() ([32]byte, error) {
	return _Ierc20byte32.Contract.Name(&_Ierc20byte32.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32CallerSession) Name() ([32]byte, error) {
	return _Ierc20byte32.Contract.Name(&_Ierc20byte32.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32Caller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32Session) Symbol() ([32]byte, error) {
	return _Ierc20byte32.Contract.Symbol(&_Ierc20byte32.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Ierc20byte32 *Ierc20byte32CallerSession) Symbol() ([32]byte, error) {
	return _Ierc20byte32.Contract.Symbol(&_Ierc20byte32.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20byte32.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32Session) TotalSupply() (*big.Int, error) {
	return _Ierc20byte32.Contract.TotalSupply(&_Ierc20byte32.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ierc20byte32 *Ierc20byte32CallerSession) TotalSupply() (*big.Int, error) {
	return _Ierc20byte32.Contract.TotalSupply(&_Ierc20byte32.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Approve(&_Ierc20byte32.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Approve(&_Ierc20byte32.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Transfer(&_Ierc20byte32.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.Transfer(&_Ierc20byte32.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.TransferFrom(&_Ierc20byte32.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Ierc20byte32 *Ierc20byte32TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Ierc20byte32.Contract.TransferFrom(&_Ierc20byte32.TransactOpts, from, to, value)
}

// Ierc20byte32ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ierc20byte32 contract.
type Ierc20byte32ApprovalIterator struct {
	Event *Ierc20byte32Approval // Event containing the contract specifics and raw log

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
func (it *Ierc20byte32ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc20byte32Approval)
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
		it.Event = new(Ierc20byte32Approval)
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
func (it *Ierc20byte32ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc20byte32ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc20byte32Approval represents a Approval event raised by the Ierc20byte32 contract.
type Ierc20byte32Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ierc20byte32 *Ierc20byte32Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Ierc20byte32ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ierc20byte32.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32ApprovalIterator{contract: _Ierc20byte32.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ierc20byte32 *Ierc20byte32Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Ierc20byte32Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ierc20byte32.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc20byte32Approval)
				if err := _Ierc20byte32.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ierc20byte32 *Ierc20byte32Filterer) ParseApproval(log types.Log) (*Ierc20byte32Approval, error) {
	event := new(Ierc20byte32Approval)
	if err := _Ierc20byte32.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ierc20byte32TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ierc20byte32 contract.
type Ierc20byte32TransferIterator struct {
	Event *Ierc20byte32Transfer // Event containing the contract specifics and raw log

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
func (it *Ierc20byte32TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ierc20byte32Transfer)
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
		it.Event = new(Ierc20byte32Transfer)
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
func (it *Ierc20byte32TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ierc20byte32TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ierc20byte32Transfer represents a Transfer event raised by the Ierc20byte32 contract.
type Ierc20byte32Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ierc20byte32 *Ierc20byte32Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Ierc20byte32TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ierc20byte32.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Ierc20byte32TransferIterator{contract: _Ierc20byte32.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ierc20byte32 *Ierc20byte32Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Ierc20byte32Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ierc20byte32.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ierc20byte32Transfer)
				if err := _Ierc20byte32.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ierc20byte32 *Ierc20byte32Filterer) ParseTransfer(log types.Log) (*Ierc20byte32Transfer, error) {
	event := new(Ierc20byte32Transfer)
	if err := _Ierc20byte32.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
