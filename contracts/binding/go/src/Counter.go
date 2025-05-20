// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Caller\",\"type\":\"event\"}]",
	ID:  "Counter",
	Bin: "0x6080604052348015600e575f5ffd5b506101918061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80633fb5c1cb146100435780638381f58a14610058578063d09de08a14610072575b5f5ffd5b6100566100513660046100e8565b61007a565b005b6100605f5481565b60405190815260200160405180910390f35b6100566100a9565b5f81815560405133917f20f727e2e534ee85bce2bd7f2517088456c3f2040194b18a6c3da3fee1a02adf91a250565b5f805490806100b7836100ff565b909155505060405133907f20f727e2e534ee85bce2bd7f2517088456c3f2040194b18a6c3da3fee1a02adf905f90a2565b5f602082840312156100f8575f5ffd5b5035919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610154577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fea2646970667358221220fb55a7d0af8510420b08c81c8ca7e8c84ba41557caf80467e9155b643e6cab2f64736f6c634300081c0033",
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	abi abi.ABI
}

// NewCounter creates a new instance of Counter.
func NewCounter() *Counter {
	parsed, err := CounterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Counter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Counter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackIncrement is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd09de08a.
//
// Solidity: function increment() returns()
func (counter *Counter) PackIncrement() []byte {
	enc, err := counter.abi.Pack("increment")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (counter *Counter) PackNumber() []byte {
	enc, err := counter.abi.Pack("number")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (counter *Counter) UnpackNumber(data []byte) (*big.Int, error) {
	out, err := counter.abi.Unpack("number", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSetNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (counter *Counter) PackSetNumber(newNumber *big.Int) []byte {
	enc, err := counter.abi.Pack("setNumber", newNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// CounterCaller represents a Caller event raised by the Counter contract.
type CounterCaller struct {
	Caller common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const CounterCallerEventName = "Caller"

// ContractEventName returns the user-defined event name.
func (CounterCaller) ContractEventName() string {
	return CounterCallerEventName
}

// UnpackCallerEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Caller(address indexed caller)
func (counter *Counter) UnpackCallerEvent(log *types.Log) (*CounterCaller, error) {
	event := "Caller"
	if log.Topics[0] != counter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CounterCaller)
	if len(log.Data) > 0 {
		if err := counter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range counter.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
