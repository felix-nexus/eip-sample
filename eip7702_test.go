package eip_sample

import (
	"context"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	binding "github.com/felix-nexus/eip-sample/contracts/binding/go/src"
	"github.com/felix-nexus/eip-sample/utils"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// https://eips.ethereum.org/EIPS/eip-7702
// https://eip7702.io
// https://hackmd.io/@rimeissner/eip7702-best-practices

// https://github.com/ethereum/go-ethereum/blob/v1.15.11/core/vm/jump_table.go#L105
// Prague 이후에 EIP-7702 를 지원한다.
// Prague 에는 solidity opcode 는 추가되지 않았다.

var (
	ctx      = context.Background()
	callOpts = &bind.CallOpts{Context: ctx}
)

// Multicall3 컨트랙트 코드를 EOA 에 등록 (eip7702)
// ERC20 토큰을 사용한다. (burn, mint, transfer, approve)
func TestEIP7702(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	chainId, err := backend.Client().ChainID(ctx)
	require.NoError(t, err)

	erc20Address, erc20ABI, erc20 := utils.DeployMockERC20(t, backend, backend.Owner(), "MockERC20", "MERC", 18)

	/// 1. MockERC20 Check
	// 1 을 민팅해 본다.
	tx, err := erc20.RawTransact(backend.Owner(), erc20ABI.PackMint(backend.Owner().From, common.Big1))
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// 1-1.
	// 1이 민팅되었는지 확인한다.
	bytes, err := erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err := erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)

	/// 2. EIP-7702 트랜잭션 실행
	// 2-1. 코드를 등록할 주소에 실제로 코드가 있는지 확인한다. (Multicall3)
	code, err := backend.Client().PendingCodeAt(ctx, utils.Multicall3Address())
	require.NoError(t, err)
	require.NotEmpty(t, code)

	// 2-2. 트랜잭션에 필요한 데이터중 체인에서 가져올수 있는 데이터를 가져온다.
	nonce, err := backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	gasTipCap, err := backend.Client().SuggestGasTipCap(ctx)
	require.NoError(t, err)

	// 2-3. 트랜잭션 생성에 필요한 데이터를 생성한다.
	prv, spender, recipient := utils.GetPrivateKey(t, backend.Owner().From), common.BytesToAddress([]byte("SPENDER")), common.BytesToAddress([]byte("RECIPIENT"))
	// 2-3-1. Multicall3 의 코드를 등록하기위한 데이터 (auth)
	auth, err := types.SignSetCode(prv, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(chainId), // 체인 ID
		Address: utils.Multicall3Address(),     // Code가 있는 컨트랙트 주소
		Nonce:   nonce + 1,                     // Tx Nonce + 1
	})
	require.NoError(t, err)
	// 2-3-2. auth 가 올바르게 생성되었는지 확인한다.
	authority, err := auth.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)

	// 2-4. 트랜잭션 데이터
	// Multicall3 를 사용하여 ERC20 의 각종 함수를 실행한다.
	data := binding.NewIMulticall3().PackAggregate3([]binding.IMulticall3Call3{
		// burn (1 case 에서 민팅을 했기때문에, 해당 벨런스 제거 가능)
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackBurn(backend.Owner().From, common.Big1),
		},
		// mint
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackMint(backend.Owner().From, common.Big2),
		},
		// transfer
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackTransfer(recipient, common.Big1),
		},
		// approve
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackApprove(spender, common.Big256),
		},
	})
	// 2-5. 트랜잭션을 생성 및 실행한다.
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(chainId),
		Nonce:     nonce,
		GasTipCap: uint256.MustFromBig(gasTipCap),
		GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
		Gas:       1e6,
		To:        backend.Owner().From, // 반드시 본인의 EOA 주소여야 한다.
		Value:     uint256.NewInt(0),
		Data:      data,
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()

	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Log("receipt.GasUsed", receipt.GasUsed)
	t.Log("receipt.Logs", len(receipt.Logs))
	for _, log := range receipt.Logs {
		if l, err := erc20ABI.UnpackTransferEvent(log); err == nil {
			t.Logf("[TRANSFER] from: %s, to: %s, value: %s", l.From.Hex(), l.To.Hex(), l.Amount.String())
		} else {
			l, err := erc20ABI.UnpackApprovalEvent(log)
			require.NoError(t, err)
			t.Logf("[APPROVAL] owner: %s, spender: %s, value: %s", l.Owner.Hex(), l.Spender.Hex(), l.Amount.String())
		}
	}

	/// 3. 결과 조회
	// 3-1. Balance: mint(1) -> burn(1) -> mint(2) -> transfer(1) = 1
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 3-2. recipient 의 balance = 1
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(recipient))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 3-3. allowance = 256
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackAllowance(backend.Owner().From, spender))
	require.NoError(t, err)
	allowance, err := erc20ABI.UnpackAllowance(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big256, allowance)

	/// 4. EIP-7702 는 해당 트랜잭션에서만 동작해야 하는데?
	// 4-1. EOA 주소에는 코드가... 남아있다?
	code, err = backend.Client().PendingCodeAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	delegation, ok := types.ParseDelegation(code)
	require.True(t, ok)
	// 4-1-1. Multicall3 의 주소가 Delegation 형태로 등록되어 있구나.
	// DelegationPrefix is used by code to denote the account is delegating to
	require.Equal(t, utils.Multicall3Address(), delegation)
	// 4-2. 코드가 남아있을꺼라 기대하고, DynamicFeeTx 을 사용할 수 없어야 하는데...?
	nonce, err = backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	address := backend.Owner().From
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: new(big.Int).Add(gasTipCap, big.NewInt(2e9)),
		Gas:       1e6,
		To:        &address, // 본인한테 보내는 트랜잭션
		Value:     common.Big0,
		Data:      data,
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// 이게 동작을 하네...
	t.Log("==============================================================================================================================")
	for _, log := range receipt.Logs {
		if l, err := erc20ABI.UnpackTransferEvent(log); err == nil {
			t.Logf("[TRANSFER] from: %s, to: %s, value: %s", l.From.Hex(), l.To.Hex(), l.Amount.String())
		} else {
			l, err := erc20ABI.UnpackApprovalEvent(log)
			require.NoError(t, err)
			t.Logf("[APPROVAL] owner: %s, spender: %s, value: %s", l.Owner.Hex(), l.Spender.Hex(), l.Amount.String())
		}
	}
	// 5. 일단 동작을 했으니 정상적으로 적용되었는지 확인해볼까
	// 5-1 beforeBalance(1) -> burn(1) -> mint(2) -> transfer(1) = 1
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 5-2. recipient 의 balance = 2
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(recipient))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big2, balance)
	// 5-3. allowance = 256
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackAllowance(backend.Owner().From, spender))
	require.NoError(t, err)
	allowance, err = erc20ABI.UnpackAllowance(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big256, allowance)
}

// TestEIP7702 에서 예상과 다른 동작을 했서, 의심되는 부분을 테스트 해본다.
// CASE: Counter 를 사용하여 nonce 를 증가시키지만  해당 트랜잭션 이후에서는 데이터가 남아있지 않아야 한다.
// Result: 남아있다...
// ==> SetCode 는 해당 트랜잭션에만 적용 된다는 멘트는 어떤 의미일까...
// ==> Simulated 환경에서는 남아있고, 실제 노드에서는
// eth_getCode 의 결과가 ParseDelegation true 인 경우에는 빈 배열을 리턴한다고 한다.
func TestEIP7702Delegation(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	chainId, err := backend.Client().ChainID(ctx)
	require.NoError(t, err)

	counterAddress, counterABI, counter := utils.DeployCounter(t, backend, backend.Owner())

	/// 1. Counter Check
	// 1 을 증가시켜본다.
	tx, err := counter.RawTransact(backend.Owner(), counterABI.PackIncrement())
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// 1-1. 1이 증가되었는지 확인한다.
	bytes, err := counter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err := counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)

	// 2. EIP-7702 트랜잭션 실행
	// 2-1. 코드를 등록할 주소에 실제로 코드가 있는지 확인한다. (Counter)
	code, err := backend.Client().PendingCodeAt(ctx, counterAddress)
	require.NoError(t, err)
	require.NotEmpty(t, code)
	// 2-2. 트랜잭션에 필요한 데이터중 체인에서 가져올수 있는 데이터를 가져온다.
	nonce, err := backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	gasTipCap, err := backend.Client().SuggestGasTipCap(ctx)
	require.NoError(t, err)
	// 2-3. 트랜잭션 생성에 필요한 데이터를 생성한다.
	prv := utils.GetPrivateKey(t, backend.Owner().From)
	// 2-3-1. Multicall3 의 코드를 등록하기위한 데이터 (auth)
	auth, err := types.SignSetCode(prv, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(chainId), // 체인 ID
		Address: counterAddress,                // Code가 있는 컨트랙트 주소
		Nonce:   nonce + 1,                     // Tx Nonce + 1
	})
	require.NoError(t, err)
	// 2-3-2. auth 가 올바르게 생성되었는지 확인한다.
	authority, err := auth.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)
	// 2-4. 트랜잭션 데이터
	// Counter 를 사용하여 1을 증가시킨다.
	data := counterABI.PackIncrement()
	// 2-5. 트랜잭션을 생성 및 실행한다.
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(chainId),
		Nonce:     nonce,
		GasTipCap: uint256.MustFromBig(gasTipCap),
		GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
		Gas:       1e6,
		To:        backend.Owner().From, // 반드시 본인의 EOA 주소여야 한다.
		Value:     uint256.NewInt(0),
		Data:      data,
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()

	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Log("receipt.GasUsed", receipt.GasUsed)
	t.Log("receipt.Logs", len(receipt.Logs))

	// 3. 결과 조회
	// 3-1. Counter 의 값이 증가했는지 확인한다.
	eoaCounter := counterABI.Instance(backend.Client(), backend.Owner().From)
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)

	// 4. Counter 는 해당 트랜잭션에서만 동작해야 하는데?
	nonce, err = backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	address := backend.Owner().From
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: new(big.Int).Add(gasTipCap, big.NewInt(2e9)),
		Gas:       1e6,
		To:        &address, // 본인한테 보내는 트랜잭션
		Value:     common.Big0,
		Data:      data,
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// 5. 일단 동작을 했으니 정상적으로 적용되었는지 확인해볼까
	// 5-1. Counter 의 값이 증가한다...?
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big2, count)

	// 5-2. DelegateCall 이니깐 배포된 컨트랙트의 number 는 그대로 1이다.
	bytes, err = counter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)
}

// TestEIP7702Delegation2 eth_getCode 의 결과가 ParseDelegation true 인 경우에는 빈 배열을 리턴한다고 한다.
// ANVIL 도 팩트라 하드포크가 아직 없어서 ethclient으로 붙어보자.
// 다를께 없다아...
func TestEIP7702Delegation2(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	client, err := ethclient.Dial(backend.IPCPath())
	require.NotNil(t, backend)
	chainId, err := client.ChainID(ctx)
	require.NoError(t, err)

	counterAddress, counterABI, counter := utils.DeployCounter(t, backend, backend.Owner())

	/// 1. Counter Check
	// 1 을 증가시켜본다.
	tx, err := counter.RawTransact(backend.Owner(), counterABI.PackIncrement())
	require.NoError(t, err)
	backend.Commit()
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// 1-1. 1이 증가되었는지 확인한다.
	bytes, err := counter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err := counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)

	// 2. EIP-7702 트랜잭션 실행
	// 2-1. 코드를 등록할 주소에 실제로 코드가 있는지 확인한다. (Counter)
	code, err := client.PendingCodeAt(ctx, counterAddress)
	require.NoError(t, err)
	require.NotEmpty(t, code)
	// 2-2. 트랜잭션에 필요한 데이터중 체인에서 가져올수 있는 데이터를 가져온다.
	nonce, err := client.PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	require.NoError(t, err)
	// 2-3. 트랜잭션 생성에 필요한 데이터를 생성한다.
	prv := utils.GetPrivateKey(t, backend.Owner().From)
	// 2-3-1. Multicall3 의 코드를 등록하기위한 데이터 (auth)
	auth, err := types.SignSetCode(prv, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(chainId), // 체인 ID
		Address: counterAddress,                // Code가 있는 컨트랙트 주소
		Nonce:   nonce + 1,                     // Tx Nonce + 1
	})
	require.NoError(t, err)
	// 2-3-2. auth 가 올바르게 생성되었는지 확인한다.
	authority, err := auth.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)
	// 2-4. 트랜잭션 데이터
	// Counter 를 사용하여 1을 증가시킨다.
	data := counterABI.PackIncrement()
	// 2-5. 트랜잭션을 생성 및 실행한다.
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(chainId),
		Nonce:     nonce,
		GasTipCap: uint256.MustFromBig(gasTipCap),
		GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
		Gas:       1e6,
		To:        backend.Owner().From, // 반드시 본인의 EOA 주소여야 한다.
		Value:     uint256.NewInt(0),
		Data:      data,
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	require.NoError(t, err)
	require.NoError(t, client.SendTransaction(ctx, tx))
	backend.Commit()

	receipt, err = client.TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Log("receipt.GasUsed", receipt.GasUsed)
	t.Log("receipt.Logs", len(receipt.Logs))

	// 3. 결과 조회
	// 3-1. Counter 의 값이 증가했는지 확인한다.
	eoaCounter := counterABI.Instance(client, backend.Owner().From)
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)

	// 4. Counter 는 해당 트랜잭션에서만 동작해야 하는데?
	nonce, err = client.PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	address := backend.Owner().From
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: new(big.Int).Add(gasTipCap, big.NewInt(2e9)),
		Gas:       1e6,
		To:        &address, // 본인한테 보내는 트랜잭션
		Value:     common.Big0,
		Data:      data,
	})
	require.NoError(t, err)
	require.NoError(t, client.SendTransaction(ctx, tx))
	backend.Commit()
	receipt, err = client.TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// 5. 일단 동작을 했으니 정상적으로 적용되었는지 확인해볼까
	// 5-1. Counter 의 값이 증가한다...?
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big2, count)

	// 5-2. DelegateCall 이니깐 배포된 컨트랙트의 number 는 그대로 1이다.
	bytes, err = counter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)
}

// TestEIP7702MultiContracts 두개 이상의 컨트랙트 코드를 등록하면 어떻게 동작 할까?
// 각각의 함수 signature 에 맞게 동작할것 같다.
func TestEIP7702MultiContracts(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	chainId, err := backend.Client().ChainID(ctx)
	require.NoError(t, err)

	erc20Address, erc20ABI, erc20 := utils.DeployMockERC20(t, backend, backend.Owner(), "MockERC20", "MERC", 18)
	counterAddress, counterABI, counter := utils.DeployCounter(t, backend, backend.Owner())

	/// 1. Contracts Check
	// 1-1. (ERC20) 을 민팅해 본다.
	tx, err := erc20.RawTransact(backend.Owner(), erc20ABI.PackMint(backend.Owner().From, common.Big1))
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// 1-1-1. 1이 민팅되었는지 확인한다.
	bytes, err := erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err := erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)

	// 1-2. (Counter) 을 증가시켜본다.
	tx, err = counter.RawTransact(backend.Owner(), counterABI.PackIncrement())
	require.NoError(t, err)
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// 1-2-1. 1이 증가되었는지 확인한다.
	bytes, err = counter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err := counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, count)

	/// 2. EIP-7702 트랜잭션 실행
	// 2-1. 코드를 등록할 주소에 실제로 코드가 있는지 확인한다. (Multicall3)
	code, err := backend.Client().PendingCodeAt(ctx, utils.Multicall3Address())
	require.NoError(t, err)
	require.NotEmpty(t, code)
	// // 2-2. 코드를 등록할 주소에 실제로 코드가 있는지 확인한다. (Counter)
	code, err = backend.Client().PendingCodeAt(ctx, counterAddress)
	require.NoError(t, err)
	require.NotEmpty(t, code)

	// 2-2. 트랜잭션에 필요한 데이터중 체인에서 가져올수 있는 데이터를 가져온다.
	nonce, err := backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	gasTipCap, err := backend.Client().SuggestGasTipCap(ctx)
	require.NoError(t, err)

	// 2-3. 트랜잭션 생성에 필요한 데이터를 생성한다.
	prv, spender, recipient := utils.GetPrivateKey(t, backend.Owner().From), common.BytesToAddress([]byte("SPENDER")), common.BytesToAddress([]byte("RECIPIENT"))
	// 2-3-1. Multicall3 의 코드를 등록하기위한 데이터 (auth)
	auth1, err := types.SignSetCode(prv,
		types.SetCodeAuthorization{
			ChainID: *uint256.MustFromBig(chainId), // 체인 ID
			Address: utils.Multicall3Address(),     // Code가 있는 컨트랙트 주소
			Nonce:   nonce + 1,                     // Tx Nonce + 1
		})
	require.NoError(t, err)
	// 2-3-2. auth 가 올바르게 생성되었는지 확인한다.
	authority, err := auth1.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)
	// 2-3-3. Counter 의 코드를 등록하기위한 데이터 (auth)
	auth2, err := types.SignSetCode(prv, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(chainId), // 체인 ID
		Address: counterAddress,                // Code가 있는 컨트랙트 주소
		Nonce:   nonce + 2,                     // Tx Nonce + 2
	})
	require.NoError(t, err)
	// 2-3-2. auth 가 올바르게 생성되었는지 확인한다.
	authority, err = auth2.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)

	// 2-4. 트랜잭션 데이터
	// Multicall3 를 사용하여 ERC20 의 각종 함수를 실행한다.
	data := binding.NewIMulticall3().PackAggregate3([]binding.IMulticall3Call3{
		// burn (1 case 에서 민팅을 했기때문에, 해당 벨런스 제거 가능)
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackBurn(backend.Owner().From, common.Big1),
		},
		// mint
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackMint(backend.Owner().From, common.Big2),
		},
		// transfer
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackTransfer(recipient, common.Big1),
		},
		// approve
		{
			Target:       erc20Address,
			AllowFailure: true,
			CallData:     erc20ABI.PackApprove(spender, common.Big256),
		},
	})
	// 2-5. 트랜잭션을 생성 및 실행한다.
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(chainId),
		Nonce:     nonce,
		GasTipCap: uint256.MustFromBig(gasTipCap),
		GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
		Gas:       1e6,
		To:        backend.Owner().From, // 반드시 본인의 EOA 주소여야 한다.
		Value:     uint256.NewInt(0),
		Data:      data,
		AuthList:  []types.SetCodeAuthorization{auth2, auth1},
		// auth1: Multicall3
		// auth2: Counter
		// {auth1, auth2} 으로 실행했을때에는 이 트랜잭션에서 실패했다.
		// {auth2, auth1} 으로 실행했을때에는 이 트랜잭션에서 성공했다.
		// Data 에 들어가는 동작은 나중에 등록된 값으로 사용하는것 같다.
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()

	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Log("receipt.GasUsed", receipt.GasUsed)
	t.Log("receipt.Logs", len(receipt.Logs))
	for _, log := range receipt.Logs {
		if l, err := erc20ABI.UnpackTransferEvent(log); err == nil {
			t.Logf("[TRANSFER] from: %s, to: %s, value: %s", l.From.Hex(), l.To.Hex(), l.Amount.String())
		} else {
			l, err := erc20ABI.UnpackApprovalEvent(log)
			require.NoError(t, err)
			t.Logf("[APPROVAL] owner: %s, spender: %s, value: %s", l.Owner.Hex(), l.Spender.Hex(), l.Amount.String())
		}
	}

	/// 3. 결과 조회
	// 3-1. Balance: mint(1) -> burn(1) -> mint(2) -> transfer(1) = 1
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 3-2. recipient 의 balance = 1
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(recipient))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 3-3. allowance = 256
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackAllowance(backend.Owner().From, spender))
	require.NoError(t, err)
	allowance, err := erc20ABI.UnpackAllowance(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big256, allowance)

	// 4-1. EOA 주소에는 Multicall3 의 코드가 남아있다
	code, err = backend.Client().PendingCodeAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	delegation, ok := types.ParseDelegation(code)
	require.True(t, ok)
	require.Equal(t, utils.Multicall3Address(), delegation)
	// 남아있는 multicall3 코드 한번더 실행
	nonce, err = backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	address := backend.Owner().From
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: new(big.Int).Add(gasTipCap, big.NewInt(2e9)),
		Gas:       1e6,
		To:        &address, // 본인한테 보내는 트랜잭션
		Value:     common.Big0,
		Data:      data,
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	t.Log("==============================================================================================================================")
	for _, log := range receipt.Logs {
		if l, err := erc20ABI.UnpackTransferEvent(log); err == nil {
			t.Logf("[TRANSFER] from: %s, to: %s, value: %s", l.From.Hex(), l.To.Hex(), l.Amount.String())
		} else {
			l, err := erc20ABI.UnpackApprovalEvent(log)
			require.NoError(t, err)
			t.Logf("[APPROVAL] owner: %s, spender: %s, value: %s", l.Owner.Hex(), l.Spender.Hex(), l.Amount.String())
		}
	}
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)
	// 5-2. recipient 의 balance = 2
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(recipient))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big2, balance)
	// 5-3. allowance = 256
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackAllowance(backend.Owner().From, spender))
	require.NoError(t, err)
	allowance, err = erc20ABI.UnpackAllowance(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big256, allowance)

	// 5-4. Counter 의 값을 증가시켜 본다. (2 에서 일괄로 추가했기 때문에 동작할껏 같다.)
	// auth2 가 등록이 되어있을줄 알았지만 아래 테스트 케이스는 실패한다.
	// 아무래도 auth1 으로 완전히 덮어씌워지고 auth2(Counter) 는 사라진것 같다.
	eoaCounter := counterABI.Instance(backend.Client(), backend.Owner().From)
	tx, err = eoaCounter.RawTransact(backend.Owner(), counterABI.PackIncrement())
	require.Error(t, err)
	// backend.Commit()
	// receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	// require.NoError(t, err)
	// require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// // 5-4-1. 1이 증가되었는지 확인한다.
	// bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	// require.NoError(t, err)
	// count, err = counterABI.UnpackNumber(bytes)
	// require.NoError(t, err)
	// require.Equal(t, common.Big1, count)

	// 6. 그렇다면 Counter 의 코드를 다시 등록해서 사용해보자
	nonce, err = backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	auth2.Nonce = nonce + 1
	auth2, err = types.SignSetCode(prv, auth2)
	require.NoError(t, err)
	// 6-1. auth2 가 올바르게 생성되었는지 확인한다.
	authority, err = auth2.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)
	// 6-2. 트랜잭션 데이터
	tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
		ChainID:   uint256.MustFromBig(chainId),
		Nonce:     nonce,
		GasTipCap: uint256.MustFromBig(gasTipCap),
		GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
		Gas:       1e6,
		To:        backend.Owner().From, // 반드시 본인의 EOA 주소여야 한다.
		Value:     uint256.NewInt(0),
		Data:      counterABI.PackSetNumber(common.Big32),
		AuthList:  []types.SetCodeAuthorization{auth2},
	})
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	// code 등록과 동시에 32 으로 설정했다.
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big32, count)

	// 7. 32로 설정된 Counter 의 값을 증가시켜 본다.
	tx, err = eoaCounter.RawTransact(backend.Owner(), counterABI.PackIncrement())
	require.NoError(t, err)
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// 7-1. 1이 증가되었는지 확인한다.
	bytes, err = eoaCounter.CallRaw(callOpts, counterABI.PackNumber())
	require.NoError(t, err)
	count, err = counterABI.UnpackNumber(bytes)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(33), count)
}

func TestEIP7702DelegateTxFee(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	chainId, err := backend.Client().ChainID(ctx)
	require.NoError(t, err)

	// Counter, ERC20 컨트랙트 배포
	counterAddress, counterABI, counter := utils.DeployCounter(t, backend, backend.Owner())
	erc20Address, erc20ABI, erc20 := utils.DeployMockERC20(t, backend, backend.Owner(), "MockERC20", "MERC", 18)

	// EntryPoint 컨트랙트 확인
	entryPointAddress, entryPointABI, entryPoint := utils.EntryPoint(backend)
	code, err := backend.Client().PendingCodeAt(ctx, entryPointAddress)
	require.NoError(t, err)
	require.NotEmpty(t, code)

	// ERC20Paymaster 컨트랙트 배포
	gasPrice, err := backend.Client().SuggestGasPrice(ctx)
	require.NoError(t, err)
	erc20PaymasterABI := binding.NewERC20Paymaster()
	erc20PaymasterAddress, tx, err := bind.DeployContract(
		backend.Owner(),
		common.Hex2Bytes(binding.ERC20PaymasterMetaData.Bin[2:]),
		backend.Client(),
		erc20PaymasterABI.PackConstructor(
			entryPointAddress,
			erc20Address,
			gasPrice,
		),
	)
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Run("deposit to erc20Paymaster", func(t *testing.T) {
		txOpts := backend.Owner()
		txOpts.Value = new(big.Int).Mul(common.Big256, big.NewInt(params.Ether))
		tx, err = erc20PaymasterABI.Instance(backend.Client(), erc20PaymasterAddress).RawTransact(txOpts, erc20PaymasterABI.PackDeposit())
		require.NoError(t, err)
		backend.Commit()
		receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		txOpts.Value = nil

		// erc20Paymaster 컨트랙트에 입금하면 erc20Paymaster 는 entryPoint 에 deposit 한다.
		balance, err := backend.Client().PendingBalanceAt(ctx, erc20PaymasterAddress)
		require.NoError(t, err)
		require.True(t, balance.Sign() == 0)

		balance, err = backend.Client().PendingBalanceAt(ctx, entryPointAddress)
		require.NoError(t, err)
		require.Equal(t, new(big.Int).Mul(common.Big256, big.NewInt(params.Ether)), balance)
	})

	// simple7702Account 컨트랙트 배포
	simple7702AccountABI := binding.NewSimple7702Account()
	simple7702AccountAddress, tx, err := bind.DeployContract(backend.Owner(), common.Hex2Bytes(binding.Simple7702AccountMetaData.Bin[2:]), backend.Client(), []byte{})
	require.NoError(t, err)
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	simple7702Account := simple7702AccountABI.Instance(backend.Client(), simple7702AccountAddress)

	// simple7702Account 컨트랙트 배포 확인
	bytes, err := simple7702Account.CallRaw(callOpts, simple7702AccountABI.PackEntryPoint())
	require.NoError(t, err)
	callEntryPointAddress, err := simple7702AccountABI.UnpackEntryPoint(bytes)
	require.NoError(t, err)
	require.Equal(t, entryPointAddress, callEntryPointAddress)

	// 테스트 계정 생성
	eoa := backend.GetEoas(t, 1)[0]
	prv := utils.GetPrivateKey(t, eoa.From)
	// 테스트 계정은 Simple7702Account 코드를 등록 해야 한다. (이때 gas 비용이 발생한다.)
	t.Run("SetCodeTx", func(t *testing.T) {
		balance, err := backend.Client().PendingBalanceAt(ctx, eoa.From)
		require.NoError(t, err)
		nonce, err := backend.Client().PendingNonceAt(ctx, eoa.From)
		require.NoError(t, err)
		auth, err := types.SignSetCode(prv, types.SetCodeAuthorization{
			ChainID: *uint256.MustFromBig(chainId), // 체인 ID
			Address: simple7702AccountAddress,      // Code가 있는 컨트랙트 주소
			Nonce:   nonce + 1,                     // Tx Nonce + 1
		})
		require.NoError(t, err)
		tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
			ChainID:   uint256.MustFromBig(chainId),
			Nonce:     nonce,
			GasTipCap: uint256.MustFromBig(big.NewInt(1e9)),
			GasFeeCap: uint256.MustFromBig(new(big.Int).Add(big.NewInt(1e9), big.NewInt(2e9))),
			Gas:       1e6,
			To:        eoa.From, // 반드시 본인의 EOA 주소여야 한다.
			Value:     uint256.NewInt(0),
			Data:      simple7702AccountABI.PackExecute(erc20Address, common.Big0, erc20ABI.PackApprove(erc20PaymasterAddress, math.MaxBig256)),
			AuthList:  []types.SetCodeAuthorization{auth},
		})
		require.NoError(t, err)
		require.NoError(t, backend.Client().SendTransaction(ctx, tx))
		backend.Commit()
		receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		afterBalance, err := backend.Client().PendingBalanceAt(ctx, eoa.From)
		require.NoError(t, err)
		require.True(t, balance.Cmp(afterBalance) > 0)
	})

	// erc20 토큰 민팅
	t.Run("Mint to EOA", func(t *testing.T) {
		tx, err = erc20.RawTransact(backend.Owner(), erc20ABI.PackMint(eoa.From, new(big.Int).Mul(common.Big256, big.NewInt(params.Ether))))
		require.NoError(t, err)
		backend.Commit()
		receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(eoa.From))
		require.NoError(t, err)
		balance, err := erc20ABI.UnpackBalanceOf(bytes)
		require.NoError(t, err)
		require.Equal(t, new(big.Int).Mul(common.Big256, big.NewInt(params.Ether)), balance)
	})
	t.Run("Make & Send UserOperation", func(t *testing.T) {
		var (
			opNonce *big.Int
			gasfee  [32]byte
		)
		bytes, err = entryPoint.CallRaw(callOpts, entryPointABI.PackGetNonce(eoa.From, common.Big0))
		require.NoError(t, err)
		opNonce, err = entryPointABI.UnpackGetNonce(bytes)
		require.NoError(t, err)

		gasfee, err = utils.MakeGasFees(ctx, backend.Client())
		require.NoError(t, err)

		var userOp binding.PackedUserOperation
		userOp, err = utils.SignPackedUserOperation(callOpts, entryPointABI, entryPoint, binding.PackedUserOperation{
			Sender:   eoa.From,
			Nonce:    opNonce,
			InitCode: []byte{},
			// Calldata
			// Simple7702Account 의 execute 함수에 전달할 calldata
			// function execute(address target, uint256 value, bytes calldata data) virtual external
			CallData: simple7702AccountABI.PackExecute(counterAddress, common.Big0, counterABI.PackIncrement()),
			// AccountGasLimits
			// abi.encodePacked(uint128(verificationGasLimit), uint128(callGasLimit))
			AccountGasLimits:   utils.MergeGasToBytes32(big.NewInt(1e6), big.NewInt(1e6)), // verificationGasLimit, callGasLimit
			PreVerificationGas: big.NewInt(1e6),
			// GasFees
			// abi.encodePacked(uint128(maxPriorityFeePerGas), uint128(maxFeePerGas))
			GasFees:          gasfee,
			PaymasterAndData: utils.MakePaymasterData(erc20PaymasterAddress, 1e6, 1e6),
			Signature:        []byte{},
		}, prv)
		require.NoError(t, err)

		// erc20PaymasterBalance 는 현재 erc20 을 들고있지 않다.
		var erc20PaymasterBalance *big.Int
		bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(erc20PaymasterAddress))
		require.NoError(t, err)
		erc20PaymasterBalance, err = erc20ABI.UnpackBalanceOf(bytes)
		require.NoError(t, err)
		require.True(t, erc20PaymasterBalance.Sign() == 0)

		txOpts := backend.GetEoas(t, 2)[1]
		require.NotEqual(t, eoa.From, txOpts.From)
		var (
			// paymaster 의 deposit amount 는 줄어든다.
			beforePaymasterDeposit, afterPaymasterDeposit binding.IStakeManagerDepositInfo
			// eoa 는 erc20 토큰을 지불한다.
			beforeErc20Balance, afterErc20Balance *big.Int
			// eoa 는 native token 으로 gas fee 를 지불하지 않는다.
			beforeTxOptsBalance, afterTxOptsBalance *big.Int

			// tx 는 실행되었다.
			beforeNumber, afterNumber *big.Int
		)
		beforeTxOptsBalance, err = backend.Client().PendingBalanceAt(ctx, txOpts.From)
		require.NoError(t, err)

		bytes, err = entryPoint.CallRaw(callOpts, entryPointABI.PackGetDepositInfo(erc20PaymasterAddress))
		require.NoError(t, err)
		beforePaymasterDeposit, err = entryPointABI.UnpackGetDepositInfo(bytes)
		require.NoError(t, err)

		bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(eoa.From))
		require.NoError(t, err)
		beforeErc20Balance, err = erc20ABI.UnpackBalanceOf(bytes)
		require.NoError(t, err)

		bytes, err = counter.CallRaw(callOpts, counterABI.PackNumber())
		require.NoError(t, err)
		beforeNumber, err = counterABI.UnpackNumber(bytes)
		require.NoError(t, err)

		// 트랜잭션은 eoa 가 아닌 계정으로 실행
		// 트랜잭션에서 erc20Paymaster 에 예치된 eth 는 txOpts에게 지급된다.?
		tx, err = entryPoint.RawTransact(txOpts, entryPointABI.PackHandleOps([]binding.PackedUserOperation{userOp}, txOpts.From))
		require.NoError(t, err)
		backend.Commit()
		receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		// Counter 의 Caller 는 eoa 이다.
		func() {
			for _, log := range receipt.Logs {
				if event, err := counterABI.UnpackCallerEvent(log); err == nil {
					require.Equal(t, eoa.From, event.Caller)
					return
				}
			}
			require.NoError(t, errors.New("no event found"))
		}()

		afterTxOptsBalance, err = backend.Client().PendingBalanceAt(ctx, txOpts.From)
		require.NoError(t, err)

		bytes, err = entryPoint.CallRaw(callOpts, entryPointABI.PackGetDepositInfo(erc20PaymasterAddress))
		require.NoError(t, err)
		afterPaymasterDeposit, err = entryPointABI.UnpackGetDepositInfo(bytes)
		require.NoError(t, err)

		bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(eoa.From))
		require.NoError(t, err)
		afterErc20Balance, err = erc20ABI.UnpackBalanceOf(bytes)
		require.NoError(t, err)

		bytes, err = counter.CallRaw(callOpts, counterABI.PackNumber())
		require.NoError(t, err)
		afterNumber, err = counterABI.UnpackNumber(bytes)
		require.NoError(t, err)

		// tx 는 실행되었다.
		require.Equal(t, new(big.Int).Add(beforeNumber, common.Big1), afterNumber)
		// paymaster 의 deposit amount 는 줄어든다.
		require.True(t, afterPaymasterDeposit.Deposit.Cmp(beforePaymasterDeposit.Deposit) < 0)
		// eoa 는 erc20 토큰을 지불한다.
		require.True(t, afterErc20Balance.Cmp(beforeErc20Balance) < 0)
		// eoa 가 지불한 가스비(erc20)은 erc20Paymaster 으로 이동된다.
		bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(erc20PaymasterAddress))
		require.NoError(t, err)
		erc20PaymasterBalance, err = erc20ABI.UnpackBalanceOf(bytes)
		require.NoError(t, err)
		require.Equal(t, new(big.Int).Sub(beforeErc20Balance, afterErc20Balance), erc20PaymasterBalance)

		subBalance := new(big.Int).Sub(beforeTxOptsBalance, afterTxOptsBalance)
		subDeposit := new(big.Int).Sub(beforePaymasterDeposit.Deposit, afterPaymasterDeposit.Deposit)
		spendGas := new(big.Int).Add(subBalance, subDeposit)
		usedGas := new(big.Int).Mul(receipt.EffectiveGasPrice, new(big.Int).SetUint64(receipt.GasUsed))
		require.Equal(t, spendGas, usedGas)
	})
}
