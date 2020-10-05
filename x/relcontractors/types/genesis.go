package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// GenesisState - all relcontractors state that must be provided at genesis
type GenesisState struct {
	RelContract RelContract `json:"rel_contract"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	contractors := make([]Contractor, 3)
	for i := 0; i < 3; i++ {
		contractors[i] = Contractor{
			ContractorAddress: getNewAccount(),
			Name:              "node1",
			OtherAddresses:    nil,
		}
	}
	return createDefaultState(10000, 2000, contractors)
}

func createDefaultState(mintedCoins int, distributedCoins int, relContractors []Contractor) GenesisState {
	return GenesisState{
		RelContract: RelContract{
			MintedCoins:          sdk.Coins{sdk.NewInt64Coin("rel", int64(mintedCoins))},
			MintedCoinsRecord:    nil,
			DistributedCoins:     sdk.Coins{sdk.NewInt64Coin("rel", int64(distributedCoins))},
			DistributedCoinsLogs: nil, //todo: add logs/records of initial distribution, if needed.
			RelContractors:       relContractors,
			VotingPolls:          nil,
		},
	}
}

func getNewAccount() (address sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return addr
}

// ValidateGenesis validates the relcontractors genesis parameters
func ValidateGenesis(data GenesisState) error {
	if data.RelContract.MintedCoins.AmountOf("rel").Int64() < int64(100) {
		return fmt.Errorf("invalid MintedCoins: Value: %s. Error: Must be greater than 100", data.RelContract.MintedCoins.String())
	}
	if data.RelContract.DistributedCoins.AmountOf("rel").Int64() > data.RelContract.MintedCoins.AmountOf("rel").Int64() {
		return fmt.Errorf("invalid DistributedCoins: Value: %s. Error: Must be less than MintedCoins", data.RelContract.DistributedCoins.String())
	}
	if len(data.RelContract.RelContractors) < 3 {
		return fmt.Errorf("invalid RelContractors:. Error:Number of RelContractors must be greater than 3")
	}
	if error := checkDistributionAddress(data); error != nil {
		return error
	}
	if data.RelContract.MintedCoinsRecord != nil {
		return fmt.Errorf("invalid RelContractors:. Error:Minted Coins record must be empty")
	}
	if data.RelContract.VotingPolls != nil {
		return fmt.Errorf("invalid RelContractors:. Error:Voting Polls record must be empty")
	}
	return nil
}

//check distribution address must be in relContractors
func checkDistributionAddress(data GenesisState) error {
	if data.RelContract.DistributedCoinsLogs != nil {
		for _, value := range data.RelContract.DistributedCoinsLogs {
			founded := false
			for _, contractor := range data.RelContract.RelContractors {
				if value.ContractorAddress.Equals(contractor.ContractorAddress) {
					founded = true
					break
				}
			}
			if founded == false {
				return fmt.Errorf("% :mismatch must be in both relcontractor and distributed record", value)
			}
		}
	}
	return nil
}
