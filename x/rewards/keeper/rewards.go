package keeper

import (
	"fmt"
	"strings"
	"time"

	"github.com/comdex-official/comdex/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k *Keeper) GetMintingRewardsID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.MintingRewardsIdKey
		value = store.Get(key)
	)
	if value == nil {
		return 0
	}
	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)
	return id.GetValue()
}

func (k *Keeper) SetMintingRewardsID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.MintingRewardsIdKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}

func (k *Keeper) SetMintingRewards(ctx sdk.Context, mintingReward types.MintingRewards) {
	var (
		store = k.Store(ctx)
		key   = types.MintingRewardsKey(mintingReward.Id)
		value = k.cdc.MustMarshal(&mintingReward)
	)
	store.Set(key, value)
}

func (k *Keeper) DeleteCollateralAuction(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.MintingRewardsKey(id)
	)
	store.Delete(key)
}

func (k *Keeper) GetMintingReward(ctx sdk.Context, id uint64) (mintingReward types.MintingRewards, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.MintingRewardsKey(id)
		value = store.Get(key)
	)
	if value == nil {
		return mintingReward, false
	}
	k.cdc.MustUnmarshal(value, &mintingReward)
	return mintingReward, true
}

func (k *Keeper) GetMintingRewards(ctx sdk.Context) (mintingRewards []types.MintingRewards) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.MintingRewardsKeyPrefix)
	)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var mintingReward types.MintingRewards
		k.cdc.MustUnmarshal(iter.Value(), &mintingReward)
		mintingRewards = append(mintingRewards, mintingReward)
	}
	return mintingRewards
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// Remove add-asset from cli
// Validate Total rewards amount
func (k Keeper) AddNewMintingRewards(ctx sdk.Context, newMintingRewardsData types.MintingRewards) error {
	availableAssets := k.asset.GetAssets(ctx)
	availableAssetsDenoms := []string{}
	for _, asset := range availableAssets {
		availableAssetsDenoms = append(availableAssetsDenoms, asset.Denom)
	}
	assetsDenomInProposal := []string{newMintingRewardsData.AllowedCasset, newMintingRewardsData.AllowedCollateral}
	invalidAssets := []string{}
	for _, asset := range assetsDenomInProposal {
		if !contains(availableAssetsDenoms, asset) {
			invalidAssets = append(invalidAssets, asset)
		}
	}
	if len(invalidAssets) > 0 {
		return sdkerrors.Wrapf(types.ErrorInvalidAssetDenoms, "invalid denoms %s", strings.Join(invalidAssets, ","))
	}
	availableMintingRewards := k.GetMintingRewards(ctx)
	for _, mintingReward := range availableMintingRewards {
		if mintingReward.AllowedCollateral == newMintingRewardsData.AllowedCollateral && mintingReward.AllowedCasset == newMintingRewardsData.AllowedCasset {
			if !mintingReward.IsActive && mintingReward.Depositor == nil {
				return sdkerrors.Wrapf(types.ErrorMintingRewardPairAlreadyExist, "reward pair %s already exists, deposit pending. ", newMintingRewardsData.AllowedCollateral, newMintingRewardsData.AllowedCasset)
			} else if !mintingReward.IsActive && mintingReward.Depositor != nil && mintingReward.StartTimestamp.After(ctx.BlockTime()) {
				return sdkerrors.Wrapf(types.ErrorMintingRewardPairAlreadyExist, "reward pair %s already exists, deposit made and waiting for start time.", newMintingRewardsData.AllowedCollateral, newMintingRewardsData.AllowedCasset)
			} else if mintingReward.IsActive && mintingReward.Depositor != nil && mintingReward.StartTimestamp.Before(ctx.BlockTime()) {
				return sdkerrors.Wrapf(types.ErrorMintingRewardPairAlreadyExist, "reward pair %s already exists, rewards are in progress", newMintingRewardsData.AllowedCollateral, newMintingRewardsData.AllowedCasset)
			}
		}
	}
	newMintingRewardsData.Id = k.GetMintingRewardsID(ctx) + 1
	k.SetMintingRewardsID(ctx, newMintingRewardsData.Id)
	k.SetMintingRewards(ctx, newMintingRewardsData)

	return nil
}

func (k Keeper) TransferDeposits(ctx sdk.Context, mintingRewardsId uint64, from sdk.AccAddress, startTimeStamp time.Time) error {
	mintingReward, found := k.GetMintingReward(ctx, mintingRewardsId)
	if !found {
		return types.ErrorMintingRewardNotFound
	}
	if mintingReward.IsActive {
		return types.ErrorMintingRewardAlreadyActive
	}
	// change minimum halt require after deposit
	if startTimeStamp.Before(ctx.BlockTime().Add(time.Minute * 1)) {
		return types.ErrorInvalidStartTime
	}
	if mintingReward.Depositor != nil {
		if mintingReward.StartTimestamp.After(ctx.BlockTime()) {
			return types.ErrorDepositAlreadyMade
		}
		return types.ErrorMintingRewardExpired
	}
	err := k.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, sdk.NewCoins(mintingReward.TotalRewards))
	if err != nil {
		return err
	}
	mintingReward.StartTimestamp = startTimeStamp
	// mintingReward.EndTimestamp = startTimeStamp.Add(time.Hour * 24 * time.Duration(mintingReward.DurationDays))
	mintingReward.EndTimestamp = startTimeStamp.Add(time.Minute * 5)
	mintingReward.AvailableRewards = mintingReward.TotalRewards
	mintingReward.Depositor = from
	k.SetMintingRewards(ctx, mintingReward)
	return nil
}

func (k Keeper) UpdateMintRewardStartTime(ctx sdk.Context, mintingRewardsId uint64, from sdk.AccAddress, newStartTimeStamp time.Time) error {
	mintingReward, found := k.GetMintingReward(ctx, mintingRewardsId)
	if !found {
		return types.ErrorMintingRewardNotFound
	}
	if !mintingReward.Depositor.Equals(from) {
		return types.ErrorUnauthorized
	}
	if mintingReward.IsActive {
		return types.ErrorMintingRewardAlreadyActive
	}
	if mintingReward.StartTimestamp.Before(ctx.BlockTime()) {
		return types.ErrorMintingRewardExpired
	}
	// change minimum halt require after deposit
	if newStartTimeStamp.Before(ctx.BlockTime().Add(time.Minute * 1)) {
		return types.ErrorInvalidStartTime
	}
	mintingReward.StartTimestamp = newStartTimeStamp
	// mintingReward.EndTimestamp = newStartTimeStamp.Add(time.Hour * 24 * time.Duration(mintingReward.DurationDays))
	mintingReward.EndTimestamp = newStartTimeStamp.Add(time.Minute * 5)
	k.SetMintingRewards(ctx, mintingReward)
	return nil
}

func (k Keeper) DisableMintingReward(ctx sdk.Context, mintingRewardsId uint64) error {
	mintingReward, found := k.GetMintingReward(ctx, mintingRewardsId)
	if !found {
		return types.ErrorMintingRewardNotFound
	}
	if !mintingReward.IsActive {
		return types.ErrorMintingRewardAlreadyDisabled
	}
	// Add Event Emitters
	err := k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mintingReward.Depositor, sdk.NewCoins(mintingReward.AvailableRewards))
	if err != nil {
		return err
	}
	mintingReward.IsActive = false
	k.SetMintingRewards(ctx, mintingReward)
	return nil
}

func (k Keeper) EnableMintingRewards(ctx sdk.Context) {
	mintingRewards := k.GetMintingRewards(ctx)
	for _, mintingReward := range mintingRewards {
		if !mintingReward.IsActive && mintingReward.Depositor != nil {
			diff := mintingReward.StartTimestamp.Sub(ctx.BlockTime()).Seconds()
			if diff >= -10 && diff <= 10 {
				mintingReward.IsActive = true
				k.SetMintingRewards(ctx, mintingReward)
			}
		}
	}
}

func (k Keeper) DisableMintingRewards(ctx sdk.Context) {
	mintingRewards := k.GetMintingRewards(ctx)
	for _, mintingReward := range mintingRewards {
		if mintingReward.IsActive && mintingReward.Depositor != nil && mintingReward.EndTimestamp.Before(ctx.BlockTime()) {
			k.DisableMintingReward(ctx, mintingReward.Id)
		}
	}
}

func (k Keeper) DistributeRewards(ctx sdk.Context, mintingReward types.MintingRewards) {
	fmt.Println("Active rewards.....", mintingReward)
}

func (k Keeper) TriggerRewards(ctx sdk.Context) {
	if ctx.BlockHeight()%20 == 0 {
		mintingRewards := k.GetMintingRewards(ctx)
		for _, mintingReward := range mintingRewards {
			if mintingReward.IsActive {
				k.DistributeRewards(ctx, mintingReward)
			}
		}
	}
}
