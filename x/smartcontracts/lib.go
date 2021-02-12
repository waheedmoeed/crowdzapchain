package smartcontracts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"time"
)

/*************************/
/***TRANSACTION METHODS***/
/*************************/
func CreateBasicContract(ctx sdk.Context, contract MsgCreateBasicContract) (string, BasicContract, error) {
	key := generateNewAddress()
	basicContract := BasicContract{
		Address: key,
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
	return key, basicContract, nil
}

/*************************/
/*****HELPER METHODS******/
/*************************/

func generateNewAddress() string {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return addr.String()
}
