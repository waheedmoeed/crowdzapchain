package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type UpdateRelContractorAddressReq struct {
	BaseReq                 rest.BaseReq `json:"base_req"`
	RelContractorAddress    string       `json:"rel_contractor_address"`
	NewRelContractorAddress string       `json:"new_rel_contractor_address"`
}

type CreatePollReq struct {
	BaseReq        rest.BaseReq   `json:"base_req"`
	PollType       string         `json:"poll_type"`
	OwnerVoterPoll sdk.AccAddress `json:"owner_voter_poll"`
	Amount         uint           `json:"amount"`
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
