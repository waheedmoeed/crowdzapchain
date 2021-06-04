package smartcontracts

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
			Registry: []InvestmentRecord{},
		},
	}
	return basicContract, nil
}

func CreateYieldContract(ctx sdk.Context, contract MsgCreateYieldContract) (YieldContract, error) {
	basicContract := YieldContract{
		Address: "",
		Contract: Contract{
			BasicDetail: BasicDetail{
				Title:       contract.Title,
				TotalSupply: contract.TotalSupply,
				SoldToken:   0,
				TokenPrice:  sdk.NewCoin("rel", sdk.NewInt(int64(contract.TokenPrice))),
				StartDate:   time.Now(),
				EndingDate:  time.Now().Add(time.Hour * 24 * 2),
			},
			Registry: []InvestmentRecord{},
		},
	}
	return basicContract, nil
}

/*************************/
/*****HELPER METHODS******/
/*************************/

//HaveTokens validate that there are equal tokens available to transfer
func HaveTokens(ctx sdk.Context, address sdk.AccAddress, amount uint, contract BasicContract) (bool, int, error) {
	founded := false
	foundedIndex := -1
	if contract.Contract.Registry != nil {
		for index, value := range contract.Contract.Registry {
			if value.InvestorAddress.Equals(address) {
				if value.OwnedTokens == amount {
					foundedIndex = index
					founded = true
				}
			}
		}
		if founded {
			return founded, foundedIndex, nil
		}
	}
	return founded, foundedIndex, errors.New("there are not investements in registry")
}
