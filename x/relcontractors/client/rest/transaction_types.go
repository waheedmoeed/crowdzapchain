package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"time"
)

type UpdateRelContractorAddressReq struct {
	BaseReq                 rest.BaseReq `json:"base_req"`
	RelContractorAddress    string       `json:"rel_contractor_address"`
	NewRelContractorAddress string       `json:"new_rel_contractor_address"`
}

type CreateVotingPollReq struct {
	BaseReq        rest.BaseReq   `json:"base_req"`
	PollType       uint           `json:"poll_type"`
	StartTime      time.Time      `json:"start_time"`
	EndTime        time.Time      `json:"end_time"`
	OwnerVoterPoll sdk.AccAddress `json:"owner_voter_poll"`
	CoinsAmount    sdk.Coin       `json:"coins_amount"`
}

type VotePollReq struct {
	BaseReq rest.BaseReq   `json:"base_req"`
	PollId  string         `json:"poll_id"`
	Vote    uint           `json:"vote"`
	Voter   sdk.AccAddress `json:"voter"`
}

type ProcessPoll struct {
	BaseReq    rest.BaseReq   `json:"base_req"`
	PollId     string         `json:"poll_id"`
	Transactor sdk.AccAddress `json:"transactor"`
}
