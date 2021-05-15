package smartcontracts

import (
	"github.com/waheedmoeed/relchain/x/smartcontracts/keeper"
	"github.com/waheedmoeed/relchain/x/smartcontracts/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	StoreKeyB    = types.StoreKeyB
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

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	//Custom MSG Types
	MsgCreateBasicContract = types.MsgCreateBasicContract
	MsgCreateYieldContract = types.MsgCreateYieldContract
	MsgInvestBasicContract = types.MsgInvestBasicContract

	//Contract types
	Contract            = types.Contract
	BasicDetail         = types.BasicDetail
	InvestmentRecord    = types.InvestmentRecord
	TokenTransferRecord = types.TokenTransferRecord
	BasicContract       = types.BasicContract
	YieldContract       = types.YieldContract
)
