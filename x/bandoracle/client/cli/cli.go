package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Oracle module sub-commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		queryMarket(),
		queryMarkets(),
		queryParams(),
		CmdGoldPriceResult(),
		CmdLastGoldPriceID(),
	)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle",
		Short: "Oracle module sub-commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		txAddMarket(),
		txUpdateMarket(),
		txRemoveMarketForAsset(),
		txFetchPrice(),
		CmdRequestGoldPriceData(),
	)

	return cmd
}
