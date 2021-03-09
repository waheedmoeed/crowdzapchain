package smartcontracts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

/*************************/
/***TRANSACTION METHODS***/
/*************************/
func CreateBasicContract(ctx sdk.Context, contract MsgCreateBasicContract) (BasicContract, error) {
	basicContract := BasicContract{
		Address: contract.ContractAddress.String(),
		Contract: Contract{
			BasicDetail: BasicDetail{
				Title:       contract.Title,
				TotalSupply: contract.TotalSupply,
				SoldToken:   0,
				TokenPrice:  sdk.NewCoin("rel", sdk.NewInt(int64(contract.TokenPrice))),
				StartDate:   time.Now(),
				EndingDate:  time.Now().Add(time.Hour * 24 * 2),
			},
			Registry: []InvestmentRecord{
				{
					InvestorAddress: nil,
					OwnedTokens:     0,
					InvestedDate:    time.Time{},
					LatestTransfer: TokenTransferRecord{
						From:         nil,
						TransferDate: time.Time{},
					},
				},
			},
		},
	}
	return basicContract, nil
}

func InvestBasic(ctx sdk.Context, contract MsgCreateBasicContract) (BasicContract, error) {
	basicContract := BasicContract{
		Address: contract.ContractAddress.String(),
		Contract: Contract{
			BasicDetail: BasicDetail{
				Title:       contract.Title,
				TotalSupply: contract.TotalSupply,
				SoldToken:   0,
				TokenPrice:  sdk.NewCoin("rel", sdk.NewInt(int64(contract.TokenPrice))),
				StartDate:   time.Now(),
				EndingDate:  time.Now().Add(time.Hour * 24 * 2),
			},
			Registry: []InvestmentRecord{
				{
					InvestorAddress: nil,
					OwnedTokens:     0,
					InvestedDate:    time.Time{},
					LatestTransfer: TokenTransferRecord{
						From:         nil,
						TransferDate: time.Time{},
					},
				},
			},
		},
	}
	return basicContract, nil
}

/*************************/
/*****HELPER METHODS******/
/*************************/
