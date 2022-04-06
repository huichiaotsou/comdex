package client

import (
	"github.com/comdex-official/comdex/x/liquidation/client/cli"
	"github.com/comdex-official/comdex/x/liquidation/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var AddUnliqPointPercentage = govclient.NewProposalHandler(cli.NewCmdSubmitUnliquidatePointPercentage, rest.AddNewUnliquidatePointPercentageProposalRESTHandler)
