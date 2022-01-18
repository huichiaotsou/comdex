package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgLockRequest)(nil)
)

func NewMsgLockRequest(from sdk.AccAddress, id uint64) *MsgLockRequest {
	return &MsgLockRequest{
		From: from.String(),
		Id: id,
	}
}

func (m *MsgLockRequest) Route() string {
	return RouterKey
}

func (m *MsgLockRequest) Type() string {
	return TypeMsgLockRequest
}

func (m *MsgLockRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidFrom, "id cannot be zero")
	}
	return nil
}

func (m *MsgLockRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgLockRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
