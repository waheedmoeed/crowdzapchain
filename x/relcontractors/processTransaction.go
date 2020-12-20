package relcontractors

import (
	"crypto/rand"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

//change contractor address with new provided address and add old one to other addresses
func handleMsgUpdateReContractorAddress(ctx sdk.Context, k Keeper, msg MsgUpdateRelContractorAddress) (*sdk.Result, error) {
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
	//TODO: Add events here
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreatePoll(ctx sdk.Context, k Keeper, msg MsgCreatePoll) (*sdk.Result, error) {
	contractor, err := k.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to read relContract from store")
	}
	//generate unique ID for every voting poll
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
		Processed:            false,
		CoinsAmount:          msg.CoinsAmount,
	}
	newContract := append(contractor.VotingPolls, poll)
	contractor.VotingPolls = newContract
	k.Set(ctx, contractor)
	//TODO: Add events here
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgVotePoll(ctx sdk.Context, k Keeper, msg MsgVotePoll) (*sdk.Result, error) {
	contractor, err := k.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to read relContract from store")
	}
	//check either the voter is contractor or  not
	contractorFounded := false
	for _, contractor := range contractor.RelContractors {
		if msg.Voter.Equals(contractor.ContractorAddress) {
			contractorFounded = true
			break
		}
	}
	if contractorFounded {
		//find poll by given Id and validate is poll and voter valid to vote
		for _, poll := range contractor.VotingPolls {
			if poll.PollId == msg.PollId {
				//time check and validate that poll is not yet processed
				if checkPollValidity(poll.StartTime, poll.EndTime) && !poll.Processed {
					if msg.Vote == 0 {
						negativeVoters := append(poll.NegativeVotesAddress, msg.Voter)
						poll.NegativeVotes = poll.NegativeVotes + 1
						poll.NegativeVotesAddress = negativeVoters
					} else {
						positiveVoters := append(poll.PositiveVotesAddress, msg.Voter)
						poll.PositiveVotes = poll.PositiveVotes + 1
						poll.PositiveVotesAddress = positiveVoters
					}
					k.Set(ctx, contractor)
					//TODO: Add events
					return &sdk.Result{Events: ctx.EventManager().Events()}, nil
				}
				//TODO: Add events
				return &sdk.Result{}, sdkerrors.Wrap(errors.New("cannot vote in this poll, already expired"), "")
			}
		}
	}
	//TODO: Add events here
	return &sdk.Result{}, sdkerrors.Wrap(errors.New("voter is not contractor"), "")
}

func handleMsgProcessPoll(ctx sdk.Context, k Keeper, msg MsgProcessPoll) (*sdk.Result, error) {
	contract, err := k.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to read relContract from store")
	}
	contractorFounded := false
	for _, contractor := range contract.RelContractors {
		if msg.Transactor.Equals(contractor.ContractorAddress) {
			contractorFounded = true
			break
		}
	}
	if contractorFounded {
		for _, poll := range contract.VotingPolls {
			if poll.PollId == msg.PollId {
				//check if all contractor voted or time ended and check this poll not yet processed
				if checkVotesProcess(contract, poll) && !poll.Processed {
					error := processPoll(poll, contract, ctx, k)
					if error != nil {
						return &sdk.Result{}, sdkerrors.Wrap(error, "Failed to process poll "+poll.PollId)
					}
				}
				return &sdk.Result{}, sdkerrors.Wrap(errors.New("cannot process this poll, to process it all contractors must vote or it reaches its expiry time"), "")
			}
		}
		return &sdk.Result{}, sdkerrors.Wrap(errors.New("failed to find any poll related to given id"), "")
	}
	//TODO: Add events here
	return &sdk.Result{}, sdkerrors.Wrap(errors.New("only rel contractors can process poll"), "")
}

//check if poll is valid to be voted in terms of both start and end time
func checkPollValidity(start time.Time, end time.Time) bool {
	now := time.Now()
	//Vote must be casted before 5 min of expiry
	if end.Sub(now).Minutes() > 5 {
		//Vote must be casted after 5 min of start
		if now.Sub(start).Minutes() > 5 {
			return true
		}
	}
	return false
}

func checkVotesProcess(contractor RelContract, poll VotingPoll) bool {
	//check if all the contractor have voted yet or not
	numContractors := len(contractor.RelContractors)
	totalVotes := int(poll.PositiveVotes + poll.NegativeVotes)
	if numContractors == totalVotes {
		return true
	}
	//check if end time of poll has reached or not
	now := time.Now()
	if poll.EndTime.Sub(now) <= 0 {
		return true
	}
	return false
}
