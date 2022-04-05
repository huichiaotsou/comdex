package rest

import (
	"github.com/comdex-official/comdex/x/liquidity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"net/http"
)

type AddUnbondingDuration struct {
	BaseReq     rest.BaseReq  `json:"base_req" yaml:"base_req"`
	Title       string        `json:"title" yaml:"title"`
	Description string        `json:"description" yaml:"description"`
	Deposit     sdk.Coins     `json:"deposit" yaml:"deposit"`
	UnbondingDuration string  `json:"unbondingDuration" yaml:"liquidity"`
}

func AddNewUnbondingPeriodProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "add-unbonding-duration",
		Handler:  AddUnbondingDurationRESTHandler(clientCtx),
	}
}

func AddUnbondingDurationRESTHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddUnbondingDuration

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewUpdateUnbondingDuration(
			req.Title,
			req.Description,
			req.UnbondingDuration)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
