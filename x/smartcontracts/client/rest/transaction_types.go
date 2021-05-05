package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"time"
)

type CreateBasicContractReq struct {
	BaseReq     rest.BaseReq   `json:"base_req"`
	Title       string         `json:"title"`
	StartTime   time.Time      `json:"start_date"`
	EndTime     time.Time      `json:"end_date"`
	TotalSupply uint           `json:"total_supply"`
	TokenPrice  uint           `json:"token_price"`
	Creator     sdk.AccAddress `json:"creator"`
}

type CreateYieldContractReq struct {
	BaseReq     rest.BaseReq   `json:"base_req"`
	Title       string         `json:"title"`
	StartTime   time.Time      `json:"start_date"`
	EndTime     time.Time      `json:"end_date"`
	TotalSupply uint           `json:"total_supply"`
	TokenPrice  uint           `json:"token_price"`
	Creator     sdk.AccAddress `json:"creator"`
}