package types

// GenesisState - all smartcontracts state that must be provided at genesis
type GenesisState struct {
	BasicContracts []BasicContract `json:"basic_contracts"`
	YieldContracts []YieldContract `json:"yield_contracts"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{
		BasicContracts: nil,
		YieldContracts: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		BasicContracts: nil,
		YieldContracts: nil,
	}
}

// ValidateGenesis validates the smartcontracts genesis parameters
func ValidateGenesis(data GenesisState) error {
	return nil
}
