package keeper_test

import (
	"testing"

	"github.com/comdex-official/comdex/app"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/vault/keeper"
	"github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

type QueryTestSuite struct {
	suite.Suite
	keeper      keeper.Keeper
	assetKeeper assetkeeper.Keeper
	app         app.TestApp
	ctx         sdk.Context
	prq         query.PageRequest
}

func (suite *QueryTestSuite) SetupTest() {
	testApp := app.NewTestApp()
	k := testApp.GetVaultKeeper()
	ak := testApp.GetAssetKeeper()
	ctx := testApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.app = testApp
	suite.keeper = k
	suite.assetKeeper = ak
	suite.ctx = ctx

	return
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

func (suite *QueryTestSuite) TestQueryAllVaults() {
	qserver := keeper.NewQueryServiceServer(suite.keeper)
	qvr := types.QueryAllVaultsRequest{
		Pagination: &suite.prq,
	}
	_, err := qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	//set a vault
	vault := types.Vault{ID: 1, PairID: 1, Owner: "abc", AmountIn: sdk.NewInt(100), AmountOut: sdk.NewInt(100)}
	suite.keeper.SetVault(suite.ctx, vault)
	_, err = qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	//set asset 1 and asset 2
	asset1 := assettypes.Asset{
		Id:       1,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset1)
	_, err = qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	asset2 := assettypes.Asset{
		Id:       2,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset2)
	_, err = qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	pair := assettypes.Pair{
		Id:               1,
		AssetIn:          1,
		AssetOut:         2,
		LiquidationRatio: sdk.Dec(sdk.MustNewDecFromStr("1.5")),
	}
	suite.assetKeeper.SetPair(suite.ctx, pair)
	allvaults, err := qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Equal(err, nil)
	suite.Equal(len(allvaults.VaultsInfo), 1)

	vault = types.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(100000),
		AmountOut: sdk.NewInt(100000),
	}

	suite.keeper.SetVault(suite.ctx, vault)
	allvaults, err = qserver.QueryAllVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Equal(err, nil)
	suite.Equal(len(allvaults.VaultsInfo), 1)

}

func (suite *QueryTestSuite) TestQueryVaults() {
	suite.SetupTest()
	qserver := keeper.NewQueryServiceServer(suite.keeper)
	qvr := types.QueryVaultsRequest{
		Owner:      "abc",
		Pagination: &suite.prq,
	}
	_, err := qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	//if there are no vaults there is no err
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.NoError(err)

	//set a vault
	vault := types.Vault{ID: 1, PairID: 1, Owner: "abc", AmountIn: sdk.NewInt(100), AmountOut: sdk.NewInt(100)}
	suite.keeper.SetVault(suite.ctx, vault)
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	//set asset 1 and asset 2
	asset1 := assettypes.Asset{
		Id:       1,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset1)
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	asset2 := assettypes.Asset{
		Id:       2,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset2)
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	pair := assettypes.Pair{
		Id:               1,
		AssetIn:          1,
		AssetOut:         2,
		LiquidationRatio: sdk.Dec(sdk.MustNewDecFromStr("1.5")),
	}
	suite.assetKeeper.SetPair(suite.ctx, pair)
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Equal(err, nil)

	vault = types.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "def",
		AmountIn:  sdk.NewInt(100000),
		AmountOut: sdk.NewInt(100000),
	}

	suite.keeper.SetVault(suite.ctx, vault)
	_, err = qserver.QueryVaults(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)
}

func (suite *QueryTestSuite) TestQueryVault() {
	qserver := keeper.NewQueryServiceServer(suite.keeper)
	qvr := types.QueryVaultRequest{Id: 1}

	_, err := qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	_, err = qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	vault := types.Vault{ID: 1, PairID: 1, Owner: "abc", AmountIn: sdk.NewInt(100), AmountOut: sdk.NewInt(100)}
	suite.keeper.SetVault(suite.ctx, vault)
	_, err = qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	//set asset 1 and asset 2
	asset1 := assettypes.Asset{
		Id:       1,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset1)
	_, err = qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	asset2 := assettypes.Asset{
		Id:       2,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 100000,
	}
	suite.assetKeeper.SetAsset(suite.ctx, asset2)
	_, err = qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Error(err)

	//completes the set vault earlier
	pair := assettypes.Pair{
		Id:               1,
		AssetIn:          1,
		AssetOut:         2,
		LiquidationRatio: sdk.Dec(sdk.MustNewDecFromStr("1.5")),
	}
	suite.assetKeeper.SetPair(suite.ctx, pair)
	_, err = qserver.QueryVault(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.Equal(err, nil)

}
