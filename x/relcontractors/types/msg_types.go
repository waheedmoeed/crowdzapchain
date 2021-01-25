package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

/*
	Define all the types needed for transactions
*/

type MsgUpdateRelContractorAddress struct {
	RelContractorAddress    sdk.AccAddress `json:"rel_contractor_address"`
	NewRelContractorAddress sdk.AccAddress `json:"new_rel_contractor_address"`
}

type MsgCreatePoll struct {
	PollType       uint           `json:"poll_type"`
	StartTime      time.Time      `json:"start_time"`
	EndTime        time.Time      `json:"end_time"`
	OwnerVoterPoll sdk.AccAddress `json:"owner_voter_poll"`
	CoinsAmount    sdk.Coin       `json:"coins_amount"`
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
