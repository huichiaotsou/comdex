package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
)

var (
	_ sdk.Msg = (*MsgAddMarketRequest)(nil)
	_ sdk.Msg = (*MsgUpdateMarketRequest)(nil)
	_ sdk.Msg = (*MsgRemoveMarketForAssetRequest)(nil)
	_ sdk.Msg = (*MsgFetchPriceRequest)(nil)
	CoinRatesClientIDKey = "coin_rates_id"
)

func NewMsgAddMarketRequest(from sdk.AccAddress, symbol string, scriptID uint64, assetID uint64) *MsgAddMarketRequest {
	return &MsgAddMarketRequest{
		From:     from.String(),
		Symbol:   symbol,
		ScriptID: scriptID,
		Id: assetID,
	}
}

func (m *MsgAddMarketRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Symbol == "" {
		return errors.Wrap(ErrorInvalidSymbol, "symbol cannot be empty")
	}
	if len(m.Symbol) > MaxMarketSymbolLength {
		return errors.Wrapf(ErrorInvalidSymbol, "symbol length cannot be greater than %d", MaxMarketSymbolLength)
	}
	if m.ScriptID == 0 {
		return errors.Wrapf(ErrorInvalidScriptID, "script_id cannot be zero")
	}

	return nil
}

func (m *MsgAddMarketRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateMarketRequest(from sdk.AccAddress, symbol string, scriptID uint64) *MsgUpdateMarketRequest {
	return &MsgUpdateMarketRequest{
		From:     from.String(),
		Symbol:   symbol,
		ScriptID: scriptID,
	}
}

func (m *MsgUpdateMarketRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Symbol != "" {
		if len(m.Symbol) > MaxMarketSymbolLength {
			return errors.Wrapf(ErrorInvalidSymbol, "symbol length cannot be greater than %d", MaxMarketSymbolLength)
		}
	}

	return nil
}

func (m *MsgUpdateMarketRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgRemoveMarketForAssetRequest(from sdk.AccAddress, id uint64, symbol string) *MsgRemoveMarketForAssetRequest {
	return &MsgRemoveMarketForAssetRequest{
		From:   from.String(),
		Id:     id,
		Symbol: symbol,
	}
}

func (m *MsgRemoveMarketForAssetRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidID, "id cannot be zero")
	}
	if m.Symbol == "" {
		return errors.Wrap(ErrorInvalidSymbol, "symbol cannot be empty")
	}
	if len(m.Symbol) > MaxMarketSymbolLength {
		return errors.Wrapf(ErrorInvalidSymbol, "symbol length cannot be greater than %d", MaxMarketSymbolLength)
	}

	return nil
}

func (m *MsgRemoveMarketForAssetRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgFetchPriceRequest(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *Calldata,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	requestKey string,
	prepareGas uint64,
	executeGas uint64,
) *MsgFetchPriceRequest {
	return &MsgFetchPriceRequest{
		ClientID:      CoinRatesClientIDKey ,
		Creator:        creator,
		OracleScriptID: uint64(oracleScriptID),
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		RequestKey:     requestKey,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

func (m *MsgFetchPriceRequest) ValidateBasic() error {
	if m.Creator == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Creator); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if err := ibchost.ChannelIdentifierValidator(m.SourceChannel); err != nil {
		return errors.Wrapf(ErrorInvalidSourceChannel, "%s", err)
	}
	if m.OracleScriptID == 0 {
		return errors.Wrapf(ErrorInvalidScriptID, "script_id cannot be zero")
	}

	return nil
}

func (m *MsgFetchPriceRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
