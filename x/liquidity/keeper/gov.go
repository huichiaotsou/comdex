package keeper

import (
	"github.com/comdex-official/comdex/x/liquidity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleProposalAddUnbondingPeriod(ctx sdk.Context, p *types.UpdateUnbondingDuration) error {
	return k.AddUnbondingDurationRecords(ctx, p.UnbondingDuration)
}
