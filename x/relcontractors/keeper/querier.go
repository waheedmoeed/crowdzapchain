package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/waheedmoeed/relchain/x/relcontractors/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	QueryDistributedMintedCoins = "minted_distributed_coins"
	QueryRelContractors         = "contractors"
	QueryRelContractorByAddress = "contractor_by_address"
	QueryDistributedRecords     = "distributed_records"
	QueryMintedRecords          = "minted_records"
	QueryPolls                  = "polls"
	QueryPollsByID              = "poll_by_id"
)

// NewQuerier creates a new querier for relcontractors clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryDistributedMintedCoins:
			return queryMintDistCoins(ctx, k)
		case QueryRelContractors:
			return queryRelContractors(ctx, k)
		case QueryRelContractorByAddress:
			return queryRelContractorByAddress(ctx, k, path[1:])
		case QueryDistributedRecords:
			return queryDistributedRecord(ctx, k)
		case QueryMintedRecords:
			return queryMintedRecord(ctx, k)
		case QueryPolls:
			return queryAllPolls(ctx, k)
		case QueryPollsByID:
			return queryPollsByID(ctx, k, path[1:])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown relcontractors query endpoint")
		}
	}
}

//return amount of minted and distributed coins
func queryMintDistCoins(ctx sdk.Context, k Keeper) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}
	type response struct {
		minted_coins      sdk.Coin
		distributed_coins sdk.Coin
	}
	output := response{
		minted_coins:      contract.MintedCoins,
		distributed_coins: contract.DistributedCoins,
	}
	res, err := codec.MarshalJSONIndent(k.cdc, output)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return all rel contractors
func queryRelContractors(ctx sdk.Context, k Keeper) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}
	res, err := codec.MarshalJSONIndent(k.cdc, contract.RelContractors)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return empty contractor object if with given address contractor is not founded
func queryRelContractorByAddress(ctx sdk.Context, k Keeper, paths []string) ([]byte, error) {
	contractorAddress, error := sdk.AccAddressFromBech32(paths[0])
	if error != nil {
		return nil, error
	}
	contractor, error := k.GetContractorByAddress(ctx, contractorAddress)
	if error != nil {
		return nil, error
	}
	res, err := codec.MarshalJSONIndent(k.cdc, contractor)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return all distribution record of coins among rel-contractors
func queryDistributedRecord(ctx sdk.Context, k Keeper) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}

	res, err := codec.MarshalJSONIndent(k.cdc, contract.DistributedCoinsLogs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return all minted coins records
func queryMintedRecord(ctx sdk.Context, k Keeper) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}
	res, err := codec.MarshalJSONIndent(k.cdc, contract.MintedCoinsRecord)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return all the voting polls
func queryAllPolls(ctx sdk.Context, k Keeper) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}

	res, err := codec.MarshalJSONIndent(k.cdc, contract.VotingPolls)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//return voting poll by given Id
func queryPollsByID(ctx sdk.Context, k Keeper, paths []string) ([]byte, error) {
	contract, error := k.Get(ctx)
	if error != nil {
		return nil, error
	}
	pollById := types.VotingPoll{}
	for _, value := range contract.VotingPolls {
		if value.PollId == paths[0] {
			pollById = value
			break
		}
	}

	res, err := codec.MarshalJSONIndent(k.cdc, pollById)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
