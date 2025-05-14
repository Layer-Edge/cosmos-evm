package types

import (
	"cosmossdk.io/math"
)

// GasToken represents a token that can be used for gas fees
type GasToken struct {
	Denom        string
	ExchangeRate math.LegacyDec // Exchange rate relative to base token
}

// AllowedGasTokens is a map of allowed token denominations that can be used for gas fees
var AllowedGasTokens = map[string]GasToken{
	"tedgen": {
		Denom:        "tedgen",
		ExchangeRate: math.LegacyNewDec(1), // 1:1 exchange rate
	},
	"sedgen": {
		Denom:        "sedgen",
		ExchangeRate: math.LegacyNewDec(1), // 1:1 exchange rate
	},
}

// IsAllowedGasToken checks if a given denomination is allowed to be used for gas fees
func IsAllowedGasToken(denom string) bool {
	_, exists := AllowedGasTokens[denom]
	return exists
}

// GetGasTokenExchangeRate returns the exchange rate for a given gas token
func GetGasTokenExchangeRate(denom string) (math.LegacyDec, bool) {
	if token, exists := AllowedGasTokens[denom]; exists {
		return token.ExchangeRate, true
	}
	return math.LegacyDec{}, false
}
