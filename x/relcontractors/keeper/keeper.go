package keeper

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/waheedmoeed/relchain/x/relcontractors/types"
)

var contractKey = []byte("cosmos1h4sf6s3xvkh04wahfg8ncm7yh9p22ds7rpyyrc")

// Keeper of the relContractors store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
	authKeeper auth.AccountKeeper
	bankKeeper types.BankKeeper
}

// NewKeeper creates a relContractors keeper
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

// Get returns the pubKey from the address-pubKey relation
func (k Keeper) Get(ctx sdk.Context) (types.RelContract, error) {
	store := ctx.KVStore(k.storeKey)
	var contract types.RelContract
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(contractKey), &contract)
	if err != nil {
		return types.RelContract{}, err
	}
	return contract, nil
}

//Set new RelContract in DB
func (k Keeper) Set(ctx sdk.Context, contract types.RelContract) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(contract)
	store.Set(contractKey, bz)
}

//Add new address to any specific contractor
func (k Keeper) UpdateContractorByAddress(ctx sdk.Context, address sdk.AccAddress, newAddress sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	var contract types.RelContract
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(contractKey), &contract)
	if err != nil {
		return err
	}

	var contractorIndex = 0
	contractor := types.Contractor{}
	for index, value := range contract.RelContractors {
		if value.ContractorAddress.Equals(address) {
			contractor = value
			contractorIndex = index
			break
		}
	}
	if !contractor.ContractorAddress.Empty() {
		errors.New("contractor not founded for respective address")
	}
	addresses := append(contractor.OtherAddresses, address)
	contractor.OtherAddresses = addresses
	contractor.ContractorAddress = newAddress

	//add to contract
	contract.RelContractors[contractorIndex] = contractor
	k.Set(ctx, contract)
	return nil
}

func (k Keeper) CreatePoll(ctx sdk.Context, pollType string, amount uint, ownerVoterPoll sdk.AccAddress) error {
	contract, err := k.Get(ctx)

	fmt.Println(contract.VotingPolls)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.New("poll creation", 234, "Poll Creation"), "Failed to read relContract from store")
	}
	//Check latest poll if it is still valid( end-time not reached or not processed

	b := make([]byte, 22)
	for i := 0; i < 22; i++ {
		b[i] = byte(97 + rand.Intn(122-97))
	}

	//time
	startTime := time.Now()
	endTime := time.Now().Add(time.Hour * 24 * 30)
	poll := types.VotingPoll{
		PollId:               string(b),
		PollType:             pollType,
		StartTime:            startTime,
		EndTime:              endTime,
		PositiveVotes:        0,
		NegativeVotes:        0,
		PositiveVotesAddress: []sdk.AccAddress{},
		NegativeVotesAddress: []sdk.AccAddress{},
		OwnerVoterPoll:       ownerVoterPoll,
		Processed:            false,
		CoinsAmount:          sdk.NewCoin("rel", sdk.NewInt(int64(amount))),
	}

	var polls []types.VotingPoll
	polls = append(contract.VotingPolls, poll)
	contract.VotingPolls = polls
	k.Set(ctx, contract)
	return nil
}

//Get contractor from store by matching addresses
func (k Keeper) GetContractorByAddress(ctx sdk.Context, address sdk.AccAddress) (types.Contractor, error) {
	contract, err := k.Get(ctx)
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
	if contractor.ContractorAddress.Empty() {
		return types.Contractor{}, errors.New("failed to find contractor with respective address")
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
		fmt.Println(newAccount.GetCoins())
		if newAccount == nil {
			return errors.New("failed to create new account with given address")
		}
	}
	coins := []sdk.Coin{amount}
	_, err := k.bankKeeper.AddCoins(ctx, address, coins)
	return err
}
