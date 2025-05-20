package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	binding "github.com/felix-nexus/eip-sample/contracts/binding/go/src"
	"github.com/pkg/errors"
	"math/big"
)

func MergeGasToBytes32(gas0, gas1 *big.Int) [32]byte {
	bytes0 := common.LeftPadBytes(gas0.Bytes(), 16)
	bytes1 := common.LeftPadBytes(gas1.Bytes(), 16)
	var result [32]byte
	copy(result[:16], bytes0)
	copy(result[16:], bytes1)
	return result
}

// MakeGasFees abi.encodePacked(uint128(maxPriorityFeePerGas), uint128(maxFeePerGas))
func MakeGasFees(ctx context.Context, client bind.ContractBackend) ([32]byte, error) {
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return [32]byte{}, err
	}
	head, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return [32]byte{}, err
	}
	baseFee := head.BaseFee
	if baseFee == nil {
		return [32]byte{}, errors.New("base fee is nil")
	}
	gasFeeCap := new(big.Int).Mul(baseFee, common.Big2)
	return MergeGasToBytes32(gasTipCap, gasFeeCap), nil
}

func EntryPoint(backend *Backend) (common.Address, *binding.IEntryPoint, *bind.BoundContract) {
	aBI := binding.NewIEntryPoint()
	address := EntryPointAddress()
	return address, aBI, aBI.Instance(backend.Client(), address)
}

func MakePaymasterData(paymaster common.Address, validationGasLimit, postOpGasLimit uint64) []byte {
	// paymasterData
	/*
	   function _copyUserOpToMemory(PackedUserOperation calldata userOp, MemoryUserOp memory mUserOp)
	       internal
	       pure
	       virtual
	   {
	       mUserOp.sender = userOp.sender;
	       mUserOp.nonce = userOp.nonce;
	       (mUserOp.verificationGasLimit, mUserOp.callGasLimit) = UserOperationLib.unpackUints(userOp.accountGasLimits);
	       mUserOp.preVerificationGas = userOp.preVerificationGas;
	       (mUserOp.maxPriorityFeePerGas, mUserOp.maxFeePerGas) = UserOperationLib.unpackUints(userOp.gasFees);
	       bytes calldata paymasterAndData = userOp.paymasterAndData;
	   --- if (paymasterAndData.length > 0) {
	   --- require(paymasterAndData.length >= UserOperationLib.PAYMASTER_DATA_OFFSET, "AA93 invalid paymasterAndData");
	   --- address paymaster;
	   --- (paymaster, mUserOp.paymasterVerificationGasLimit, mUserOp.paymasterPostOpGasLimit) =
	   --- UserOperationLib.unpackPaymasterStaticFields(paymasterAndData);
	   --- require(paymaster != address(0), "AA98 invalid paymaster");
	   --- mUserOp.paymaster = paymaster;
	   --- }
	   }
	*/
	/*
		uint256 public constant PAYMASTER_VALIDATION_GAS_OFFSET = 20;
		uint256 public constant PAYMASTER_POSTOP_GAS_OFFSET = 36;
		uint256 public constant PAYMASTER_DATA_OFFSET = 52;
		   function unpackPaymasterStaticFields(
		       bytes calldata paymasterAndData
		   ) internal pure returns (address paymaster, uint256 validationGasLimit, uint256 postOpGasLimit) {
		       return (
		           address(bytes20(paymasterAndData[: PAYMASTER_VALIDATION_GAS_OFFSET])),
		           uint128(bytes16(paymasterAndData[PAYMASTER_VALIDATION_GAS_OFFSET : PAYMASTER_POSTOP_GAS_OFFSET])),
		           uint128(bytes16(paymasterAndData[PAYMASTER_POSTOP_GAS_OFFSET : PAYMASTER_DATA_OFFSET]))
		       );
		   }
	*/
	/*
		[]byte{pamymaster, mUserOp.paymasterVerificationGasLimit, mUserOp.paymasterPostOpGasLimit}
	*/

	address := common.LeftPadBytes(paymaster.Bytes(), 20)
	gas := MergeGasToBytes32(new(big.Int).SetUint64(validationGasLimit), new(big.Int).SetUint64(postOpGasLimit))
	return append(address, gas[:]...)
}

func SignPackedUserOperation(callOpts *bind.CallOpts, entryPointABI *binding.IEntryPoint, entryPoint *bind.BoundContract, userOp binding.PackedUserOperation, pk *ecdsa.PrivateKey) (binding.PackedUserOperation, error) {
	bytes, err := entryPoint.CallRaw(callOpts, entryPointABI.PackGetUserOpHash(userOp))
	if err != nil {
		return binding.PackedUserOperation{}, err
	}
	userOpHash, err := entryPointABI.UnpackGetUserOpHash(bytes)
	if err != nil {
		return binding.PackedUserOperation{}, err
	}
	signature, err := crypto.Sign(userOpHash[:], pk)
	if err != nil {
		return binding.PackedUserOperation{}, err
	}
	signature[64] += 27
	userOp.Signature = signature
	return userOp, nil
}
