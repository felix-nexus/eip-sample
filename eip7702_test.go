package eip_sample

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	binding "github.com/felix-nexus/eip-sample/contracts/binding/go/src"
	"github.com/felix-nexus/eip-sample/utils"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	ctx      = context.Background()
	callOpts = &bind.CallOpts{Context: ctx}
	eip7702  = true
)

func TestEIP7702(t *testing.T) {
	backend := utils.NewBackend(t)
	require.NotNil(t, backend)
	chainId, err := backend.Client().ChainID(ctx)
	require.NoError(t, err)

	erc20Address, erc20ABI, erc20 := utils.DeployMockERC20(t, backend, backend.Owner(), "MockERC20", "MERC", 18)
	_, _, _ = erc20Address, erc20ABI, erc20

	tx, err := erc20.RawTransact(backend.Owner(), erc20ABI.PackMint(backend.Owner().From, common.Big1))
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	bytes, err := erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err := erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big1, balance)

	nonce, err := backend.Client().PendingNonceAt(ctx, backend.Owner().From)
	require.NoError(t, err)
	gasTipCap, err := backend.Client().SuggestGasTipCap(ctx)
	require.NoError(t, err)

	code, err := backend.Client().PendingCodeAt(ctx, utils.Multicall3Address())
	require.NoError(t, err)
	require.NotEmpty(t, code)
	prv, spender := utils.GetPrivateKey(t, backend.Owner().From), common.BytesToAddress([]byte("SPENDER"))
	auth, err := types.SignSetCode(prv, types.SetCodeAuthorization{
		ChainID: *uint256.MustFromBig(chainId),
		Address: utils.Multicall3Address(),
		Nonce:   nonce + 1,
	})
	require.NoError(t, err)
	authority, err := auth.Authority()
	require.NoError(t, err)
	require.Equal(t, backend.Owner().From, authority)

	data := binding.NewIMulticall3().PackAggregate3([]binding.IMulticall3Call3{
		// burn
		{
			Target:       erc20Address,
			AllowFailure: false,
			CallData:     erc20ABI.PackBurn(backend.Owner().From, common.Big1),
		},
		// mint
		{
			Target:       erc20Address,
			AllowFailure: false,
			CallData:     erc20ABI.PackMint(backend.Owner().From, common.Big2),
		},
		// approve
		{
			Target:       erc20Address,
			AllowFailure: false,
			CallData:     erc20ABI.PackApprove(spender, common.Big256),
		},
	})
	var owner common.Address
	if !eip7702 {
		owner = utils.Multicall3Address()
		multicall3 := utils.Multicall3Address()
		tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.DynamicFeeTx{
			ChainID:   chainId,
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: new(big.Int).Add(gasTipCap, big.NewInt(2e9)),
			Gas:       1e6,
			To:        &multicall3,
			Value:     nil,
			Data:      data,
		})
	} else {
		owner = backend.Owner().From
		tx, err = types.SignNewTx(prv, types.NewPragueSigner(chainId), &types.SetCodeTx{
			ChainID:   uint256.MustFromBig(chainId),
			Nonce:     nonce,
			GasTipCap: uint256.MustFromBig(gasTipCap),
			GasFeeCap: uint256.MustFromBig(new(big.Int).Add(gasTipCap, big.NewInt(2e9))),
			Gas:       1e6,
			To:        backend.Owner().From,
			Value:     uint256.NewInt(0),
			Data:      data,
			AuthList:  []types.SetCodeAuthorization{auth},
		})
	}
	require.NoError(t, err)
	require.NoError(t, backend.Client().SendTransaction(ctx, tx))
	backend.Commit()

	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	t.Log("receipt.GasUsed", receipt.GasUsed)
	t.Log("receipt.Logs", len(receipt.Logs))

	// check balance
	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackBalanceOf(backend.Owner().From))
	require.NoError(t, err)
	balance, err = erc20ABI.UnpackBalanceOf(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big2, balance)

	bytes, err = erc20.CallRaw(callOpts, erc20ABI.PackAllowance(owner, spender))
	require.NoError(t, err)
	allowance, err := erc20ABI.UnpackAllowance(bytes)
	require.NoError(t, err)
	require.Equal(t, common.Big256, allowance)
}
