package types

import govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

const (
	ProposalAddUnliqpointpercentage = "AddUnliqpointpercentage"
)

func init() {
	govtypes.RegisterProposalType(ProposalAddUnliqpointpercentage)
	govtypes.RegisterProposalTypeCodec(&UpdateUnliquidatePointPercent{}, "comdex/UpdateUnliquidatePointPercent")
}

var (
	_ govtypes.Content = &UpdateUnliquidatePointPercent{}
)

func NewUpdateUnliquidatePointPercentage(title, description, Unliqpointpercentage string) govtypes.Content {
	return &UpdateUnliquidatePointPercent{
		Title:             title,
		Description:       description,
		Unliqpointpercentage: Unliqpointpercentage,
	}

}

func (m *UpdateUnliquidatePointPercent) ProposalRoute() string {
	return RouterKey
}

func (m *UpdateUnliquidatePointPercent) ProposalType() string {
	return ProposalAddUnliqpointpercentage
}

func (m *UpdateUnliquidatePointPercent) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	if len(m.Unliqpointpercentage) == 0 {
		return ErrorEmptyProposalUnliqpointpercentage
	}
	return nil
}

func (m *UpdateUnliquidatePointPercent) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateUnliquidatePointPercent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}
