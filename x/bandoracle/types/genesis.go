package types

func NewGenesisState(markets []Market, params Params, portID string) *GenesisState {
	return &GenesisState{
		Markets: markets,
		Params:  params,
		PortId: portID,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		nil,
		DefaultParams(),
		"",
	)
}

func ValidateGenesis(_ *GenesisState) error {
	return nil
}
