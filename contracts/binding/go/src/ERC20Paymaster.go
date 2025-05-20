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

// ERC20PaymasterMetaData contains all meta data concerning the ERC20Paymaster contract.
var ERC20PaymasterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenPerGas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setTokenPerGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPerGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedUserOps\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	ID:  "ERC20Paymaster",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161165038038061165083398101604081905261002e916101be565b82338061005557604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005e81610086565b50610068816100a2565b6001600160a01b039081166080529190911660a05260025550610224565b600180546001600160a01b031916905561009f8161015b565b50565b6040516301ffc9a760e01b8152631313998b60e31b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156100eb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061010f91906101fe565b61009f5760405162461bcd60e51b815260206004820152601e60248201527f49456e747279506f696e7420696e74657266616365206d69736d617463680000604482015260640161004c565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461009f575f5ffd5b5f5f5f606084860312156101d0575f5ffd5b83516101db816101aa565b60208501519093506101ec816101aa565b80925050604084015190509250925092565b5f6020828403121561020e575f5ffd5b8151801515811461021d575f5ffd5b9392505050565b60805160a0516113be6102925f395f81816103b30152818161061d015281816106ea01528181610cf90152610eef01525f818161028b01528181610411015281816104d7015281816107ca0152818161089f01528181610922015281816109d30152610b4401526113be5ff3fe60806040526004361061013d575f3560e01c8063b0d691fe116100bb578063c399ec8811610071578063e30c397811610057578063e30c397814610359578063f2fde38b14610383578063fc0c546a146103a2575f5ffd5b8063c399ec881461033d578063d0e30db014610351575f5ffd5b8063bc32e1e2116100a1578063bc32e1e2146102c1578063bfaaa7e7146102ff578063c23a5cea1461031e575f5ffd5b8063b0d691fe1461027a578063bb9fe6bf146102ad575f5ffd5b8063715018a6116101105780637c627b21116100f65780637c627b21146101f257806389476069146102115780638da5cb5b14610230575f5ffd5b8063715018a6146101ca57806379ba5097146101de575f5ffd5b80630396cb6014610141578063205c2878146101565780634e40f6341461017557806352b7512c1461019d575b5f5ffd5b61015461014f3660046110c7565b6103d5565b005b348015610161575f5ffd5b50610154610170366004611112565b610483565b348015610180575f5ffd5b5061018a60025481565b6040519081526020015b60405180910390f35b3480156101a8575f5ffd5b506101bc6101b736600461113c565b610518565b60405161019492919061118b565b3480156101d5575f5ffd5b5061015461053a565b3480156101e9575f5ffd5b5061015461054d565b3480156101fd575f5ffd5b5061015461020c3660046111e5565b6105c9565b34801561021c575f5ffd5b5061015461022b366004611279565b6105e5565b34801561023b575f5ffd5b505f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610194565b348015610285575f5ffd5b506102557f000000000000000000000000000000000000000000000000000000000000000081565b3480156102b8575f5ffd5b506101546107c0565b3480156102cc575f5ffd5b506102ef6102db366004611294565b60036020525f908152604090205460ff1681565b6040519015158152602001610194565b34801561030a575f5ffd5b50610154610319366004611294565b610845565b348015610329575f5ffd5b50610154610338366004611279565b610852565b348015610348575f5ffd5b5061018a6108f2565b6101546109a5565b348015610364575f5ffd5b5060015473ffffffffffffffffffffffffffffffffffffffff16610255565b34801561038e575f5ffd5b5061015461039d366004611279565b610a2b565b3480156103ad575f5ffd5b506102557f000000000000000000000000000000000000000000000000000000000000000081565b6103dd610ada565b6040517f0396cb6000000000000000000000000000000000000000000000000000000000815263ffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690630396cb609034906024015f604051808303818588803b158015610469575f5ffd5b505af115801561047b573d5f5f3e3d5ffd5b505050505050565b61048b610ada565b6040517f205c287800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063205c2878906044015f604051808303815f87803b158015610469575f5ffd5b60605f610523610b2c565b61052e858585610bcb565b91509150935093915050565b610542610ada565b61054b5f610e2e565b565b600154339073ffffffffffffffffffffffffffffffffffffffff1681146105bd576040517f118cdaa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b6105c681610e2e565b50565b6105d1610b2c565b6105de8585858585610e5f565b5050505050565b6105ed610ada565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610677573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061069b91906112ab565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303815f875af1158015610732573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061075691906112c2565b6107bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5769746864726177206661696c6564000000000000000000000000000000000060448201526064016105b4565b5050565b6107c8610ada565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bb9fe6bf6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561082d575f5ffd5b505af115801561083f573d5f5f3e3d5ffd5b50505050565b61084d610ada565b600255565b61085a610ada565b6040517fc23a5cea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c23a5cea906024015f604051808303815f87803b1580156108e0575f5ffd5b505af11580156105de573d5f5f3e3d5ffd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561097c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a091906112ab565b905090565b6040517fb760faf90000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b760faf99034906024015f604051808303818588803b1580156108e0575f5ffd5b610a33610ada565b6001805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155610a955f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461054b576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105b4565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461054b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f53656e646572206e6f7420456e747279506f696e74000000000000000000000060448201526064016105b4565b5f828152600360205260408120546060919060ff1615610c47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f557365724f7020616c726561647920757365640000000000000000000000000060448201526064016105b4565b5f848152600360209081526040822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610c8b90870187611279565b90505f670de0b6b3a764000060025486610ca591906112e1565b610caf9190611323565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291925082917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610d3e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6291906112ab565b1015610dca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e7420746f6b656e000000000000000000000000000060448201526064016105b4565b6040805173ffffffffffffffffffffffffffffffffffffffff84166020820152908101829052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052975f975095505050505050565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556105c681610f22565b5f80610e6d85870187611112565b90925090505f610e7d84866112e1565b90505f670de0b6b3a764000060025483610e9791906112e1565b610ea19190611323565b90506001896002811115610eb757610eb761135b565b03610eca57610ec7600282611323565b90505b82811115610ed55750815b610f1773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016853084610f96565b505050505050505050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152841660448201526064808201849052825180830390910181526084909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000178152825161083f93889390925f9283929183919082885af180611047576040513d5f823e3d81fd5b50505f513d9150811561105e578060011415611078565b73ffffffffffffffffffffffffffffffffffffffff84163b155b1561083f576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016105b4565b5f602082840312156110d7575f5ffd5b813563ffffffff811681146110ea575f5ffd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811681146105c6575f5ffd5b5f5f60408385031215611123575f5ffd5b823561112e816110f1565b946020939093013593505050565b5f5f5f6060848603121561114e575f5ffd5b833567ffffffffffffffff811115611164575f5ffd5b84016101208187031215611176575f5ffd5b95602085013595506040909401359392505050565b604081525f83518060408401528060208601606085015e5f6060828501015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168401019150508260208301529392505050565b5f5f5f5f5f608086880312156111f9575f5ffd5b853560038110611207575f5ffd5b9450602086013567ffffffffffffffff811115611222575f5ffd5b8601601f81018813611232575f5ffd5b803567ffffffffffffffff811115611248575f5ffd5b886020828401011115611259575f5ffd5b959860209190910197509495604081013595606090910135945092505050565b5f60208284031215611289575f5ffd5b81356110ea816110f1565b5f602082840312156112a4575f5ffd5b5035919050565b5f602082840312156112bb575f5ffd5b5051919050565b5f602082840312156112d2575f5ffd5b815180151581146110ea575f5ffd5b808202811582820484141761131d577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b5f82611356577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea264697066735822122055ebb3ade57e244e34a83d30054f408ad615f77738e3ffe6ea94e4ce1fafb18c64736f6c634300081c0033",
}

// ERC20Paymaster is an auto generated Go binding around an Ethereum contract.
type ERC20Paymaster struct {
	abi abi.ABI
}

// NewERC20Paymaster creates a new instance of ERC20Paymaster.
func NewERC20Paymaster() *ERC20Paymaster {
	parsed, err := ERC20PaymasterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Paymaster{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Paymaster) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _entryPoint, address _token, uint256 _tokenPerGas) returns()
func (eRC20Paymaster *ERC20Paymaster) PackConstructor(_entryPoint common.Address, _token common.Address, _tokenPerGas *big.Int) []byte {
	enc, err := eRC20Paymaster.abi.Pack("", _entryPoint, _token, _tokenPerGas)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAcceptOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (eRC20Paymaster *ERC20Paymaster) PackAcceptOwnership() []byte {
	enc, err := eRC20Paymaster.abi.Pack("acceptOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (eRC20Paymaster *ERC20Paymaster) PackAddStake(unstakeDelaySec uint32) []byte {
	enc, err := eRC20Paymaster.abi.Pack("addStake", unstakeDelaySec)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (eRC20Paymaster *ERC20Paymaster) PackDeposit() []byte {
	enc, err := eRC20Paymaster.abi.Pack("deposit")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEntryPoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) PackEntryPoint() []byte {
	enc, err := eRC20Paymaster.abi.Pack("entryPoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEntryPoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) UnpackEntryPoint(data []byte) (common.Address, error) {
	out, err := eRC20Paymaster.abi.Unpack("entryPoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (eRC20Paymaster *ERC20Paymaster) PackGetDeposit() []byte {
	enc, err := eRC20Paymaster.abi.Pack("getDeposit")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (eRC20Paymaster *ERC20Paymaster) UnpackGetDeposit(data []byte) (*big.Int, error) {
	out, err := eRC20Paymaster.abi.Unpack("getDeposit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) PackOwner() []byte {
	enc, err := eRC20Paymaster.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Paymaster.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) PackPendingOwner() []byte {
	enc, err := eRC20Paymaster.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := eRC20Paymaster.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPostOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (eRC20Paymaster *ERC20Paymaster) PackPostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) []byte {
	enc, err := eRC20Paymaster.abi.Pack("postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (eRC20Paymaster *ERC20Paymaster) PackRenounceOwnership() []byte {
	enc, err := eRC20Paymaster.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTokenPerGas is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfaaa7e7.
//
// Solidity: function setTokenPerGas(uint256 _price) returns()
func (eRC20Paymaster *ERC20Paymaster) PackSetTokenPerGas(price *big.Int) []byte {
	enc, err := eRC20Paymaster.abi.Pack("setTokenPerGas", price)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) PackToken() []byte {
	enc, err := eRC20Paymaster.abi.Pack("token")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (eRC20Paymaster *ERC20Paymaster) UnpackToken(data []byte) (common.Address, error) {
	out, err := eRC20Paymaster.abi.Unpack("token", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTokenPerGas is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e40f634.
//
// Solidity: function tokenPerGas() view returns(uint256)
func (eRC20Paymaster *ERC20Paymaster) PackTokenPerGas() []byte {
	enc, err := eRC20Paymaster.abi.Pack("tokenPerGas")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenPerGas is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e40f634.
//
// Solidity: function tokenPerGas() view returns(uint256)
func (eRC20Paymaster *ERC20Paymaster) UnpackTokenPerGas(data []byte) (*big.Int, error) {
	out, err := eRC20Paymaster.abi.Unpack("tokenPerGas", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (eRC20Paymaster *ERC20Paymaster) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := eRC20Paymaster.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnlockStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (eRC20Paymaster *ERC20Paymaster) PackUnlockStake() []byte {
	enc, err := eRC20Paymaster.abi.Pack("unlockStake")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUsedUserOps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc32e1e2.
//
// Solidity: function usedUserOps(bytes32 ) view returns(bool)
func (eRC20Paymaster *ERC20Paymaster) PackUsedUserOps(arg0 [32]byte) []byte {
	enc, err := eRC20Paymaster.abi.Pack("usedUserOps", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUsedUserOps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc32e1e2.
//
// Solidity: function usedUserOps(bytes32 ) view returns(bool)
func (eRC20Paymaster *ERC20Paymaster) UnpackUsedUserOps(data []byte) (bool, error) {
	out, err := eRC20Paymaster.abi.Unpack("usedUserOps", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackValidatePaymasterUserOp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (eRC20Paymaster *ERC20Paymaster) PackValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) []byte {
	enc, err := eRC20Paymaster.abi.Pack("validatePaymasterUserOp", userOp, userOpHash, maxCost)
	if err != nil {
		panic(err)
	}
	return enc
}

// ValidatePaymasterUserOpOutput serves as a container for the return parameters of contract
// method ValidatePaymasterUserOp.
type ValidatePaymasterUserOpOutput struct {
	Context        []byte
	ValidationData *big.Int
}

// UnpackValidatePaymasterUserOp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (eRC20Paymaster *ERC20Paymaster) UnpackValidatePaymasterUserOp(data []byte) (ValidatePaymasterUserOpOutput, error) {
	out, err := eRC20Paymaster.abi.Unpack("validatePaymasterUserOp", data)
	outstruct := new(ValidatePaymasterUserOpOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackWithdrawStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (eRC20Paymaster *ERC20Paymaster) PackWithdrawStake(withdrawAddress common.Address) []byte {
	enc, err := eRC20Paymaster.abi.Pack("withdrawStake", withdrawAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (eRC20Paymaster *ERC20Paymaster) PackWithdrawTo(withdrawAddress common.Address, amount *big.Int) []byte {
	enc, err := eRC20Paymaster.abi.Pack("withdrawTo", withdrawAddress, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdrawToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89476069.
//
// Solidity: function withdrawToken(address to) returns()
func (eRC20Paymaster *ERC20Paymaster) PackWithdrawToken(to common.Address) []byte {
	enc, err := eRC20Paymaster.abi.Pack("withdrawToken", to)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20PaymasterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ERC20Paymaster contract.
type ERC20PaymasterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ERC20PaymasterOwnershipTransferStartedEventName = "OwnershipTransferStarted"

// ContractEventName returns the user-defined event name.
func (ERC20PaymasterOwnershipTransferStarted) ContractEventName() string {
	return ERC20PaymasterOwnershipTransferStartedEventName
}

// UnpackOwnershipTransferStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (eRC20Paymaster *ERC20Paymaster) UnpackOwnershipTransferStartedEvent(log *types.Log) (*ERC20PaymasterOwnershipTransferStarted, error) {
	event := "OwnershipTransferStarted"
	if log.Topics[0] != eRC20Paymaster.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20PaymasterOwnershipTransferStarted)
	if len(log.Data) > 0 {
		if err := eRC20Paymaster.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Paymaster.abi.Events[event].Inputs {
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

// ERC20PaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Paymaster contract.
type ERC20PaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ERC20PaymasterOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ERC20PaymasterOwnershipTransferred) ContractEventName() string {
	return ERC20PaymasterOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (eRC20Paymaster *ERC20Paymaster) UnpackOwnershipTransferredEvent(log *types.Log) (*ERC20PaymasterOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != eRC20Paymaster.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20PaymasterOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := eRC20Paymaster.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Paymaster.abi.Events[event].Inputs {
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
func (eRC20Paymaster *ERC20Paymaster) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Paymaster.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return eRC20Paymaster.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Paymaster.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Paymaster.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Paymaster.abi.Errors["SafeERC20FailedOperation"].ID.Bytes()[:4]) {
		return eRC20Paymaster.UnpackSafeERC20FailedOperationError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20PaymasterOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the ERC20Paymaster contract.
type ERC20PaymasterOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func ERC20PaymasterOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (eRC20Paymaster *ERC20Paymaster) UnpackOwnableInvalidOwnerError(raw []byte) (*ERC20PaymasterOwnableInvalidOwner, error) {
	out := new(ERC20PaymasterOwnableInvalidOwner)
	if err := eRC20Paymaster.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PaymasterOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the ERC20Paymaster contract.
type ERC20PaymasterOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func ERC20PaymasterOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (eRC20Paymaster *ERC20Paymaster) UnpackOwnableUnauthorizedAccountError(raw []byte) (*ERC20PaymasterOwnableUnauthorizedAccount, error) {
	out := new(ERC20PaymasterOwnableUnauthorizedAccount)
	if err := eRC20Paymaster.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20PaymasterSafeERC20FailedOperation represents a SafeERC20FailedOperation error raised by the ERC20Paymaster contract.
type ERC20PaymasterSafeERC20FailedOperation struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeERC20FailedOperation(address token)
func ERC20PaymasterSafeERC20FailedOperationErrorID() common.Hash {
	return common.HexToHash("0x5274afe73c98b4749fc91ffae6b7b574e7842cb2144a159e9377a5f20b32edf9")
}

// UnpackSafeERC20FailedOperationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeERC20FailedOperation(address token)
func (eRC20Paymaster *ERC20Paymaster) UnpackSafeERC20FailedOperationError(raw []byte) (*ERC20PaymasterSafeERC20FailedOperation, error) {
	out := new(ERC20PaymasterSafeERC20FailedOperation)
	if err := eRC20Paymaster.abi.UnpackIntoInterface(out, "SafeERC20FailedOperation", raw); err != nil {
		return nil, err
	}
	return out, nil
}
