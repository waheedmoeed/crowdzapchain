package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

var name = "maTurtle"

func TestLOgic(t *testing.T) {
	coins := sdk.Coins{sdk.NewInt64Coin("rel", 100)}

	require.Equal(t, coins.AmountOf("rel"), 100)
}

func TestTime(t *testing.T) {
	bytes := make([]byte, 22)

	for i := 0; i < 22; i++ {
		bytes[i] = byte(97 + rand.Intn(122-97))
	}
	fmt.Println(string(bytes))
	require.Equal(t, true, true)
}
