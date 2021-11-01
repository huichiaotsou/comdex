package bandoracle

import (
	"fmt"
	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	"github.com/comdex-official/comdex/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"strconv"
)

// handleOraclePacket handles the result of the received BandChain oracles
// packet and saves the data into the KV database
func (am AppModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData packet.OracleResponsePacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}

	id, err := strconv.ParseUint(modulePacketData.ClientID, 10, 64)
	if err != nil {
		return channeltypes.Acknowledgement{},err
	}

	if modulePacketData.ResolveStatus == packet.RESOLVE_STATUS_SUCCESS {
		calldata, found := am.keeper.GetCalldata(ctx, id)
		if !found {
			return channeltypes.Acknowledgement{},fmt.Errorf("calldata does not exist for id %d", id)
		}

		var result types.Result
		if err := obi.Decode(modulePacketData.Result, &result); err != nil {
			return channeltypes.Acknowledgement{}, err
		}

		for i := range calldata.Symbols {
			am.keeper.SetPriceForMarket(ctx, calldata.Symbols[i], result.Rates[i])
		}
	}

	am.keeper.DeleteCalldata(ctx, id)
	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			packet.NewOracleRequestPacketAcknowledgement(modulePacketData.RequestID),
		),
	)
	return ack, nil
}

// handleOracleAcknowledgment handles the acknowledgment result from the BandChain
// request and saves the request-id into the KV database
func (am AppModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck packet.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		var data packet.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}
		requestID := types.OracleRequestID(oracleAck.RequestID)

		switch data.GetClientID() {

		case types.CoinRatesClientIDKey:
			var coinRatesData types.CoinRatesCallData
			if err = obi.Decode(data.GetCalldata(), &coinRatesData); err != nil {
				return nil, sdkerrors.Wrap(err,
					"cannot decode the coinRates oracle acknowledgment packet")
			}
			am.keeper.SetLastCoinRatesID(ctx, requestID)
			return &sdk.Result{}, nil
			// this line is used by starport scaffolding # oracle/module/ack

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"oracle acknowledgment packet not found: %s", data.GetClientID())
		}
	}
	return nil, nil
}
