// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// PingMetaData contains all meta data concerning the Ping contract.
var PingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"multiSend\",\"inputs\":[{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ping\",\"inputs\":[{\"name\":\"createdTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pingCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Pong\",\"inputs\":[{\"name\":\"createdTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pingCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// PingABI is the input ABI used to generate the binding from.
// Deprecated: Use PingMetaData.ABI instead.
var PingABI = PingMetaData.ABI

// Ping is an auto generated Go binding around an Ethereum contract.
type Ping struct {
	PingCaller     // Read-only binding to the contract
	PingTransactor // Write-only binding to the contract
	PingFilterer   // Log filterer for contract events
}

// PingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingSession struct {
	Contract     *Ping             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingCallerSession struct {
	Contract *PingCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingTransactorSession struct {
	Contract     *PingTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingRaw struct {
	Contract *Ping // Generic contract binding to access the raw methods on
}

// PingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingCallerRaw struct {
	Contract *PingCaller // Generic read-only contract binding to access the raw methods on
}

// PingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingTransactorRaw struct {
	Contract *PingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPing creates a new instance of Ping, bound to a specific deployed contract.
func NewPing(address common.Address, backend bind.ContractBackend) (*Ping, error) {
	contract, err := bindPing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ping{PingCaller: PingCaller{contract: contract}, PingTransactor: PingTransactor{contract: contract}, PingFilterer: PingFilterer{contract: contract}}, nil
}

// NewPingCaller creates a new read-only instance of Ping, bound to a specific deployed contract.
func NewPingCaller(address common.Address, caller bind.ContractCaller) (*PingCaller, error) {
	contract, err := bindPing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingCaller{contract: contract}, nil
}

// NewPingTransactor creates a new write-only instance of Ping, bound to a specific deployed contract.
func NewPingTransactor(address common.Address, transactor bind.ContractTransactor) (*PingTransactor, error) {
	contract, err := bindPing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingTransactor{contract: contract}, nil
}

// NewPingFilterer creates a new log filterer instance of Ping, bound to a specific deployed contract.
func NewPingFilterer(address common.Address, filterer bind.ContractFilterer) (*PingFilterer, error) {
	contract, err := bindPing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingFilterer{contract: contract}, nil
}

// bindPing binds a generic wrapper to an already deployed contract.
func bindPing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.PingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transact(opts, method, params...)
}

// PingCount is a free data retrieval call binding the contract method 0x87704569.
//
// Solidity: function pingCount() view returns(uint256)
func (_Ping *PingCaller) PingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ping.contract.Call(opts, &out, "pingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PingCount is a free data retrieval call binding the contract method 0x87704569.
//
// Solidity: function pingCount() view returns(uint256)
func (_Ping *PingSession) PingCount() (*big.Int, error) {
	return _Ping.Contract.PingCount(&_Ping.CallOpts)
}

// PingCount is a free data retrieval call binding the contract method 0x87704569.
//
// Solidity: function pingCount() view returns(uint256)
func (_Ping *PingCallerSession) PingCount() (*big.Int, error) {
	return _Ping.Contract.PingCount(&_Ping.CallOpts)
}

// MultiSend is a paid mutator transaction binding the contract method 0xde8d3262.
//
// Solidity: function multiSend(address[] recipients, uint256 amount) payable returns()
func (_Ping *PingTransactor) MultiSend(opts *bind.TransactOpts, recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "multiSend", recipients, amount)
}

// MultiSend is a paid mutator transaction binding the contract method 0xde8d3262.
//
// Solidity: function multiSend(address[] recipients, uint256 amount) payable returns()
func (_Ping *PingSession) MultiSend(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.MultiSend(&_Ping.TransactOpts, recipients, amount)
}

// MultiSend is a paid mutator transaction binding the contract method 0xde8d3262.
//
// Solidity: function multiSend(address[] recipients, uint256 amount) payable returns()
func (_Ping *PingTransactorSession) MultiSend(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.MultiSend(&_Ping.TransactOpts, recipients, amount)
}

// Ping is a paid mutator transaction binding the contract method 0x773acdef.
//
// Solidity: function ping(uint256 createdTimestamp) returns()
func (_Ping *PingTransactor) Ping(opts *bind.TransactOpts, createdTimestamp *big.Int) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "ping", createdTimestamp)
}

// Ping is a paid mutator transaction binding the contract method 0x773acdef.
//
// Solidity: function ping(uint256 createdTimestamp) returns()
func (_Ping *PingSession) Ping(createdTimestamp *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.Ping(&_Ping.TransactOpts, createdTimestamp)
}

// Ping is a paid mutator transaction binding the contract method 0x773acdef.
//
// Solidity: function ping(uint256 createdTimestamp) returns()
func (_Ping *PingTransactorSession) Ping(createdTimestamp *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.Ping(&_Ping.TransactOpts, createdTimestamp)
}

// PingPongIterator is returned from FilterPong and is used to iterate over the raw logs and unpacked data for Pong events raised by the Ping contract.
type PingPongIterator struct {
	Event *PingPong // Event containing the contract specifics and raw log

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
func (it *PingPongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPong)
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
		it.Event = new(PingPong)
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
func (it *PingPongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPong represents a Pong event raised by the Ping contract.
type PingPong struct {
	CreatedTimestamp *big.Int
	PingCount        *big.Int
	BlockTimestamp   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPong is a free log retrieval operation binding the contract event 0x911212e82523572b4350e315606f05d2fa0abf7a009ace0d0ac208ace681a535.
//
// Solidity: event Pong(uint256 createdTimestamp, uint256 pingCount, uint256 blockTimestamp)
func (_Ping *PingFilterer) FilterPong(opts *bind.FilterOpts) (*PingPongIterator, error) {

	logs, sub, err := _Ping.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &PingPongIterator{contract: _Ping.contract, event: "Pong", logs: logs, sub: sub}, nil
}

// WatchPong is a free log subscription operation binding the contract event 0x911212e82523572b4350e315606f05d2fa0abf7a009ace0d0ac208ace681a535.
//
// Solidity: event Pong(uint256 createdTimestamp, uint256 pingCount, uint256 blockTimestamp)
func (_Ping *PingFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *PingPong) (event.Subscription, error) {

	logs, sub, err := _Ping.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPong)
				if err := _Ping.contract.UnpackLog(event, "Pong", log); err != nil {
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

// ParsePong is a log parse operation binding the contract event 0x911212e82523572b4350e315606f05d2fa0abf7a009ace0d0ac208ace681a535.
//
// Solidity: event Pong(uint256 createdTimestamp, uint256 pingCount, uint256 blockTimestamp)
func (_Ping *PingFilterer) ParsePong(log types.Log) (*PingPong, error) {
	event := new(PingPong)
	if err := _Ping.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
