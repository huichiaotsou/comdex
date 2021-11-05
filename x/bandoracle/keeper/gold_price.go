package keeper

import (
	"github.com/comdex-official/comdex/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"


)

func (k Keeper) SetGoldPriceResult(ctx sdk.Context, requestID types.OracleRequestID, result types.GoldPriceResult) {
	store := ctx.KVStore(k.key)
	store.Set(types.GoldPriceResultStoreKey(requestID), k.cdc.MustMarshal(&result))
}

// GetGoldPriceResult returns the GoldPrice by requestId
func (k Keeper) GetGoldPriceResult(ctx sdk.Context, id types.OracleRequestID) (types.GoldPriceResult, error) {
	bz := ctx.KVStore(k.key).Get(types.GoldPriceResultStoreKey(id))
	if bz == nil {
		return types.GoldPriceResult{}, sdkerrors.Wrapf(types.ErrSample,
			"GetResult: Result for request ID %d is not available.", id,
		)
	}
	var result types.GoldPriceResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, nil
}

// GetLastGoldPriceID return the id from the last GoldPrice request
func (k Keeper) GetLastGoldPriceID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.key).Get(types.KeyPrefix(types.LastGoldPriceIDKey))
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

// SetLastGoldPriceID saves the id from the last GoldPrice request
func (k Keeper) SetLastGoldPriceID(ctx sdk.Context, id types.OracleRequestID) {
	store := ctx.KVStore(k.key)
	store.Set(types.KeyPrefix(types.LastGoldPriceIDKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: int64(id)}))
}
