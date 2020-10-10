package relcontractors

import (
	"crypto/rand"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//change contractor address with new provided address and add old one to other addresses
func handleMsgUpdateReContractorAddress(ctx sdk.Context, k Keeper, msg MsgUpdateRelContractorAddress) (*sdk.Result, error) {
	//Todo define logic to update contract and store it in db
	contractor, err := k.GetContractorByAddress(ctx, msg.RelContractorAddress)
	if err != nil {
		return nil, err
	}
	for _, value := range contractor.OtherAddresses {
		if value.Equals(msg.NewRelContractorAddress) {
			return nil, errors.New("address already present in contractor")
		}
	}
	addresses := append(contractor.OtherAddresses, msg.RelContractorAddress)
	//update contractor values
	contractor.OtherAddresses = addresses
	contractor.ContractorAddress = msg.NewRelContractorAddress
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreatePoll(ctx sdk.Context, k Keeper, msg MsgCreatePoll) (*sdk.Result, error) {
	//Todo define logic to update contract and store it in db
	contractor, err := k.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to read relContract from store")
	}
	//generate unique ID for every voting polll
	b := make([]byte, 16)
	_, error := rand.Read(b)
	if err != nil {
		return &sdk.Result{}, sdkerrors.Wrap(error, "Failed to generate new Id for voting poll")
	}
	poll := VotingPoll{
		PollId:               string(b),
		PollType:             msg.PollType,
		StartTime:            msg.StartTime,
		EndTime:              msg.EndTime,
		PositiveVotes:        0,
		NegativeVotes:        0,
		PositiveVotesAddress: nil,
		NegativeVotesAddress: nil,
		OwnerVoterPoll:       msg.OwnerVoterPoll,
	}
	newContract := append(contractor.VotingPolls, poll)
	contractor.VotingPolls = newContract
	k.Set(ctx, contractor)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
