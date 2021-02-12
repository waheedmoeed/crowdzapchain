package smartcontracts

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, contract := range data.BasicContracts {
		k.SetBasicContract(ctx, contract.Address, contract)
	}

	for _, contract := range data.YieldContracts {
		k.SetYieldContract(ctx, contract.Address, contract)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	basicContracts, err := k.GetAllBasicContracts(ctx)
	if err != nil {
		fmt.Println(err[0])
		panic("Failed to export genesis state in rel contractor state")
	}

	yieldContracts, err := k.GetAllYieldContracts(ctx)
	if err != nil {
		fmt.Println(err[0])
		panic("Failed to export genesis state in rel contractor state")
	}

	genesisState := NewGenesisState()
	genesisState.BasicContracts = basicContracts
	genesisState.YieldContracts = yieldContracts

	return genesisState
}
