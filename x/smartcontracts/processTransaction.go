package smartcontracts

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//change contractor address with new provided address and add old one to other addresses
func handleMsgCreateBasicContract(ctx sdk.Context, k Keeper, msg MsgCreateBasicContract) (*sdk.Result, error) {
	key, basicContract, err := CreateBasicContract(ctx, msg)
	fmt.Println("HHHHHHEEELLLOOO")
	fmt.Println(basicContract)
	if err != nil {
		return nil, err
	}
	k.SetBasicContract(ctx, key, basicContract)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
