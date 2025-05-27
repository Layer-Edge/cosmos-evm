package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/evm/x/precisebank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestFractionalBalanceKey(t *testing.T) {
	addr := sdk.AccAddress([]byte("edge-address"))

	key := types.FractionalBalanceKey(addr)
	require.Equal(t, addr.Bytes(), key)
	require.Equal(t, addr, sdk.AccAddress(key), "key should be able to be converted back to address")
}
