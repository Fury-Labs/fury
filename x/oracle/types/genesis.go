package types

func NewGenesisState(markets []Market, params Params) *GenesisState {
	return &GenesisState{
		Markets: markets,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		nil,
		DefaultParams(),
	)
}

func ValidateGenesis(_ *GenesisState) error {
	return nil
}