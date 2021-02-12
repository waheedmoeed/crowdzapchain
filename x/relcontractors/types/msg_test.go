package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"testing"
	"time"
)

var name = "maTurtle"

func TestLOgic(t *testing.T) {
	coins := sdk.Coins{sdk.NewInt64Coin("rel", 100)}

	require.Equal(t, coins.AmountOf("rel"), 100)
}

func TestTime(t *testing.T) {
	fg := time.Now()
	fmt.Println(fg)
	endTime := time.Now().Add(time.Hour * 24 * 2)
	fmt.Println(endTime)
	require.Equal(t, true, true)
}

func TestGenerateNewAddress(t *testing.T) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	fmt.Println(addr.String())
	require.Equal(t, true, true)
}
