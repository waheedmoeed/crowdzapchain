package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type MsgUpdateRelContractorAddress struct {
	RelContractorAddress    sdk.AccAddress `json:"rel_contractor_address"`
	NewRelContractorAddress sdk.AccAddress `json:"new_rel_contractor_address"`
}
