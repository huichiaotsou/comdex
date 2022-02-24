package keeper_test

import (
	"testing"

	"github.com/comdex-official/comdex/app"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	"github.com/comdex-official/comdex/x/oracle/keeper"
	"github.com/comdex-official/comdex/x/oracle/types"
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
	k := testApp.GetOracleKeeper()
	ak := testApp.GetAssetKeeper()
	ctx := testApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.app = testApp
	suite.keeper = k
	suite.assetKeeper = ak
	suite.ctx = ctx
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

func (suite *QueryTestSuite) TestQueryMarkets() {
	qserver := keeper.NewQueryServiceServer(suite.keeper)

	_, err := qserver.QueryMarkets(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	qvr := types.QueryMarketsRequest{
		Pagination: &suite.prq,
	}
	_, err = qserver.QueryMarkets(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.NoError(err)

}

func (suite *QueryTestSuite) TestQueryMarket() {
	qserver := keeper.NewQueryServiceServer(suite.keeper)

	_, err := qserver.QueryMarket(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	qvr := types.QueryMarketRequest{
		Symbol: "ucmdx",
	}
	_, err = qserver.QueryMarket(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.NoError(err)

}
func (suite *QueryTestSuite) TestQueryParams() {
	qserver := keeper.NewQueryServiceServer(suite.keeper)

	_, err := qserver.QueryParams(sdk.WrapSDKContext(suite.ctx), nil)
	suite.Error(err)

	qvr := types.QueryParamsRequest{}
	_, err = qserver.QueryParams(sdk.WrapSDKContext(suite.ctx), &qvr)
	suite.NoError(err)

}
