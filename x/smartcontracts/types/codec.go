package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// TODO: Register the modules msgs
	cdc.RegisterConcrete(MsgCreateBasicContract{}, "smartcontracts/create_basic_contract", nil)
	cdc.RegisterConcrete(MsgCreateYieldContract{}, "smartcontracts/create_yield_contract", nil)
	cdc.RegisterConcrete(MsgInvestBasicContract{}, "smartcontracts/invest_basic_contract", nil)
	cdc.RegisterConcrete(MsgTransferBasicContract{}, "smartcontracts/transfer_basic_contract", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
