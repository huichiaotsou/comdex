package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&UpdateUnliquidatePointPercent{}, "comdex/UpdateUnliquidatePointPercent", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdateUnliquidatePointPercent{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
