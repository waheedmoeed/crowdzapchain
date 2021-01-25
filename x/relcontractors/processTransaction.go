package relcontractors

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/rand"
	"time"
)

//change contractor address with new provided address and add old one to other addresses
func handleMsgUpdateReContractorAddress(ctx sdk.Context, k Keeper, msg MsgUpdateRelContractorAddress) (*sdk.Result, error) {
	err := k.UpdateContractorByAddress(ctx, msg.RelContractorAddress, msg.NewRelContractorAddress)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

//Create new voting poll in contract
func handleMsgCreatePoll(ctx sdk.Context, k Keeper, msg MsgCreatePoll) (*sdk.Result, error) {
	contract, latestPoll, err := k.GetLatestPollAndContract(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to read relContract from store")
	}
	//Check latest poll if it is still valid( end-time not reached or not processed )
	if !(latestPoll.PollId == "") {
		if latestPoll.EndTime.Equal(time.Now()) || !latestPoll.Processed {
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("poll creation", 234, "Poll Creation"), "Already another poll is valid")
		}
	}
	//generate unique ID for every voting poll
	b := make([]byte, 22)
	for i := 0; i < 22; i++ {
		b[i] = byte(97 + rand.Intn(122-97))
	}
	//check validity of:
	//start(must be greater than now time)
	//end time (must be greater than 24 hour )
	if msg.StartTime.Before(time.Now()) || msg.EndTime.Before(time.Now().Add(time.Hour*24)) {
		return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("poll creation", 234, "Poll Creation"), "Time of poll start or end invalid")
	}
	poll := VotingPoll{
		PollId:               string(b),
		PollType:             msg.PollType,
		StartTime:            msg.StartTime,
		EndTime:              msg.EndTime,
		PositiveVotes:        0,
		NegativeVotes:        0,
		PositiveVotesAddress: []sdk.AccAddress{},
		NegativeVotesAddress: []sdk.AccAddress{},
		OwnerVoterPoll:       msg.OwnerVoterPoll,
		Processed:            false,
		CoinsAmount:          msg.CoinsAmount,
	}
	newContract := append(contract.VotingPolls, poll)
	contract.VotingPolls = newContract
	k.Set(ctx, contract)
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
		if len(contractor.VotingPolls) > 0 {
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
					return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("vote poll", 235, "POll Voting"), "cannot vote in this poll, already expired")
				}
			}
		} else {
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("vote poll", 235, "POll Voting"), "there is no poll to vote")
		}
	}
	//TODO: Add events here
	return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("vote poll", 235, "POll Voting"), "voter is not contractor")
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
		if len(contract.VotingPolls) > 0 {
			for _, poll := range contract.VotingPolls {
				if poll.PollId == msg.PollId {
					//check if all contractor voted or time ended and check this poll not yet processed
					if checkVotesProcess(contract, poll) && !poll.Processed {
						error := processPoll(poll, contract, ctx, k)
						if error != nil {
							return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("process poll", 236, "Process Poll"), fmt.Sprintf("Failed to process poll %s %s", poll.PollId, error))
						}
						return &sdk.Result{Events: ctx.EventManager().Events()}, nil
					} else {
						return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("process poll", 236, "Process Poll"), "cannot process this poll, to process it all contractors must vote or it reaches its expiry time")
					}
				}
			}
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("process poll", 236, "Process Poll"), "failed to find any poll related to given id")
		} else {
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("process poll", 236, "Process Poll"), "there is no poll to process")
		}
	}
	//TODO: Add events here
	return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("process poll", 236, "Process Poll"), "only rel contractors can process poll")
}

///////////////////////////
/////HELPER METHODS///////
/////////////////////////

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
