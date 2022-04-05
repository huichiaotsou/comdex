package types

import govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

const (
	ProposalAddUnbondingDuration = "AddUnbondingDuration"
)

func init() {
	govtypes.RegisterProposalType(ProposalAddUnbondingDuration)
	govtypes.RegisterProposalTypeCodec(&UpdateUnbondingDuration{}, "comdex/UpdateUnbondingDuration")
}

var (
	_ govtypes.Content = &UpdateUnbondingDuration{}
)

func NewUpdateUnbondingDuration(title, description, UnbondingDuration string) govtypes.Content {
	return &UpdateUnbondingDuration{
		Title:             title,
		Description:       description,
		UnbondingDuration: UnbondingDuration,
	}

}

func (m *UpdateUnbondingDuration) ProposalRoute() string {
	return RouterKey
}

func (m *UpdateUnbondingDuration) ProposalType() string {
	return ProposalAddUnbondingDuration
}

func (m *UpdateUnbondingDuration) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	if len(m.UnbondingDuration) == 0 {
		return ErrorEmptyProposalUnbondingDuration
	}
	return nil
}
