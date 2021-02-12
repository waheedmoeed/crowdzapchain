package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/waheedmoeed/relchain/x/smartcontracts/types"
)

// Keeper of the smartcontracts store
type Keeper struct {
	storeKey   sdk.StoreKey
	storeKeyB  sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
	authKeeper auth.AccountKeeper
	bankKeeper types.BankKeeper
}

/*************************/
/**GENESIS STATE METHODS**/
/*************************/

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetAllBasicContracts(ctx sdk.Context) ([]types.BasicContract, []error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	var basicContracts []types.BasicContract
	var errors []error

	for ; iterator.Valid(); iterator.Next() {
		address := string(iterator.Key())
		basicContract, error := k.GetBasicContract(ctx, address)
		if error != nil {
			errors = append(errors, error)
		}
		basicContracts = append(basicContracts, basicContract)
	}
	return basicContracts, errors
}

func (k Keeper) GetAllYieldContracts(ctx sdk.Context) ([]types.YieldContract, []error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	var basicContracts []types.YieldContract
	var errors []error

	for ; iterator.Valid(); iterator.Next() {
		address := string(iterator.Key())
		basicContract, error := k.GetYieldContract(ctx, address)
		if error != nil {
			errors = append(errors, error)
		}
		basicContracts = append(basicContracts, basicContract)
	}
	return basicContracts, errors
}

/*************************/
/*******BASIC METHODS*****/
/*************************/
// NewKeeper creates a smartcontracts keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, keyB sdk.StoreKey, authKeeper auth.AccountKeeper, bankKeeper types.BankKeeper) Keeper {
	keeper := Keeper{
		storeKey:   key,
		storeKeyB:  keyB,
		cdc:        cdc,
		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the simple contract
func (k Keeper) GetBasicContract(ctx sdk.Context, key string) (types.BasicContract, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.BasicContract
	byteKey := []byte(key)

	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return types.BasicContract{}, err
	}
	return item, nil
}

func (k Keeper) SetBasicContract(ctx sdk.Context, key string, contract types.BasicContract) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(contract)
	store.Set([]byte(key), bz)
}

/************/
/***********/

// Get returns the yielding contract
func (k Keeper) GetYieldContract(ctx sdk.Context, key string) (types.YieldContract, error) {
	store := ctx.KVStore(k.storeKeyB)
	var item types.YieldContract
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return types.YieldContract{}, err
	}
	return item, nil
}

func (k Keeper) SetYieldContract(ctx sdk.Context, key string, contract types.YieldContract) {
	store := ctx.KVStore(k.storeKeyB)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(contract)
	store.Set([]byte(key), bz)
}

/*************************/
/****HELPER METHODS*******/
/*************************/
