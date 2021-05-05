package keeper

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/waheedmoeed/relchain/x/smartcontracts/types"
)

const (
	QueryBasicContract = "get_basic_contract"
	QueryYieldContract = "get_yield_contract"
)

// NewQuerier creates a new querier for smartcontracts clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryBasicContract:
			return queryBasicContract(ctx, k, path[1:])
		case QueryYieldContract:
			return queryYieldContract(ctx, k, path[1:])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown smartcontracts query endpoint")
		}
	}
}

func queryBasicContract(ctx sdk.Context, k Keeper, paths []string) ([]byte, error) {
	contract, err := k.GetBasicContract(ctx, paths[0])
	if err != nil || contract.Contract.BasicDetail.Title == "" {
		return nil, sdkerrors.Wrap(err, "failed to get contract")
	}
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("Failed to get contract for address %s", paths[0]))
	}

	return res, nil
}

func queryYieldContract(ctx sdk.Context, k Keeper, paths []string) ([]byte, error) {
	contract, err := k.GetYieldContract(
		ctx, paths[0])
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get contract")
	}
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
