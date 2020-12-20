package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice is Initial Starting Price for a name that was never previously owned
//var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("rel", 1)}

// Whois is a struct that contains all the metadata of a name
type RelContract struct {
	//to keep track of total coins and distributed coins among contractors.
	MintedCoins       sdk.Coin            `json:"minted_coins"`
	MintedCoinsRecord []CoinsMintedRecord `json:"minted_coins_record"`

	DistributedCoins     sdk.Coin                 `json:distributed_coins`
	DistributedCoinsLogs []DistributedCoinsRecord `json:"rel_contractors_coins"`

	RelContractors []Contractor `json:"rel_contractors"`

	VotingPolls []VotingPoll `json:"voting_polls"`
}

type Contractor struct {
	ContractorAddress sdk.AccAddress   `json:contractor_address`
	Name              string           `json:contractor_name`
	OtherAddresses    []sdk.AccAddress `json:"other_addresses"`
}

type DistributedCoinsRecord struct {
	ContractorAddress      sdk.AccAddress `json:contractor_address`
	Coins                  sdk.Coin       `json:"coins"`
	IssuedDate             time.Time      `json:"issued_date"`
	DistributedCoinsAmount sdk.Coin       `json:"distributed_coins_amount"` //amount of distributed coins at the time of distribution
}

//Keeps track of coins minted and positive votes to that minted coins
type CoinsMintedRecord struct {
	Coins           sdk.Coin         `json:"coins"`
	Date            time.Time        `json:"date"`
	VoterContractor []sdk.AccAddress `json:contractor_addresses`
}

type VotingPoll struct {
	PollId               string           `json:"poll_id"`
	PollType             uint             `json:"type"`
	StartTime            time.Time        `json:"start_time"`
	EndTime              time.Time        `json:"end_time"`
	PositiveVotes        uint             `json:"positive_votes"`
	NegativeVotes        uint             `json:"negative_votes"`
	PositiveVotesAddress []sdk.AccAddress `json:"positive_votes_address"`
	NegativeVotesAddress []sdk.AccAddress `json:"negative_votes_address"`
	OwnerVoterPoll       sdk.AccAddress   `json:"owner_voter_poll"`
	Processed            bool             `json:"processed"`
	CoinsAmount          sdk.Coin         `json:"coins_amount"`
}

// NewWhois returns a new Whois with the  as the price
func NewRelContract() RelContract {
	return RelContract{}
}

// implement fmt.Stringer
func (contract RelContract) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`Minted_Coins: %s
				Minted_Coins_Record: %s
				Distributed_Coins: %s
				DistributedCoinsLogs: %s
				RelContractors: %s
				VotingPolls: %s`,
		contract.MintedCoins.String(),
		coinsMintedRecordString(contract.MintedCoinsRecord),
		contract.DistributedCoins.String(),
		distributedCoinsLogsString(contract.DistributedCoinsLogs),
		relContractorString(contract.RelContractors),
		votingPollsString(contract.VotingPolls),
	))
}

func coinsMintedRecordString(records []CoinsMintedRecord) string {
	coinsMintedRecord := ""
	for _, value := range records {
		coinsMintedRecord = coinsMintedRecord + fmt.Sprintf(
			`Coins: %s
			Date: %d
			Contractrors_Positive: %s`,
			value.Coins.String(),
			value.Date,
			addressesString(value.VoterContractor, "Positive_Voter"),
		)
	}
	return coinsMintedRecord
}

func distributedCoinsLogsString(records []DistributedCoinsRecord) string {
	distributedCoinsLogs := ""
	for _, value := range records {
		distributedCoinsLogs = distributedCoinsLogs + fmt.Sprintf(
			`ContractorAddress: %s
					Coins: %s
					IssuedDate: %d
					DistributedCoinsAmount: %s`,
			value.ContractorAddress.String(),
			value.Coins.String(),
			value.IssuedDate,
			value.DistributedCoinsAmount.String(),
		)
	}
	return strings.TrimSpace(distributedCoinsLogs)
}

func relContractorString(contracters []Contractor) string {
	relContractors := ""
	for _, value := range contracters {
		relContractors = relContractors + fmt.Sprintf(
			`ContractorAddress: %s
					Name: %s
					OtherAdressess: %s`,
			value.ContractorAddress.String(),
			value.Name,
			addressesString(value.OtherAddresses, "Other_Address"),
		)
	}
	return strings.TrimSpace(relContractors)
}

func votingPollsString(poll []VotingPoll) string {
	votingPolls := ""
	for _, value := range poll {
		votingPolls = votingPolls + fmt.Sprintf(`
			PollType: %d
			StartTime: %d
			EndTime: %d 
			PositiveVotes: %d
			NegativeVotes: %d
			PositiveVotesAddress %s
			NegativeVotesAddress: %s
			OwnerVoterPoll: %s`,
			value.PollType,
			value.StartTime,
			value.EndTime,
			value.PositiveVotes,
			value.NegativeVotes,
			addressesString(value.PositiveVotesAddress, "Positive Voters"),
			addressesString(value.NegativeVotesAddress, "Negative Voters"),
			value.OwnerVoterPoll.String(),
		)
	}
	return votingPolls
}

//to stringify addresses
func addressesString(addresses []sdk.AccAddress, tag string) string {
	addressesString := ""
	for index, address := range addresses {
		addressesString = addressesString + fmt.Sprintf("%d %s: %s", index, tag, address.String())
	}
	return addressesString
}
