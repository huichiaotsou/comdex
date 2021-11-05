package cli

import (
	"github.com/comdex-official/comdex/x/bandoracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcchannelclientutils "github.com/cosmos/ibc-go/modules/core/04-channel/client/utils"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

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

/*func txAddMarketForAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-market-for-asset [asset] [symbol]",
		Short: "Add a market for asset",
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

			msg := types.NewMsgAddMarketForAssetRequest(
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
}*/

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

func txFetchPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch-price [source-port] [source-channel] [symbols] [script-id]",
		Short: "Fetch price from Oracle",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			scriptID, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			timeoutHeight, err := GetPacketTimeoutHeight(cmd)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			absoluteTimeouts, err := cmd.Flags().GetBool(flagAbsoluteTimeouts)
			if err != nil {
				return err
			}

			feeLimit, err := GetFeeLimit(cmd)
			if err != nil {
				return err
			}

			prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
			if err != nil {
				return err
			}

			executeGas, err := cmd.Flags().GetUint64(flagExecuteGas)
			if err != nil {
				return err
			}

			if !absoluteTimeouts {
				state, height, _, err := ibcchannelclientutils.QueryLatestConsensusState(ctx, args[0], args[1])
				if err != nil {
					return err
				}

				if !timeoutHeight.IsZero() {
					timeoutHeight.RevisionHeight += height.RevisionHeight
					timeoutHeight.RevisionNumber += height.RevisionNumber
				}
				if timeoutTimestamp != 0 {
					timeoutTimestamp += state.GetTimestamp()
				}
			}

			msg := types.NewMsgFetchPriceRequest(
				ctx.FromAddress,
				args[0],
				args[1],
				timeoutHeight,
				timeoutTimestamp,
				strings.Split(args[2], ","),
				scriptID,
				feeLimit,
				prepareGas,
				executeGas,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagPacketTimeoutHeight, "0-1000", "packet timeout block height")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, 300, "packet timeout timestamp")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "timeout flags are used as absolute timeouts")
	cmd.Flags().String(flagFeeLimit, "", "fee limit")
	cmd.Flags().Uint64(flagPrepareGas, 0, "prepare gas")
	cmd.Flags().Uint64(flagExecuteGas, 0, "execute gas")

	return cmd
}

func CmdRequestGoldPriceData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gold-price-data [oracle-script-id] [requested-validator-count] [sufficient-validator-count]",
		Short: "Make a new GoldPrice query request via an existing BandChain oracle script",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// retrieve the oracle script id.
			uint64OracleScriptID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			oracleScriptID := types.OracleScriptID(uint64OracleScriptID)

			// retrieve the requested validator count.
			askCount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// retrieve the sufficient(minimum) validator count.
			minCount, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			channel, err := cmd.Flags().GetString(flagChannel)
			if err != nil {
				return err
			}

			// retrieve the multiplier for the symbols' price.
			multiplier, err := cmd.Flags().GetUint64(flagMultiplier)
			if err != nil {
				return err
			}

			calldata := &types.GoldPriceCallData{
				Multiplier: multiplier,
			}

			// retrieve the amount of coins allowed to be paid for oracle request fee from the pool account.
			coinStr, err := cmd.Flags().GetString(flagFeeLimit)
			if err != nil {
				return err
			}
			feeLimit, err := sdk.ParseCoinsNormalized(coinStr)
			if err != nil {
				return err
			}

			// retrieve the amount of gas allowed for the prepare step of the oracle script.
			prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
			if err != nil {
				return err
			}

			// retrieve the amount of gas allowed for the execute step of the oracle script.
			executeGas, err := cmd.Flags().GetUint64(flagExecuteGas)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGoldPriceData(
				clientCtx.GetFromAddress().String(),
				oracleScriptID,
				channel,
				calldata,
				askCount,
				minCount,
				feeLimit,
				prepareGas,
				executeGas,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagChannel, "", "The channel id")
	cmd.MarkFlagRequired(flagChannel)
	cmd.Flags().StringSlice(flagSymbols, nil, "Symbols used in calling the oracle script")
	cmd.Flags().Uint64(flagMultiplier, 1000000, "Multiplier used in calling the oracle script")
	cmd.Flags().String(flagFeeLimit, "", "the maximum tokens that will be paid to all data source providers")
	cmd.Flags().String(flagRequestkey, "", "Key for generating escrow address")
	cmd.Flags().Uint64(flagPrepareGas, 200000, "Prepare gas used in fee counting for prepare request")
	cmd.Flags().Uint64(flagExecuteGas, 200000, "Execute gas used in fee counting for execute request")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}