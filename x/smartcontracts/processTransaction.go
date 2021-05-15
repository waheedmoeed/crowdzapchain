package smartcontracts

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgCreateBasicContract(ctx sdk.Context, k Keeper, msg MsgCreateBasicContract) (*sdk.Result, error) {
	basicContract, err := CreateBasicContract(ctx, msg)
	if err != nil {
		return nil, err
	}
	k.SetBasicContract(ctx, basicContract.Address, basicContract)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateYieldContract(ctx sdk.Context, k Keeper, msg MsgCreateYieldContract) (*sdk.Result, error) {
	basicContract, err := CreateYieldContract(ctx, msg)
	if err != nil {
		return nil, err
	}
	k.SetYieldContract(ctx, "", basicContract)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgInvestBasicContract(ctx sdk.Context, k Keeper, msg MsgInvestBasicContract) (*sdk.Result, error) {
	basicContract, err := k.GetBasicContract(ctx, msg.ContractAddress.String())
	fmt.Println(basicContract)
	fmt.Println(err)
	if err != nil {
		return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("invest basic contract", 235, "Invest Basic"), "cannot invest in basic, failed to get basic contract")
	}

	//check if there are enough available tokens to purchase
	if (basicContract.Contract.BasicDetail.TotalSupply - basicContract.Contract.BasicDetail.SoldToken) >= msg.Amount {
		haveBalance, err := k.CheckRequiredBalance(ctx, msg.Investor, msg.Amount)
		if err != nil {
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("invest basic contract", 235, "Invest Basic"), "investor account not founded")
		}
		if haveBalance {
			err = k.DeductCoins(ctx, msg.ContractAddress, msg.Investor, msg.Amount)
			if err != nil {
				return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("invest basic contract", 235, "Invest Basic"), "cannot invest in basic, failed to deduct balance from investor account")
			}

			investmentRecord := InvestmentRecord{
				InvestorAddress: msg.Investor,
				OwnedTokens:     msg.Amount,
				InvestedDate:    time.Now(),
				LatestTransfer:  TokenTransferRecord{},
			}

			basicContract.Contract.BasicDetail.SoldToken = basicContract.Contract.BasicDetail.SoldToken + msg.Amount
			basicContract.Contract.Registry = append(basicContract.Contract.Registry, investmentRecord)
			//TODO: Add events
			k.SetBasicContract(ctx, msg.ContractAddress.String(), basicContract)
			return &sdk.Result{Events: ctx.EventManager().Events()}, nil
		} else {
			return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("invest basic contract", 235, "Invest Basic"), "not enough balnce to purchase tokens in this contract")
		}
	} else {
		return &sdk.Result{}, sdkerrors.Wrap(sdkerrors.New("invest basic contract", 235, "Invest Basic"), "not enough token available to purchase")
	}

}
