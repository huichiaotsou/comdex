package keeper

import (
	"github.com/comdex-official/comdex/x/loan/types"
)

var _ types.QueryServer = Keeper{}
