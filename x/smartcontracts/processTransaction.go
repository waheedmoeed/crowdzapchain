package smartcontracts

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateBasicContract(ctx sdk.Context, k Keeper, msg MsgCreateBasicContract) (*sdk.Result, error) {
	basicContract, err := CreateBasicContract(ctx, msg)
	fmt.Println("HELLO BASIC CONTRACT")
	fmt.Println(basicContract)
	if err != nil {
		return nil, err
	}
	k.SetBasicContract(ctx, basicContract.Address, basicContract)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateYieldContract(ctx sdk.Context, k Keeper, msg MsgCreateYieldContract) (*sdk.Result, error) {
	basicContract, err := CreateYieldContract(ctx, msg)
	fmt.Println("HELLO YIELD CONTRACT")
	fmt.Println(basicContract)
	if err != nil {
		return nil, err
	}
	k.SetYieldContract(ctx, "", basicContract)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
