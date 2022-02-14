package keeper_test

// account expected.AccountKeeper,
// 	bank expected.BankKeeper,
// 	asset expected.AssetKeeper,
// 	vault expected.VaultKeeper,
// 	oracle expected.OracleKeeper,

import (
	"testing"

	"github.com/comdex-official/comdex/app"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	types3 "github.com/comdex-official/comdex/x/asset/types"
	types1 "github.com/comdex-official/comdex/x/liquidation/types"
	oraclekeeper "github.com/comdex-official/comdex/x/oracle/keeper"
	vaultkeeper "github.com/comdex-official/comdex/x/vault/keeper"
	types2 "github.com/comdex-official/comdex/x/vault/types"
	types "github.com/cosmos/cosmos-sdk/x/capability/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"time"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	"github.com/stretchr/testify/suite"
)

type LiquidateVaultTestSuite struct {
	suite.Suite
	keeper        keeper.Keeper
	assetKeeper   assetkeeper.Keeper
	oracleKeeper  oraclekeeper.Keeper
	vaultKeeper   vaultkeeper.Keeper
	bankKeeper    bankkeeper.Keeper
	accountKeeper accountkeeper.AccountKeeper
	app           app.TestApp
	ctx           sdk.Context
}

func (suite *LiquidateVaultTestSuite) SetupTest() {
	testApp := app.NewTestApp()
	k := testApp.GetLiquidationKeeper()
	ak := testApp.GetAssetKeeper()
	ok := testApp.GetOracleKeeper()
	vk := testApp.GetVaultKeeper()
	acck := testApp.GetAccountKeeper()
	bk := testApp.GetBankKeeper()
	ctx := testApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.app = testApp
	suite.keeper = k
	suite.assetKeeper = ak
	suite.oracleKeeper = ok
	suite.vaultKeeper = vk
	suite.bankKeeper = bk
	suite.accountKeeper = acck
	suite.ctx = ctx
}

func TestVaultTestSuite(t *testing.T) {
	suite.Run(t, new(LiquidateVaultTestSuite))
}

func (suite *LiquidateVaultTestSuite) TestLockedVault_SetGet() {
	suite.SetupTest()
	locked_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, locked_vault)

	t, found := suite.keeper.GetLockedVault(suite.ctx, types.DefaultIndex)
	suite.True(found)
	suite.Equal(locked_vault, t)
	_, found = suite.keeper.GetLockedVault(suite.ctx, 100)
	suite.False(found)
	suite.keeper.DeleteLockedVault(suite.ctx, 100)
	_, found = suite.keeper.GetLockedVault(suite.ctx, 100)
	suite.False(found)

}

func (suite *LiquidateVaultTestSuite) TestLockedVaultID_SetGet() {
	suite.SetupTest()
	suite.keeper.SetLockedVaultID(suite.ctx, types.DefaultIndex)
	id := suite.keeper.GetLockedVaultID(suite.ctx)
	suite.Equal(types.DefaultIndex, id)
}

func (suite *LiquidateVaultTestSuite) TestGetLockedVaults() {
	suite.SetupTest()
	locked_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, locked_vault)

	locked_vaults := suite.keeper.GetLockedVaults(suite.ctx)
	if locked_vaults == nil {
		print(suite)
	}
}

func (suite *LiquidateVaultTestSuite) TestGetLockedVault() {
	suite.SetupTest()
	locked_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, locked_vault)
	_, found := suite.keeper.GetLockedVault(suite.ctx, locked_vault.LockedVaultId)
	suite.NotEqual(found, false)
}

func (suite *LiquidateVaultTestSuite) TestDeleteLockedVault() {
	suite.SetupTest()
	locked_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, locked_vault)
	suite.keeper.DeleteLockedVault(suite.ctx, types.DefaultIndex)
	_, found := suite.keeper.GetLockedVault(suite.ctx, types.DefaultIndex)
	suite.False(found)
}

// func (k *Keeper) SetFlagIsAuctionInProgress(ctx sdk.Context, id uint64, flag bool) error {

// 	locked_vault, found := k.GetLockedVault(ctx, id)
// 	if !found {
// 		return types.LockedVaultDoesNotExist
// 	}
// 	locked_vault.IsAuctionInProgress = flag
// 	k.SetLockedVault(ctx, locked_vault)
// 	return nil
// }
func (suite *LiquidateVaultTestSuite) TestSetFlagIsAuction_InProgressComplete() {
	suite.SetupTest()
	//define dummy locked_vault
	locked_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	//set the vault
	suite.keeper.SetLockedVault(suite.ctx, locked_vault)

	//change the param for the gotten locked vault
	err := suite.keeper.SetFlagIsAuctionInProgress(suite.ctx, locked_vault.LockedVaultId, true)
	suite.Equal(err, nil)
	err = suite.keeper.SetFlagIsAuctionComplete(suite.ctx, locked_vault.LockedVaultId, true)
	suite.Equal(err, nil)

	//get the vault
	locked_vault1, found := suite.keeper.GetLockedVault(suite.ctx, types.DefaultIndex)
	suite.True(found)

	//change the resp param in our dummy locked_vault
	locked_vault.IsAuctionInProgress = true
	locked_vault.IsAuctionComplete = true

	//getlockedvault and dummy vault should be equal
	suite.Equal(locked_vault, locked_vault1)

	//setting flag for a random vault id
	err = suite.keeper.SetFlagIsAuctionInProgress(suite.ctx, 100, true)
	suite.NotEqual(err, nil)

	err = suite.keeper.SetFlagIsAuctionComplete(suite.ctx, 100, true)
	suite.NotEqual(err, nil)
}

func (suite *LiquidateVaultTestSuite) TestCreateLockedVault() {
	suite.SetupTest()
	vault := types2.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(1000000),
		AmountOut: sdk.NewInt(1000000),
	}

	locked_vaultId := suite.keeper.GetLockedVaultID(suite.ctx)
	suite.keeper.SetLockedVaultID(suite.ctx, locked_vaultId+1)

	err := suite.keeper.CreateLockedVault(suite.ctx, vault, sdk.MustNewDecFromStr("1.0"))
	suite.Equal(err, nil)

	locked_vault, found := suite.keeper.GetLockedVault(suite.ctx, locked_vaultId+1)
	suite.Equal(found, true)

	dummy_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "abc",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.NewDec(1.0),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}
	suite.Equal(locked_vault, dummy_vault)

	suite.Equal(suite.keeper.GetLockedVaultID(suite.ctx), locked_vaultId+2)

}

func (suite *LiquidateVaultTestSuite) TestLiquidateVaults() {
	suite.SetupTest()
	assetin := types3.Asset{
		Id:       1,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 1000000,
	}

	assetout := types3.Asset{
		Id:       2,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 500000,
	}

	suite.assetKeeper.SetAsset(suite.ctx, assetin)
	suite.assetKeeper.SetAsset(suite.ctx, assetout)

	pair := types3.Pair{
		Id:               1,
		AssetIn:          1,
		AssetOut:         2,
		LiquidationRatio: sdk.MustNewDecFromStr("1.5"),
	}

	suite.assetKeeper.SetPair(suite.ctx, pair)

	vault := types2.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(1000000),
		AmountOut: sdk.NewInt(1000000),
	}

	suite.vaultKeeper.SetVault(suite.ctx, vault)

	lockedvaultId := suite.keeper.GetLockedVaultID(suite.ctx)

	err := suite.keeper.LiquidateVaults(suite.ctx)
	suite.Equal(err, nil)

	locked_vault1, found := suite.keeper.GetLockedVault(suite.ctx, lockedvaultId)
	suite.Equal(found, true)

	crAtLiquidation := locked_vault1.CrAtLiquidation
	suite.Equal(sdk.MustNewDecFromStr("1.0"), crAtLiquidation)

	_, found = suite.vaultKeeper.GetVault(suite.ctx, 1)
	suite.Equal(found, false)

	suite.SetupTest()

	assetin = types3.Asset{
		Id:       3,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 1000000,
	}

	assetout = types3.Asset{
		Id:       4,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 1000000,
	}

	suite.assetKeeper.SetAsset(suite.ctx, assetin)
	suite.assetKeeper.SetAsset(suite.ctx, assetout)

	pair = types3.Pair{
		Id:               2,
		AssetIn:          3,
		AssetOut:         4,
		LiquidationRatio: sdk.MustNewDecFromStr("1.5"),
	}

	suite.assetKeeper.SetPair(suite.ctx, pair)

	vault = types2.Vault{
		ID:        2,
		PairID:    2,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(1000000),
		AmountOut: sdk.NewInt(1000000),
	}

	suite.vaultKeeper.SetVault(suite.ctx, vault)

	lockedvaultId = suite.keeper.GetLockedVaultID(suite.ctx)
	err = suite.keeper.LiquidateVaults(suite.ctx)
	suite.Equal(err, nil)

	_, found = suite.keeper.GetLockedVault(suite.ctx, lockedvaultId)
	suite.Equal(found, false)

	_, found = suite.vaultKeeper.GetVault(suite.ctx, 2)
	suite.Equal(found, true)

}

func (suite *LiquidateVaultTestSuite) TestUnliquidateLockedVaults() {
	suite.SetupTest()
	err := suite.keeper.UnliquidateLockedVaults(suite.ctx)
	suite.Equal(err, nil)

	dummy_vault := types1.LockedVault{
		LockedVaultId:                types.DefaultIndex,
		OriginalVaultId:              types.DefaultIndex,
		PairId:                       1,
		Owner:                        "comdex1pwu5sjk2lje94",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            true,
		IsAuctionInProgress:          true,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.MustNewDecFromStr("1.6"),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, dummy_vault)
	suite.keeper.SetLockedVaultID(suite.ctx, 1)

	vault_id := suite.keeper.GetVaultID(suite.ctx)
	err = suite.keeper.UnliquidateLockedVaults(suite.ctx)
	suite.Equal(err, nil)

	suite.Equal(vault_id+1, suite.keeper.GetVaultID(suite.ctx))

	vault_we_should_get := types2.Vault{
		ID:        vault_id + 1,
		PairID:    dummy_vault.PairId,
		Owner:     dummy_vault.Owner,
		AmountIn:  dummy_vault.AmountIn,
		AmountOut: dummy_vault.AmountOut,
	}

	vaults_set_by_func := suite.keeper.GetVaults(suite.ctx)
	suite.Equal(len(vaults_set_by_func), 1)

	suite.Equal(vaults_set_by_func[0], vault_we_should_get)

	_, found := suite.keeper.GetLockedVault(suite.ctx, types.DefaultIndex)
	suite.False(found)

	found = suite.keeper.HasVaultForAddressByPair(suite.ctx, sdk.AccAddress(dummy_vault.Owner), dummy_vault.PairId)
	suite.True(found)

	//tests for the false case
	suite.SetupTest()

	locked_vault_id := suite.keeper.GetLockedVaultID(suite.ctx)
	dummy_vault = types1.LockedVault{
		LockedVaultId:                1,
		OriginalVaultId:              1,
		PairId:                       locked_vault_id,
		Owner:                        "comdex1pwu5sjk2lje94",
		AmountIn:                     sdk.NewInt(1000000),
		AmountOut:                    sdk.NewInt(1000000),
		Initiator:                    types1.ModuleName,
		IsAuctionComplete:            false,
		IsAuctionInProgress:          false,
		CrAtLiquidation:              sdk.NewDec(1.0),
		CurrentCollaterlisationRatio: sdk.MustNewDecFromStr("1.6"),
		CollateralToBeAuctioned:      nil,
		LiquidationTimestamp:         time.Time{},
		SellOffHistory:               nil,
	}

	suite.keeper.SetLockedVault(suite.ctx, dummy_vault)

	vault_id = suite.keeper.GetVaultID(suite.ctx)
	err = suite.keeper.UnliquidateLockedVaults(suite.ctx)
	suite.Equal(err, nil)

	suite.Equal(vault_id, suite.keeper.GetVaultID(suite.ctx))

	vaults_set_by_func = suite.keeper.GetVaults(suite.ctx)
	suite.Equal(len(vaults_set_by_func), 0)

	lockedVaults := suite.keeper.GetLockedVaults(suite.ctx)
	suite.Equal(lockedVaults[0], dummy_vault)

	found = suite.keeper.HasVaultForAddressByPair(suite.ctx, sdk.AccAddress(dummy_vault.Owner), 1)
	suite.False(found)

}

func (suite *LiquidateVaultTestSuite) TestUpdateLockedVaults() {
	suite.SetupTest()
	//empty keeper
	vault_dummy := types2.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(120),
		AmountOut: sdk.NewInt(500),
	}
	_ = suite.keeper.CreateLockedVault(suite.ctx, vault_dummy, sdk.MustNewDecFromStr("1.5"))
	LV_before, _ := suite.keeper.GetLockedVault(suite.ctx, 0)
	err := suite.keeper.UpdateLockedVaults(suite.ctx)
	suite.Equal(err, nil)
	LV_after, _ := suite.keeper.GetLockedVault(suite.ctx, 0)
	suite.Equal(LV_before.CollateralToBeAuctioned, LV_after.CollateralToBeAuctioned)

	//setup assets and vaults
	suite.SetupTest()
	assetin := types3.Asset{
		Id:       1,
		Name:     "GOLD",
		Denom:    "ucGOLD",
		Decimals: 100,
	}

	assetout := types3.Asset{
		Id:       2,
		Name:     "SILVER",
		Denom:    "ucSILVER",
		Decimals: 500,
	}

	suite.assetKeeper.SetAsset(suite.ctx, assetin)
	suite.assetKeeper.SetAsset(suite.ctx, assetout)

	pair := types3.Pair{
		Id:               1,
		AssetIn:          1,
		AssetOut:         2,
		LiquidationRatio: sdk.MustNewDecFromStr("1.5"),
	}

	suite.assetKeeper.SetPair(suite.ctx, pair)

	vault := types2.Vault{
		ID:        1,
		PairID:    1,
		Owner:     "abc",
		AmountIn:  sdk.NewInt(120),
		AmountOut: sdk.NewInt(500),
	}
	suite.vaultKeeper.SetVault(suite.ctx, vault)

	locked_vault_id := suite.keeper.GetLockedVaultID(suite.ctx)
	err = suite.keeper.CreateLockedVault(suite.ctx, vault, sdk.MustNewDecFromStr("1.5"))
	suite.Equal(err, nil)

	err = suite.keeper.UpdateLockedVaults(suite.ctx)
	suite.Equal(err, nil)
	locked_vault, _ := suite.keeper.GetLockedVault(suite.ctx, locked_vault_id)
	selloff_expected := sdk.MustNewDecFromStr("12000")
	suite.Equal(locked_vault.CollateralToBeAuctioned, &selloff_expected)

}
