package keeper

import (
	"github.com/comdex-official/comdex/x/liquidation/types"
)

var _ types.QueryServer = Keeper{}
