package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	LockedVaultDoesNotExist                = sdkerrors.Register(ModuleName, 201, "locked vault does not exist with given id")
	ErrorEmptyProposalAssets               = sdkerrors.Register(ModuleName, 202, "empty assets in proposal")
	ErrorEmptyProposalUnliqpointpercentage = sdkerrors.Register(ModuleName, 203, "empty Unliquidate point percentage in proposal")
	ErrorUnknownProposalType               = sdkerrors.Register(ModuleName, 204, "unknown proposal type")
)
