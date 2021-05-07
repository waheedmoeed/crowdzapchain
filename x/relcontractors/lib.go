package relcontractors

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MINTCOIN         = "1"
	DISTRIBUTECOIN   = "2"
	ADDRELCONTRACTOR = "3"
)

////TODO:::: Recover from panic state when keeper set method called

func processPoll(poll VotingPoll, contract RelContract, ctx sdk.Context, k Keeper) error {
	var err error
	switch poll.PollType {
	case MINTCOIN:
		processMintPoll(poll, &contract) //  add/mint new coins in contract
		k.Set(ctx, contract)             // add new contract in store
		break
	case DISTRIBUTECOIN:
		//Check if there are enough coins to distribute to contractor
		if (contract.MintedCoins.Amount.Int64() - contract.DistributedCoins.Amount.Int64()) < poll.CoinsAmount.Amount.Int64() {
			err = errors.New("new Coins needed to be minted to distribute coins among nodes")
			break
		}
		//Add coins into poll_owner/contractor
		err = k.SendCoinsToContractor(ctx, poll.OwnerVoterPoll, poll.CoinsAmount)
		if err != nil {
			return err
		}
		processDistributionPoll(poll, &contract)
		k.Set(ctx, contract)
		break
	case ADDRELCONTRACTOR:
		err = processAddRelContractor(poll, ctx, k)
		break
	default:
		err = sdkerrors.Wrap(errors.New("cannot process this poll, invalid type of poll"), " at lib.go line 27")
		break
	}
	return err
}

func processMintPoll(poll VotingPoll, contract *RelContract) {
	contract.MintedCoins.Add(poll.CoinsAmount)
	records := append(contract.MintedCoinsRecord, CoinsMintedRecord{
		Coins:           poll.CoinsAmount,
		Date:            time.Now(),
		VoterContractor: poll.PositiveVotesAddress,
	})
	contract.MintedCoinsRecord = records
}

//Update contract status like distributedRecord and DistributedCoinsAmount
func processDistributionPoll(poll VotingPoll, contract *RelContract) {
	contract.DistributedCoins.Add(poll.CoinsAmount)
	records := append(contract.DistributedCoinsLogs, DistributedCoinsRecord{
		ContractorAddress:      poll.OwnerVoterPoll,
		Coins:                  poll.CoinsAmount,
		IssuedDate:             time.Now(),
		DistributedCoinsAmount: contract.DistributedCoins,
	})
	contract.DistributedCoinsLogs = records
}

func processAddRelContractor(poll VotingPoll, ctx sdk.Context, k Keeper) error {
	return nil
}
