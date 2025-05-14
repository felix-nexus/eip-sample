package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	binding "github.com/felix-nexus/eip-sample/contracts/binding/go/src"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func DeployMockERC20(t *testing.T, backend *Backend, auth *bind.TransactOpts, name, symbol string, decimals uint8) (common.Address, *binding.MockERC20, *bind.BoundContract) {
	ctx := context.Background()
	bin := common.Hex2Bytes(strings.TrimPrefix(binding.MockERC20MetaData.Bin, "0x"))

	ABI := binding.NewMockERC20()
	addr, tx, err := bind.DeployContract(auth, bin, backend.Client(), ABI.PackConstructor(name, symbol, decimals))
	require.NoError(t, err)
	backend.Commit()

	address, err := bind.WaitDeployed(ctx, backend.Client(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, addr, address)

	require.NoError(t, err)
	return addr, ABI, ABI.Instance(backend.Client(), address)
}

func DeployCounter(t *testing.T, backend *Backend, auth *bind.TransactOpts) (common.Address, *binding.Counter, *bind.BoundContract) {
	ctx := context.Background()
	bin := common.Hex2Bytes(strings.TrimPrefix(binding.CounterMetaData.Bin, "0x"))

	ABI := binding.NewCounter()
	addr, tx, err := bind.DeployContract(auth, bin, backend.Client(), []byte{})
	require.NoError(t, err)
	backend.Commit()

	address, err := bind.WaitDeployed(ctx, backend.Client(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, addr, address)

	require.NoError(t, err)
	return addr, ABI, ABI.Instance(backend.Client(), address)
}
