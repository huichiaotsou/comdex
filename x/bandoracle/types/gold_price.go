package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgGoldPriceData{}

	// GoldPriceResultStoreKeyPrefix is a prefix for storing result
	GoldPriceResultStoreKeyPrefix = "gold_price_result"

	// LastGoldPriceIDKey is the key for the last request id
	LastGoldPriceIDKey = "gold_price_last_id"

	// GoldPriceClientIDKey is query request identifier
	GoldPriceClientIDKey = "gold_price_id"
)

// NewMsgGoldPriceData creates a new GoldPrice message
func NewMsgGoldPriceData(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *GoldPriceCallData,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgGoldPriceData {
	return &MsgGoldPriceData{
		ClientID:       GoldPriceClientIDKey,
		Creator:        creator,
		OracleScriptID: uint64(oracleScriptID),
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// Route returns the message route
func (m *MsgGoldPriceData) Route() string {
	return RouterKey
}

// Type returns the message type
func (m *MsgGoldPriceData) Type() string {
	return "GoldPriceData"
}

// GetSigners returns the message signers
func (m *MsgGoldPriceData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message
func (m *MsgGoldPriceData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation
func (m *MsgGoldPriceData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	return nil
}

// GoldPriceResultStoreKey is a function to generate key for each result in store
func GoldPriceResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(GoldPriceResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}