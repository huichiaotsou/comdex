package cli

import (
	"context"
	"github.com/comdex-official/comdex/x/bandoracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdGoldPriceResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gold-price-result [request-id]",
		Short: "Query the GoldPrice result data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryServiceClient(clientCtx)
			r, err := queryClient.GoldPriceResult(context.Background(), &types.QueryGoldPriceRequest{RequestId: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdLastGoldPriceID queries latest request
func CmdLastGoldPriceID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last-gold-price-id",
		Short: "Query the last request id returned by GoldPrice ack packet",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryServiceClient(clientCtx)
			r, err := queryClient.LastGoldPriceId(context.Background(), &types.QueryLastGoldPriceIdRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}