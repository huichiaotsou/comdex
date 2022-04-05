package liquidity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/pkg/errors"

	"github.com/comdex-official/comdex/x/liquidity/keeper"
	"github.com/comdex-official/comdex/x/liquidity/types"
)

// NewHandler returns a handler for all "liquidity" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreatePool:
			res, err := msgServer.CreatePool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDepositWithinBatch:
			res, err := msgServer.DepositWithinBatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawWithinBatch:
			res, err := msgServer.WithdrawWithinBatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSwapWithinBatch:
			res, err := msgServer.Swap(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBondPoolTokens:
			res,err  := msgServer.BondPoolTokens(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx,res,err)
		case *types.MsgUnbondPoolTokens:
			res,err  := msgServer.UnbondPoolTokens(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx,res,err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

func NewAddUnbondingPeriodProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateUnbondingDuration:
			return handleAddUnbondingPeriodProposal(ctx, k, c)
		default:
			return errors.Wrapf(types.ErrorUnknownProposalType, "%T", c)
		}
	}
}

func handleAddUnbondingPeriodProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateUnbondingDuration) error {
	return k.HandleProposalAddUnbondingPeriod(ctx, p)
}
