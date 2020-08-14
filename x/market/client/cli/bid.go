package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ovrclk/akash/x/market/types"
	"github.com/spf13/cobra"
)

func cmdGetBids() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query for all bids",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			bfilters, state, err := BidFiltersFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			// checking state flag
			stateVal, ok := types.Bid_State_value[state]

			if (!ok && (state != "")) || state == "invalid" {
				return ErrStateValue
			}

			bfilters.State = types.Bid_State(stateVal)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryBidsRequest{
				Filters:    bfilters,
				Pagination: pageReq,
			}

			res, err := queryClient.Bids(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutputLegacy(res.Bids)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "bids")
	AddBidFilterFlags(cmd.Flags())
	return cmd
}

func cmdGetBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Query order",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			bidID, err := BidIDFromFlagsWithoutCtx(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Bid(context.Background(), &types.QueryBidRequest{ID: bidID})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res.Bid)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	AddQueryBidIDFlags(cmd.Flags())
	MarkReqBidIDFlags(cmd)

	return cmd
}
