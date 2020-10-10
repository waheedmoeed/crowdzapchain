package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

type MsgUpdateRelContractorAddress struct {
	RelContractorAddress    sdk.AccAddress `json:"rel_contractor_address"`
	NewRelContractorAddress sdk.AccAddress `json:"new_rel_contractor_address"`
}

func NewMsgUpdateRelContractorAddress(relContractorAdd sdk.AccAddress, newRelContractAdd sdk.AccAddress) MsgUpdateRelContractorAddress {
	return MsgUpdateRelContractorAddress{
		RelContractorAddress:    relContractorAdd,
		NewRelContractorAddress: newRelContractAdd,
	}
}

func (msg MsgUpdateRelContractorAddress) Route() string { return RouterKey }
func (msg MsgUpdateRelContractorAddress) Type() string  { return "update_relcontractor_address" }
func (msg MsgUpdateRelContractorAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.RelContractorAddress)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUpdateRelContractorAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUpdateRelContractorAddress) ValidateBasic() error {
	if msg.RelContractorAddress.Empty() || msg.NewRelContractorAddress.Empty() {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Missing now RelContractor Address")
	}
	return nil
}

type MsgCreatePoll struct {
	PollType       uint           `json:"type"`
	StartTime      time.Time      `json:"start_time"`
	EndTime        time.Time      `json:"end_time"`
	OwnerVoterPoll sdk.AccAddress `json:"owner_voter_poll"`
}

const (
	MINT_COINS           = 1
	DISTRIBUTE_COINS     = 2
	BLACKLIST_CONTRACTOR = 3
)

func NewMsgCreatePoll(pollType uint, startTime time.Time, endTime time.Time, ownerVoterPoll sdk.AccAddress) MsgCreatePoll {
	return MsgCreatePoll{
		PollType:       pollType,
		StartTime:      startTime,
		EndTime:        endTime,
		OwnerVoterPoll: ownerVoterPoll,
	}
}
func (msg MsgCreatePoll) Route() string { return RouterKey }
func (msg MsgCreatePoll) Type() string  { return "update_relcontractor_address" }
func (msg MsgCreatePoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.OwnerVoterPoll)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

//TODO: do data validation
// ValidateBasic validity check for the AnteHandler
func (msg MsgCreatePoll) ValidateBasic() error {
	if msg.GetSigners()[0].Equals(msg.OwnerVoterPoll) {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Mismatch RelContractor Address and signer")
	}
	if !(msg.PollType > 0 && msg.PollType < 4) {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid type of pool")
	}
	if msg.EndTime.Sub(msg.StartTime).Hours() < 24 {
		return sdkErrors.Wrap(errors.New("Time Duration for poll must be greater than 24 hours"), "")
	}
	return nil
}

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
/*
// verify interface at compile time
var _ sdk.Msg = &Msg<Action>{}

// Msg<Action> - struct for unjailing jailed validator
type Msg<Action> struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
	return Msg<Action>{
		ValidatorAddr: validatorAddr,
	}
}

const <action>Const = "<action>"

// nolint
func (msg Msg<Action>) Route() string { return RouterKey }
func (msg Msg<Action>) Type() string  { return <action>Const }
func (msg Msg<Action>) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg Msg<Action>) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg Msg<Action>) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
*/
