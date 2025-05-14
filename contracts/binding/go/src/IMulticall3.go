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

// IMulticall3Call is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call struct {
	Target   common.Address
	CallData []byte
}

// IMulticall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// IMulticall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// IMulticall3Result is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Result struct {
	Success    bool
	ReturnData []byte
}

// IMulticall3MetaData contains all meta data concerning the IMulticall3 contract.
var IMulticall3MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "IMulticall3",
}

// IMulticall3 is an auto generated Go binding around an Ethereum contract.
type IMulticall3 struct {
	abi abi.ABI
}

// NewIMulticall3 creates a new instance of IMulticall3.
func NewIMulticall3() *IMulticall3 {
	parsed, err := IMulticall3MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IMulticall3{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IMulticall3) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (iMulticall3 *IMulticall3) PackAggregate(calls []IMulticall3Call) []byte {
	enc, err := iMulticall3.abi.Pack("aggregate", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// AggregateOutput serves as a container for the return parameters of contract
// method Aggregate.
type AggregateOutput struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}

// UnpackAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (iMulticall3 *IMulticall3) UnpackAggregate(data []byte) (AggregateOutput, error) {
	out, err := iMulticall3.abi.Unpack("aggregate", data)
	outstruct := new(AggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	return *outstruct, err

}

// PackAggregate3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) PackAggregate3(calls []IMulticall3Call3) []byte {
	enc, err := iMulticall3.abi.Pack("aggregate3", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAggregate3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) UnpackAggregate3(data []byte) ([]IMulticall3Result, error) {
	out, err := iMulticall3.abi.Unpack("aggregate3", data)
	if err != nil {
		return *new([]IMulticall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]IMulticall3Result)).(*[]IMulticall3Result)
	return out0, err
}

// PackAggregate3Value is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) PackAggregate3Value(calls []IMulticall3Call3Value) []byte {
	enc, err := iMulticall3.abi.Pack("aggregate3Value", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAggregate3Value is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) UnpackAggregate3Value(data []byte) ([]IMulticall3Result, error) {
	out, err := iMulticall3.abi.Unpack("aggregate3Value", data)
	if err != nil {
		return *new([]IMulticall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]IMulticall3Result)).(*[]IMulticall3Result)
	return out0, err
}

// PackBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) PackBlockAndAggregate(calls []IMulticall3Call) []byte {
	enc, err := iMulticall3.abi.Pack("blockAndAggregate", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// BlockAndAggregateOutput serves as a container for the return parameters of contract
// method BlockAndAggregate.
type BlockAndAggregateOutput struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []IMulticall3Result
}

// UnpackBlockAndAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) UnpackBlockAndAggregate(data []byte) (BlockAndAggregateOutput, error) {
	out, err := iMulticall3.abi.Unpack("blockAndAggregate", data)
	outstruct := new(BlockAndAggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]IMulticall3Result)).(*[]IMulticall3Result)
	return *outstruct, err

}

// PackGetBasefee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (iMulticall3 *IMulticall3) PackGetBasefee() []byte {
	enc, err := iMulticall3.abi.Pack("getBasefee")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBasefee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (iMulticall3 *IMulticall3) UnpackGetBasefee(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getBasefee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (iMulticall3 *IMulticall3) PackGetBlockHash(blockNumber *big.Int) []byte {
	enc, err := iMulticall3.abi.Pack("getBlockHash", blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBlockHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (iMulticall3 *IMulticall3) UnpackGetBlockHash(data []byte) ([32]byte, error) {
	out, err := iMulticall3.abi.Unpack("getBlockHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (iMulticall3 *IMulticall3) PackGetBlockNumber() []byte {
	enc, err := iMulticall3.abi.Pack("getBlockNumber")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBlockNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (iMulticall3 *IMulticall3) UnpackGetBlockNumber(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getBlockNumber", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (iMulticall3 *IMulticall3) PackGetChainId() []byte {
	enc, err := iMulticall3.abi.Pack("getChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (iMulticall3 *IMulticall3) UnpackGetChainId(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetCurrentBlockCoinbase is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (iMulticall3 *IMulticall3) PackGetCurrentBlockCoinbase() []byte {
	enc, err := iMulticall3.abi.Pack("getCurrentBlockCoinbase")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentBlockCoinbase is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (iMulticall3 *IMulticall3) UnpackGetCurrentBlockCoinbase(data []byte) (common.Address, error) {
	out, err := iMulticall3.abi.Unpack("getCurrentBlockCoinbase", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetCurrentBlockDifficulty is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (iMulticall3 *IMulticall3) PackGetCurrentBlockDifficulty() []byte {
	enc, err := iMulticall3.abi.Pack("getCurrentBlockDifficulty")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentBlockDifficulty is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (iMulticall3 *IMulticall3) UnpackGetCurrentBlockDifficulty(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getCurrentBlockDifficulty", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetCurrentBlockGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (iMulticall3 *IMulticall3) PackGetCurrentBlockGasLimit() []byte {
	enc, err := iMulticall3.abi.Pack("getCurrentBlockGasLimit")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentBlockGasLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (iMulticall3 *IMulticall3) UnpackGetCurrentBlockGasLimit(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getCurrentBlockGasLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetCurrentBlockTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (iMulticall3 *IMulticall3) PackGetCurrentBlockTimestamp() []byte {
	enc, err := iMulticall3.abi.Pack("getCurrentBlockTimestamp")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentBlockTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (iMulticall3 *IMulticall3) UnpackGetCurrentBlockTimestamp(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getCurrentBlockTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetEthBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (iMulticall3 *IMulticall3) PackGetEthBalance(addr common.Address) []byte {
	enc, err := iMulticall3.abi.Pack("getEthBalance", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEthBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (iMulticall3 *IMulticall3) UnpackGetEthBalance(data []byte) (*big.Int, error) {
	out, err := iMulticall3.abi.Unpack("getEthBalance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetLastBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (iMulticall3 *IMulticall3) PackGetLastBlockHash() []byte {
	enc, err := iMulticall3.abi.Pack("getLastBlockHash")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetLastBlockHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (iMulticall3 *IMulticall3) UnpackGetLastBlockHash(data []byte) ([32]byte, error) {
	out, err := iMulticall3.abi.Unpack("getLastBlockHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackTryAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) PackTryAggregate(requireSuccess bool, calls []IMulticall3Call) []byte {
	enc, err := iMulticall3.abi.Pack("tryAggregate", requireSuccess, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTryAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) UnpackTryAggregate(data []byte) ([]IMulticall3Result, error) {
	out, err := iMulticall3.abi.Unpack("tryAggregate", data)
	if err != nil {
		return *new([]IMulticall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]IMulticall3Result)).(*[]IMulticall3Result)
	return out0, err
}

// PackTryBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) PackTryBlockAndAggregate(requireSuccess bool, calls []IMulticall3Call) []byte {
	enc, err := iMulticall3.abi.Pack("tryBlockAndAggregate", requireSuccess, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryBlockAndAggregateOutput serves as a container for the return parameters of contract
// method TryBlockAndAggregate.
type TryBlockAndAggregateOutput struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []IMulticall3Result
}

// UnpackTryBlockAndAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (iMulticall3 *IMulticall3) UnpackTryBlockAndAggregate(data []byte) (TryBlockAndAggregateOutput, error) {
	out, err := iMulticall3.abi.Unpack("tryBlockAndAggregate", data)
	outstruct := new(TryBlockAndAggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]IMulticall3Result)).(*[]IMulticall3Result)
	return *outstruct, err

}
