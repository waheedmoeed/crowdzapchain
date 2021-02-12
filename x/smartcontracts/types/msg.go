package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func NewMsgCreateBasicContract(Creator sdk.AccAddress, Title string, TotalSupply uint, TokenPrice uint, StartDate time.Time, EndDate time.Time) MsgCreateBasicContract {
	return MsgCreateBasicContract{
		Creator:     Creator,
		Title:       Title,
		StartTime:   StartDate,
		EndTime:     EndDate,
		TotalSupply: TotalSupply,
		TokenPrice:  TokenPrice,
	}
}

// nolint
func (msg MsgCreateBasicContract) Route() string { return RouterKey }

func (msg MsgCreateBasicContract) Type() string {
	return "create_basic_contract"
}

func (msg MsgCreateBasicContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateBasicContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateBasicContract) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	return nil
}
