package keeper

import (
	"github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleProposalAddUnliquidatePointPercentage(ctx sdk.Context, p *types.UpdateUnliquidatePointPercent) error {
	return k.AddUnliquidatePointPercentageRecords(ctx, p.Unliqpointpercentage)
}
