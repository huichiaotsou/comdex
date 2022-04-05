package client

import (
	"github.com/comdex-official/comdex/x/liquidity/client/cli"
	"github.com/comdex-official/comdex/x/liquidity/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var AddUnbondingDurationHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUnbondingDuration, rest.AddNewUnbondingPeriodProposalRESTHandler)
