package testutil

import (
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Coin(t testing.TB) sdk.Coin {
	t.Helper()
	return sdk.NewCoin("testcoin", sdk.NewInt(int64(rand.Intn(999)+1))) // nolint: gosec
}

// AkashCoin provides simple interface to the Akash sdk.Coin type.
func AkashCoinRandom(t testing.TB) sdk.Coin {
  t.Helper()
  amt := sdk.NewInt(int64(rand.Intn(999)+1))
  return sdk.NewCoin(CoinDenom, amt)
}

// AkashCoin provides simple interface to the Akash sdk.Coin type.
func AkashCoin(t testing.TB, amount int64) sdk.Coin {
	t.Helper()
	amt := sdk.NewInt(amount)
	return sdk.NewCoin(CoinDenom, amt)
}
