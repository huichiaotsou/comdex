package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidation module sentinel errors
var (
	ErrSample		      = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrorInvalidFrom      = sdkerrors.Register(ModuleName, 1200, "invalid from")

)
