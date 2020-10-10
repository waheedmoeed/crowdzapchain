package relcontractors

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, authKeeper auth.AccountKeeper, data GenesisState) {
	// TODO: Define logic for when you would like to initialize a new genesis
	if validateRelDistribution(data, authKeeper, ctx) {
		k.InitContract(ctx, data.RelContract)
	}
}

//validate that coins distributed to RelContractors are same defined in RelDistributionLogs
func validateRelDistribution(data GenesisState, authKeeper auth.AccountKeeper, ctx sdk.Context) bool {
	for _, value := range data.RelContract.DistributedCoinsLogs {
		if authKeeper.GetAccount(ctx, value.ContractorAddress).GetCoins().AmountOf("rel").Int64() != value.Coins.Amount.Int64() {
			panic("Amount of coins must be equal in both account and log at \"validateRelDistribution\" for address: " + value.ContractorAddress.String())
		}
	}
	return true
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	state, err := k.Get(ctx)
	if err != nil {
		panic("Failed to export genesis state in rel contractor state")
	}
	genesisState := NewGenesisState()
	genesisState.RelContract = state
	return genesisState
}
