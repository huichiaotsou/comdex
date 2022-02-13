package app

import (
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	liquidationkeeper "github.com/comdex-official/comdex/x/liquidation/keeper"
	oraclekeeper "github.com/comdex-official/comdex/x/oracle/keeper"
	vaultkeeper "github.com/comdex-official/comdex/x/vault/keeper"
	"github.com/cosmos/cosmos-sdk/simapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

type TestApp struct {
	App
}

// account expected.AccountKeeper,
// 	bank expected.BankKeeper,
// 	asset expected.AssetKeeper,
// 	vault expected.VaultKeeper,
// 	oracle expected.OracleKeeper,

func NewTestApp() TestApp {

	encoding := MakeEncodingConfig()
	db := tmdb.NewMemDB()
	app := New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, string(" "), 0, encoding, simapp.EmptyAppOptions{})
	return TestApp{App: *app}
}

func (tApp TestApp) GetVaultKeeper() vaultkeeper.Keeper             { return tApp.vaultKeeper }
func (tApp TestApp) GetAssetKeeper() assetkeeper.Keeper             { return tApp.assetKeeper }
func (tApp TestApp) GetOracleKeeper() oraclekeeper.Keeper           { return tApp.oracleKeeper }
func (tApp TestApp) GetBankKeeper() bankkeeper.Keeper               { return tApp.bankKeeper }
func (tApp TestApp) GetAccountKeeper() authkeeper.AccountKeeper     { return tApp.accountKeeper }
func (tApp TestApp) GetLiquidationKeeper() liquidationkeeper.Keeper { return tApp.liquidationKeeper }
