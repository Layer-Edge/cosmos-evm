package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	teststaking "github.com/cosmos/cosmos-sdk/x/staking/testutil"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// CreateValidator creates a validator with the given amount of staked tokens in the bond denomination set
// in the staking keeper.
func CreateValidator(t *testing.T, ctx sdk.Context, pubKey cryptotypes.PubKey, sk stakingkeeper.Keeper, stakeAmt math.Int) {
	t.Helper()
	zeroDec := math.LegacyZeroDec()
	stakingParams, err := sk.GetParams(ctx)
	require.NoError(t, err)
	stakingParams.BondDenom, err = sk.BondDenom(ctx)
	require.NoError(t, err)
	stakingParams.MinCommissionRate = zeroDec
	err = sk.SetParams(ctx, stakingParams)
	require.NoError(t, err)

	stakingHelper := teststaking.NewHelper(t, ctx, &sk)
	stakingHelper.Commission = stakingtypes.NewCommissionRates(zeroDec, zeroDec, zeroDec)
	stakingHelper.Denom, err = sk.BondDenom(ctx)
	require.NoError(t, err)

	valAddr := sdk.ValAddress(pubKey.Address())
	stakingHelper.CreateValidator(valAddr, pubKey, stakeAmt, true)
}
