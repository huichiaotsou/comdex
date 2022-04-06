package liquidation

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/pkg/errors"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	"github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func NewAddUnliquidatePointPercentageProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateUnliquidatePointPercent:
			return handleAddUnliquidatePointPercentageProposal(ctx, k, c)
		default:
			return errors.Wrapf(types.ErrorUnknownProposalType, "%T", c)
		}
	}
}

func handleAddUnliquidatePointPercentageProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateUnliquidatePointPercent) error {
	return k.HandleProposalAddUnliquidatePointPercentage(ctx, p)
}
