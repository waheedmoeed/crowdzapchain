package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	Define all the types needed for transactions
*/

type MsgUpdateRelContractorAddress struct {
	RelContractorAddress    sdk.AccAddress `json:"rel_contractor_address"`
	NewRelContractorAddress sdk.AccAddress `json:"new_rel_contractor_address"`
}

type MsgCreateVotePoll struct {
	PollType       string         `json:"poll_type"`
	OwnerVoterPoll sdk.AccAddress `json:"owner_voter_poll"`
}

type MsgVotePoll struct {
	PollId string         `json:"poll_id"`
	Vote   uint           `json:"vote"`
	Voter  sdk.AccAddress `json:"voter"`
}

type MsgProcessPoll struct {
	PollId     string         `json:"poll_id"`
	Transactor sdk.AccAddress `json:"transactor"`
}
