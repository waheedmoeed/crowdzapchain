package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type Contract struct {
	//to keep track of total coins and distributed coins among contractors.
	BasicDetail BasicDetail        `json:"basic_detail"`
	Registry    []InvestmentRecord `json:"registry"`
}

type BasicDetail struct {
	Title       string    `json:"title"`
	TotalSupply uint      `json:"total_supply"`
	SoldToken   uint      `json:"sold_token"`
	TokenPrice  sdk.Coin  `json:"token_price"`
	StartDate   time.Time `json:"start_date"`
	EndingDate  time.Time `json:"ending_date"`
}

type InvestmentRecord struct {
	InvestorAddress sdk.AccAddress      `json:"investor_address"`
	OwnedTokens     uint                `json:"owned_token"`
	InvestedDate    time.Time           `json:"invested_date"`
	LatestTransfer  TokenTransferRecord `json:"token_transfer"`
}

type TokenTransferRecord struct {
	From         sdk.AccAddress `json:"transfer_from"`
	TransferDate time.Time      `json:"transfer_date"`
}

/**************/
/**************/

type BasicContract struct {
	Address  string   `json:"contract_address"`
	Contract Contract `json:"basic_contract"`
	//TODO :: define other attributes for this contract
}

type YieldContract struct {
	Address  string   `json:"contract_address"`
	Contract Contract `json:"basic_contract"`
	//TODO :: define other attributes for this contract
}

///////////Transaction method
//send (to other user)
//invest (from escrow contract)
//

/////////Query method
//////Get token price
//////Get start date
//////Get end date
//////Get Basic Detail
//////Get
