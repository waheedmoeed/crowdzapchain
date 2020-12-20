package keeper

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
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
	authKeeper auth.AccountKeeper
	bankKeeper types.BankKeeper
}

// NewKeeper creates a relcontractors keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, authKeeper auth.AccountKeeper, bankKeeper types.BankKeeper) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
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

//Get contractor from store by matching addresses
func (k Keeper) GetContractorByAddress(ctx sdk.Context, address sdk.AccAddress) (types.Contractor, error) {
	store := ctx.KVStore(k.storeKey)
	var contract types.RelContract
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(contractKey), &contract)
	if err != nil {
		return types.Contractor{}, err
	}
	contractor := types.Contractor{}
	for _, value := range contract.RelContractors {
		if value.ContractorAddress.Equals(address) {
			contractor = value
			break
		}
	}
	return contractor, nil
}

/*
	1) Check first if account with given address already existed.
	2) If there is no account, create new account with given address.
	3) Send defined coins in poll to given address
*/

func (k Keeper) SendCoinsToContractor(ctx sdk.Context, address sdk.AccAddress, amount sdk.Coin) error {
	account := k.authKeeper.GetAccount(ctx, address)
	if account == nil {
		newAccount := k.authKeeper.NewAccountWithAddress(ctx, address)
		if newAccount == nil {
			errors.New("failed to create new account with given address")
		}
	}
	coins := []sdk.Coin{amount}
	_, err := k.bankKeeper.AddCoins(ctx, address, coins)
	return err
}
