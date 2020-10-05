package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	QueryRelContractor = "relContractor"
	QueryWhois         = "whois"
	QueryNames         = "names"
)

// NewQuerier creates a new querier for relcontractors clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryRelContractor:
			return queryParams(ctx, k, path[1:])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown relcontractors query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper, paths []string) ([]byte, error) {
	//todo:create methods in keeper as you needed
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}
	res, err := codec.MarshalJSONIndent(k.cdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
