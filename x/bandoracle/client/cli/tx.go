package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/comdex-official/comdex/x/bandoracle/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdRequestFetchPriceData(),
		txAddMarket(),
		txUpdateMarket(),
		txRemoveMarketForAsset(),
		)
	// this line is used by starport scaffolding # 1

	return cmd
}

func txAddMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-market [symbol] [script-id] [asset]",
		Short: "Add a market",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			scriptID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			asset, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMarketRequest(
				ctx.FromAddress,
				args[0],
				scriptID,
				asset,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txUpdateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-market [symbol]",
		Short: "Update a market",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			scriptID, err := cmd.Flags().GetUint64(flagScriptID)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMarketRequest(
				ctx.FromAddress,
				args[0],
				scriptID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint64(flagScriptID, 0, "script identity")

	return cmd
}

func txRemoveMarketForAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-market-for-asset [asset] [symbol]",
		Short: "Remove a market for asset",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveMarketForAssetRequest(
				ctx.FromAddress,
				asset,
				args[1],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}