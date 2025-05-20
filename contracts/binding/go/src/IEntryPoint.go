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

// IEntryPointUserOpsPerAggregator is an auto generated low-level Go binding around an user-defined struct.
type IEntryPointUserOpsPerAggregator struct {
	UserOps    []PackedUserOperation
	Aggregator common.Address
	Signature  []byte
}

// IStakeManagerDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerDepositInfo struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// IEntryPointMetaData contains all meta data concerning the IEntryPoint contract.
var IEntryPointMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senderCreator\",\"outputs\":[{\"internalType\":\"contractISenderCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"PostOpRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"UserOperationPrefundTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"name\":\"DelegateAndRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"inner\",\"type\":\"bytes\"}],\"name\":\"FailedOpWithRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"PostOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"}]",
	ID:  "IEntryPoint",
}

// IEntryPoint is an auto generated Go binding around an Ethereum contract.
type IEntryPoint struct {
	abi abi.ABI
}

// NewIEntryPoint creates a new instance of IEntryPoint.
func NewIEntryPoint() *IEntryPoint {
	parsed, err := IEntryPointMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IEntryPoint{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IEntryPoint) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (iEntryPoint *IEntryPoint) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := iEntryPoint.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iEntryPoint *IEntryPoint) PackBalanceOf(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iEntryPoint *IEntryPoint) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iEntryPoint.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDelegateAndRevert is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x850aaf62.
//
// Solidity: function delegateAndRevert(address target, bytes data) returns()
func (iEntryPoint *IEntryPoint) PackDelegateAndRevert(target common.Address, data []byte) []byte {
	enc, err := iEntryPoint.abi.Pack("delegateAndRevert", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDepositTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (iEntryPoint *IEntryPoint) PackDepositTo(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("depositTo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDepositInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iEntryPoint *IEntryPoint) PackGetDepositInfo(account common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("getDepositInfo", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDepositInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint256,bool,uint112,uint32,uint48) info)
func (iEntryPoint *IEntryPoint) UnpackGetDepositInfo(data []byte) (IStakeManagerDepositInfo, error) {
	out, err := iEntryPoint.abi.Unpack("getDepositInfo", data)
	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}
	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)
	return out0, err
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iEntryPoint *IEntryPoint) PackGetNonce(sender common.Address, key *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("getNonce", sender, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (iEntryPoint *IEntryPoint) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := iEntryPoint.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetSenderAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (iEntryPoint *IEntryPoint) PackGetSenderAddress(initCode []byte) []byte {
	enc, err := iEntryPoint.abi.Pack("getSenderAddress", initCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetUserOpHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (iEntryPoint *IEntryPoint) PackGetUserOpHash(userOp PackedUserOperation) []byte {
	enc, err := iEntryPoint.abi.Pack("getUserOpHash", userOp)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserOpHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22cdde4c.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp) view returns(bytes32)
func (iEntryPoint *IEntryPoint) UnpackGetUserOpHash(data []byte) ([32]byte, error) {
	out, err := iEntryPoint.abi.Unpack("getUserOpHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackHandleAggregatedOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbed18e0.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (iEntryPoint *IEntryPoint) PackHandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("handleAggregatedOps", opsPerAggregator, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (iEntryPoint *IEntryPoint) PackHandleOps(ops []PackedUserOperation, beneficiary common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("handleOps", ops, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (iEntryPoint *IEntryPoint) PackIncrementNonce(key *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("incrementNonce", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSenderCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (iEntryPoint *IEntryPoint) PackSenderCreator() []byte {
	enc, err := iEntryPoint.abi.Pack("senderCreator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSenderCreator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (iEntryPoint *IEntryPoint) UnpackSenderCreator(data []byte) (common.Address, error) {
	out, err := iEntryPoint.abi.Unpack("senderCreator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (iEntryPoint *IEntryPoint) PackUnlockStake() []byte {
	enc, err := iEntryPoint.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (iEntryPoint *IEntryPoint) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := iEntryPoint.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (iEntryPoint *IEntryPoint) PackWithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) []byte {
	enc, err := iEntryPoint.abi.Pack("withdrawTo", withdrawAddress, withdrawAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// IEntryPointAccountDeployed represents a AccountDeployed event raised by the IEntryPoint contract.
type IEntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointAccountDeployedEventName = "AccountDeployed"

// ContractEventName returns the user-defined event name.
func (IEntryPointAccountDeployed) ContractEventName() string {
	return IEntryPointAccountDeployedEventName
}

// UnpackAccountDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (iEntryPoint *IEntryPoint) UnpackAccountDeployedEvent(log *types.Log) (*IEntryPointAccountDeployed, error) {
	event := "AccountDeployed"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointAccountDeployed)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointBeforeExecution represents a BeforeExecution event raised by the IEntryPoint contract.
type IEntryPointBeforeExecution struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const IEntryPointBeforeExecutionEventName = "BeforeExecution"

// ContractEventName returns the user-defined event name.
func (IEntryPointBeforeExecution) ContractEventName() string {
	return IEntryPointBeforeExecutionEventName
}

// UnpackBeforeExecutionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeforeExecution()
func (iEntryPoint *IEntryPoint) UnpackBeforeExecutionEvent(log *types.Log) (*IEntryPointBeforeExecution, error) {
	event := "BeforeExecution"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointBeforeExecution)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointDeposited represents a Deposited event raised by the IEntryPoint contract.
type IEntryPointDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointDepositedEventName = "Deposited"

// ContractEventName returns the user-defined event name.
func (IEntryPointDeposited) ContractEventName() string {
	return IEntryPointDepositedEventName
}

// UnpackDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (iEntryPoint *IEntryPoint) UnpackDepositedEvent(log *types.Log) (*IEntryPointDeposited, error) {
	event := "Deposited"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointDeposited)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointPostOpRevertReason represents a PostOpRevertReason event raised by the IEntryPoint contract.
type IEntryPointPostOpRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointPostOpRevertReasonEventName = "PostOpRevertReason"

// ContractEventName returns the user-defined event name.
func (IEntryPointPostOpRevertReason) ContractEventName() string {
	return IEntryPointPostOpRevertReasonEventName
}

// UnpackPostOpRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PostOpRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (iEntryPoint *IEntryPoint) UnpackPostOpRevertReasonEvent(log *types.Log) (*IEntryPointPostOpRevertReason, error) {
	event := "PostOpRevertReason"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointPostOpRevertReason)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the IEntryPoint contract.
type IEntryPointSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointSignatureAggregatorChangedEventName = "SignatureAggregatorChanged"

// ContractEventName returns the user-defined event name.
func (IEntryPointSignatureAggregatorChanged) ContractEventName() string {
	return IEntryPointSignatureAggregatorChangedEventName
}

// UnpackSignatureAggregatorChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (iEntryPoint *IEntryPoint) UnpackSignatureAggregatorChangedEvent(log *types.Log) (*IEntryPointSignatureAggregatorChanged, error) {
	event := "SignatureAggregatorChanged"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointSignatureAggregatorChanged)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointStakeLocked represents a StakeLocked event raised by the IEntryPoint contract.
type IEntryPointStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeLockedEventName = "StakeLocked"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeLocked) ContractEventName() string {
	return IEntryPointStakeLockedEventName
}

// UnpackStakeLockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (iEntryPoint *IEntryPoint) UnpackStakeLockedEvent(log *types.Log) (*IEntryPointStakeLocked, error) {
	event := "StakeLocked"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeLocked)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointStakeUnlocked represents a StakeUnlocked event raised by the IEntryPoint contract.
type IEntryPointStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeUnlockedEventName = "StakeUnlocked"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeUnlocked) ContractEventName() string {
	return IEntryPointStakeUnlockedEventName
}

// UnpackStakeUnlockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (iEntryPoint *IEntryPoint) UnpackStakeUnlockedEvent(log *types.Log) (*IEntryPointStakeUnlocked, error) {
	event := "StakeUnlocked"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeUnlocked)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointStakeWithdrawn represents a StakeWithdrawn event raised by the IEntryPoint contract.
type IEntryPointStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointStakeWithdrawnEventName = "StakeWithdrawn"

// ContractEventName returns the user-defined event name.
func (IEntryPointStakeWithdrawn) ContractEventName() string {
	return IEntryPointStakeWithdrawnEventName
}

// UnpackStakeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iEntryPoint *IEntryPoint) UnpackStakeWithdrawnEvent(log *types.Log) (*IEntryPointStakeWithdrawn, error) {
	event := "StakeWithdrawn"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointStakeWithdrawn)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointUserOperationEvent represents a UserOperationEvent event raised by the IEntryPoint contract.
type IEntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationEventEventName = "UserOperationEvent"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationEvent) ContractEventName() string {
	return IEntryPointUserOperationEventEventName
}

// UnpackUserOperationEventEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (iEntryPoint *IEntryPoint) UnpackUserOperationEventEvent(log *types.Log) (*IEntryPointUserOperationEvent, error) {
	event := "UserOperationEvent"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationEvent)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointUserOperationPrefundTooLow represents a UserOperationPrefundTooLow event raised by the IEntryPoint contract.
type IEntryPointUserOperationPrefundTooLow struct {
	UserOpHash [32]byte
	Sender     common.Address
	Nonce      *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationPrefundTooLowEventName = "UserOperationPrefundTooLow"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationPrefundTooLow) ContractEventName() string {
	return IEntryPointUserOperationPrefundTooLowEventName
}

// UnpackUserOperationPrefundTooLowEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationPrefundTooLow(bytes32 indexed userOpHash, address indexed sender, uint256 nonce)
func (iEntryPoint *IEntryPoint) UnpackUserOperationPrefundTooLowEvent(log *types.Log) (*IEntryPointUserOperationPrefundTooLow, error) {
	event := "UserOperationPrefundTooLow"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationPrefundTooLow)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the IEntryPoint contract.
type IEntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const IEntryPointUserOperationRevertReasonEventName = "UserOperationRevertReason"

// ContractEventName returns the user-defined event name.
func (IEntryPointUserOperationRevertReason) ContractEventName() string {
	return IEntryPointUserOperationRevertReasonEventName
}

// UnpackUserOperationRevertReasonEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (iEntryPoint *IEntryPoint) UnpackUserOperationRevertReasonEvent(log *types.Log) (*IEntryPointUserOperationRevertReason, error) {
	event := "UserOperationRevertReason"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointUserOperationRevertReason)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// IEntryPointWithdrawn represents a Withdrawn event raised by the IEntryPoint contract.
type IEntryPointWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const IEntryPointWithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (IEntryPointWithdrawn) ContractEventName() string {
	return IEntryPointWithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (iEntryPoint *IEntryPoint) UnpackWithdrawnEvent(log *types.Log) (*IEntryPointWithdrawn, error) {
	event := "Withdrawn"
	if log.Topics[0] != iEntryPoint.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IEntryPointWithdrawn)
	if len(log.Data) > 0 {
		if err := iEntryPoint.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iEntryPoint.abi.Events[event].Inputs {
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

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iEntryPoint *IEntryPoint) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["DelegateAndRevert"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackDelegateAndRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["FailedOp"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackFailedOpError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["FailedOpWithRevert"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackFailedOpWithRevertError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["PostOpReverted"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackPostOpRevertedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["SenderAddressResult"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackSenderAddressResultError(raw[4:])
	}
	if bytes.Equal(raw[:4], iEntryPoint.abi.Errors["SignatureValidationFailed"].ID.Bytes()[:4]) {
		return iEntryPoint.UnpackSignatureValidationFailedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IEntryPointDelegateAndRevert represents a DelegateAndRevert error raised by the IEntryPoint contract.
type IEntryPointDelegateAndRevert struct {
	Success bool
	Ret     []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func IEntryPointDelegateAndRevertErrorID() common.Hash {
	return common.HexToHash("0x9941055444a6b8379a4c66beac4880d5e96ca9c2fff188903cb0bc945a4dae05")
}

// UnpackDelegateAndRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelegateAndRevert(bool success, bytes ret)
func (iEntryPoint *IEntryPoint) UnpackDelegateAndRevertError(raw []byte) (*IEntryPointDelegateAndRevert, error) {
	out := new(IEntryPointDelegateAndRevert)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "DelegateAndRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointFailedOp represents a FailedOp error raised by the IEntryPoint contract.
type IEntryPointFailedOp struct {
	OpIndex *big.Int
	Reason  string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func IEntryPointFailedOpErrorID() common.Hash {
	return common.HexToHash("0x220266b6eadfd2488e398d00abc1c765e1f119da6226c1b55ec9cc0186560475")
}

// UnpackFailedOpError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOp(uint256 opIndex, string reason)
func (iEntryPoint *IEntryPoint) UnpackFailedOpError(raw []byte) (*IEntryPointFailedOp, error) {
	out := new(IEntryPointFailedOp)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "FailedOp", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointFailedOpWithRevert represents a FailedOpWithRevert error raised by the IEntryPoint contract.
type IEntryPointFailedOpWithRevert struct {
	OpIndex *big.Int
	Reason  string
	Inner   []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func IEntryPointFailedOpWithRevertErrorID() common.Hash {
	return common.HexToHash("0x65c8fd4dd32f2bb83f7d57728e1afa231e9956aaf4bdeea76c78b967426acb8c")
}

// UnpackFailedOpWithRevertError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedOpWithRevert(uint256 opIndex, string reason, bytes inner)
func (iEntryPoint *IEntryPoint) UnpackFailedOpWithRevertError(raw []byte) (*IEntryPointFailedOpWithRevert, error) {
	out := new(IEntryPointFailedOpWithRevert)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "FailedOpWithRevert", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointPostOpReverted represents a PostOpReverted error raised by the IEntryPoint contract.
type IEntryPointPostOpReverted struct {
	ReturnData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PostOpReverted(bytes returnData)
func IEntryPointPostOpRevertedErrorID() common.Hash {
	return common.HexToHash("0xad7954bcae0d69c56f2323097b645bc6b6033c9f3777375271d77a608fc206ae")
}

// UnpackPostOpRevertedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PostOpReverted(bytes returnData)
func (iEntryPoint *IEntryPoint) UnpackPostOpRevertedError(raw []byte) (*IEntryPointPostOpReverted, error) {
	out := new(IEntryPointPostOpReverted)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "PostOpReverted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointSenderAddressResult represents a SenderAddressResult error raised by the IEntryPoint contract.
type IEntryPointSenderAddressResult struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SenderAddressResult(address sender)
func IEntryPointSenderAddressResultErrorID() common.Hash {
	return common.HexToHash("0x6ca7b806055bb56d141e3a30604d4c6540fe6f010574d1b40dfd951da77b956d")
}

// UnpackSenderAddressResultError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SenderAddressResult(address sender)
func (iEntryPoint *IEntryPoint) UnpackSenderAddressResultError(raw []byte) (*IEntryPointSenderAddressResult, error) {
	out := new(IEntryPointSenderAddressResult)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "SenderAddressResult", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IEntryPointSignatureValidationFailed represents a SignatureValidationFailed error raised by the IEntryPoint contract.
type IEntryPointSignatureValidationFailed struct {
	Aggregator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func IEntryPointSignatureValidationFailedErrorID() common.Hash {
	return common.HexToHash("0x86a9f750e187f100e9e8a18befbbf44e6f512d6f78cfe16a061541aaaad52523")
}

// UnpackSignatureValidationFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignatureValidationFailed(address aggregator)
func (iEntryPoint *IEntryPoint) UnpackSignatureValidationFailedError(raw []byte) (*IEntryPointSignatureValidationFailed, error) {
	out := new(IEntryPointSignatureValidationFailed)
	if err := iEntryPoint.abi.UnpackIntoInterface(out, "SignatureValidationFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}
