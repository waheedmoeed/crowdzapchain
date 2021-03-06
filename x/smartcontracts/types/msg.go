package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateBasicContract(ContractAddress sdk.AccAddress, Creator sdk.AccAddress, Title string, TotalSupply, TokenPrice uint, StartDate, EndDate time.Time) MsgCreateBasicContract {
	return MsgCreateBasicContract{
		ContractAddress: ContractAddress,
		Creator:         Creator,
		Title:           Title,
		StartTime:       StartDate,
		EndTime:         EndDate,
		TotalSupply:     TotalSupply,
		TokenPrice:      TokenPrice,
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

func NewMsgCreateYieldContract(Creator sdk.AccAddress, Title string, TotalSupply uint, TokenPrice uint, StartDate time.Time, EndDate time.Time) MsgCreateBasicContract {
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
func (msg MsgCreateYieldContract) Route() string { return RouterKey }

func (msg MsgCreateYieldContract) Type() string {
	return "create_basic_contract"
}

func (msg MsgCreateYieldContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateYieldContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateYieldContract) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	return nil
}

////////////////////////////////////
///////////////////////////////////
func NewMsgInvestBasicContract(Creator sdk.AccAddress, Title string, TotalSupply uint, TokenPrice uint, StartDate time.Time, EndDate time.Time) MsgCreateBasicContract {
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
func (msg MsgInvestBasicContract) Route() string { return RouterKey }

func (msg MsgInvestBasicContract) Type() string {
	return "invest_basic_contract"
}

func (msg MsgInvestBasicContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Investor)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgInvestBasicContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgInvestBasicContract) ValidateBasic() error {
	if msg.Investor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	return nil
}

/////////////////////////////////
func (msg MsgTransferBasicContract) Route() string { return RouterKey }

func (msg MsgTransferBasicContract) Type() string {
	return "transfer_basic_contract"
}

func (msg MsgTransferBasicContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.From)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgTransferBasicContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgTransferBasicContract) ValidateBasic() error {
	if msg.To.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing to address")
	}
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
	}
	return nil
}
