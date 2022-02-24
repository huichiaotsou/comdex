package keeper_test

import (
	"testing"

	"github.com/comdex-official/comdex/app"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	"github.com/comdex-official/comdex/x/oracle/keeper"
	"github.com/comdex-official/comdex/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

type MsgTestSuite struct {
	suite.Suite
	keeper      keeper.Keeper
	assetKeeper assetkeeper.Keeper
	app         app.TestApp
	ctx         sdk.Context
}

func (suite *MsgTestSuite) SetupTest() {
	testApp := app.NewTestApp()
	k := testApp.GetOracleKeeper()
	ak := testApp.GetAssetKeeper()
	ctx := testApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.app = testApp
	suite.keeper = k
	suite.assetKeeper = ak
	suite.ctx = ctx
}

func TestMsgTestSuite(t *testing.T) {
	suite.Run(t, new(MsgTestSuite))
}

func (suite *MsgTestSuite) TestMsgRemoveMarketForAsset() {
	app.SetAccountAddressPrefixes()
	msgreq := types.MsgRemoveMarketForAssetRequest{
		From:   suite.assetKeeper.Admin(suite.ctx),
		Id:     1,
		Symbol: "ucmdx",
	}

	msgServer := keeper.NewMsgServiceServer(suite.keeper)
	_, err := msgServer.MsgRemoveMarketForAsset(sdk.WrapSDKContext(suite.ctx), &msgreq)
	suite.Error(err)

	suite.keeper.SetMarketForAsset(suite.ctx, 1, "ucmdx")
	_, err = msgServer.MsgRemoveMarketForAsset(sdk.WrapSDKContext(suite.ctx), &msgreq)
	suite.NoError(err)

}

func (suite *MsgTestSuite) TestMsgAddMarket() {
	app.SetAccountAddressPrefixes()
	msgreq := types.MsgAddMarketRequest{
		From:     suite.assetKeeper.Admin(suite.ctx),
		Id:       1,
		Symbol:   "ucmdx",
		ScriptID: 1,
		Rates:    100,
	}

	msgServer := keeper.NewMsgServiceServer(suite.keeper)
	_, err := msgServer.MsgAddMarket(sdk.WrapSDKContext(suite.ctx), &msgreq)
	suite.NoError(err)

}

func (suite *MsgTestSuite) TestMsgUpdateMarket() {
	app.SetAccountAddressPrefixes()
	msgreq := types.MsgUpdateMarketRequest{
		From:     suite.assetKeeper.Admin(suite.ctx),
		Symbol:   "ucmdx",
		ScriptID: 1,
		Rates:    100,
	}

	msgServer := keeper.NewMsgServiceServer(suite.keeper)
	_, err := msgServer.MsgUpdateMarket(sdk.WrapSDKContext(suite.ctx), &msgreq)
	suite.NoError(err)

}
