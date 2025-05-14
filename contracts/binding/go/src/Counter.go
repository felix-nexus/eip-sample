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
	ABI: "[{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Counter",
	Bin: "0x6080604052348015600e575f5ffd5b506101e18061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80633fb5c1cb146100435780638381f58a1461005f578063d09de08a1461007d575b5f5ffd5b61005d600480360381019061005891906100e4565b610087565b005b610067610090565b604051610074919061011e565b60405180910390f35b610085610095565b005b805f8190555050565b5f5481565b5f5f8154809291906100a690610164565b9190505550565b5f5ffd5b5f819050919050565b6100c3816100b1565b81146100cd575f5ffd5b50565b5f813590506100de816100ba565b92915050565b5f602082840312156100f9576100f86100ad565b5b5f610106848285016100d0565b91505092915050565b610118816100b1565b82525050565b5f6020820190506101315f83018461010f565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61016e826100b1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036101a05761019f610137565b5b60018201905091905056fea2646970667358221220a9c86ead1357f44ff78489c5c5e3ab80ac3d6e74f2e85f187f070a88ba34f6b364736f6c634300081c0033",
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
