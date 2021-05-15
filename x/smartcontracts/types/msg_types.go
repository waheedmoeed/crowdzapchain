package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateBasicContract struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
	Title           string         `json:"title"`
	StartTime       time.Time      `json:"start_date"`
	EndTime         time.Time      `json:"end_date"`
	TotalSupply     uint           `json:"total_supply"`
	TokenPrice      uint           `json:"token_price"`
	Creator         sdk.AccAddress `json:"creator"`
}

type MsgCreateYieldContract struct {
	Title       string         `json:"title"`
	StartTime   time.Time      `json:"start_date"`
	EndTime     time.Time      `json:"end_date"`
	TotalSupply uint           `json:"total_supply"`
	TokenPrice  uint           `json:"token_price"`
	Creator     sdk.AccAddress `json:"creator"`
}

type MsgInvestBasicContract struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
	Amount          uint           `json:"amount"`
	Investor        sdk.AccAddress `json:"investor"`
}
