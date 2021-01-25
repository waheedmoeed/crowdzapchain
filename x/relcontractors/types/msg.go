package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

func NewMsgUpdateRelContractorAddress(relContractorAdd sdk.AccAddress, newRelContractAdd sdk.AccAddress) MsgUpdateRelContractorAddress {
	return MsgUpdateRelContractorAddress{
		RelContractorAddress:    relContractorAdd,
		NewRelContractorAddress: newRelContractAdd,
	}
}

func (msg MsgUpdateRelContractorAddress) Route() string { return RouterKey }
func (msg MsgUpdateRelContractorAddress) Type() string  { return "update_relcontractor_address" }
func (msg MsgUpdateRelContractorAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.RelContractorAddress}
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

/*



 */

func NewMsgCreatePoll(pollType uint, startTime time.Time, endTime time.Time, ownerVoterPoll sdk.AccAddress, coinsAmount sdk.Coin) MsgCreatePoll {
	return MsgCreatePoll{
		PollType:       pollType,
		StartTime:      startTime,
		EndTime:        endTime,
		OwnerVoterPoll: ownerVoterPoll,
		CoinsAmount:    coinsAmount,
	}
}
func (msg MsgCreatePoll) Route() string { return RouterKey }
func (msg MsgCreatePoll) Type() string  { return "create_poll" }
func (msg MsgCreatePoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.OwnerVoterPoll}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreatePoll) ValidateBasic() error {
	if !(msg.PollType > 0 && msg.PollType < 4) {
		return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "Invalid type of pool")
	}
	if msg.EndTime.Sub(msg.StartTime).Hours() < 24 {
		return sdkErrors.Wrap(sdkErrors.New("poll creation", 234, "POll Validation"), "Time to vote is less than 24 hours")
	}
	return nil
}

/*



 */
func NewMsgVotePoll(pollId string, vote uint, voter sdk.AccAddress) MsgVotePoll {
	return MsgVotePoll{
		PollId: pollId,
		Vote:   vote,
		Voter:  voter,
	}
}
func (msg MsgVotePoll) Route() string { return RouterKey }
func (msg MsgVotePoll) Type() string  { return "vote_poll" }
func (msg MsgVotePoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgVotePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgVotePoll) ValidateBasic() error {
	//Vote must be either 0 or 1
	if msg.Vote == 0 || msg.Vote == 1 {
		return nil
	}
	return sdkErrors.Wrap(sdkErrors.New("poll vote", 234, "POll Voting"), "Invalid vote")
}

/*



 */
func NewProcessPoll(pollId string, transactor sdk.AccAddress) MsgProcessPoll {
	return MsgProcessPoll{
		PollId:     pollId,
		Transactor: transactor,
	}
}
func (msg MsgProcessPoll) Route() string { return RouterKey }
func (msg MsgProcessPoll) Type() string  { return "process_poll" }
func (msg MsgProcessPoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Transactor}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgProcessPoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgProcessPoll) ValidateBasic() error {
	return nil
}

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
