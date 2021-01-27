package relcontractors

import (
	"github.com/waheedmoeed/relchain/x/relcontractors/keeper"
	"github.com/waheedmoeed/relchain/x/relcontractors/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases
	NewRelContract = types.NewRelContract()
	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases

	//TODO:: Custom Error registration
)

type (
	Keeper                        = keeper.Keeper
	GenesisState                  = types.GenesisState
	Params                        = types.Params
	RelContract                   = types.RelContract
	VotingPoll                    = types.VotingPoll
	CoinsMintedRecord             = types.CoinsMintedRecord
	DistributedCoinsRecord        = types.DistributedCoinsRecord
	MsgUpdateRelContractorAddress = types.MsgUpdateRelContractorAddress
	MsgCreateVotePoll             = types.MsgCreateVotePoll
	MsgVotePoll                   = types.MsgVotePoll
	MsgProcessPoll                = types.MsgProcessPoll
	BankKeeper                    = types.BankKeeper
)
