package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/waheedmoeed/relchain/x/relcontractors/types"
)

var contractKey = []byte("cosmos1h4sf6s3xvkh04wahfg8ncm7yh9p22ds7rpyyrc")

// Keeper of the relcontractors store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a relcontractors keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// store/create smart contract in DB >>> Used at the time Genesis Process
func (k Keeper) InitContract(ctx sdk.Context, relContract types.RelContract) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(relContract)
	store.Set(contractKey, bz)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) Get(ctx sdk.Context) (types.RelContract, error) {
	store := ctx.KVStore(k.storeKey)
	var contract types.RelContract
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(contractKey), &contract)
	if err != nil {
		return types.RelContract{}, err
	}
	return contract, nil
}

func (k Keeper) Set(ctx sdk.Context, contract types.RelContract) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(contract)
	store.Set(contractKey, bz)
}
