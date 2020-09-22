package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var name = "maTurtle"

func TestLOgic(t *testing.T) {
	coinss := sdk.Coins{sdk.NewInt64Coin("rel", 100)}

	require.Equal(t, coinss.AmountOf("rel"), 100)
}
