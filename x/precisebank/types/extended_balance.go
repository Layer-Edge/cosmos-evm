package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SumExtendedCoin returns a sdk.Coin of extended coin denomination
// with all integer and fractional amounts combined. e.g. if amount contains
// both coins of integer denom and extended denom, this will return the total
// amount in extended coins. This is intended to get the full value to emit in
// events.
func SumExtendedCoin(amt sdk.Coins) sdk.Coin {
	// uedgens converted to aedgens
	integerAmount := amt.AmountOf(IntegerCoinDenom()).Mul(ConversionFactor())
	// aedgens as is
	extendedAmount := amt.AmountOf(ExtendedCoinDenom())

	// total of uedgens and aedgens amounts
	fullEmissionAmount := integerAmount.Add(extendedAmount)

	return sdk.NewCoin(
		ExtendedCoinDenom(),
		fullEmissionAmount,
	)
}
