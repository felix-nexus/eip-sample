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
	erc20Bin := common.Hex2Bytes(strings.TrimPrefix(binding.MockERC20MetaData.Bin, "0x"))

	erc20ABI := binding.NewMockERC20()
	erc20Address, tx, err := bind.DeployContract(auth, erc20Bin, backend.Client(), erc20ABI.PackConstructor(name, symbol, decimals))
	require.NoError(t, err)
	backend.Commit()

	address, err := bind.WaitDeployed(ctx, backend.Client(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, erc20Address, address)

	require.NoError(t, err)
	return erc20Address, erc20ABI, erc20ABI.Instance(backend.Client(), address)
}
